// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Makes an authorization decision about a service request described in the
// parameters. The principal in this request comes from an external identity source
// in the form of an identity token formatted as a [JSON web token (JWT)]. The information in the
// parameters can also define additional context that Verified Permissions can
// include in the evaluation. The request is evaluated against all matching
// policies in the specified policy store. The result of the decision is either
// Allow or Deny , along with a list of the policies that resulted in the decision.
//
// At this time, Verified Permissions accepts tokens from only Amazon Cognito.
//
// Verified Permissions validates each token that is specified in a request by
// checking its expiration date and its signature.
//
// If you delete a Amazon Cognito user pool or user, tokens from that deleted pool
// or that deleted user continue to be usable until they expire.
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
func (c *Client) IsAuthorizedWithToken(ctx context.Context, params *IsAuthorizedWithTokenInput, optFns ...func(*Options)) (*IsAuthorizedWithTokenOutput, error) {
	if params == nil {
		params = &IsAuthorizedWithTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IsAuthorizedWithToken", params, optFns, c.addOperationIsAuthorizedWithTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IsAuthorizedWithTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IsAuthorizedWithTokenInput struct {

	// Specifies the ID of the policy store. Policies in this policy store will be
	// used to make an authorization decision for the input.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies an access token for the principal to be authorized. This token is
	// provided to you by the identity provider (IdP) associated with the specified
	// identity source. You must specify either an accessToken , an identityToken , or
	// both.
	//
	// Must be an access token. Verified Permissions returns an error if the token_use
	// claim in the submitted token isn't access .
	AccessToken *string

	// Specifies the requested action to be authorized. Is the specified principal
	// authorized to perform this action on the specified resource.
	Action *types.ActionIdentifier

	// Specifies additional context that can be used to make more granular
	// authorization decisions.
	Context types.ContextDefinition

	// Specifies the list of resources and their associated attributes that Verified
	// Permissions can examine when evaluating the policies.
	//
	// You can't include principals in this parameter, only resource and action
	// entities. This parameter can't include any entities of a type that matches the
	// user or group entity types that you defined in your identity source.
	//
	//   - The IsAuthorizedWithToken operation takes principal attributes from only the
	//   identityToken or accessToken passed to the operation.
	//
	//   - For action entities, you can include only their Identifier and EntityType .
	Entities types.EntitiesDefinition

	// Specifies an identity token for the principal to be authorized. This token is
	// provided to you by the identity provider (IdP) associated with the specified
	// identity source. You must specify either an accessToken , an identityToken , or
	// both.
	//
	// Must be an ID token. Verified Permissions returns an error if the token_use
	// claim in the submitted token isn't id .
	IdentityToken *string

	// Specifies the resource for which the authorization decision is made. For
	// example, is the principal allowed to perform the action on the resource?
	Resource *types.EntityIdentifier

	noSmithyDocumentSerde
}

type IsAuthorizedWithTokenOutput struct {

	// An authorization decision that indicates if the authorization request should be
	// allowed or denied.
	//
	// This member is required.
	Decision types.Decision

	// The list of determining policies used to make the authorization decision. For
	// example, if there are multiple matching policies, where at least one is a forbid
	// policy, then because forbid always overrides permit the forbid policies are the
	// determining policies. If all matching policies are permit policies, then those
	// policies are the determining policies. When no policies match and the response
	// is the default DENY, there are no determining policies.
	//
	// This member is required.
	DeterminingPolicies []types.DeterminingPolicyItem

	// Errors that occurred while making an authorization decision. For example, a
	// policy references an entity or entity attribute that does not exist in the
	// slice.
	//
	// This member is required.
	Errors []types.EvaluationErrorItem

	// The identifier of the principal in the ID or access token.
	Principal *types.EntityIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIsAuthorizedWithTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpIsAuthorizedWithToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpIsAuthorizedWithToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "IsAuthorizedWithToken"); err != nil {
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
	if err = addOpIsAuthorizedWithTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIsAuthorizedWithToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIsAuthorizedWithToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IsAuthorizedWithToken",
	}
}
