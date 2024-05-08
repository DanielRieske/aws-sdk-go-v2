// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the revocation status of one or more of the signing profile, signing
// job, and signing certificate.
func (c *Client) GetRevocationStatus(ctx context.Context, params *GetRevocationStatusInput, optFns ...func(*Options)) (*GetRevocationStatusOutput, error) {
	if params == nil {
		params = &GetRevocationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRevocationStatus", params, optFns, c.addOperationGetRevocationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRevocationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRevocationStatusInput struct {

	// A list of composite signed hashes that identify certificates.
	//
	// A certificate identifier consists of a subject certificate TBS hash (signed by
	// the parent CA) combined with a parent CA TBS hash (signed by the parent CA’s
	// CA). Root certificates are defined as their own CA.
	//
	// The following example shows how to calculate a hash for this parameter using
	// OpenSSL commands:
	//
	//     openssl asn1parse -in childCert.pem -strparse 4 -out childCert.tbs
	//
	//     openssl sha384 < childCert.tbs -binary > childCertTbsHash
	//
	//     openssl asn1parse -in parentCert.pem -strparse 4 -out parentCert.tbs
	//
	//     openssl sha384 < parentCert.tbs -binary > parentCertTbsHash xxd -p
	//     childCertTbsHash > certificateHash.hex xxd -p parentCertTbsHash >>
	//     certificateHash.hex
	//
	//     cat certificateHash.hex | tr -d '\n'
	//
	// This member is required.
	CertificateHashes []string

	// The ARN of a signing job.
	//
	// This member is required.
	JobArn *string

	// The ID of a signing platform.
	//
	// This member is required.
	PlatformId *string

	// The version of a signing profile.
	//
	// This member is required.
	ProfileVersionArn *string

	// The timestamp of the signature that validates the profile or job.
	//
	// This member is required.
	SignatureTimestamp *time.Time

	noSmithyDocumentSerde
}

type GetRevocationStatusOutput struct {

	// A list of revoked entities (including zero or more of the signing profile ARN,
	// signing job ARN, and certificate hashes) supplied as input to the API.
	RevokedEntities []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRevocationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRevocationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRevocationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRevocationStatus"); err != nil {
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
	if err = addEndpointPrefix_opGetRevocationStatusMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetRevocationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRevocationStatus(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetRevocationStatusMiddleware struct {
}

func (*endpointPrefix_opGetRevocationStatusMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetRevocationStatusMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "verification." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetRevocationStatusMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetRevocationStatusMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetRevocationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRevocationStatus",
	}
}
