// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation is not supported by directory buckets.
//
// Gets the tags of Amazon S3 Storage Lens configuration. For more information
// about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:GetStorageLensConfigurationTagging action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the
// Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func (c *Client) GetStorageLensConfigurationTagging(ctx context.Context, params *GetStorageLensConfigurationTaggingInput, optFns ...func(*Options)) (*GetStorageLensConfigurationTaggingOutput, error) {
	if params == nil {
		params = &GetStorageLensConfigurationTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStorageLensConfigurationTagging", params, optFns, c.addOperationGetStorageLensConfigurationTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStorageLensConfigurationTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStorageLensConfigurationTaggingInput struct {

	// The account ID of the requester.
	//
	// This member is required.
	AccountId *string

	// The ID of the Amazon S3 Storage Lens configuration.
	//
	// This member is required.
	ConfigId *string

	noSmithyDocumentSerde
}

func (in *GetStorageLensConfigurationTaggingInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetStorageLensConfigurationTaggingOutput struct {

	// The tags of S3 Storage Lens configuration requested.
	Tags []types.StorageLensTag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStorageLensConfigurationTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetStorageLensConfigurationTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetStorageLensConfigurationTagging{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetStorageLensConfigurationTagging"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetStorageLensConfigurationTaggingMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetStorageLensConfigurationTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStorageLensConfigurationTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetStorageLensConfigurationTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetStorageLensConfigurationTaggingMiddleware struct {
}

func (*endpointPrefix_opGetStorageLensConfigurationTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetStorageLensConfigurationTaggingMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetStorageLensConfigurationTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetStorageLensConfigurationTaggingMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetStorageLensConfigurationTaggingMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetStorageLensConfigurationTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetStorageLensConfigurationTagging",
	}
}

func copyGetStorageLensConfigurationTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetStorageLensConfigurationTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetStorageLensConfigurationTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetStorageLensConfigurationTaggingInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetStorageLensConfigurationTaggingAccountID(input interface{}, v string) error {
	in := input.(*GetStorageLensConfigurationTaggingInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetStorageLensConfigurationTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetStorageLensConfigurationTaggingInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
