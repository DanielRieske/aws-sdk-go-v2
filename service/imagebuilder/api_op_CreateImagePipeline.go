// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new image pipeline. Image pipelines enable you to automate the
// creation and distribution of images.
func (c *Client) CreateImagePipeline(ctx context.Context, params *CreateImagePipelineInput, optFns ...func(*Options)) (*CreateImagePipelineOutput, error) {
	if params == nil {
		params = &CreateImagePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImagePipeline", params, optFns, c.addOperationCreateImagePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImagePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImagePipelineInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the infrastructure configuration that will be
	// used to build images created by this image pipeline.
	//
	// This member is required.
	InfrastructureConfigurationArn *string

	// The name of the image pipeline.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the container recipe that is used to
	// configure images created by this container pipeline.
	ContainerRecipeArn *string

	// The description of the image pipeline.
	Description *string

	// The Amazon Resource Name (ARN) of the distribution configuration that will be
	// used to configure and distribute images created by this image pipeline.
	DistributionConfigurationArn *string

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled *bool

	// The name or Amazon Resource Name (ARN) for the IAM role you create that grants
	// Image Builder access to perform workflow actions.
	ExecutionRole *string

	// The Amazon Resource Name (ARN) of the image recipe that will be used to
	// configure images created by this image pipeline.
	ImageRecipeArn *string

	// Contains settings for vulnerability scans.
	ImageScanningConfiguration *types.ImageScanningConfiguration

	// The image test configuration of the image pipeline.
	ImageTestsConfiguration *types.ImageTestsConfiguration

	// The schedule of the image pipeline.
	Schedule *types.Schedule

	// The status of the image pipeline.
	Status types.PipelineStatus

	// The tags of the image pipeline.
	Tags map[string]string

	// Contains an array of workflow configuration objects.
	Workflows []types.WorkflowConfiguration

	noSmithyDocumentSerde
}

type CreateImagePipelineOutput struct {

	// The client token that uniquely identifies the request.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image pipeline that was created by this
	// request.
	ImagePipelineArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateImagePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateImagePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateImagePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateImagePipeline"); err != nil {
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
	if err = addIdempotencyToken_opCreateImagePipelineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateImagePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImagePipeline(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateImagePipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateImagePipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateImagePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateImagePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateImagePipelineInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateImagePipelineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateImagePipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateImagePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateImagePipeline",
	}
}
