// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists one or more Amazon Kendra experiences. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see [Building a search experience with no code].
//
// [Building a search experience with no code]: https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html
func (c *Client) ListExperiences(ctx context.Context, params *ListExperiencesInput, optFns ...func(*Options)) (*ListExperiencesOutput, error) {
	if params == nil {
		params = &ListExperiencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperiences", params, optFns, c.addOperationListExperiencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperiencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperiencesInput struct {

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// The maximum number of returned Amazon Kendra experiences.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of Amazon Kendra experiences.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExperiencesOutput struct {

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in a later request to retrieve the next set of Amazon Kendra experiences.
	NextToken *string

	// An array of summary information for one or more Amazon Kendra experiences.
	SummaryItems []types.ExperiencesSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExperiencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExperiences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExperiences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExperiences"); err != nil {
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
	if err = addOpListExperiencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperiences(options.Region), middleware.Before); err != nil {
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

// ListExperiencesAPIClient is a client that implements the ListExperiences
// operation.
type ListExperiencesAPIClient interface {
	ListExperiences(context.Context, *ListExperiencesInput, ...func(*Options)) (*ListExperiencesOutput, error)
}

var _ ListExperiencesAPIClient = (*Client)(nil)

// ListExperiencesPaginatorOptions is the paginator options for ListExperiences
type ListExperiencesPaginatorOptions struct {
	// The maximum number of returned Amazon Kendra experiences.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExperiencesPaginator is a paginator for ListExperiences
type ListExperiencesPaginator struct {
	options   ListExperiencesPaginatorOptions
	client    ListExperiencesAPIClient
	params    *ListExperiencesInput
	nextToken *string
	firstPage bool
}

// NewListExperiencesPaginator returns a new ListExperiencesPaginator
func NewListExperiencesPaginator(client ListExperiencesAPIClient, params *ListExperiencesInput, optFns ...func(*ListExperiencesPaginatorOptions)) *ListExperiencesPaginator {
	if params == nil {
		params = &ListExperiencesInput{}
	}

	options := ListExperiencesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExperiencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExperiencesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExperiences page.
func (p *ListExperiencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExperiencesOutput, error) {
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

	result, err := p.client.ListExperiences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExperiences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExperiences",
	}
}
