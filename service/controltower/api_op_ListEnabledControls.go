// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/controltower/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the controls enabled by Amazon Web Services Control Tower on the
// specified organizational unit and the accounts it contains. For usage examples,
// see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/control-api-examples-short.html
func (c *Client) ListEnabledControls(ctx context.Context, params *ListEnabledControlsInput, optFns ...func(*Options)) (*ListEnabledControlsOutput, error) {
	if params == nil {
		params = &ListEnabledControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnabledControls", params, optFns, c.addOperationListEnabledControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnabledControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnabledControlsInput struct {

	// The ARN of the organizational unit. For information on how to find the
	// targetIdentifier , see [the overview page].
	//
	// [the overview page]: https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html
	//
	// This member is required.
	TargetIdentifier *string

	// How many results to return per API call.
	MaxResults *int32

	// The token to continue the list from a previous API call with the same
	// parameters.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnabledControlsOutput struct {

	// Lists the controls enabled by Amazon Web Services Control Tower on the
	// specified organizational unit and the accounts it contains.
	//
	// This member is required.
	EnabledControls []types.EnabledControlSummary

	// Retrieves the next page of results. If the string is empty, the response is the
	// end of the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnabledControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnabledControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnabledControls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnabledControls"); err != nil {
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
	if err = addOpListEnabledControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnabledControls(options.Region), middleware.Before); err != nil {
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

// ListEnabledControlsAPIClient is a client that implements the
// ListEnabledControls operation.
type ListEnabledControlsAPIClient interface {
	ListEnabledControls(context.Context, *ListEnabledControlsInput, ...func(*Options)) (*ListEnabledControlsOutput, error)
}

var _ ListEnabledControlsAPIClient = (*Client)(nil)

// ListEnabledControlsPaginatorOptions is the paginator options for
// ListEnabledControls
type ListEnabledControlsPaginatorOptions struct {
	// How many results to return per API call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnabledControlsPaginator is a paginator for ListEnabledControls
type ListEnabledControlsPaginator struct {
	options   ListEnabledControlsPaginatorOptions
	client    ListEnabledControlsAPIClient
	params    *ListEnabledControlsInput
	nextToken *string
	firstPage bool
}

// NewListEnabledControlsPaginator returns a new ListEnabledControlsPaginator
func NewListEnabledControlsPaginator(client ListEnabledControlsAPIClient, params *ListEnabledControlsInput, optFns ...func(*ListEnabledControlsPaginatorOptions)) *ListEnabledControlsPaginator {
	if params == nil {
		params = &ListEnabledControlsInput{}
	}

	options := ListEnabledControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnabledControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnabledControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnabledControls page.
func (p *ListEnabledControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnabledControlsOutput, error) {
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

	result, err := p.client.ListEnabledControls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnabledControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnabledControls",
	}
}
