// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get a step in the migration workflow.
func (c *Client) GetWorkflowStep(ctx context.Context, params *GetWorkflowStepInput, optFns ...func(*Options)) (*GetWorkflowStepOutput, error) {
	if params == nil {
		params = &GetWorkflowStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowStep", params, optFns, c.addOperationGetWorkflowStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowStepInput struct {

	// The ID of the step.
	//
	// This member is required.
	Id *string

	// desThe ID of the step group.
	//
	// This member is required.
	StepGroupId *string

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	noSmithyDocumentSerde
}

type GetWorkflowStepOutput struct {

	// The time at which the step was created.
	CreationTime *time.Time

	// The description of the step.
	Description *string

	// The time at which the step ended.
	EndTime *time.Time

	// The time at which the workflow was last started.
	LastStartTime *time.Time

	// The name of the step.
	Name *string

	// The next step.
	Next []string

	// The number of servers that have been migrated.
	NoOfSrvCompleted *int32

	// The number of servers that have failed to migrate.
	NoOfSrvFailed *int32

	// The outputs of the step.
	Outputs []types.WorkflowStepOutput

	// The owner of the step.
	Owner types.Owner

	// The previous step.
	Previous []string

	// The output location of the script.
	ScriptOutputLocation *string

	// The status of the step.
	Status types.StepStatus

	// The status message of the migration workflow.
	StatusMessage *string

	// The action type of the step. You must run and update the status of a manual step
	// for the workflow to continue after the completion of the step.
	StepActionType types.StepActionType

	// The ID of the step group.
	StepGroupId *string

	// The ID of the step.
	StepId *string

	// The servers on which a step will be run.
	StepTarget []string

	// The total number of servers that have been migrated.
	TotalNoOfSrv *int32

	// The ID of the migration workflow.
	WorkflowId *string

	// The custom script to run tests on source or target environments.
	WorkflowStepAutomationConfiguration *types.WorkflowStepAutomationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflowStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflowStep{}, middleware.After)
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
	if err = addOpGetWorkflowStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowStep(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflowStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "GetWorkflowStep",
	}
}