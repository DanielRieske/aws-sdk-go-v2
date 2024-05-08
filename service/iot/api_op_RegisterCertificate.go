// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a device certificate with IoT in the same [certificate mode] as the signing CA. If you
// have more than one CA certificate that has the same subject field, you must
// specify the CA certificate that was used to sign the device certificate being
// registered.
//
// Requires permission to access the [RegisterCertificate] action.
//
// [RegisterCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [certificate mode]: https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode
func (c *Client) RegisterCertificate(ctx context.Context, params *RegisterCertificateInput, optFns ...func(*Options)) (*RegisterCertificateOutput, error) {
	if params == nil {
		params = &RegisterCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCertificate", params, optFns, c.addOperationRegisterCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the RegisterCertificate operation.
type RegisterCertificateInput struct {

	// The certificate data, in PEM format.
	//
	// This member is required.
	CertificatePem *string

	// The CA certificate used to sign the device certificate being registered.
	CaCertificatePem *string

	// A boolean value that specifies if the certificate is set to active.
	//
	// Valid values: ACTIVE | INACTIVE
	//
	// Deprecated: This member has been deprecated.
	SetAsActive *bool

	// The status of the register certificate request. Valid values that you can use
	// include ACTIVE , INACTIVE , and REVOKED .
	Status types.CertificateStatus

	noSmithyDocumentSerde
}

// The output from the RegisterCertificate operation.
type RegisterCertificateOutput struct {

	// The certificate ARN.
	CertificateArn *string

	// The certificate identifier.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterCertificate"); err != nil {
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
	if err = addOpRegisterCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterCertificate",
	}
}
