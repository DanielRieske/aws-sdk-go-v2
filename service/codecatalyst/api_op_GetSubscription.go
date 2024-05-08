// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the Amazon Web Services account used for billing
// purposes and the billing plan for the space.
func (c *Client) GetSubscription(ctx context.Context, params *GetSubscriptionInput, optFns ...func(*Options)) (*GetSubscriptionOutput, error) {
	if params == nil {
		params = &GetSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSubscription", params, optFns, c.addOperationGetSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionInput struct {

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	noSmithyDocumentSerde
}

type GetSubscriptionOutput struct {

	// The display name of the Amazon Web Services account used for billing for the
	// space.
	AwsAccountName *string

	// The day and time the pending change will be applied to the space, in
	// coordinated universal time (UTC) timestamp format as specified in [RFC 3339].
	//
	// [RFC 3339]: https://www.rfc-editor.org/rfc/rfc3339#section-5.6
	PendingSubscriptionStartTime *time.Time

	// The type of the billing plan that the space will be changed to at the start of
	// the next billing cycle. This applies only to changes that reduce the
	// functionality available for the space. Billing plan changes that increase
	// functionality are applied immediately. For more information, see [Pricing].
	//
	// [Pricing]: https://codecatalyst.aws/explore/pricing
	PendingSubscriptionType *string

	// The type of the billing plan for the space.
	SubscriptionType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSubscription"); err != nil {
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
	if err = addOpGetSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSubscription",
	}
}
