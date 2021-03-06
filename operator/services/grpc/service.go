package grpc

import (
	go_context "context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	"github.com/containers-ai/alameda/operator/pkg/controller/alamedascaler"
	"github.com/containers-ai/alameda/operator/pkg/kubernetes/metrics"
	"github.com/containers-ai/alameda/operator/pkg/kubernetes/metrics/prometheus"
	logUtil "github.com/containers-ai/alameda/operator/pkg/utils/log"
	operator_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/operator"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Service struct {
	err    chan error
	server *grpc.Server

	Config    Config
	Manager   manager.Manager
	MetricsDB metrics.MetricsDB
}

var (
	scope = logUtil.RegisterScope("gRPC", "gRPC server log", 0)
)

func NewService(c *Config, manager manager.Manager) *Service {

	s := &Service{
		err: make(chan error),

		Config:  *c,
		Manager: manager,
	}

	// New Prometheus as metrics database
	db := prometheus.New(*c.Prometheus)
	s.MetricsDB = db

	return s
}

func (s *Service) Open() error {

	// Open metrics database
	if err := s.MetricsDB.Connect(); err != nil {
		return err
	}

	// build server listener
	scope.Info(("starting gRPC server"))
	ln, err := net.Listen("tcp", s.Config.BindAddress)
	if err != nil {
		scope.Error("gRPC server failed listen: " + err.Error())
		return fmt.Errorf("GRPC server failed to bind address: %s", s.Config.BindAddress)
	}
	scope.Info("gRPC server listening on " + s.Config.BindAddress)

	// build gRPC server
	server, err := s.newGRPCServer()
	if err != nil {
		scope.Error(err.Error())
		return err
	}
	s.server = server

	// register gRPC server
	s.registGRPCServer(server)
	reflection.Register(server)

	// run gRPC server
	if err := server.Serve(ln); err != nil {
		s.err <- fmt.Errorf("GRPC server failed to serve: %s", err.Error())
	}

	return nil
}

func (s *Service) newGRPCServer() (*grpc.Server, error) {

	var (
		server *grpc.Server
	)

	server = grpc.NewServer()

	return server, nil
}

func (s *Service) registGRPCServer(server *grpc.Server) {

	operator_v1alpha1.RegisterOperatorServiceServer(server, s)
}

func (s *Service) Close() error {

	if err := s.MetricsDB.Close(); err != nil {
		return err
	}

	s.server.Stop()

	return nil
}

func (s *Service) Err() <-chan error {
	return s.err
}

func (s *Service) ListMetrics(ctx context.Context, in *operator_v1alpha1.ListMetricsRequest) (*operator_v1alpha1.ListMetricsResponse, error) {

	var resp *operator_v1alpha1.ListMetricsResponse

	// Validate request
	err := ValidateListMetricsRequest(*in)
	if err != nil {
		resp = &operator_v1alpha1.ListMetricsResponse{}
		resp.Status = &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}
		return resp, nil
	}

	// build query instance to query metrics db
	q := buildMetircQuery(in)

	// query to metrics db
	quertResp, err := s.MetricsDB.Query(q)
	if err != nil {
		resp = &operator_v1alpha1.ListMetricsResponse{}
		resp.Status = &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}
		return resp, nil
	}

	// convert response of query metrics db to containers-ai.operator.v1alpha1.ListMetricssResposne
	resp = convertMetricsQueryResponseToProtoResponse(&quertResp)
	resp.Status = &status.Status{
		Code: int32(code.Code_OK),
	}
	return resp, nil
}

func (s *Service) ListMetricsSum(ctx context.Context, in *operator_v1alpha1.ListMetricsSumRequest) (*operator_v1alpha1.ListMetricsSumResponse, error) {

	return &operator_v1alpha1.ListMetricsSumResponse{
		Status: &status.Status{
			Code:    int32(code.Code_UNIMPLEMENTED),
			Message: "Not implemented",
		},
	}, nil
}

func (s *Service) CreatePredictResult(ctx context.Context, in *operator_v1alpha1.CreatePredictResultRequest) (*operator_v1alpha1.CreatePredictResultResponse, error) {
	// 1. Get namespace list information from predicted pods
	nsRange := map[string]bool{}
	for _, predictPod := range in.GetPredictPods() {
		if _, ok := nsRange[predictPod.GetNamespace()]; !ok {
			nsRange[predictPod.GetNamespace()] = true
		}
	}
	// 2. Get AlamedaScaler list from namespace list
	alaListRange := []autoscalingv1alpha1.AlamedaScaler{}
	for namespace, _ := range nsRange {
		alamedascalerList := &autoscalingv1alpha1.AlamedaScalerList{}
		err := s.Manager.GetClient().List(go_context.TODO(), client.InNamespace(namespace), alamedascalerList)
		if err == nil {
			alaListRange = append(alaListRange, alamedascalerList.Items...)
		}
	}
	if len(alaListRange) == 0 {
		return &operator_v1alpha1.CreatePredictResultResponse{
			Status: &status.Status{
				Code:    int32(code.Code_NOT_FOUND),
				Message: "AlamedaScaler not found.",
			},
		}, nil
	}
	for _, ala := range alaListRange {
		alaAnno := ala.GetAnnotations()
		if alaAnno == nil {
			scope.Warnf(fmt.Sprintf("No annotation found in AlamedaScaler %s in namespace %s in AlamedaScaler list, try searching next item", ala.GetName(), ala.GetNamespace()))
			continue
		}
		if _, ok := alaAnno[alamedascaler.AlamedaK8sController]; !ok {
			scope.Warnf(fmt.Sprintf("No k8s controller annotation key found in AlamedaResouce %s in namespace %s in AlamedaScaler list, try searching next item", ala.GetName(), ala.GetNamespace()))
			continue
		}
		scope.Infof(fmt.Sprintf("K8s controller annotation found %s in AlamedaResouce %s in namespace %s in AlamedaScaler list", alaAnno[alamedascaler.AlamedaK8sController], ala.GetName(), ala.GetNamespace()))
	}
	inBin, _ := json.Marshal(*in)
	return &operator_v1alpha1.CreatePredictResultResponse{
		Status: &status.Status{
			Code:    int32(code.Code_OK),
			Message: string(inBin),
		},
	}, nil
}

