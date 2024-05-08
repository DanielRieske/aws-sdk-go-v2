// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all entities in a workspace.
func (c *Client) ListEntities(ctx context.Context, params *ListEntitiesInput, optFns ...func(*Options)) (*ListEntitiesOutput, error) {
	if params == nil {
		params = &ListEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntities", params, optFns, c.addOperationListEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesInput struct {

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// A list of objects that filter the request.
	//
	// Only one object is accepted as a valid input.
	Filters []types.ListEntitiesFilter

	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntitiesOutput struct {

	// A list of objects that contain information about the entities.
	EntitySummaries []types.EntitySummary

	// The string that specifies the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEntities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEntities"); err != nil {
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
	if err = addEndpointPrefix_opListEntitiesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntities(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListEntitiesMiddleware struct {
}

func (*endpointPrefix_opListEntitiesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListEntitiesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListEntitiesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListEntitiesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListEntitiesAPIClient is a client that implements the ListEntities operation.
type ListEntitiesAPIClient interface {
	ListEntities(context.Context, *ListEntitiesInput, ...func(*Options)) (*ListEntitiesOutput, error)
}

var _ ListEntitiesAPIClient = (*Client)(nil)

// ListEntitiesPaginatorOptions is the paginator options for ListEntities
type ListEntitiesPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntitiesPaginator is a paginator for ListEntities
type ListEntitiesPaginator struct {
	options   ListEntitiesPaginatorOptions
	client    ListEntitiesAPIClient
	params    *ListEntitiesInput
	nextToken *string
	firstPage bool
}

// NewListEntitiesPaginator returns a new ListEntitiesPaginator
func NewListEntitiesPaginator(client ListEntitiesAPIClient, params *ListEntitiesInput, optFns ...func(*ListEntitiesPaginatorOptions)) *ListEntitiesPaginator {
	if params == nil {
		params = &ListEntitiesInput{}
	}

	options := ListEntitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEntitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEntities page.
func (p *ListEntitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntitiesOutput, error) {
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

	result, err := p.client.ListEntities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEntities",
	}
}
