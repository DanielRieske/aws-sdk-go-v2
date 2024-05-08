// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks out the specified license for offline use.
func (c *Client) CheckoutBorrowLicense(ctx context.Context, params *CheckoutBorrowLicenseInput, optFns ...func(*Options)) (*CheckoutBorrowLicenseOutput, error) {
	if params == nil {
		params = &CheckoutBorrowLicenseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckoutBorrowLicense", params, optFns, c.addOperationCheckoutBorrowLicenseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckoutBorrowLicenseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckoutBorrowLicenseInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// Digital signature method. The possible value is JSON Web Signature (JWS)
	// algorithm PS384. For more information, see [RFC 7518 Digital Signature with RSASSA-PSS].
	//
	// [RFC 7518 Digital Signature with RSASSA-PSS]: https://tools.ietf.org/html/rfc7518#section-3.5
	//
	// This member is required.
	DigitalSignatureMethod types.DigitalSignatureMethod

	// License entitlements. Partial checkouts are not supported.
	//
	// This member is required.
	Entitlements []types.EntitlementData

	// Amazon Resource Name (ARN) of the license. The license must use the borrow
	// consumption configuration.
	//
	// This member is required.
	LicenseArn *string

	// Information about constraints.
	CheckoutMetadata []types.Metadata

	// Node ID.
	NodeId *string

	noSmithyDocumentSerde
}

type CheckoutBorrowLicenseOutput struct {

	// Information about constraints.
	CheckoutMetadata []types.Metadata

	// Allowed license entitlements.
	EntitlementsAllowed []types.EntitlementData

	// Date and time at which the license checkout expires.
	Expiration *string

	// Date and time at which the license checkout is issued.
	IssuedAt *string

	// Amazon Resource Name (ARN) of the license.
	LicenseArn *string

	// License consumption token.
	LicenseConsumptionToken *string

	// Node ID.
	NodeId *string

	// Signed token.
	SignedToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCheckoutBorrowLicenseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCheckoutBorrowLicense{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckoutBorrowLicense{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CheckoutBorrowLicense"); err != nil {
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
	if err = addOpCheckoutBorrowLicenseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckoutBorrowLicense(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckoutBorrowLicense(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CheckoutBorrowLicense",
	}
}
