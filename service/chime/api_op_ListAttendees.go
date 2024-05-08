// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Lists the attendees for the specified Amazon Chime SDK meeting. For more
//
// information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime SDK Developer
// Guide.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [ListAttendees], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by ListAttendees in the Amazon Chime SDK Meetings Namespace
//
// [ListAttendees]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_ListAttendees.html
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func (c *Client) ListAttendees(ctx context.Context, params *ListAttendeesInput, optFns ...func(*Options)) (*ListAttendeesOutput, error) {
	if params == nil {
		params = &ListAttendeesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttendees", params, optFns, c.addOperationListAttendeesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttendeesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttendeesInput struct {

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttendeesOutput struct {

	// The Amazon Chime SDK attendee information.
	Attendees []types.Attendee

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttendeesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttendees{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttendees{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAttendees"); err != nil {
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
	if err = addOpListAttendeesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttendees(options.Region), middleware.Before); err != nil {
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

// ListAttendeesAPIClient is a client that implements the ListAttendees operation.
type ListAttendeesAPIClient interface {
	ListAttendees(context.Context, *ListAttendeesInput, ...func(*Options)) (*ListAttendeesOutput, error)
}

var _ ListAttendeesAPIClient = (*Client)(nil)

// ListAttendeesPaginatorOptions is the paginator options for ListAttendees
type ListAttendeesPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttendeesPaginator is a paginator for ListAttendees
type ListAttendeesPaginator struct {
	options   ListAttendeesPaginatorOptions
	client    ListAttendeesAPIClient
	params    *ListAttendeesInput
	nextToken *string
	firstPage bool
}

// NewListAttendeesPaginator returns a new ListAttendeesPaginator
func NewListAttendeesPaginator(client ListAttendeesAPIClient, params *ListAttendeesInput, optFns ...func(*ListAttendeesPaginatorOptions)) *ListAttendeesPaginator {
	if params == nil {
		params = &ListAttendeesInput{}
	}

	options := ListAttendeesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttendeesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttendeesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttendees page.
func (p *ListAttendeesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttendeesOutput, error) {
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

	result, err := p.client.ListAttendees(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttendees(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAttendees",
	}
}
