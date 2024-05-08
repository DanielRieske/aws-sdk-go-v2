// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration of a specified Neptune Analytics graph
func (c *Client) UpdateGraph(ctx context.Context, params *UpdateGraphInput, optFns ...func(*Options)) (*UpdateGraphOutput, error) {
	if params == nil {
		params = &UpdateGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGraph", params, optFns, c.addOperationUpdateGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGraphInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// A value that indicates whether the graph has deletion protection enabled. The
	// graph can't be deleted when deletion protection is enabled.
	DeletionProtection *bool

	// The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the
	// graph.
	//
	// Min = 128
	ProvisionedMemory *int32

	// Specifies whether or not the graph can be reachable over the internet. All
	// access to graphs is IAM authenticated. ( true to enable, or false to disable.
	PublicConnectivity *bool

	noSmithyDocumentSerde
}

func (in *UpdateGraphInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type UpdateGraphOutput struct {

	// The ARN associated with the graph.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph.
	//
	// This member is required.
	Id *string

	// The name of the graph.
	//
	// This member is required.
	Name *string

	// The build number of the graph.
	BuildNumber *string

	// The time at which the graph was created.
	CreateTime *time.Time

	// If true , deletion protection is enabled for the graph.
	DeletionProtection *bool

	// The graph endpoint.
	Endpoint *string

	// The ID of the KMS key used to encrypt and decrypt graph data.
	KmsKeyIdentifier *string

	// The number of memory-optimized Neptune Capacity Units (m-NCUs) allocated to the
	// graph.
	ProvisionedMemory *int32

	// If true , the graph has a public endpoint, otherwise not.
	PublicConnectivity *bool

	// The number of replicas for the graph.
	ReplicaCount *int32

	// The ID of the snapshot from which the graph was created, if any.
	SourceSnapshotId *string

	// The status of the graph.
	Status types.GraphStatus

	// The reason that the graph has this status.
	StatusReason *string

	// Specifies the number of dimensions for vector embeddings loaded into the graph.
	// Max = 65535
	VectorSearchConfiguration *types.VectorSearchConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGraph{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGraph"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateGraphValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGraph(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGraph",
	}
}
