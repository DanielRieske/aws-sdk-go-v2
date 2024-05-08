// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Amazon CloudWatch log settings for a playback configuration.
func (c *Client) ConfigureLogsForPlaybackConfiguration(ctx context.Context, params *ConfigureLogsForPlaybackConfigurationInput, optFns ...func(*Options)) (*ConfigureLogsForPlaybackConfigurationOutput, error) {
	if params == nil {
		params = &ConfigureLogsForPlaybackConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfigureLogsForPlaybackConfiguration", params, optFns, c.addOperationConfigureLogsForPlaybackConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfigureLogsForPlaybackConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Configures Amazon CloudWatch log settings for a playback configuration.
type ConfigureLogsForPlaybackConfigurationInput struct {

	// The percentage of session logs that MediaTailor sends to your Cloudwatch Logs
	// account. For example, if your playback configuration has 1000 sessions and
	// percentEnabled is set to 60 , MediaTailor sends logs for 600 of the sessions to
	// CloudWatch Logs. MediaTailor decides at random which of the playback
	// configuration sessions to send logs for. If you want to view logs for a specific
	// session, you can use the [debug log mode].
	//
	// Valid values: 0 - 100
	//
	// [debug log mode]: https://docs.aws.amazon.com/mediatailor/latest/ug/debug-log-mode.html
	//
	// This member is required.
	PercentEnabled int32

	// The name of the playback configuration.
	//
	// This member is required.
	PlaybackConfigurationName *string

	noSmithyDocumentSerde
}

type ConfigureLogsForPlaybackConfigurationOutput struct {

	// The percentage of session logs that MediaTailor sends to your Cloudwatch Logs
	// account.
	//
	// This member is required.
	PercentEnabled int32

	// The name of the playback configuration.
	PlaybackConfigurationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfigureLogsForPlaybackConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConfigureLogsForPlaybackConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConfigureLogsForPlaybackConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ConfigureLogsForPlaybackConfiguration"); err != nil {
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
	if err = addOpConfigureLogsForPlaybackConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfigureLogsForPlaybackConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConfigureLogsForPlaybackConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConfigureLogsForPlaybackConfiguration",
	}
}
