// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/metric.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetricType int32

const (
	MetricType_UNDEFINED                              MetricType = 0
	MetricType_CONTAINER_CPU_USAGE_SECONDS_PERCENTAGE MetricType = 1
	MetricType_CONTAINER_MEMORY_USAGE_BYTES           MetricType = 2
	MetricType_NODE_CPU_USAGE_SECONDS_PERCENTAGE      MetricType = 3
	MetricType_NODE_MEMORY_USAGE_BYTES                MetricType = 4
)

var MetricType_name = map[int32]string{
	0: "UNDEFINED",
	1: "CONTAINER_CPU_USAGE_SECONDS_PERCENTAGE",
	2: "CONTAINER_MEMORY_USAGE_BYTES",
	3: "NODE_CPU_USAGE_SECONDS_PERCENTAGE",
	4: "NODE_MEMORY_USAGE_BYTES",
}
var MetricType_value = map[string]int32{
	"UNDEFINED":                              0,
	"CONTAINER_CPU_USAGE_SECONDS_PERCENTAGE": 1,
	"CONTAINER_MEMORY_USAGE_BYTES":           2,
	"NODE_CPU_USAGE_SECONDS_PERCENTAGE":      3,
	"NODE_MEMORY_USAGE_BYTES":                4,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{0}
}

type StrOp int32

const (
	StrOp_EQUAL     StrOp = 0
	StrOp_NOT_EQUAL StrOp = 1
)

var StrOp_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
}
var StrOp_value = map[string]int32{
	"EQUAL":     0,
	"NOT_EQUAL": 1,
}

func (x StrOp) String() string {
	return proto.EnumName(StrOp_name, int32(x))
}
func (StrOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{1}
}

type ContainerMetric struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           []*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{0}
}
func (m *ContainerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerMetric.Unmarshal(m, b)
}
func (m *ContainerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerMetric.Marshal(b, m, deterministic)
}
func (dst *ContainerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerMetric.Merge(dst, src)
}
func (m *ContainerMetric) XXX_Size() int {
	return xxx_messageInfo_ContainerMetric.Size(m)
}
func (m *ContainerMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerMetric proto.InternalMessageInfo

func (m *ContainerMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerMetric) GetMetricData() []*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

type PodMetric struct {
	NamespacedName       *NamespacedName    `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ContainerMetrics     []*ContainerMetric `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PodMetric) Reset()         { *m = PodMetric{} }
func (m *PodMetric) String() string { return proto.CompactTextString(m) }
func (*PodMetric) ProtoMessage()    {}
func (*PodMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{1}
}
func (m *PodMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodMetric.Unmarshal(m, b)
}
func (m *PodMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodMetric.Marshal(b, m, deterministic)
}
func (dst *PodMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodMetric.Merge(dst, src)
}
func (m *PodMetric) XXX_Size() int {
	return xxx_messageInfo_PodMetric.Size(m)
}
func (m *PodMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_PodMetric.DiscardUnknown(m)
}

var xxx_messageInfo_PodMetric proto.InternalMessageInfo

func (m *PodMetric) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodMetric) GetContainerMetrics() []*ContainerMetric {
	if m != nil {
		return m.ContainerMetrics
	}
	return nil
}

type NodeMetric struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           []*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeMetric) Reset()         { *m = NodeMetric{} }
func (m *NodeMetric) String() string { return proto.CompactTextString(m) }
func (*NodeMetric) ProtoMessage()    {}
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{2}
}
func (m *NodeMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetric.Unmarshal(m, b)
}
func (m *NodeMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetric.Marshal(b, m, deterministic)
}
func (dst *NodeMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetric.Merge(dst, src)
}
func (m *NodeMetric) XXX_Size() int {
	return xxx_messageInfo_NodeMetric.Size(m)
}
func (m *NodeMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetric.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetric proto.InternalMessageInfo

func (m *NodeMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetric) GetMetricData() []*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

type Sample struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	NumValue             string               `protobuf:"bytes,2,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{3}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (dst *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(dst, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Sample) GetNumValue() string {
	if m != nil {
		return m.NumValue
	}
	return ""
}

type MetricResult struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Samples              []*Sample         `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricResult) Reset()         { *m = MetricResult{} }
func (m *MetricResult) String() string { return proto.CompactTextString(m) }
func (*MetricResult) ProtoMessage()    {}
func (*MetricResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{4}
}
func (m *MetricResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricResult.Unmarshal(m, b)
}
func (m *MetricResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricResult.Marshal(b, m, deterministic)
}
func (dst *MetricResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricResult.Merge(dst, src)
}
func (m *MetricResult) XXX_Size() int {
	return xxx_messageInfo_MetricResult.Size(m)
}
func (m *MetricResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricResult.DiscardUnknown(m)
}

