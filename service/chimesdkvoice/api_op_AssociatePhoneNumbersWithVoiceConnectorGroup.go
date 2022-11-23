// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorGroupInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
	if params == nil {
		params = &AssociatePhoneNumbersWithVoiceConnectorGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePhoneNumbersWithVoiceConnectorGroup", params, optFns, c.addOperationAssociatePhoneNumbersWithVoiceConnectorGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePhoneNumbersWithVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePhoneNumbersWithVoiceConnectorGroupInput struct {

	// This member is required.
	E164PhoneNumbers []string

	// This member is required.
	VoiceConnectorGroupId *string

	ForceAssociate *bool

	noSmithyDocumentSerde
}

type AssociatePhoneNumbersWithVoiceConnectorGroupOutput struct {
	PhoneNumberErrors []types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePhoneNumbersWithVoiceConnectorGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociatePhoneNumbersWithVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociatePhoneNumbersWithVoiceConnectorGroup{}, middleware.After)
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
	if err = addOpAssociatePhoneNumbersWithVoiceConnectorGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnectorGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePhoneNumbersWithVoiceConnectorGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "AssociatePhoneNumbersWithVoiceConnectorGroup",
	}
}