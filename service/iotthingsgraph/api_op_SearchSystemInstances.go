// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for system instances in the user's account.
//
// Deprecated: since: 2022-08-30
func (c *Client) SearchSystemInstances(ctx context.Context, params *SearchSystemInstancesInput, optFns ...func(*Options)) (*SearchSystemInstancesOutput, error) {
	if params == nil {
		params = &SearchSystemInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSystemInstances", params, optFns, c.addOperationSearchSystemInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSystemInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSystemInstancesInput struct {

	// Optional filter to apply to the search. Valid filters are SYSTEM_TEMPLATE_ID ,
	// STATUS , and GREENGRASS_GROUP_NAME .
	//
	// Multiple filters function as OR criteria in the query. Multiple values passed
	// inside the filter function as AND criteria.
	Filters []types.SystemInstanceFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchSystemInstancesOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary data abour the system instances in the
	// result set.
	Summaries []types.SystemInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchSystemInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSystemInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSystemInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchSystemInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSystemInstances(options.Region), middleware.Before); err != nil {
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

// SearchSystemInstancesAPIClient is a client that implements the
// SearchSystemInstances operation.
type SearchSystemInstancesAPIClient interface {
	SearchSystemInstances(context.Context, *SearchSystemInstancesInput, ...func(*Options)) (*SearchSystemInstancesOutput, error)
}

var _ SearchSystemInstancesAPIClient = (*Client)(nil)

// SearchSystemInstancesPaginatorOptions is the paginator options for
// SearchSystemInstances
type SearchSystemInstancesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchSystemInstancesPaginator is a paginator for SearchSystemInstances
type SearchSystemInstancesPaginator struct {
	options   SearchSystemInstancesPaginatorOptions
	client    SearchSystemInstancesAPIClient
	params    *SearchSystemInstancesInput
	nextToken *string
	firstPage bool
}

// NewSearchSystemInstancesPaginator returns a new SearchSystemInstancesPaginator
func NewSearchSystemInstancesPaginator(client SearchSystemInstancesAPIClient, params *SearchSystemInstancesInput, optFns ...func(*SearchSystemInstancesPaginatorOptions)) *SearchSystemInstancesPaginator {
	if params == nil {
		params = &SearchSystemInstancesInput{}
	}

	options := SearchSystemInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchSystemInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchSystemInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchSystemInstances page.
func (p *SearchSystemInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchSystemInstancesOutput, error) {
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

	result, err := p.client.SearchSystemInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchSystemInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchSystemInstances",
	}
}
