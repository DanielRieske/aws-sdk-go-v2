// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about your channel's schedule.
func (c *Client) GetChannelSchedule(ctx context.Context, params *GetChannelScheduleInput, optFns ...func(*Options)) (*GetChannelScheduleOutput, error) {
	if params == nil {
		params = &GetChannelScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChannelSchedule", params, optFns, c.addOperationGetChannelScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChannelScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChannelScheduleInput struct {

	// The name of the channel associated with this Channel Schedule.
	//
	// This member is required.
	ChannelName *string

	// The single audience for GetChannelScheduleRequest.
	Audience *string

	// The duration in minutes of the channel schedule.
	DurationMinutes *string

	// The maximum number of channel schedules that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults channel
	// schedules, use the value of NextToken in the response to get the next page of
	// results.
	MaxResults *int32

	// (Optional) If the playback configuration has more than MaxResults channel
	// schedules, use NextToken to get the second and subsequent pages of results.
	//
	// For the first GetChannelScheduleRequest request, omit this value.
	//
	// For the second and subsequent requests, get the value of NextToken from the
	// previous response and specify that value for NextToken in the request.
	//
	// If the previous response didn't include a NextToken element, there are no more
	// channel schedules to get.
	NextToken *string

	noSmithyDocumentSerde
}

type GetChannelScheduleOutput struct {

	// A list of schedule entries for the channel.
	Items []types.ScheduleEntry

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChannelScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChannelSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChannelSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetChannelSchedule"); err != nil {
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
	if err = addOpGetChannelScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChannelSchedule(options.Region), middleware.Before); err != nil {
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

// GetChannelScheduleAPIClient is a client that implements the GetChannelSchedule
// operation.
type GetChannelScheduleAPIClient interface {
	GetChannelSchedule(context.Context, *GetChannelScheduleInput, ...func(*Options)) (*GetChannelScheduleOutput, error)
}

var _ GetChannelScheduleAPIClient = (*Client)(nil)

// GetChannelSchedulePaginatorOptions is the paginator options for
// GetChannelSchedule
type GetChannelSchedulePaginatorOptions struct {
	// The maximum number of channel schedules that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults channel
	// schedules, use the value of NextToken in the response to get the next page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetChannelSchedulePaginator is a paginator for GetChannelSchedule
type GetChannelSchedulePaginator struct {
	options   GetChannelSchedulePaginatorOptions
	client    GetChannelScheduleAPIClient
	params    *GetChannelScheduleInput
	nextToken *string
	firstPage bool
}

// NewGetChannelSchedulePaginator returns a new GetChannelSchedulePaginator
func NewGetChannelSchedulePaginator(client GetChannelScheduleAPIClient, params *GetChannelScheduleInput, optFns ...func(*GetChannelSchedulePaginatorOptions)) *GetChannelSchedulePaginator {
	if params == nil {
		params = &GetChannelScheduleInput{}
	}

	options := GetChannelSchedulePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetChannelSchedulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetChannelSchedulePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetChannelSchedule page.
func (p *GetChannelSchedulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetChannelScheduleOutput, error) {
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

	result, err := p.client.GetChannelSchedule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetChannelSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetChannelSchedule",
	}
}
