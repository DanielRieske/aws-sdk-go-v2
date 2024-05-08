// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests the details of a job for a third party action. Used for partner
// actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
func (c *Client) GetThirdPartyJobDetails(ctx context.Context, params *GetThirdPartyJobDetailsInput, optFns ...func(*Options)) (*GetThirdPartyJobDetailsOutput, error) {
	if params == nil {
		params = &GetThirdPartyJobDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetThirdPartyJobDetails", params, optFns, c.addOperationGetThirdPartyJobDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetThirdPartyJobDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetThirdPartyJobDetails action.
type GetThirdPartyJobDetailsInput struct {

	// The clientToken portion of the clientId and clientToken pair used to verify
	// that the calling entity is allowed access to the job and its details.
	//
	// This member is required.
	ClientToken *string

	// The unique system-generated ID used for identifying the job.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

// Represents the output of a GetThirdPartyJobDetails action.
type GetThirdPartyJobDetailsOutput struct {

	// The details of the job, including any protected values defined for the job.
	JobDetails *types.ThirdPartyJobDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetThirdPartyJobDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetThirdPartyJobDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetThirdPartyJobDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetThirdPartyJobDetails"); err != nil {
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
	if err = addOpGetThirdPartyJobDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetThirdPartyJobDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetThirdPartyJobDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetThirdPartyJobDetails",
	}
}
