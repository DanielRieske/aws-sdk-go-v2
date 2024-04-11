// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists cloudwatch alarm template groups.
func (c *Client) ListCloudWatchAlarmTemplateGroups(ctx context.Context, params *ListCloudWatchAlarmTemplateGroupsInput, optFns ...func(*Options)) (*ListCloudWatchAlarmTemplateGroupsOutput, error) {
	if params == nil {
		params = &ListCloudWatchAlarmTemplateGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCloudWatchAlarmTemplateGroups", params, optFns, c.addOperationListCloudWatchAlarmTemplateGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCloudWatchAlarmTemplateGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListCloudWatchAlarmTemplateGroupsRequest
type ListCloudWatchAlarmTemplateGroupsInput struct {

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// A token used to retrieve the next set of results in paginated list responses.
	NextToken *string

	// Represents the scope of a resource, with options for all scopes, AWS provided
	// resources, or local resources.
	Scope *string

	// A signal map's identifier. Can be either be its id or current name.
	SignalMapIdentifier *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListCloudWatchAlarmTemplateGroupsResponse
type ListCloudWatchAlarmTemplateGroupsOutput struct {

	// Placeholder documentation for __listOfCloudWatchAlarmTemplateGroupSummary
	CloudWatchAlarmTemplateGroups []types.CloudWatchAlarmTemplateGroupSummary

	// A token used to retrieve the next set of results in paginated list responses.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCloudWatchAlarmTemplateGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCloudWatchAlarmTemplateGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCloudWatchAlarmTemplateGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCloudWatchAlarmTemplateGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCloudWatchAlarmTemplateGroups(options.Region), middleware.Before); err != nil {
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

// ListCloudWatchAlarmTemplateGroupsAPIClient is a client that implements the
// ListCloudWatchAlarmTemplateGroups operation.
type ListCloudWatchAlarmTemplateGroupsAPIClient interface {
	ListCloudWatchAlarmTemplateGroups(context.Context, *ListCloudWatchAlarmTemplateGroupsInput, ...func(*Options)) (*ListCloudWatchAlarmTemplateGroupsOutput, error)
}

var _ ListCloudWatchAlarmTemplateGroupsAPIClient = (*Client)(nil)

// ListCloudWatchAlarmTemplateGroupsPaginatorOptions is the paginator options for
// ListCloudWatchAlarmTemplateGroups
type ListCloudWatchAlarmTemplateGroupsPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCloudWatchAlarmTemplateGroupsPaginator is a paginator for
// ListCloudWatchAlarmTemplateGroups
type ListCloudWatchAlarmTemplateGroupsPaginator struct {
	options   ListCloudWatchAlarmTemplateGroupsPaginatorOptions
	client    ListCloudWatchAlarmTemplateGroupsAPIClient
	params    *ListCloudWatchAlarmTemplateGroupsInput
	nextToken *string
	firstPage bool
}

// NewListCloudWatchAlarmTemplateGroupsPaginator returns a new
// ListCloudWatchAlarmTemplateGroupsPaginator
func NewListCloudWatchAlarmTemplateGroupsPaginator(client ListCloudWatchAlarmTemplateGroupsAPIClient, params *ListCloudWatchAlarmTemplateGroupsInput, optFns ...func(*ListCloudWatchAlarmTemplateGroupsPaginatorOptions)) *ListCloudWatchAlarmTemplateGroupsPaginator {
	if params == nil {
		params = &ListCloudWatchAlarmTemplateGroupsInput{}
	}

	options := ListCloudWatchAlarmTemplateGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCloudWatchAlarmTemplateGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCloudWatchAlarmTemplateGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCloudWatchAlarmTemplateGroups page.
func (p *ListCloudWatchAlarmTemplateGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCloudWatchAlarmTemplateGroupsOutput, error) {
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

	result, err := p.client.ListCloudWatchAlarmTemplateGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCloudWatchAlarmTemplateGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCloudWatchAlarmTemplateGroups",
	}
}
