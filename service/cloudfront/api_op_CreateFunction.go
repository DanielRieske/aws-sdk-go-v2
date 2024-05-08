// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a CloudFront function.
//
// To create a function, you provide the function code and some configuration
// information about the function. The response contains an Amazon Resource Name
// (ARN) that uniquely identifies the function.
//
// When you create a function, it's in the DEVELOPMENT stage. In this stage, you
// can test the function with TestFunction , and update it with UpdateFunction .
//
// When you're ready to use your function with a CloudFront distribution, use
// PublishFunction to copy the function from the DEVELOPMENT stage to LIVE . When
// it's live, you can attach the function to a distribution's cache behavior, using
// the function's ARN.
func (c *Client) CreateFunction(ctx context.Context, params *CreateFunctionInput, optFns ...func(*Options)) (*CreateFunctionOutput, error) {
	if params == nil {
		params = &CreateFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFunction", params, optFns, c.addOperationCreateFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFunctionInput struct {

	// The function code. For more information about writing a CloudFront function,
	// see [Writing function code for CloudFront Functions]in the Amazon CloudFront Developer Guide.
	//
	// [Writing function code for CloudFront Functions]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html
	//
	// This member is required.
	FunctionCode []byte

	// Configuration information about the function, including an optional comment and
	// the function's runtime.
	//
	// This member is required.
	FunctionConfig *types.FunctionConfig

	// A name to identify the function.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateFunctionOutput struct {

	// The version identifier for the current version of the CloudFront function.
	ETag *string

	// Contains configuration information and metadata about a CloudFront function.
	FunctionSummary *types.FunctionSummary

	// The URL of the CloudFront function. Use the URL to manage the function with the
	// CloudFront API.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFunction"); err != nil {
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
	if err = addOpCreateFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFunction",
	}
}
