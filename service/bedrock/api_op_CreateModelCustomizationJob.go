// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a fine-tuning job to customize a base model.
//
// You specify the base foundation model and the location of the training data.
// After the model-customization job completes successfully, your custom model
// resource will be ready to use. Amazon Bedrock returns validation loss metrics
// and output generations after the job completes.
//
// For information on the format of training and validation data, see [Prepare the datasets].
//
// Model-customization jobs are asynchronous and the completion time depends on
// the base model and the training/validation data size. To monitor a job, use the
// GetModelCustomizationJob operation to retrieve the job status.
//
// For more information, see [Custom models] in the Amazon Bedrock User Guide.
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Prepare the datasets]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-prepare.html
func (c *Client) CreateModelCustomizationJob(ctx context.Context, params *CreateModelCustomizationJobInput, optFns ...func(*Options)) (*CreateModelCustomizationJobOutput, error) {
	if params == nil {
		params = &CreateModelCustomizationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelCustomizationJob", params, optFns, c.addOperationCreateModelCustomizationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelCustomizationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelCustomizationJobInput struct {

	// Name of the base model.
	//
	// This member is required.
	BaseModelIdentifier *string

	// A name for the resulting custom model.
	//
	// This member is required.
	CustomModelName *string

	// Parameters related to tuning the model. For details on the format for different
	// models, see [Custom model hyperparameters].
	//
	// [Custom model hyperparameters]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models-hp.html
	//
	// This member is required.
	HyperParameters map[string]string

	// A name for the fine-tuning job.
	//
	// This member is required.
	JobName *string

	// S3 location for the output data.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The Amazon Resource Name (ARN) of an IAM service role that Amazon Bedrock can
	// assume to perform tasks on your behalf. For example, during model training,
	// Amazon Bedrock needs your permission to read input data from an S3 bucket, write
	// model artifacts to an S3 bucket. To pass this role to Amazon Bedrock, the caller
	// of this API must have the iam:PassRole permission.
	//
	// This member is required.
	RoleArn *string

	// Information about the training dataset.
	//
	// This member is required.
	TrainingDataConfig *types.TrainingDataConfig

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientRequestToken *string

	// The custom model is encrypted at rest using this key.
	CustomModelKmsKeyId *string

	// Tags to attach to the resulting custom model.
	CustomModelTags []types.Tag

	// The customization type.
	CustomizationType types.CustomizationType

	// Tags to attach to the job.
	JobTags []types.Tag

	// Information about the validation dataset.
	ValidationDataConfig *types.ValidationDataConfig

	// VPC configuration (optional). Configuration parameters for the private Virtual
	// Private Cloud (VPC) that contains the resources you are using for this job.
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateModelCustomizationJobOutput struct {

	// Amazon Resource Name (ARN) of the fine tuning job
	//
	// This member is required.
	JobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelCustomizationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateModelCustomizationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateModelCustomizationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModelCustomizationJob"); err != nil {
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
	if err = addIdempotencyToken_opCreateModelCustomizationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateModelCustomizationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelCustomizationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateModelCustomizationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateModelCustomizationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateModelCustomizationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateModelCustomizationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateModelCustomizationJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateModelCustomizationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateModelCustomizationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateModelCustomizationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModelCustomizationJob",
	}
}
