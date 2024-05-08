// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about playback key pairs. For more information, see [Setting Up Private Channels]
// in the Amazon IVS User Guide.
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func (c *Client) ListPlaybackKeyPairs(ctx context.Context, params *ListPlaybackKeyPairsInput, optFns ...func(*Options)) (*ListPlaybackKeyPairsOutput, error) {
	if params == nil {
		params = &ListPlaybackKeyPairsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlaybackKeyPairs", params, optFns, c.addOperationListPlaybackKeyPairsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlaybackKeyPairsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaybackKeyPairsInput struct {

	// Maximum number of key pairs to return. Default: your service quota or 100,
	// whichever is smaller.
	MaxResults *int32

	// The first key pair to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlaybackKeyPairsOutput struct {

	// List of key pairs.
	//
	// This member is required.
	KeyPairs []types.PlaybackKeyPairSummary

	// If there are more key pairs than maxResults , use nextToken in the request to
	// get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlaybackKeyPairsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlaybackKeyPairs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaybackKeyPairs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPlaybackKeyPairs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlaybackKeyPairs(options.Region), middleware.Before); err != nil {
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

// ListPlaybackKeyPairsAPIClient is a client that implements the
// ListPlaybackKeyPairs operation.
type ListPlaybackKeyPairsAPIClient interface {
	ListPlaybackKeyPairs(context.Context, *ListPlaybackKeyPairsInput, ...func(*Options)) (*ListPlaybackKeyPairsOutput, error)
}

var _ ListPlaybackKeyPairsAPIClient = (*Client)(nil)

// ListPlaybackKeyPairsPaginatorOptions is the paginator options for
// ListPlaybackKeyPairs
type ListPlaybackKeyPairsPaginatorOptions struct {
	// Maximum number of key pairs to return. Default: your service quota or 100,
	// whichever is smaller.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlaybackKeyPairsPaginator is a paginator for ListPlaybackKeyPairs
type ListPlaybackKeyPairsPaginator struct {
	options   ListPlaybackKeyPairsPaginatorOptions
	client    ListPlaybackKeyPairsAPIClient
	params    *ListPlaybackKeyPairsInput
	nextToken *string
	firstPage bool
}

// NewListPlaybackKeyPairsPaginator returns a new ListPlaybackKeyPairsPaginator
func NewListPlaybackKeyPairsPaginator(client ListPlaybackKeyPairsAPIClient, params *ListPlaybackKeyPairsInput, optFns ...func(*ListPlaybackKeyPairsPaginatorOptions)) *ListPlaybackKeyPairsPaginator {
	if params == nil {
		params = &ListPlaybackKeyPairsInput{}
	}

	options := ListPlaybackKeyPairsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlaybackKeyPairsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlaybackKeyPairsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlaybackKeyPairs page.
func (p *ListPlaybackKeyPairsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlaybackKeyPairsOutput, error) {
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

	result, err := p.client.ListPlaybackKeyPairs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlaybackKeyPairs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPlaybackKeyPairs",
	}
}
