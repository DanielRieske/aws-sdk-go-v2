// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all attributes that are associated with an object.
func (c *Client) ListObjectAttributes(ctx context.Context, params *ListObjectAttributesInput, optFns ...func(*Options)) (*ListObjectAttributesOutput, error) {
	if params == nil {
		params = &ListObjectAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectAttributes", params, optFns, c.addOperationListObjectAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectAttributesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory where the object
	// resides. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the object whose attributes will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// Used to filter the list of object attributes that are associated with a certain
	// facet.
	FacetFilter *types.SchemaFacet

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListObjectAttributesOutput struct {

	// Attributes map that is associated with the object. AttributeArn is the key, and
	// attribute value is the value.
	Attributes []types.AttributeKeyAndValue

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListObjectAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListObjectAttributes"); err != nil {
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
	if err = addOpListObjectAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectAttributes(options.Region), middleware.Before); err != nil {
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

// ListObjectAttributesAPIClient is a client that implements the
// ListObjectAttributes operation.
type ListObjectAttributesAPIClient interface {
	ListObjectAttributes(context.Context, *ListObjectAttributesInput, ...func(*Options)) (*ListObjectAttributesOutput, error)
}

var _ ListObjectAttributesAPIClient = (*Client)(nil)

// ListObjectAttributesPaginatorOptions is the paginator options for
// ListObjectAttributes
type ListObjectAttributesPaginatorOptions struct {
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListObjectAttributesPaginator is a paginator for ListObjectAttributes
type ListObjectAttributesPaginator struct {
	options   ListObjectAttributesPaginatorOptions
	client    ListObjectAttributesAPIClient
	params    *ListObjectAttributesInput
	nextToken *string
	firstPage bool
}

// NewListObjectAttributesPaginator returns a new ListObjectAttributesPaginator
func NewListObjectAttributesPaginator(client ListObjectAttributesAPIClient, params *ListObjectAttributesInput, optFns ...func(*ListObjectAttributesPaginatorOptions)) *ListObjectAttributesPaginator {
	if params == nil {
		params = &ListObjectAttributesInput{}
	}

	options := ListObjectAttributesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListObjectAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListObjectAttributesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListObjectAttributes page.
func (p *ListObjectAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListObjectAttributesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListObjectAttributes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListObjectAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListObjectAttributes",
	}
}