var xxx_messageInfo_MetricResult proto.InternalMessageInfo

func (m *MetricResult) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricResult) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type TimeRange struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration   `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{5}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (dst *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(dst, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

type LabelSelector struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op                   StrOp    `protobuf:"varint,2,opt,name=op,proto3,enum=containers_ai.alameda.v1alpha1.datahub.StrOp" json:"op,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{6}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelector.Unmarshal(m, b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
}
func (dst *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(dst, src)
}
func (m *LabelSelector) XXX_Size() int {
	return xxx_messageInfo_LabelSelector.Size(m)
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelector) GetOp() StrOp {
	if m != nil {
		return m.Op
	}
	return StrOp_EQUAL
}

func (m *LabelSelector) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MetricData struct {
	MetricType MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containers_ai.alameda.v1alpha1.datahub.MetricType" json:"metric_type,omitempty"`
	// data can be time series or non-time series
	Data                 []*Sample `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_a4fd053830513e7e, []int{7}
}
func (m *MetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricData.Unmarshal(m, b)
}
func (m *MetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricData.Marshal(b, m, deterministic)
}
func (dst *MetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricData.Merge(dst, src)
}
func (m *MetricData) XXX_Size() int {
	return xxx_messageInfo_MetricData.Size(m)
}
func (m *MetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricData proto.InternalMessageInfo

func (m *MetricData) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_UNDEFINED
}

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.ContainerMetric")
	proto.RegisterType((*PodMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.PodMetric")
	proto.RegisterType((*NodeMetric)(nil), "containers_ai.alameda.v1alpha1.datahub.NodeMetric")
	proto.RegisterType((*Sample)(nil), "containers_ai.alameda.v1alpha1.datahub.Sample")
	proto.RegisterType((*MetricResult)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricResult")
	proto.RegisterMapType((map[string]string)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry")
	proto.RegisterType((*TimeRange)(nil), "containers_ai.alameda.v1alpha1.datahub.TimeRange")
	proto.RegisterType((*LabelSelector)(nil), "containers_ai.alameda.v1alpha1.datahub.LabelSelector")
	proto.RegisterType((*MetricData)(nil), "containers_ai.alameda.v1alpha1.datahub.MetricData")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.MetricType", MetricType_name, MetricType_value)
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.StrOp", StrOp_name, StrOp_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/metric.proto", fileDescriptor_metric_a4fd053830513e7e)
}

var fileDescriptor_metric_a4fd053830513e7e = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xf4, 0xe2, 0x09, 0x6d, 0xc3, 0x0a, 0x89, 0x90, 0x22, 0x68, 0x2d, 0x51, 0x95,
	0xa2, 0x3a, 0x6a, 0x10, 0x97, 0x22, 0x21, 0x91, 0x26, 0x4b, 0xa9, 0xd4, 0x3a, 0x65, 0x9d, 0x20,
	0xfa, 0xb4, 0xda, 0xc4, 0x4b, 0x1b, 0xe1, 0x9b, 0xec, 0x75, 0xa5, 0xf0, 0x3f, 0xbc, 0xc1, 0xcf,
	0xf0, 0x1f, 0xfc, 0x03, 0xf2, 0xda, 0x4e, 0xd3, 0x52, 0x51, 0xe7, 0x85, 0xb7, 0x99, 0x64, 0xce,
	0x99, 0x33, 0xc7, 0x67, 0xe1, 0x29, 0x73, 0x98, 0xcb, 0x6d, 0x46, 0x59, 0x30, 0x6a, 0x5c, 0xec,
	0x32, 0x27, 0x38, 0x67, 0xbb, 0x0d, 0x9b, 0x09, 0x76, 0x1e, 0x0f, 0x1a, 0x2e, 0x17, 0xe1, 0x68,
	0x68, 0x04, 0xa1, 0x2f, 0x7c, 0xb4, 0x39, 0xf4, 0x3d, 0xc1, 0x46, 0x1e, 0x0f, 0x23, 0xca, 0x46,
	0x46, 0x06, 0x34, 0x72, 0x90, 0x91, 0x81, 0xea, 0x8f, 0xcf, 0x7c, 0xff, 0xcc, 0xe1, 0x0d, 0x89,
	0x1a, 0xc4, 0x5f, 0x1a, 0x62, 0xe4, 0xf2, 0x48, 0x30, 0x37, 0x48, 0x89, 0xea, 0x8f, 0xae, 0x0f,
	0xd8, 0x71, 0xc8, 0xc4, 0xc8, 0xf7, 0xb2, 0xff, 0x9f, 0xdd, 0xa6, 0x89, 0x25, 0x75, 0x3a, 0xac,
	0x7f, 0x83, 0xd5, 0x76, 0xae, 0xeb, 0x58, 0xca, 0x45, 0x08, 0xca, 0x1e, 0x73, 0x79, 0x4d, 0x59,
	0x57, 0xb6, 0x34, 0x22, 0x6b, 0x64, 0x41, 0x25, 0x3d, 0x86, 0x26, 0xd8, 0x9a, 0xba, 0x5e, 0xda,
	0xaa, 0x34, 0x9b, 0x46, 0xb1, 0x93, 0x8c, 0x94, 0xb8, 0xc3, 0x04, 0x23, 0xe0, 0x4e, 0x6a, 0xfd,
	0x97, 0x02, 0xda, 0x89, 0x6f, 0x67, 0x6b, 0x29, 0xac, 0x26, 0xab, 0xa2, 0x80, 0x0d, 0xb9, 0x4d,
	0x27, 0x0a, 0x2a, 0xcd, 0x97, 0x45, 0xd7, 0x98, 0x13, 0x78, 0x52, 0x91, 0x15, 0xef, 0x4a, 0x8f,
	0x6c, 0xb8, 0x3b, 0x21, 0xa2, 0xa9, 0x8c, 0x28, 0xbb, 0xe4, 0x55, 0xd1, 0x15, 0xd7, 0xbc, 0x22,
	0xd5, 0xe1, 0xd5, 0x1f, 0x22, 0x3d, 0x06, 0x30, 0x7d, 0x9b, 0xff, 0x6f, 0x2f, 0xfb, 0xb0, 0x60,
	0x31, 0x37, 0x70, 0x38, 0x32, 0xa0, 0x9c, 0x24, 0x26, 0x33, 0xaf, 0x6e, 0xa4, 0x69, 0x31, 0xf2,
	0xb4, 0x18, 0xbd, 0x3c, 0x4e, 0x44, 0xce, 0xa1, 0x35, 0xd0, 0xbc, 0xd8, 0xa5, 0x17, 0xcc, 0x89,
	0x79, 0x4d, 0x95, 0x3a, 0x97, 0xbc, 0xd8, 0xfd, 0x94, 0xf4, 0xfa, 0x6f, 0x05, 0xee, 0x64, 0xa7,
	0xf2, 0x28, 0x76, 0x04, 0xfa, 0x0c, 0x0b, 0x0e, 0x1b, 0x70, 0x27, 0xaa, 0x29, 0x52, 0xf7, 0xbb,
	0xd9, 0x74, 0xa7, 0x2c, 0xc6, 0x91, 0xa4, 0xc0, 0x9e, 0x08, 0xc7, 0x24, 0xe3, 0x43, 0x1f, 0x60,
	0x31, 0x92, 0x17, 0xe4, 0x1f, 0xc5, 0x28, 0x4a, 0x9d, 0x1e, 0x4e, 0x72, 0x78, 0x7d, 0x0f, 0x2a,
	0x53, 0x0b, 0x50, 0x15, 0x4a, 0x5f, 0xf9, 0x38, 0xfb, 0x04, 0x49, 0x89, 0xee, 0xc1, 0xfc, 0xf4,
	0xb9, 0x69, 0xf3, 0x46, 0x7d, 0xad, 0xe8, 0x3f, 0x14, 0xd0, 0x12, 0x83, 0x08, 0xf3, 0xce, 0x38,
	0xda, 0x03, 0x88, 0x04, 0x0b, 0x05, 0x2d, 0x68, 0xa8, 0x26, 0xa7, 0x93, 0x1e, 0xbd, 0x80, 0x25,
	0xee, 0xd9, 0x29, 0x50, 0xbd, 0x15, 0xb8, 0xc8, 0x3d, 0x5b, 0xc2, 0x76, 0xa0, 0x1c, 0x09, 0x1e,
	0xd4, 0x4a, 0x12, 0xf2, 0xe0, 0x2f, 0x48, 0x27, 0x7b, 0xea, 0x44, 0x8e, 0xe9, 0x17, 0xb0, 0x2c,
	0x2f, 0xb5, 0xb8, 0xc3, 0x87, 0xc2, 0x0f, 0x6f, 0xb8, 0xf5, 0x2d, 0xa8, 0x7e, 0x20, 0x25, 0xac,
	0x34, 0x77, 0x0a, 0x3b, 0x2a, 0xc2, 0x6e, 0x40, 0x54, 0x3f, 0xb8, 0xb4, 0xaa, 0x34, 0x65, 0x95,
	0xfe, 0x5d, 0x01, 0xb8, 0x0c, 0xe2, 0x54, 0xa2, 0xc5, 0x38, 0x48, 0x8d, 0x5a, 0x99, 0x35, 0xd1,
	0xbd, 0x71, 0xc0, 0xf3, 0x44, 0x27, 0x35, 0xda, 0x87, 0xf2, 0xd4, 0xfb, 0x98, 0x35, 0x0c, 0x12,
	0xbb, 0xfd, 0x73, 0xa2, 0x53, 0x52, 0x2e, 0x83, 0xd6, 0x37, 0x3b, 0xf8, 0xfd, 0xa1, 0x89, 0x3b,
	0xd5, 0x39, 0xb4, 0x0d, 0x9b, 0xed, 0xae, 0xd9, 0x6b, 0x1d, 0x9a, 0x98, 0xd0, 0xf6, 0x49, 0x9f,
	0xf6, 0xad, 0xd6, 0x01, 0xa6, 0x16, 0x6e, 0x77, 0xcd, 0x8e, 0x45, 0x4f, 0x30, 0x69, 0x63, 0xb3,
	0xd7, 0x3a, 0xc0, 0x55, 0x05, 0xad, 0xc3, 0xc3, 0xcb, 0xd9, 0x63, 0x7c, 0xdc, 0x25, 0xa7, 0xd9,
	0xf8, 0xfe, 0x69, 0x0f, 0x5b, 0x55, 0x15, 0x3d, 0x81, 0x0d, 0xb3, 0xdb, 0xc1, 0xff, 0x26, 0x2a,
	0xa1, 0x35, 0xb8, 0x2f, 0xc7, 0x6e, 0xe0, 0x28, 0x6f, 0x6f, 0xc0, 0xbc, 0xb4, 0x1e, 0x69, 0x30,
	0x8f, 0x3f, 0xf6, 0x5b, 0x47, 0xd5, 0xb9, 0x44, 0xb4, 0xd9, 0xed, 0xd1, 0xb4, 0x55, 0x06, 0x0b,
	0x32, 0x0b, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x1e, 0x61, 0xa5, 0x7a, 0x06, 0x00,
	0x00,
}