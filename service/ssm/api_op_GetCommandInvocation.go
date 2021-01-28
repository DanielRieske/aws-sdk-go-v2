// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns detailed information about command execution for an invocation or
// plugin.
func (c *Client) GetCommandInvocation(ctx context.Context, params *GetCommandInvocationInput, optFns ...func(*Options)) (*GetCommandInvocationOutput, error) {
	if params == nil {
		params = &GetCommandInvocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCommandInvocation", params, optFns, addOperationGetCommandInvocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommandInvocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommandInvocationInput struct {

	// (Required) The parent command ID of the invocation plugin.
	//
	// This member is required.
	CommandId *string

	// (Required) The ID of the managed instance targeted by the command. A managed
	// instance can be an EC2 instance or an instance in your hybrid environment that
	// is configured for Systems Manager.
	//
	// This member is required.
	InstanceId *string

	// (Optional) The name of the plugin for which you want detailed results. If the
	// document contains only one plugin, the name can be omitted and the details will
	// be returned. Plugin names are also referred to as step names in Systems Manager
	// documents.
	PluginName *string
}

type GetCommandInvocationOutput struct {

	// CloudWatch Logs information where Systems Manager sent the command output.
	CloudWatchOutputConfig *types.CloudWatchOutputConfig

	// The parent command ID of the invocation plugin.
	CommandId *string

	// The comment text for the command.
	Comment *string

	// The name of the document that was run. For example, AWS-RunShellScript.
	DocumentName *string

	// The SSM document version used in the request.
	DocumentVersion *string

	// Duration since ExecutionStartDateTime.
	ExecutionElapsedTime *string

	// The date and time the plugin was finished running. Date and time are written in
	// ISO 8601 format. For example, June 7, 2017 is represented as 2017-06-7. The
	// following sample AWS CLI command uses the InvokedAfter filter. aws ssm
	// list-commands --filters key=InvokedAfter,value=2017-06-07T00:00:00Z If the
	// plugin has not started to run, the string is empty.
	ExecutionEndDateTime *string

	// The date and time the plugin started running. Date and time are written in ISO
	// 8601 format. For example, June 7, 2017 is represented as 2017-06-7. The
	// following sample AWS CLI command uses the InvokedBefore filter. aws ssm
	// list-commands --filters key=InvokedBefore,value=2017-06-07T00:00:00Z If the
	// plugin has not started to run, the string is empty.
	ExecutionStartDateTime *string

	// The ID of the managed instance targeted by the command. A managed instance can
	// be an EC2 instance or an instance in your hybrid environment that is configured
	// for Systems Manager.
	InstanceId *string

	// The name of the plugin for which you want detailed results. For example,
	// aws:RunShellScript is a plugin.
	PluginName *string

	// The error level response code for the plugin script. If the response code is -1,
	// then the command has not started running on the instance, or it was not received
	// by the instance.
	ResponseCode int32

	// The first 8,000 characters written by the plugin to stderr. If the command has
	// not finished running, then this string is empty.
	StandardErrorContent *string

	// The URL for the complete text written by the plugin to stderr. If the command
	// has not finished running, then this string is empty.
	StandardErrorUrl *string

	// The first 24,000 characters written by the plugin to stdout. If the command has
	// not finished running, if ExecutionStatus is neither Succeeded nor Failed, then
	// this string is empty.
	StandardOutputContent *string

	// The URL for the complete text written by the plugin to stdout in Amazon S3. If
	// an S3 bucket was not specified, then this string is empty.
	StandardOutputUrl *string

	// The status of this invocation plugin. This status can be different than
	// StatusDetails.
	Status types.CommandInvocationStatus

	// A detailed status of the command execution for an invocation. StatusDetails
	// includes more information than Status because it includes states resulting from
	// error and concurrency control parameters. StatusDetails can show different
	// results than Status. For more information about these statuses, see
	// Understanding command statuses
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/monitor-commands.html)
	// in the AWS Systems Manager User Guide. StatusDetails can be one of the following
	// values:
	//
	// * Pending: The command has not been sent to the instance.
	//
	// * In
	// Progress: The command has been sent to the instance but has not reached a
	// terminal state.
	//
	// * Delayed: The system attempted to send the command to the
	// target, but the target was not available. The instance might not be available
	// because of network issues, because the instance was stopped, or for similar
	// reasons. The system will try to send the command again.
	//
	// * Success: The command
	// or plugin ran successfully. This is a terminal state.
	//
	// * Delivery Timed Out: The
	// command was not delivered to the instance before the delivery timeout expired.
	// Delivery timeouts do not count against the parent command's MaxErrors limit, but
	// they do contribute to whether the parent command status is Success or
	// Incomplete. This is a terminal state.
	//
	// * Execution Timed Out: The command
	// started to run on the instance, but the execution was not complete before the
	// timeout expired. Execution timeouts count against the MaxErrors limit of the
	// parent command. This is a terminal state.
	//
	// * Failed: The command wasn't run
	// successfully on the instance. For a plugin, this indicates that the result code
	// was not zero. For a command invocation, this indicates that the result code for
	// one or more plugins was not zero. Invocation failures count against the
	// MaxErrors limit of the parent command. This is a terminal state.
	//
	// * Canceled:
	// The command was terminated before it was completed. This is a terminal state.
	//
	// *
	// Undeliverable: The command can't be delivered to the instance. The instance
	// might not exist or might not be responding. Undeliverable invocations don't
	// count against the parent command's MaxErrors limit and don't contribute to
	// whether the parent command status is Success or Incomplete. This is a terminal
	// state.
	//
	// * Terminated: The parent command exceeded its MaxErrors limit and
	// subsequent command invocations were canceled by the system. This is a terminal
	// state.
	StatusDetails *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCommandInvocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommandInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommandInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetCommandInvocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommandInvocation(options.Region), middleware.Before); err != nil {
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
	return nil
}

// GetCommandInvocationAPIClient is a client that implements the
// GetCommandInvocation operation.
type GetCommandInvocationAPIClient interface {
	GetCommandInvocation(context.Context, *GetCommandInvocationInput, ...func(*Options)) (*GetCommandInvocationOutput, error)
}

var _ GetCommandInvocationAPIClient = (*Client)(nil)

// CommandExecutedWaiterOptions are waiter options for CommandExecutedWaiter
type CommandExecutedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CommandExecutedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, CommandExecutedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetCommandInvocationInput, *GetCommandInvocationOutput, error) (bool, error)
}

// CommandExecutedWaiter defines the waiters for CommandExecuted
type CommandExecutedWaiter struct {
	client GetCommandInvocationAPIClient

	options CommandExecutedWaiterOptions
}

// NewCommandExecutedWaiter constructs a CommandExecutedWaiter.
func NewCommandExecutedWaiter(client GetCommandInvocationAPIClient, optFns ...func(*CommandExecutedWaiterOptions)) *CommandExecutedWaiter {
	options := CommandExecutedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = commandExecutedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CommandExecutedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CommandExecuted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *CommandExecutedWaiter) Wait(ctx context.Context, params *GetCommandInvocationInput, maxWaitDur time.Duration, optFns ...func(*CommandExecutedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetCommandInvocation(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for CommandExecuted waiter")
}

func commandExecutedStateRetryable(ctx context.Context, input *GetCommandInvocationInput, output *GetCommandInvocationOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Pending"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InProgress"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Delayed"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Success"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Cancelled"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TimedOut"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Cancelling"
		value, ok := pathValue.(types.CommandInvocationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CommandInvocationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetCommandInvocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetCommandInvocation",
	}
}
