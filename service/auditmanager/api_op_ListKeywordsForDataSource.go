// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of keywords that are pre-mapped to the specified control data
//
// source.
func (c *Client) ListKeywordsForDataSource(ctx context.Context, params *ListKeywordsForDataSourceInput, optFns ...func(*Options)) (*ListKeywordsForDataSourceOutput, error) {
	if params == nil {
		params = &ListKeywordsForDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeywordsForDataSource", params, optFns, c.addOperationListKeywordsForDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeywordsForDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeywordsForDataSourceInput struct {

	//  The control mapping data source that the keywords apply to.
	//
	// This member is required.
	Source types.SourceType

	//  Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKeywordsForDataSourceOutput struct {

	//  The list of keywords for the event mapping source.
	Keywords []string

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeywordsForDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKeywordsForDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKeywordsForDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKeywordsForDataSource"); err != nil {
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
	if err = addOpListKeywordsForDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeywordsForDataSource(options.Region), middleware.Before); err != nil {
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

// ListKeywordsForDataSourceAPIClient is a client that implements the
// ListKeywordsForDataSource operation.
type ListKeywordsForDataSourceAPIClient interface {
	ListKeywordsForDataSource(context.Context, *ListKeywordsForDataSourceInput, ...func(*Options)) (*ListKeywordsForDataSourceOutput, error)
}

var _ ListKeywordsForDataSourceAPIClient = (*Client)(nil)

// ListKeywordsForDataSourcePaginatorOptions is the paginator options for
// ListKeywordsForDataSource
type ListKeywordsForDataSourcePaginatorOptions struct {
	//  Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeywordsForDataSourcePaginator is a paginator for ListKeywordsForDataSource
type ListKeywordsForDataSourcePaginator struct {
	options   ListKeywordsForDataSourcePaginatorOptions
	client    ListKeywordsForDataSourceAPIClient
	params    *ListKeywordsForDataSourceInput
	nextToken *string
	firstPage bool
}

// NewListKeywordsForDataSourcePaginator returns a new
// ListKeywordsForDataSourcePaginator
func NewListKeywordsForDataSourcePaginator(client ListKeywordsForDataSourceAPIClient, params *ListKeywordsForDataSourceInput, optFns ...func(*ListKeywordsForDataSourcePaginatorOptions)) *ListKeywordsForDataSourcePaginator {
	if params == nil {
		params = &ListKeywordsForDataSourceInput{}
	}

	options := ListKeywordsForDataSourcePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeywordsForDataSourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeywordsForDataSourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeywordsForDataSource page.
func (p *ListKeywordsForDataSourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeywordsForDataSourceOutput, error) {
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

	result, err := p.client.ListKeywordsForDataSource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKeywordsForDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKeywordsForDataSource",
	}
}
