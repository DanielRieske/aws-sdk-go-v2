// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides one or more conformance packs deployment status.
//
// If there are no conformance packs then you will see an empty result.
func (c *Client) DescribeConformancePackStatus(ctx context.Context, params *DescribeConformancePackStatusInput, optFns ...func(*Options)) (*DescribeConformancePackStatusOutput, error) {
	if params == nil {
		params = &DescribeConformancePackStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConformancePackStatus", params, optFns, c.addOperationDescribeConformancePackStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConformancePackStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConformancePackStatusInput struct {

	// Comma-separated list of conformance pack names.
	ConformancePackNames []string

	// The maximum number of conformance packs status returned on each page.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConformancePackStatusOutput struct {

	// A list of ConformancePackStatusDetail objects.
	ConformancePackStatusDetails []types.ConformancePackStatusDetail

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConformancePackStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConformancePackStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConformancePackStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConformancePackStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConformancePackStatus(options.Region), middleware.Before); err != nil {
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

// DescribeConformancePackStatusAPIClient is a client that implements the
// DescribeConformancePackStatus operation.
type DescribeConformancePackStatusAPIClient interface {
	DescribeConformancePackStatus(context.Context, *DescribeConformancePackStatusInput, ...func(*Options)) (*DescribeConformancePackStatusOutput, error)
}

var _ DescribeConformancePackStatusAPIClient = (*Client)(nil)

// DescribeConformancePackStatusPaginatorOptions is the paginator options for
// DescribeConformancePackStatus
type DescribeConformancePackStatusPaginatorOptions struct {
	// The maximum number of conformance packs status returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConformancePackStatusPaginator is a paginator for
// DescribeConformancePackStatus
type DescribeConformancePackStatusPaginator struct {
	options   DescribeConformancePackStatusPaginatorOptions
	client    DescribeConformancePackStatusAPIClient
	params    *DescribeConformancePackStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeConformancePackStatusPaginator returns a new
// DescribeConformancePackStatusPaginator
func NewDescribeConformancePackStatusPaginator(client DescribeConformancePackStatusAPIClient, params *DescribeConformancePackStatusInput, optFns ...func(*DescribeConformancePackStatusPaginatorOptions)) *DescribeConformancePackStatusPaginator {
	if params == nil {
		params = &DescribeConformancePackStatusInput{}
	}

	options := DescribeConformancePackStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConformancePackStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConformancePackStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConformancePackStatus page.
func (p *DescribeConformancePackStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConformancePackStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConformancePackStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConformancePackStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConformancePackStatus",
	}
}
