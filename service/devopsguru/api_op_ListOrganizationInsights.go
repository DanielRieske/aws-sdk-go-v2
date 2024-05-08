// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of insights associated with the account or OU Id.
func (c *Client) ListOrganizationInsights(ctx context.Context, params *ListOrganizationInsightsInput, optFns ...func(*Options)) (*ListOrganizationInsightsOutput, error) {
	if params == nil {
		params = &ListOrganizationInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationInsights", params, optFns, c.addOperationListOrganizationInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationInsightsInput struct {

	//  A filter used by ListInsights to specify which insights to return.
	//
	// This member is required.
	StatusFilter *types.ListInsightsStatusFilter

	// The ID of the Amazon Web Services account.
	AccountIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The ID of the organizational unit.
	OrganizationalUnitIds []string

	noSmithyDocumentSerde
}

type ListOrganizationInsightsOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// An integer that specifies the number of open proactive insights in your Amazon
	// Web Services account.
	ProactiveInsights []types.ProactiveOrganizationInsightSummary

	// An integer that specifies the number of open reactive insights in your Amazon
	// Web Services account.
	ReactiveInsights []types.ReactiveOrganizationInsightSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationInsights{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOrganizationInsights"); err != nil {
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
	if err = addOpListOrganizationInsightsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationInsights(options.Region), middleware.Before); err != nil {
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

// ListOrganizationInsightsAPIClient is a client that implements the
// ListOrganizationInsights operation.
type ListOrganizationInsightsAPIClient interface {
	ListOrganizationInsights(context.Context, *ListOrganizationInsightsInput, ...func(*Options)) (*ListOrganizationInsightsOutput, error)
}

var _ ListOrganizationInsightsAPIClient = (*Client)(nil)

// ListOrganizationInsightsPaginatorOptions is the paginator options for
// ListOrganizationInsights
type ListOrganizationInsightsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationInsightsPaginator is a paginator for ListOrganizationInsights
type ListOrganizationInsightsPaginator struct {
	options   ListOrganizationInsightsPaginatorOptions
	client    ListOrganizationInsightsAPIClient
	params    *ListOrganizationInsightsInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationInsightsPaginator returns a new
// ListOrganizationInsightsPaginator
func NewListOrganizationInsightsPaginator(client ListOrganizationInsightsAPIClient, params *ListOrganizationInsightsInput, optFns ...func(*ListOrganizationInsightsPaginatorOptions)) *ListOrganizationInsightsPaginator {
	if params == nil {
		params = &ListOrganizationInsightsInput{}
	}

	options := ListOrganizationInsightsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationInsightsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationInsightsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationInsights page.
func (p *ListOrganizationInsightsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationInsightsOutput, error) {
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

	result, err := p.client.ListOrganizationInsights(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrganizationInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOrganizationInsights",
	}
}
