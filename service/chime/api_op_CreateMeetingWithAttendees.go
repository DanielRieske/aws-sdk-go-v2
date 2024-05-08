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

//	Creates a new Amazon Chime SDK meeting in the specified media Region, with
//
// attendees. For more information about specifying media Regions, see [Amazon Chime SDK Media Regions]in the
// Amazon Chime SDK Developer Guide . For more information about the Amazon Chime
// SDK, see [Using the Amazon Chime SDK]in the Amazon Chime SDK Developer Guide .
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [CreateMeetingWithAttendees], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by CreateMeetingWithAttendees in the Amazon Chime SDK
// Meetings Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [Amazon Chime SDK Media Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/chime-sdk-meetings-regions.html
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
// [CreateMeetingWithAttendees]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_CreateMeetingWithAttendees.html
func (c *Client) CreateMeetingWithAttendees(ctx context.Context, params *CreateMeetingWithAttendeesInput, optFns ...func(*Options)) (*CreateMeetingWithAttendeesOutput, error) {
	if params == nil {
		params = &CreateMeetingWithAttendeesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMeetingWithAttendees", params, optFns, c.addOperationCreateMeetingWithAttendeesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeetingWithAttendeesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingWithAttendeesInput struct {

	// The unique identifier for the client request. Use a different token for
	// different meetings.
	//
	// This member is required.
	ClientRequestToken *string

	// The request containing the attendees to create.
	Attendees []types.CreateAttendeeRequestItem

	// The external meeting ID.
	ExternalMeetingId *string

	//  The Region in which to create the meeting. Default: us-east-1 .
	//
	// Available values: af-south-1 , ap-northeast-1 , ap-northeast-2 , ap-south-1 ,
	// ap-southeast-1 , ap-southeast-2 , ca-central-1 , eu-central-1 , eu-north-1 ,
	// eu-south-1 , eu-west-1 , eu-west-2 , eu-west-3 , sa-east-1 , us-east-1 ,
	// us-east-2 , us-west-1 , us-west-2 .
	MediaRegion *string

	// Reserved.
	MeetingHostId *string

	// The resource target configurations for receiving Amazon Chime SDK meeting and
	// attendee event notifications. The Amazon Chime SDK supports resource targets
	// located in the US East (N. Virginia) AWS Region (us-east-1).
	NotificationsConfiguration *types.MeetingNotificationConfiguration

	// The tag key-value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMeetingWithAttendeesOutput struct {

	// The attendee information, including attendees IDs and join tokens.
	Attendees []types.Attendee

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []types.CreateAttendeeError

	// A meeting created using the Amazon Chime SDK.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMeetingWithAttendeesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeetingWithAttendees{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeetingWithAttendees{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMeetingWithAttendees"); err != nil {
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
	if err = addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMeetingWithAttendeesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeetingWithAttendees(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMeetingWithAttendees struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMeetingWithAttendees) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMeetingWithAttendees) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMeetingWithAttendeesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMeetingWithAttendeesInput ")
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
func addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMeetingWithAttendees{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMeetingWithAttendees(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMeetingWithAttendees",
	}
}
