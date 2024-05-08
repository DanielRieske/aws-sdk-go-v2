// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a budget action.
func (c *Client) UpdateBudgetAction(ctx context.Context, params *UpdateBudgetActionInput, optFns ...func(*Options)) (*UpdateBudgetActionOutput, error) {
	if params == nil {
		params = &UpdateBudgetActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBudgetAction", params, optFns, c.addOperationUpdateBudgetActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBudgetActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBudgetActionInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	//  A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	//  A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// The trigger threshold of the action.
	ActionThreshold *types.ActionThreshold

	//  This specifies if the action needs manual or automatic approval.
	ApprovalModel types.ApprovalModel

	// Specifies all of the type-specific parameters.
	Definition *types.Definition

	//  The role passed for action execution and reversion. Roles and actions must be
	// in the same account.
	ExecutionRoleArn *string

	//  The type of a notification. It must be ACTUAL or FORECASTED.
	NotificationType types.NotificationType

	//  A list of subscribers.
	Subscribers []types.Subscriber

	noSmithyDocumentSerde
}

type UpdateBudgetActionOutput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	//  A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	//  The updated action resource information.
	//
	// This member is required.
	NewAction *types.Action

	//  The previous action resource information.
	//
	// This member is required.
	OldAction *types.Action

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBudgetActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBudgetAction"); err != nil {
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
	if err = addOpUpdateBudgetActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBudgetAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBudgetAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBudgetAction",
	}
}
