// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration for an action group for an agent.
func (c *Client) UpdateAgentActionGroup(ctx context.Context, params *UpdateAgentActionGroupInput, optFns ...func(*Options)) (*UpdateAgentActionGroupOutput, error) {
	if params == nil {
		params = &UpdateAgentActionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAgentActionGroup", params, optFns, c.addOperationUpdateAgentActionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAgentActionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAgentActionGroupInput struct {

	// The unique identifier of the action group.
	//
	// This member is required.
	ActionGroupId *string

	// Specifies a new name for the action group.
	//
	// This member is required.
	ActionGroupName *string

	// The unique identifier of the agent for which to update the action group.
	//
	// This member is required.
	AgentId *string

	// The unique identifier of the agent version for which to update the action group.
	//
	// This member is required.
	AgentVersion *string

	// The Amazon Resource Name (ARN) of the Lambda function containing the business
	// logic that is carried out upon invoking the action.
	ActionGroupExecutor types.ActionGroupExecutor

	// Specifies whether the action group is available for the agent to invoke or not
	// when sending an InvokeAgent (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html)
	// request.
	ActionGroupState types.ActionGroupState

	// Contains either details about the S3 object containing the OpenAPI schema for
	// the action group or the JSON or YAML-formatted payload defining the schema. For
	// more information, see Action group OpenAPI schemas (https://docs.aws.amazon.com/bedrock/latest/userguide/agents-api-schema.html)
	// .
	ApiSchema types.APISchema

	// Specifies a new name for the action group.
	Description *string

	// To allow your agent to request the user for additional information when trying
	// to complete a task, set this field to AMAZON.UserInput . You must leave the
	// description , apiSchema , and actionGroupExecutor fields blank for this action
	// group. During orchestration, if your agent determines that it needs to invoke an
	// API in an action group, but doesn't have enough information to complete the API
	// request, it will invoke this action group instead and return an Observation (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html)
	// reprompting the user for more information.
	ParentActionGroupSignature types.ActionGroupSignature

	noSmithyDocumentSerde
}

type UpdateAgentActionGroupOutput struct {

	// Contains details about the action group that was updated.
	//
	// This member is required.
	AgentActionGroup *types.AgentActionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAgentActionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAgentActionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAgentActionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAgentActionGroup"); err != nil {
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
	if err = addOpUpdateAgentActionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAgentActionGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAgentActionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAgentActionGroup",
	}
}
