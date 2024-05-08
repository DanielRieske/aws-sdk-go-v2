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

// Starts transcription for the specified meetingId . For more information, refer
// to [Using Amazon Chime SDK live transcription]in the Amazon Chime SDK Developer Guide.
//
// If you specify an invalid configuration, a TranscriptFailed event will be sent
// with the contents of the BadRequestException generated by Amazon Transcribe.
// For more information on each parameter and which combinations are valid, refer
// to the [StartStreamTranscription]API in the Amazon Transcribe Developer Guide.
//
// Amazon Chime SDK live transcription is powered by Amazon Transcribe. Use of
// Amazon Transcribe is subject to the [AWS Service Terms], including the terms specific to the AWS
// Machine Learning and Artificial Intelligence Services.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [StartMeetingTranscription], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by StartMeetingTranscription in the Amazon Chime SDK
// Meetings Namespace
//
// [AWS Service Terms]: https://aws.amazon.com/service-terms/
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [Using Amazon Chime SDK live transcription]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meeting-transcription.html
// [StartMeetingTranscription]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_StartMeetingTranscription.html
// [StartStreamTranscription]: https://docs.aws.amazon.com/transcribe/latest/APIReference/API_streaming_StartStreamTranscription.html
func (c *Client) StartMeetingTranscription(ctx context.Context, params *StartMeetingTranscriptionInput, optFns ...func(*Options)) (*StartMeetingTranscriptionOutput, error) {
	if params == nil {
		params = &StartMeetingTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMeetingTranscription", params, optFns, c.addOperationStartMeetingTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMeetingTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMeetingTranscriptionInput struct {

	// The unique ID of the meeting being transcribed.
	//
	// This member is required.
	MeetingId *string

	// The configuration for the current transcription operation. Must contain
	// EngineTranscribeSettings or EngineTranscribeMedicalSettings .
	//
	// This member is required.
	TranscriptionConfiguration *types.TranscriptionConfiguration

	noSmithyDocumentSerde
}

type StartMeetingTranscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMeetingTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMeetingTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMeetingTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMeetingTranscription"); err != nil {
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
	if err = addOpStartMeetingTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMeetingTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMeetingTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMeetingTranscription",
	}
}
