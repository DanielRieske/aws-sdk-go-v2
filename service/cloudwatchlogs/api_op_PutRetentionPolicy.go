// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the retention of the specified log group. With a retention policy, you can
// configure the number of days for which to retain log events in the specified log
// group.
//
// CloudWatch Logs doesn’t immediately delete log events when they reach their
// retention setting. It typically takes up to 72 hours after that before log
// events are deleted, but in rare situations might take longer.
//
// To illustrate, imagine that you change a log group to have a longer retention
// setting when it contains log events that are past the expiration date, but
// haven’t been deleted. Those log events will take up to 72 hours to be deleted
// after the new retention date is reached. To make sure that log data is deleted
// permanently, keep a log group at its lower retention setting until 72 hours
// after the previous retention period ends. Alternatively, wait to change the
// retention setting until you confirm that the earlier log events are deleted.
//
// When log events reach their retention setting they are marked for deletion.
// After they are marked for deletion, they do not add to your archival storage
// costs anymore, even if they are not actually deleted until later. These log
// events marked for deletion are also not included when you use an API to retrieve
// the storedBytes value to see how many bytes a log group is storing.
func (c *Client) PutRetentionPolicy(ctx context.Context, params *PutRetentionPolicyInput, optFns ...func(*Options)) (*PutRetentionPolicyOutput, error) {
	if params == nil {
		params = &PutRetentionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRetentionPolicy", params, optFns, c.addOperationPutRetentionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRetentionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRetentionPolicyInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The number of days to retain the log events in the specified log group.
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545,
	// 731, 1096, 1827, 2192, 2557, 2922, 3288, and 3653.
	//
	// To set a log group so that its log events do not expire, use [DeleteRetentionPolicy].
	//
	// [DeleteRetentionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html
	//
	// This member is required.
	RetentionInDays *int32

	noSmithyDocumentSerde
}

type PutRetentionPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRetentionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRetentionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRetentionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRetentionPolicy"); err != nil {
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
	if err = addOpPutRetentionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRetentionPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRetentionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRetentionPolicy",
	}
}
