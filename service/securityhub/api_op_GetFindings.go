// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of findings that match the specified criteria.
//
// If finding aggregation is enabled, then when you call GetFindings from the
// aggregation Region, the results include all of the matching findings from both
// the aggregation Region and the linked Regions.
func (c *Client) GetFindings(ctx context.Context, params *GetFindingsInput, optFns ...func(*Options)) (*GetFindingsOutput, error) {
	if params == nil {
		params = &GetFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindings", params, optFns, c.addOperationGetFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsInput struct {

	// The finding attributes used to define a condition to filter the returned
	// findings.
	//
	// You can filter by up to 10 finding attributes. For each attribute, you can
	// provide up to 20 filter values.
	//
	// Note that in the available filter fields, WorkflowState is deprecated. To
	// search for a finding based on its workflow status, use WorkflowStatus .
	Filters *types.AwsSecurityFindingFilters

	// The maximum number of findings to return.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the GetFindings
	// operation, set the value of this parameter to NULL .
	//
	// For subsequent calls to the operation, to continue listing data, set the value
	// of this parameter to the value returned from the previous response.
	NextToken *string

	// The finding attributes used to sort the list of returned findings.
	SortCriteria []types.SortCriterion

	noSmithyDocumentSerde
}

type GetFindingsOutput struct {

	// The findings that matched the filters specified in the request.
	//
	// This member is required.
	Findings []types.AwsSecurityFinding

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFindings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindings(options.Region), middleware.Before); err != nil {
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

// GetFindingsAPIClient is a client that implements the GetFindings operation.
type GetFindingsAPIClient interface {
	GetFindings(context.Context, *GetFindingsInput, ...func(*Options)) (*GetFindingsOutput, error)
}

var _ GetFindingsAPIClient = (*Client)(nil)

// GetFindingsPaginatorOptions is the paginator options for GetFindings
type GetFindingsPaginatorOptions struct {
	// The maximum number of findings to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFindingsPaginator is a paginator for GetFindings
type GetFindingsPaginator struct {
	options   GetFindingsPaginatorOptions
	client    GetFindingsAPIClient
	params    *GetFindingsInput
	nextToken *string
	firstPage bool
}

// NewGetFindingsPaginator returns a new GetFindingsPaginator
func NewGetFindingsPaginator(client GetFindingsAPIClient, params *GetFindingsInput, optFns ...func(*GetFindingsPaginatorOptions)) *GetFindingsPaginator {
	if params == nil {
		params = &GetFindingsInput{}
	}

	options := GetFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetFindings page.
func (p *GetFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFindingsOutput, error) {
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

	result, err := p.client.GetFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFindings",
	}
}
