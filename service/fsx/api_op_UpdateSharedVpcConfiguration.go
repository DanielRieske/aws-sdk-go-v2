// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures whether participant accounts in your organization can create Amazon
// FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a
// virtual private cloud (VPC) owner. For more information, see the [Amazon FSx for NetApp ONTAP User Guide].
//
// We strongly recommend that participant-created Multi-AZ file systems in the
// shared VPC are deleted before you disable this feature. Once the feature is
// disabled, these file systems will enter a MISCONFIGURED state and behave like
// Single-AZ file systems. For more information, see [Important considerations before disabling shared VPC support for Multi-AZ file systems].
//
// [Amazon FSx for NetApp ONTAP User Guide]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html
// [Important considerations before disabling shared VPC support for Multi-AZ file systems]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html#disabling-maz-vpc-sharing
func (c *Client) UpdateSharedVpcConfiguration(ctx context.Context, params *UpdateSharedVpcConfigurationInput, optFns ...func(*Options)) (*UpdateSharedVpcConfigurationOutput, error) {
	if params == nil {
		params = &UpdateSharedVpcConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSharedVpcConfiguration", params, optFns, c.addOperationUpdateSharedVpcConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSharedVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSharedVpcConfigurationInput struct {

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// Specifies whether participant accounts can create FSx for ONTAP Multi-AZ file
	// systems in shared subnets. Set to true to enable or false to disable.
	EnableFsxRouteTableUpdatesFromParticipantAccounts *string

	noSmithyDocumentSerde
}

type UpdateSharedVpcConfigurationOutput struct {

	// Indicates whether participant accounts can create FSx for ONTAP Multi-AZ file
	// systems in shared subnets.
	EnableFsxRouteTableUpdatesFromParticipantAccounts *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSharedVpcConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSharedVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSharedVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSharedVpcConfiguration"); err != nil {
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
	if err = addIdempotencyToken_opUpdateSharedVpcConfigurationMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSharedVpcConfiguration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSharedVpcConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSharedVpcConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSharedVpcConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSharedVpcConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSharedVpcConfigurationInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSharedVpcConfigurationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSharedVpcConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSharedVpcConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSharedVpcConfiguration",
	}
}
