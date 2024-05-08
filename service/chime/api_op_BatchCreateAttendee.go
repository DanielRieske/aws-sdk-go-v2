// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates up to 100 new attendees for an active Amazon Chime SDK meeting.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [BatchCreateAttendee], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// For more information about the Amazon Chime SDK, see [Using the Amazon Chime SDK] in the Amazon Chime SDK
// Developer Guide.
//
// Deprecated: Replaced by BatchCreateAttendee in the Amazon Chime SDK Meetings
// Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [BatchCreateAttendee]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_BatchCreateAttendee.html
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func (c *Client) BatchCreateAttendee(ctx context.Context, params *BatchCreateAttendeeInput, optFns ...func(*Options)) (*BatchCreateAttendeeOutput, error) {
	if params == nil {
		params = &BatchCreateAttendeeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateAttendee", params, optFns, c.addOperationBatchCreateAttendeeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateAttendeeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateAttendeeInput struct {

	// The request containing the attendees to create.
	//
	// This member is required.
	Attendees []types.CreateAttendeeRequestItem

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string

	noSmithyDocumentSerde
}

type BatchCreateAttendeeOutput struct {

	// The attendee information, including attendees IDs and join tokens.
	Attendees []types.Attendee

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []types.CreateAttendeeError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateAttendeeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchCreateAttendee"); err != nil {
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
	if err = addOpBatchCreateAttendeeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateAttendee(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateAttendee(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchCreateAttendee",
	}
}
