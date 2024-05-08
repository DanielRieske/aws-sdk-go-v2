// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the default patch baseline. Amazon Web Services Systems Manager
// supports creating multiple default patch baselines. For example, you can create
// a default patch baseline for each operating system.
//
// If you don't specify an operating system value, the default patch baseline for
// Windows is returned.
func (c *Client) GetDefaultPatchBaseline(ctx context.Context, params *GetDefaultPatchBaselineInput, optFns ...func(*Options)) (*GetDefaultPatchBaselineOutput, error) {
	if params == nil {
		params = &GetDefaultPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDefaultPatchBaseline", params, optFns, c.addOperationGetDefaultPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDefaultPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDefaultPatchBaselineInput struct {

	// Returns the default patch baseline for the specified operating system.
	OperatingSystem types.OperatingSystem

	noSmithyDocumentSerde
}

type GetDefaultPatchBaselineOutput struct {

	// The ID of the default patch baseline.
	BaselineId *string

	// The operating system for the returned patch baseline.
	OperatingSystem types.OperatingSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDefaultPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDefaultPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDefaultPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDefaultPatchBaseline"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDefaultPatchBaseline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDefaultPatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDefaultPatchBaseline",
	}
}
