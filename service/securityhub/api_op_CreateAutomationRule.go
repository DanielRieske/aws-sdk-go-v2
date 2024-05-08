// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an automation rule based on input parameters.
func (c *Client) CreateAutomationRule(ctx context.Context, params *CreateAutomationRuleInput, optFns ...func(*Options)) (*CreateAutomationRuleOutput, error) {
	if params == nil {
		params = &CreateAutomationRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutomationRule", params, optFns, c.addOperationCreateAutomationRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutomationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutomationRuleInput struct {

	//  One or more actions to update finding fields if a finding matches the
	// conditions specified in Criteria .
	//
	// This member is required.
	Actions []types.AutomationRulesAction

	//  A set of ASFF finding field attributes and corresponding expected values that
	// Security Hub uses to filter findings. If a rule is enabled and a finding matches
	// the conditions specified in this parameter, Security Hub applies the rule action
	// to the finding.
	//
	// This member is required.
	Criteria *types.AutomationRulesFindingFilters

	//  A description of the rule.
	//
	// This member is required.
	Description *string

	//  The name of the rule.
	//
	// This member is required.
	RuleName *string

	// An integer ranging from 1 to 1000 that represents the order in which the rule
	// action is applied to findings. Security Hub applies rules with lower values for
	// this parameter first.
	//
	// This member is required.
	RuleOrder *int32

	// Specifies whether a rule is the last to be applied with respect to a finding
	// that matches the rule criteria. This is useful when a finding matches the
	// criteria for multiple rules, and each rule has different actions. If a rule is
	// terminal, Security Hub applies the rule action to a finding that matches the
	// rule criteria and doesn't evaluate other rules for the finding. By default, a
	// rule isn't terminal.
	IsTerminal *bool

	//  Whether the rule is active after it is created. If this parameter is equal to
	// ENABLED , Security Hub starts applying the rule to findings and finding updates
	// after the rule is created. To change the value of this parameter after creating
	// a rule, use [BatchUpdateAutomationRules]BatchUpdateAutomationRules .
	//
	// [BatchUpdateAutomationRules]: https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_BatchUpdateAutomationRules.html
	RuleStatus types.RuleStatus

	//  User-defined tags associated with an automation rule.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAutomationRuleOutput struct {

	//  The Amazon Resource Name (ARN) of the automation rule that you created.
	RuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutomationRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAutomationRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAutomationRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAutomationRule"); err != nil {
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
	if err = addOpCreateAutomationRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutomationRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutomationRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAutomationRule",
	}
}
