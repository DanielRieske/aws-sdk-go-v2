// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an account-level monthly spending limit override for sending text
// messages. Deleting a spend limit override will set the EnforcedLimit to equal
// the MaxLimit , which is controlled by Amazon Web Services. For more information
// on spend limits (quotas) see [Amazon Pinpoint quotas]in the Amazon Pinpoint Developer Guide.
//
// [Amazon Pinpoint quotas]: https://docs.aws.amazon.com/pinpoint/latest/developerguide/quotas.html
func (c *Client) DeleteTextMessageSpendLimitOverride(ctx context.Context, params *DeleteTextMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*DeleteTextMessageSpendLimitOverrideOutput, error) {
	if params == nil {
		params = &DeleteTextMessageSpendLimitOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTextMessageSpendLimitOverride", params, optFns, c.addOperationDeleteTextMessageSpendLimitOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTextMessageSpendLimitOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTextMessageSpendLimitOverrideInput struct {
	noSmithyDocumentSerde
}

type DeleteTextMessageSpendLimitOverrideOutput struct {

	// The current monthly limit, in US dollars.
	MonthlyLimit *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTextMessageSpendLimitOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteTextMessageSpendLimitOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteTextMessageSpendLimitOverride{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTextMessageSpendLimitOverride"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTextMessageSpendLimitOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTextMessageSpendLimitOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTextMessageSpendLimitOverride",
	}
}
