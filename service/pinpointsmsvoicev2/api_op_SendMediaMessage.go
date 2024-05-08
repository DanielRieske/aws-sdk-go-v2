// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new multimedia message (MMS) and sends it to a recipient's phone
// number.
func (c *Client) SendMediaMessage(ctx context.Context, params *SendMediaMessageInput, optFns ...func(*Options)) (*SendMediaMessageOutput, error) {
	if params == nil {
		params = &SendMediaMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendMediaMessage", params, optFns, c.addOperationSendMediaMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendMediaMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMediaMessageInput struct {

	// The destination phone number in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The origination identity of the message. This can be either the PhoneNumber,
	// PhoneNumberId, PhoneNumberArn, SenderId, SenderIdArn, PoolId, or PoolArn.
	//
	// This member is required.
	OriginationIdentity *string

	// The name of the configuration set to use. This can be either the
	// ConfigurationSetName or ConfigurationSetArn.
	ConfigurationSetName *string

	// You can specify custom data in this field. If you do, that data is logged to
	// the event destination.
	Context map[string]string

	// When set to true, the message is checked and validated, but isn't sent to the
	// end recipient.
	DryRun bool

	// The maximum amount that you want to spend, in US dollars, per each MMS message.
	MaxPrice *string

	// An array of URLs to each media file to send.
	//
	// The media files have to be stored in a publicly available S3 bucket. Supported
	// media file formats are listed in [MMS file types, size and character limits]. For more information on creating an S3
	// bucket and managing objects, see [Creating a bucket]and [Uploading objects] in the S3 user guide.
	//
	// [Creating a bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html
	// [MMS file types, size and character limits]: https://docs.aws.amazon.com/sms-voice/latest/userguide/mms-limitations-character.html
	// [Uploading objects]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html
	MediaUrls []string

	// The text body of the message.
	MessageBody *string

	// The unique identifier of the protect configuration to use.
	ProtectConfigurationId *string

	// How long the text message is valid for. By default this is 72 hours.
	TimeToLive *int32

	noSmithyDocumentSerde
}

type SendMediaMessageOutput struct {

	// The unique identifier for the message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendMediaMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendMediaMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendMediaMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendMediaMessage"); err != nil {
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
	if err = addOpSendMediaMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendMediaMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendMediaMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendMediaMessage",
	}
}
