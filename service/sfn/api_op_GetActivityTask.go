// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by workers to retrieve a task (with the specified activity ARN) which has
// been scheduled for execution by a running state machine. This initiates a long
// poll, where the service holds the HTTP connection open and responds as soon as a
// task becomes available (i.e. an execution of a task of this type is needed.) The
// maximum time the service holds on to the request before responding is 60
// seconds. If no task is available within 60 seconds, the poll returns a taskToken
// with a null string.
//
// This API action isn't logged in CloudTrail.
//
// Workers should set their client side socket timeout to at least 65 seconds (5
// seconds higher than the maximum time the service may hold the poll request).
//
// Polling with GetActivityTask can cause latency in some implementations. See [Avoid Latency When Polling for Activity Tasks] in
// the Step Functions Developer Guide.
//
// [Avoid Latency When Polling for Activity Tasks]: https://docs.aws.amazon.com/step-functions/latest/dg/bp-activity-pollers.html
func (c *Client) GetActivityTask(ctx context.Context, params *GetActivityTaskInput, optFns ...func(*Options)) (*GetActivityTaskOutput, error) {
	if params == nil {
		params = &GetActivityTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetActivityTask", params, optFns, c.addOperationGetActivityTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetActivityTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetActivityTaskInput struct {

	// The Amazon Resource Name (ARN) of the activity to retrieve tasks from (assigned
	// when you create the task using CreateActivity.)
	//
	// This member is required.
	ActivityArn *string

	// You can provide an arbitrary name in order to identify the worker that the task
	// is assigned to. This name is used when it is logged in the execution history.
	WorkerName *string

	noSmithyDocumentSerde
}

type GetActivityTaskOutput struct {

	// The string that contains the JSON input data for the task. Length constraints
	// apply to the payload size, and are expressed as bytes in UTF-8 encoding.
	Input *string

	// A token that identifies the scheduled task. This token must be copied and
	// included in subsequent calls to SendTaskHeartbeat, SendTaskSuccess or SendTaskFailure in order to report the progress or
	// completion of the task.
	TaskToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetActivityTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetActivityTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetActivityTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetActivityTask"); err != nil {
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
	if err = addOpGetActivityTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetActivityTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetActivityTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetActivityTask",
	}
}
