// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package vision

import (
	"time"

	"cloud.google.com/go/internal/version"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// ImageAnnotatorCallOptions contains the retry settings for each method of ImageAnnotatorClient.
type ImageAnnotatorCallOptions struct {
	BatchAnnotateImages     []gax.CallOption
	AsyncBatchAnnotateFiles []gax.CallOption
}

func defaultImageAnnotatorClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("vision.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultImageAnnotatorCallOptions() *ImageAnnotatorCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &ImageAnnotatorCallOptions{
		BatchAnnotateImages:     retry[[2]string{"default", "idempotent"}],
		AsyncBatchAnnotateFiles: retry[[2]string{"default", "idempotent"}],
	}
}

// ImageAnnotatorClient is a client for interacting with Cloud Vision API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ImageAnnotatorClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	imageAnnotatorClient visionpb.ImageAnnotatorClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *ImageAnnotatorCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewImageAnnotatorClient creates a new image annotator client.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
func NewImageAnnotatorClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultImageAnnotatorClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ImageAnnotatorClient{
		conn:        conn,
		CallOptions: defaultImageAnnotatorCallOptions(),

		imageAnnotatorClient: visionpb.NewImageAnnotatorClient(conn),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, option.WithGRPCConn(conn))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection
		// and never actually need to dial.
		// If this does happen, we could leak conn. However, we cannot close conn:
		// If the user invoked the function with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO(pongad): investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *ImageAnnotatorClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageAnnotatorClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageAnnotatorClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", version.Go()}, keyval...)
	kv = append(kv, "gapic", version.Repo, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// BatchAnnotateImages run image detection and annotation for a batch of images.
func (c *ImageAnnotatorClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.BatchAnnotateImages[0:len(c.CallOptions.BatchAnnotateImages):len(c.CallOptions.BatchAnnotateImages)], opts...)
	var resp *visionpb.BatchAnnotateImagesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AsyncBatchAnnotateFiles run asynchronous image detection and annotation for a list of generic
// files, such as PDF files, which may contain multiple pages and multiple
// images per page. Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateFilesResponse (results).
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.AsyncBatchAnnotateFiles[0:len(c.CallOptions.AsyncBatchAnnotateFiles):len(c.CallOptions.AsyncBatchAnnotateFiles)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.AsyncBatchAnnotateFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// AsyncBatchAnnotateFilesOperation manages a long-running operation from AsyncBatchAnnotateFiles.
type AsyncBatchAnnotateFilesOperation struct {
	lro *longrunning.Operation
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AsyncBatchAnnotateFilesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateFilesResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, 45000*time.Millisecond, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *AsyncBatchAnnotateFilesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateFilesResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *AsyncBatchAnnotateFilesOperation) Metadata() (*visionpb.OperationMetadata, error) {
	var meta visionpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AsyncBatchAnnotateFilesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AsyncBatchAnnotateFilesOperation) Name() string {
	return op.lro.Name()
}