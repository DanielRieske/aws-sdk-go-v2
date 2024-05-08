// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This IAM-authenticated API operation provides a code that Amazon Cognito sent
// to your user when they signed up in your user pool. After your user enters their
// code, they confirm ownership of the email address or phone number that they
// provided, and their user account becomes active. Depending on your user pool
// configuration, your users will receive their confirmation code in an email or
// SMS message.
//
// Local users who signed up in your user pool are the only type of user who can
// confirm sign-up with a code. Users who federate through an external identity
// provider (IdP) have already been confirmed by their IdP. Administrator-created
// users confirm their accounts when they respond to their invitation email message
// and choose a password.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) AdminConfirmSignUp(ctx context.Context, params *AdminConfirmSignUpInput, optFns ...func(*Options)) (*AdminConfirmSignUpOutput, error) {
	if params == nil {
		params = &AdminConfirmSignUpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminConfirmSignUp", params, optFns, c.addOperationAdminConfirmSignUpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminConfirmSignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Confirm a user's registration as a user pool administrator.
type AdminConfirmSignUpInput struct {

	// The user pool ID for which you want to confirm user registration.
	//
	// This member is required.
	UserPoolId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// If your user pool configuration includes triggers, the AdminConfirmSignUp API
	// action invokes the Lambda function that is specified for the post confirmation
	// trigger. When Amazon Cognito invokes this function, it passes a JSON payload,
	// which the function receives as input. In this payload, the clientMetadata
	// attribute provides the data that you assigned to the ClientMetadata parameter in
	// your AdminConfirmSignUp request. In your function code in Lambda, you can
	// process the ClientMetadata value to enhance your workflow for your specific
	// needs.
	//
	// For more information, see [Customizing user pool Workflows with Lambda Triggers] in the Amazon Cognito Developer Guide.
	//
	// When you use the ClientMetadata parameter, remember that Amazon Cognito won't
	// do the following:
	//
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//
	//   - Validate the ClientMetadata value.
	//
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	//
	// [Customizing user pool Workflows with Lambda Triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// Represents the response from the server for the request to confirm registration.
type AdminConfirmSignUpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminConfirmSignUpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminConfirmSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminConfirmSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminConfirmSignUp"); err != nil {
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
	if err = addOpAdminConfirmSignUpValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminConfirmSignUp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminConfirmSignUp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminConfirmSignUp",
	}
}
