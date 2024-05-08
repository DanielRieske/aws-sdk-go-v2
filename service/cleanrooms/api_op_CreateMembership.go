// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a membership for a specific collaboration identifier and joins the
// collaboration.
func (c *Client) CreateMembership(ctx context.Context, params *CreateMembershipInput, optFns ...func(*Options)) (*CreateMembershipOutput, error) {
	if params == nil {
		params = &CreateMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembership", params, optFns, c.addOperationCreateMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembershipInput struct {

	// The unique ID for the associated collaboration.
	//
	// This member is required.
	CollaborationIdentifier *string

	// An indicator as to whether query logging has been enabled or disabled for the
	// membership.
	//
	// This member is required.
	QueryLogStatus types.MembershipQueryLogStatus

	// The default protected query result configuration as specified by the member who
	// can receive results.
	DefaultResultConfiguration *types.MembershipProtectedQueryResultConfiguration

	// The payment responsibilities accepted by the collaboration member.
	//
	// Not required if the collaboration member has the member ability to run queries.
	//
	// Required if the collaboration member doesn't have the member ability to run
	// queries but is configured as a payer by the collaboration creator.
	PaymentConfiguration *types.MembershipPaymentConfiguration

	// An optional label that you can assign to a resource when you create it. Each
	// tag consists of a key and an optional value, both of which you define. When you
	// use tagging, you can also use tag-based access control in IAM policies to
	// control access to this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateMembershipOutput struct {

	// The membership that was created.
	//
	// This member is required.
	Membership *types.Membership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMembership"); err != nil {
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
	if err = addOpCreateMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMembership",
	}
}
