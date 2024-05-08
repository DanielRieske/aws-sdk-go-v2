// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of open workflow executions within the given domain that
// meet the specified filtering criteria.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//
//   - Use an Action element to allow or deny permission to call this action.
//
//   - Constrain the following parameters by using a Condition element with the
//     appropriate keys.
//
//   - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//
//   - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//
//   - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func (c *Client) CountOpenWorkflowExecutions(ctx context.Context, params *CountOpenWorkflowExecutionsInput, optFns ...func(*Options)) (*CountOpenWorkflowExecutionsOutput, error) {
	if params == nil {
		params = &CountOpenWorkflowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CountOpenWorkflowExecutions", params, optFns, c.addOperationCountOpenWorkflowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CountOpenWorkflowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CountOpenWorkflowExecutionsInput struct {

	// The name of the domain containing the workflow executions to count.
	//
	// This member is required.
	Domain *string

	// Specifies the start time criteria that workflow executions must meet in order
	// to be counted.
	//
	// This member is required.
	StartTimeFilter *types.ExecutionTimeFilter

	// If specified, only workflow executions matching the WorkflowId in the filter
	// are counted.
	//
	// executionFilter , typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	ExecutionFilter *types.WorkflowExecutionFilter

	// If specified, only executions that have a tag that matches the filter are
	// counted.
	//
	// executionFilter , typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TagFilter *types.TagFilter

	// Specifies the type of the workflow executions to be counted.
	//
	// executionFilter , typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TypeFilter *types.WorkflowTypeFilter

	noSmithyDocumentSerde
}

// Contains the count of workflow executions returned from CountOpenWorkflowExecutions or CountClosedWorkflowExecutions
type CountOpenWorkflowExecutionsOutput struct {

	// The number of workflow executions.
	//
	// This member is required.
	Count int32

	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCountOpenWorkflowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCountOpenWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCountOpenWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CountOpenWorkflowExecutions"); err != nil {
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
	if err = addOpCountOpenWorkflowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCountOpenWorkflowExecutions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCountOpenWorkflowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CountOpenWorkflowExecutions",
	}
}
