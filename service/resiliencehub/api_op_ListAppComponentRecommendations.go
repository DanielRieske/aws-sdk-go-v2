// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the recommendations for an Resilience Hub Application Component.
func (c *Client) ListAppComponentRecommendations(ctx context.Context, params *ListAppComponentRecommendationsInput, optFns ...func(*Options)) (*ListAppComponentRecommendationsOutput, error) {
	if params == nil {
		params = &ListAppComponentRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppComponentRecommendations", params, optFns, c.addOperationListAppComponentRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppComponentRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppComponentRecommendationsInput struct {

	// Amazon Resource Name (ARN) of the assessment. The format for this ARN is: arn:
	// partition :resiliencehub: region : account :app-assessment/ app-id . For more
	// information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	AssessmentArn *string

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppComponentRecommendationsOutput struct {

	// The recommendations for an Resilience Hub Application Component, returned as an
	// object. This object contains the names of the Application Components,
	// configuration recommendations, and recommendation statuses.
	//
	// This member is required.
	ComponentRecommendations []types.ComponentRecommendation

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppComponentRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppComponentRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppComponentRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppComponentRecommendations"); err != nil {
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
	if err = addOpListAppComponentRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppComponentRecommendations(options.Region), middleware.Before); err != nil {
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

// ListAppComponentRecommendationsAPIClient is a client that implements the
// ListAppComponentRecommendations operation.
type ListAppComponentRecommendationsAPIClient interface {
	ListAppComponentRecommendations(context.Context, *ListAppComponentRecommendationsInput, ...func(*Options)) (*ListAppComponentRecommendationsOutput, error)
}

var _ ListAppComponentRecommendationsAPIClient = (*Client)(nil)

// ListAppComponentRecommendationsPaginatorOptions is the paginator options for
// ListAppComponentRecommendations
type ListAppComponentRecommendationsPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppComponentRecommendationsPaginator is a paginator for
// ListAppComponentRecommendations
type ListAppComponentRecommendationsPaginator struct {
	options   ListAppComponentRecommendationsPaginatorOptions
	client    ListAppComponentRecommendationsAPIClient
	params    *ListAppComponentRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListAppComponentRecommendationsPaginator returns a new
// ListAppComponentRecommendationsPaginator
func NewListAppComponentRecommendationsPaginator(client ListAppComponentRecommendationsAPIClient, params *ListAppComponentRecommendationsInput, optFns ...func(*ListAppComponentRecommendationsPaginatorOptions)) *ListAppComponentRecommendationsPaginator {
	if params == nil {
		params = &ListAppComponentRecommendationsInput{}
	}

	options := ListAppComponentRecommendationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppComponentRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppComponentRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppComponentRecommendations page.
func (p *ListAppComponentRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppComponentRecommendationsOutput, error) {
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

	result, err := p.client.ListAppComponentRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppComponentRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppComponentRecommendations",
	}
}
