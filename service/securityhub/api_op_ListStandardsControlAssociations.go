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

//	Specifies whether a control is currently enabled or disabled in each enabled
//
// standard in the calling account.
func (c *Client) ListStandardsControlAssociations(ctx context.Context, params *ListStandardsControlAssociationsInput, optFns ...func(*Options)) (*ListStandardsControlAssociationsOutput, error) {
	if params == nil {
		params = &ListStandardsControlAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStandardsControlAssociations", params, optFns, c.addOperationListStandardsControlAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStandardsControlAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStandardsControlAssociationsInput struct {

	//  The identifier of the control (identified with SecurityControlId ,
	// SecurityControlArn , or a mix of both parameters) that you want to determine the
	// enablement status of in each enabled standard.
	//
	// This member is required.
	SecurityControlId *string

	//  An optional parameter that limits the total results of the API response to the
	// specified number. If this parameter isn't provided in the request, the results
	// include the first 25 standard and control associations. The results also include
	// a NextToken parameter that you can use in a subsequent API call to get the next
	// 25 associations. This repeats until all associations for the specified control
	// are returned. The number of results is limited by the number of supported
	// Security Hub standards that you've enabled in the calling account.
	MaxResults *int32

	//  Optional pagination parameter.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStandardsControlAssociationsOutput struct {

	//  An array that provides the enablement status and other details for each
	// security control that applies to each enabled standard.
	//
	// This member is required.
	StandardsControlAssociationSummaries []types.StandardsControlAssociationSummary

	//  A pagination parameter that's included in the response only if it was included
	// in the request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStandardsControlAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStandardsControlAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStandardsControlAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStandardsControlAssociations"); err != nil {
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
	if err = addOpListStandardsControlAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStandardsControlAssociations(options.Region), middleware.Before); err != nil {
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

// ListStandardsControlAssociationsAPIClient is a client that implements the
// ListStandardsControlAssociations operation.
type ListStandardsControlAssociationsAPIClient interface {
	ListStandardsControlAssociations(context.Context, *ListStandardsControlAssociationsInput, ...func(*Options)) (*ListStandardsControlAssociationsOutput, error)
}

var _ ListStandardsControlAssociationsAPIClient = (*Client)(nil)

// ListStandardsControlAssociationsPaginatorOptions is the paginator options for
// ListStandardsControlAssociations
type ListStandardsControlAssociationsPaginatorOptions struct {
	//  An optional parameter that limits the total results of the API response to the
	// specified number. If this parameter isn't provided in the request, the results
	// include the first 25 standard and control associations. The results also include
	// a NextToken parameter that you can use in a subsequent API call to get the next
	// 25 associations. This repeats until all associations for the specified control
	// are returned. The number of results is limited by the number of supported
	// Security Hub standards that you've enabled in the calling account.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStandardsControlAssociationsPaginator is a paginator for
// ListStandardsControlAssociations
type ListStandardsControlAssociationsPaginator struct {
	options   ListStandardsControlAssociationsPaginatorOptions
	client    ListStandardsControlAssociationsAPIClient
	params    *ListStandardsControlAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListStandardsControlAssociationsPaginator returns a new
// ListStandardsControlAssociationsPaginator
func NewListStandardsControlAssociationsPaginator(client ListStandardsControlAssociationsAPIClient, params *ListStandardsControlAssociationsInput, optFns ...func(*ListStandardsControlAssociationsPaginatorOptions)) *ListStandardsControlAssociationsPaginator {
	if params == nil {
		params = &ListStandardsControlAssociationsInput{}
	}

	options := ListStandardsControlAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStandardsControlAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStandardsControlAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStandardsControlAssociations page.
func (p *ListStandardsControlAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStandardsControlAssociationsOutput, error) {
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

	result, err := p.client.ListStandardsControlAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStandardsControlAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStandardsControlAssociations",
	}
}