func (s *Service) GetResourceInfo(ctx context.Context, in *operator_v1alpha1.GetResourceInfoRequest) (*operator_v1alpha1.GetResourceInfoResponse, error) {
	return nil, nil
}

func (s *Service) GetResourceRecommendation(ctx context.Context, in *operator_v1alpha1.GetResourceRecommendationRequest) (*operator_v1alpha1.GetResourceRecommendationResponse, error) {
	return nil, nil
}

func buildMetircQuery(req *operator_v1alpha1.ListMetricsRequest) metrics.Query {

	var q = metrics.Query{}

	switch req.GetMetricType() {
	case operator_v1alpha1.MetricType_CONTAINER_CPU_USAGE_TOTAL:
		q.Metric = metrics.MetricTypeContainerCPUUsageTotal
	case operator_v1alpha1.MetricType_CONTAINER_CPU_USAGE_TOTAL_RATE:
		q.Metric = metrics.MetricTypeContainerCPUUsageTotalRate
	case operator_v1alpha1.MetricType_CONTAINER_MEMORY_USAGE:
		q.Metric = metrics.MetricTypeContainerMemoryUsage
	case operator_v1alpha1.MetricType_NODE_CPU_USAGE_SECONDS_AVG1M:
		q.Metric = metrics.MetricTypeNodeCPUUsageSecondsAvg1M
	case operator_v1alpha1.MetricType_NODE_MEMORY_USAGE_BYTES:
		q.Metric = metrics.MetricTypeNodeMemoryUsageBytes
	}

	for _, labelSelector := range req.GetConditions() {

		k := labelSelector.GetKey()
		v := labelSelector.GetValue()
		var op metrics.StringOperator
		switch labelSelector.GetOp() {
		case operator_v1alpha1.StrOp_EQUAL:
			op = metrics.StringOperatorEqueal
		case operator_v1alpha1.StrOp_NOT_EQUAL:
			op = metrics.StringOperatorNotEqueal
		}

		q.LabelSelectors = append(q.LabelSelectors, metrics.LabelSelector{Key: k, Op: op, Value: v})
	}

	// assign difference type of time to query instance by type of gRPC request time
	switch req.TimeSelector.(type) {
	case nil:
		q.TimeSelector = nil
	case *operator_v1alpha1.ListMetricsRequest_Time:
		q.TimeSelector = &metrics.Timestamp{T: time.Unix(req.GetTime().GetSeconds(), int64(req.GetTime().GetNanos()))}
	case *operator_v1alpha1.ListMetricsRequest_Duration:
		d, _ := ptypes.Duration(req.GetDuration())
		q.TimeSelector = &metrics.Since{
			Duration: d,
		}
	case *operator_v1alpha1.ListMetricsRequest_TimeRange:
		startTime := req.GetTimeRange().GetStartTime()
		endTime := req.GetTimeRange().GetEndTime()
		step, _ := ptypes.Duration(req.GetTimeRange().GetStep())
		q.TimeSelector = &metrics.TimeRange{
			StartTime: time.Unix(startTime.GetSeconds(), int64(startTime.GetNanos())),
			EndTime:   time.Unix(endTime.GetSeconds(), int64(endTime.GetNanos())),
			Step:      step,
		}
	}

	return q
}

func convertMetricsQueryResponseToProtoResponse(resp *metrics.QueryResponse) *operator_v1alpha1.ListMetricsResponse {

	// initiallize proto response
	ListMetricssResponse := &operator_v1alpha1.ListMetricsResponse{}
	ListMetricssResponse.Metrics = []*operator_v1alpha1.MetricResult{}

	for _, result := range resp.Results {
		series := &operator_v1alpha1.MetricResult{}

		series.Labels = result.Labels
		for _, sample := range result.Samples {
			s := &operator_v1alpha1.Sample{}

			timestampProto, err := ptypes.TimestampProto(sample.Time)
			if err != nil {
				scope.Error("convert time.Time to google.protobuf.Timestamp failed")
			}
			s.Time = timestampProto
			s.Value = sample.Value
			series.Samples = append(series.Samples, s)
		}
		ListMetricssResponse.Metrics = append(ListMetricssResponse.Metrics, series)
	}

	return ListMetricssResponse
}
