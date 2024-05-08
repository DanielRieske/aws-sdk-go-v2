// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an Config rule for your entire organization to evaluate if your
// Amazon Web Services resources comply with your desired configurations. For
// information on how many organization Config rules you can have per account, see [Service Limits]
// in the Config Developer Guide.
//
// Only a management account and a delegated administrator can create or update an
// organization Config rule. When calling this API with a delegated administrator,
// you must ensure Organizations ListDelegatedAdministrator permissions are added.
// An organization can have up to 3 delegated administrators.
//
// This API enables organization service access through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. Config verifies the
// existence of role with GetRole action.
//
// To use this API with delegated administrator, register a delegated
// administrator by calling Amazon Web Services Organization
// register-delegated-administrator for config-multiaccountsetup.amazonaws.com .
//
// There are two types of rules: Config Managed Rules and Config Custom Rules. You
// can use PutOrganizationConfigRule to create both Config Managed Rules and
// Config Custom Rules.
//
// Config Managed Rules are predefined, customizable rules created by Config. For
// a list of managed rules, see [List of Config Managed Rules]. If you are adding an Config managed rule, you
// must specify the rule's identifier for the RuleIdentifier key.
//
// Config Custom Rules are rules that you create from scratch. There are two ways
// to create Config custom rules: with Lambda functions ([Lambda Developer Guide] ) and with Guard ([Guard GitHub Repository] ), a
// policy-as-code language.
//
// Config custom rules created with Lambda are called Config Custom Lambda Rules
// and Config custom rules created with Guard are called Config Custom Policy
// Rules.
//
// If you are adding a new Config Custom Lambda rule, you first need to create an
// Lambda function in the management account or a delegated administrator that the
// rule invokes to evaluate your resources. You also need to create an IAM role in
// the managed account that can be assumed by the Lambda function. When you use
// PutOrganizationConfigRule to add a Custom Lambda rule to Config, you must
// specify the Amazon Resource Name (ARN) that Lambda assigns to the function.
//
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in
// an organization.
//
// Make sure to specify one of either OrganizationCustomPolicyRuleMetadata for
// Custom Policy rules, OrganizationCustomRuleMetadata for Custom Lambda rules, or
// OrganizationManagedRuleMetadata for managed rules.
//
// [List of Config Managed Rules]: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
// [Lambda Developer Guide]: https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
// [Guard GitHub Repository]: https://github.com/aws-cloudformation/cloudformation-guard
func (c *Client) PutOrganizationConfigRule(ctx context.Context, params *PutOrganizationConfigRuleInput, optFns ...func(*Options)) (*PutOrganizationConfigRuleOutput, error) {
	if params == nil {
		params = &PutOrganizationConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutOrganizationConfigRule", params, optFns, c.addOperationPutOrganizationConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutOrganizationConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutOrganizationConfigRuleInput struct {

	// The name that you assign to an organization Config rule.
	//
	// This member is required.
	OrganizationConfigRuleName *string

	// A comma-separated list of accounts that you want to exclude from an
	// organization Config rule.
	ExcludedAccounts []string

	// An OrganizationCustomPolicyRuleMetadata object. This object specifies metadata
	// for your organization's Config Custom Policy rule. The metadata includes the
	// runtime system in use, which accounts have debug logging enabled, and other
	// custom rule metadata, such as resource type, resource ID of Amazon Web Services
	// resource, and organization trigger types that initiate Config to evaluate Amazon
	// Web Services resources against a rule.
	OrganizationCustomPolicyRuleMetadata *types.OrganizationCustomPolicyRuleMetadata

	// An OrganizationCustomRuleMetadata object. This object specifies organization
	// custom rule metadata such as resource type, resource ID of Amazon Web Services
	// resource, Lambda function ARN, and organization trigger types that trigger
	// Config to evaluate your Amazon Web Services resources against a rule. It also
	// provides the frequency with which you want Config to run evaluations for the
	// rule if the trigger type is periodic.
	OrganizationCustomRuleMetadata *types.OrganizationCustomRuleMetadata

	// An OrganizationManagedRuleMetadata object. This object specifies organization
	// managed rule metadata such as resource type and ID of Amazon Web Services
	// resource along with the rule identifier. It also provides the frequency with
	// which you want Config to run evaluations for the rule if the trigger type is
	// periodic.
	OrganizationManagedRuleMetadata *types.OrganizationManagedRuleMetadata

	noSmithyDocumentSerde
}

type PutOrganizationConfigRuleOutput struct {

	// The Amazon Resource Name (ARN) of an organization Config rule.
	OrganizationConfigRuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutOrganizationConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutOrganizationConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutOrganizationConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutOrganizationConfigRule"); err != nil {
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
	if err = addOpPutOrganizationConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutOrganizationConfigRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutOrganizationConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutOrganizationConfigRule",
	}
}
