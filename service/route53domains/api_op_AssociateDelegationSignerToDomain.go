// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates a delegation signer (DS) record in the registry zone for this domain
//
// name.
//
// Note that creating DS record at the registry impacts DNSSEC validation of your
// DNS records. This action may render your domain name unavailable on the internet
// if the steps are completed in the wrong order, or with incorrect timing. For
// more information about DNSSEC signing, see [Configuring DNSSEC signing]in the Route 53 developer guide.
//
// [Configuring DNSSEC signing]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html
func (c *Client) AssociateDelegationSignerToDomain(ctx context.Context, params *AssociateDelegationSignerToDomainInput, optFns ...func(*Options)) (*AssociateDelegationSignerToDomainOutput, error) {
	if params == nil {
		params = &AssociateDelegationSignerToDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDelegationSignerToDomain", params, optFns, c.addOperationAssociateDelegationSignerToDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDelegationSignerToDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDelegationSignerToDomainInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The information about a key, including the algorithm, public key-value, and
	// flags.
	//
	// This member is required.
	SigningAttributes *types.DnssecSigningAttributes

	noSmithyDocumentSerde
}

type AssociateDelegationSignerToDomainOutput struct {

	// The identifier for tracking the progress of the request. To query the operation
	// status, use [GetOperationDetail].
	//
	// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDelegationSignerToDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDelegationSignerToDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDelegationSignerToDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateDelegationSignerToDomain"); err != nil {
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
	if err = addOpAssociateDelegationSignerToDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDelegationSignerToDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateDelegationSignerToDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateDelegationSignerToDomain",
	}
}
