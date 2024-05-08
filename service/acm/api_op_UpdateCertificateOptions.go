// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a certificate. Currently, you can use this function to specify whether
// to opt in to or out of recording your certificate in a certificate transparency
// log. For more information, see [Opting Out of Certificate Transparency Logging].
//
// [Opting Out of Certificate Transparency Logging]: https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency
func (c *Client) UpdateCertificateOptions(ctx context.Context, params *UpdateCertificateOptionsInput, optFns ...func(*Options)) (*UpdateCertificateOptionsOutput, error) {
	if params == nil {
		params = &UpdateCertificateOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCertificateOptions", params, optFns, c.addOperationUpdateCertificateOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCertificateOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCertificateOptionsInput struct {

	// ARN of the requested certificate to update. This must be of the form:
	//
	//     arn:aws:acm:us-east-1:account:certificate/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateArn *string

	// Use to update the options for your certificate. Currently, you can specify
	// whether to add your certificate to a transparency log. Certificate transparency
	// makes it possible to detect SSL/TLS certificates that have been mistakenly or
	// maliciously issued. Certificates that have not been logged typically produce an
	// error message in a browser.
	//
	// This member is required.
	Options *types.CertificateOptions

	noSmithyDocumentSerde
}

type UpdateCertificateOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCertificateOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCertificateOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCertificateOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCertificateOptions"); err != nil {
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
	if err = addOpUpdateCertificateOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCertificateOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCertificateOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCertificateOptions",
	}
}
