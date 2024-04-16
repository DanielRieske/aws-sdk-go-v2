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

// Creates an agent that orchestrates interactions between foundation models, data
// sources, software applications, user conversations, and APIs to carry out tasks
// to help customers.
//   - Specify the following fields for security purposes.
//   - agentResourceRoleArn – The Amazon Resource Name (ARN) of the role with
//     permissions to invoke API operations on an agent.
//   - (Optional) customerEncryptionKeyArn – The Amazon Resource Name (ARN) of a
//     KMS key to encrypt the creation of the agent.
//   - (Optional) idleSessionTTLinSeconds – Specify the number of seconds for which
//     the agent should maintain session information. After this time expires, the
//     subsequent InvokeAgent request begins a new session.
//   - To override the default prompt behavior for agent orchestration and to use
//     advanced prompts, include a promptOverrideConfiguration object. For more
//     information, see Advanced prompts (https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html)
//     .
//   - If you agent fails to be created, the response returns a list of
//     failureReasons alongside a list of recommendedActions for you to troubleshoot.
func (c *Client) CreateAgent(ctx context.Context, params *CreateAgentInput, optFns ...func(*Options)) (*CreateAgentOutput, error) {
	if params == nil {
		params = &CreateAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAgent", params, optFns, c.addOperationCreateAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAgentInput struct {

	// A name for the agent that you create.
	//
	// This member is required.
	AgentName *string

	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API
	// operations on the agent.
	AgentResourceRoleArn *string

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see
	// Ensuring idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The Amazon Resource Name (ARN) of the KMS key with which to encrypt the agent.
	CustomerEncryptionKeyArn *string

	// A description of the agent.
	Description *string

	// The foundation model to be used for orchestration by the agent you create.
	FoundationModel *string

	// The number of seconds for which Amazon Bedrock keeps information about a user's
	// conversation with the agent. A user interaction remains active for the amount of
	// time specified. If no conversation occurs during this time, the session expires
	// and Amazon Bedrock deletes any data provided before the timeout.
	IdleSessionTTLInSeconds *int32

	// Instructions that tell the agent what it should do and how it should interact
	// with users.
	Instruction *string

	// Contains configurations to override prompts in different parts of an agent
	// sequence. For more information, see Advanced prompts (https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html)
	// .
	PromptOverrideConfiguration *types.PromptOverrideConfiguration

	// Any tags that you want to attach to the agent.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAgentOutput struct {

	// Contains details about the agent created.
	//
	// This member is required.
	Agent *types.Agent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAgent"); err != nil {
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
	if err = addIdempotencyToken_opCreateAgentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAgent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAgent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAgent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAgent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAgentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAgentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAgentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAgent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAgent",
	}
}
