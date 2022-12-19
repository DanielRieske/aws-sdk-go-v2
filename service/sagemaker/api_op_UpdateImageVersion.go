// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the properties of a SageMaker image version.
func (c *Client) UpdateImageVersion(ctx context.Context, params *UpdateImageVersionInput, optFns ...func(*Options)) (*UpdateImageVersionOutput, error) {
	if params == nil {
		params = &UpdateImageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateImageVersion", params, optFns, c.addOperationUpdateImageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateImageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateImageVersionInput struct {

	// The name of the image.
	//
	// This member is required.
	ImageName *string

	// The alias of the image version.
	Alias *string

	// A list of aliases to add.
	AliasesToAdd []string

	// A list of aliases to delete.
	AliasesToDelete []string

	// Indicates Horovod compatibility.
	Horovod bool

	// Indicates SageMaker job type compatibility.
	//
	// * TRAINING: The image version is
	// compatible with SageMaker training jobs.
	//
	// * INFERENCE: The image version is
	// compatible with SageMaker inference jobs.
	//
	// * NOTEBOOK_KERNEL: The image version
	// is compatible with SageMaker notebook kernels.
	JobType types.JobType

	// The machine learning framework vended in the image version.
	MLFramework *string

	// Indicates CPU or GPU compatibility.
	//
	// * CPU: The image version is compatible with
	// CPU.
	//
	// * GPU: The image version is compatible with GPU.
	Processor types.Processor

	// The supported programming language and its version.
	ProgrammingLang *string

	// The maintainer description of the image version.
	ReleaseNotes *string

	// The availability of the image version specified by the maintainer.
	//
	// *
	// NOT_PROVIDED: The maintainers did not provide a status for image version
	// stability.
	//
	// * STABLE: The image version is stable.
	//
	// * TO_BE_ARCHIVED: The image
	// version is set to be archived. Custom image versions that are set to be archived
	// are automatically archived after three months.
	//
	// * ARCHIVED: The image version is
	// archived. Archived image versions are not searchable and are no longer actively
	// supported.
	VendorGuidance types.VendorGuidance

	// The version of the image.
	Version *int32

	noSmithyDocumentSerde
}

type UpdateImageVersionOutput struct {

	// The ARN of the image version.
	ImageVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateImageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateImageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateImageVersion(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateImageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateImageVersion",
	}
}
