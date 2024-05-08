// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or modifies account-level configurations in ACM.
//
// The supported configuration option is DaysBeforeExpiry . This option specifies
// the number of days prior to certificate expiration when ACM starts generating
// EventBridge events. ACM sends one event per day per certificate until the
// certificate expires. By default, accounts receive events starting 45 days before
// certificate expiration.
func (c *Client) PutAccountConfiguration(ctx context.Context, params *PutAccountConfigurationInput, optFns ...func(*Options)) (*PutAccountConfigurationOutput, error) {
	if params == nil {
		params = &PutAccountConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountConfiguration", params, optFns, c.addOperationPutAccountConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountConfigurationInput struct {

	// Customer-chosen string used to distinguish between calls to
	// PutAccountConfiguration . Idempotency tokens time out after one hour. If you
	// call PutAccountConfiguration multiple times with the same unexpired idempotency
	// token, ACM treats it as the same request and returns the original result. If you
	// change the idempotency token for each call, ACM treats each call as a new
	// request.
	//
	// This member is required.
	IdempotencyToken *string

	// Specifies expiration events associated with an account.
	ExpiryEvents *types.ExpiryEventsConfiguration

	noSmithyDocumentSerde
}

type PutAccountConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountConfiguration"); err != nil {
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
	if err = addOpPutAccountConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountConfiguration",
	}
}
