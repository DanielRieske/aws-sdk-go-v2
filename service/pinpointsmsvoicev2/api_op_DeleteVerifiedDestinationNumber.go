// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Delete a verified destination phone number.
func (c *Client) DeleteVerifiedDestinationNumber(ctx context.Context, params *DeleteVerifiedDestinationNumberInput, optFns ...func(*Options)) (*DeleteVerifiedDestinationNumberOutput, error) {
	if params == nil {
		params = &DeleteVerifiedDestinationNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVerifiedDestinationNumber", params, optFns, c.addOperationDeleteVerifiedDestinationNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVerifiedDestinationNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVerifiedDestinationNumberInput struct {

	// The unique identifier for the verified destination phone number.
	//
	// This member is required.
	VerifiedDestinationNumberId *string

	noSmithyDocumentSerde
}

type DeleteVerifiedDestinationNumberOutput struct {

	// The time when the destination phone number was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The verified destination phone number, in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The Amazon Resource Name (ARN) for the verified destination phone number.
	//
	// This member is required.
	VerifiedDestinationNumberArn *string

	// The unique identifier for the verified destination phone number.
	//
	// This member is required.
	VerifiedDestinationNumberId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVerifiedDestinationNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteVerifiedDestinationNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteVerifiedDestinationNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteVerifiedDestinationNumber"); err != nil {
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
	if err = addOpDeleteVerifiedDestinationNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVerifiedDestinationNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVerifiedDestinationNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteVerifiedDestinationNumber",
	}
}
