// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of datashares when the account identifier being called is a
// producer account identifier.
func (c *Client) DescribeDataSharesForProducer(ctx context.Context, params *DescribeDataSharesForProducerInput, optFns ...func(*Options)) (*DescribeDataSharesForProducerOutput, error) {
	if params == nil {
		params = &DescribeDataSharesForProducerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataSharesForProducer", params, optFns, c.addOperationDescribeDataSharesForProducerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataSharesForProducerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSharesForProducerInput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDataSharesForProducerrequest exceed the value specified in
	// MaxRecords , Amazon Web Services returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) of the producer namespace that returns in the
	// list of datashares.
	ProducerArn *string

	// An identifier giving the status of a datashare in the producer. If this field
	// is specified, Amazon Redshift returns the list of datashares that have the
	// specified status.
	Status types.DataShareStatusForProducer

	noSmithyDocumentSerde
}

type DescribeDataSharesForProducerOutput struct {

	// Shows the results of datashares available for producers.
	DataShares []types.DataShare

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDataSharesForProducerrequest exceed the value specified in
	// MaxRecords , Amazon Web Services returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataSharesForProducerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDataSharesForProducer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDataSharesForProducer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDataSharesForProducer"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSharesForProducer(options.Region), middleware.Before); err != nil {
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

// DescribeDataSharesForProducerAPIClient is a client that implements the
// DescribeDataSharesForProducer operation.
type DescribeDataSharesForProducerAPIClient interface {
	DescribeDataSharesForProducer(context.Context, *DescribeDataSharesForProducerInput, ...func(*Options)) (*DescribeDataSharesForProducerOutput, error)
}

var _ DescribeDataSharesForProducerAPIClient = (*Client)(nil)

// DescribeDataSharesForProducerPaginatorOptions is the paginator options for
// DescribeDataSharesForProducer
type DescribeDataSharesForProducerPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDataSharesForProducerPaginator is a paginator for
// DescribeDataSharesForProducer
type DescribeDataSharesForProducerPaginator struct {
	options   DescribeDataSharesForProducerPaginatorOptions
	client    DescribeDataSharesForProducerAPIClient
	params    *DescribeDataSharesForProducerInput
	nextToken *string
	firstPage bool
}

// NewDescribeDataSharesForProducerPaginator returns a new
// DescribeDataSharesForProducerPaginator
func NewDescribeDataSharesForProducerPaginator(client DescribeDataSharesForProducerAPIClient, params *DescribeDataSharesForProducerInput, optFns ...func(*DescribeDataSharesForProducerPaginatorOptions)) *DescribeDataSharesForProducerPaginator {
	if params == nil {
		params = &DescribeDataSharesForProducerInput{}
	}

	options := DescribeDataSharesForProducerPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDataSharesForProducerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDataSharesForProducerPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDataSharesForProducer page.
func (p *DescribeDataSharesForProducerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDataSharesForProducerOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDataSharesForProducer(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDataSharesForProducer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDataSharesForProducer",
	}
}
