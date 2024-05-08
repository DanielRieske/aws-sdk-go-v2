// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of an existing data repository association on an
// Amazon FSx for Lustre file system. Data repository associations are supported on
// all FSx for Lustre 2.12 and 2.15 file systems, excluding scratch_1 deployment
// type.
func (c *Client) UpdateDataRepositoryAssociation(ctx context.Context, params *UpdateDataRepositoryAssociationInput, optFns ...func(*Options)) (*UpdateDataRepositoryAssociationOutput, error) {
	if params == nil {
		params = &UpdateDataRepositoryAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataRepositoryAssociation", params, optFns, c.addOperationUpdateDataRepositoryAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataRepositoryAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataRepositoryAssociationInput struct {

	// The ID of the data repository association that you are updating.
	//
	// This member is required.
	AssociationId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// For files imported from a data repository, this value determines the stripe
	// count and maximum amount of data per file (in MiB) stored on a single physical
	// disk. The maximum number of disks that a single file can be striped across is
	// limited by the total number of disks that make up the file system.
	//
	// The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB
	// (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	ImportedFileChunkSize *int32

	// The configuration for an Amazon S3 data repository linked to an Amazon FSx
	// Lustre file system with a data repository association. The configuration defines
	// which file events (new, changed, or deleted files or directories) are
	// automatically imported from the linked data repository to the file system or
	// automatically exported from the file system to the data repository.
	S3 *types.S3DataRepositoryConfiguration

	noSmithyDocumentSerde
}

type UpdateDataRepositoryAssociationOutput struct {

	// The response object returned after the data repository association is updated.
	Association *types.DataRepositoryAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataRepositoryAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDataRepositoryAssociation"); err != nil {
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
	if err = addIdempotencyToken_opUpdateDataRepositoryAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateDataRepositoryAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataRepositoryAssociation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateDataRepositoryAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateDataRepositoryAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateDataRepositoryAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateDataRepositoryAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateDataRepositoryAssociationInput ")
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
func addIdempotencyToken_opUpdateDataRepositoryAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateDataRepositoryAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateDataRepositoryAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDataRepositoryAssociation",
	}
}
