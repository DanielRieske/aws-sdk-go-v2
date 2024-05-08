// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the subtasks for an Inference Recommender job.
//
// The supported subtasks are benchmarks, which evaluate the performance of your
// model on different instance types.
func (c *Client) ListInferenceRecommendationsJobSteps(ctx context.Context, params *ListInferenceRecommendationsJobStepsInput, optFns ...func(*Options)) (*ListInferenceRecommendationsJobStepsOutput, error) {
	if params == nil {
		params = &ListInferenceRecommendationsJobStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceRecommendationsJobSteps", params, optFns, c.addOperationListInferenceRecommendationsJobStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceRecommendationsJobStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceRecommendationsJobStepsInput struct {

	// The name for the Inference Recommender job.
	//
	// This member is required.
	JobName *string

	// The maximum number of results to return.
	MaxResults *int32

	// A token that you can specify to return more results from the list. Specify this
	// field if you have a token that was returned from a previous request.
	NextToken *string

	// A filter to return benchmarks of a specified status. If this field is left
	// empty, then all benchmarks are returned.
	Status types.RecommendationJobStatus

	// A filter to return details about the specified type of subtask.
	//
	// BENCHMARK : Evaluate the performance of your model on different instance types.
	StepType types.RecommendationStepType

	noSmithyDocumentSerde
}

type ListInferenceRecommendationsJobStepsOutput struct {

	// A token that you can specify in your next request to return more results from
	// the list.
	NextToken *string

	// A list of all subtask details in Inference Recommender.
	Steps []types.InferenceRecommendationsJobStep

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceRecommendationsJobStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInferenceRecommendationsJobSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInferenceRecommendationsJobSteps{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInferenceRecommendationsJobSteps"); err != nil {
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
	if err = addOpListInferenceRecommendationsJobStepsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceRecommendationsJobSteps(options.Region), middleware.Before); err != nil {
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

// ListInferenceRecommendationsJobStepsAPIClient is a client that implements the
// ListInferenceRecommendationsJobSteps operation.
type ListInferenceRecommendationsJobStepsAPIClient interface {
	ListInferenceRecommendationsJobSteps(context.Context, *ListInferenceRecommendationsJobStepsInput, ...func(*Options)) (*ListInferenceRecommendationsJobStepsOutput, error)
}

var _ ListInferenceRecommendationsJobStepsAPIClient = (*Client)(nil)

// ListInferenceRecommendationsJobStepsPaginatorOptions is the paginator options
// for ListInferenceRecommendationsJobSteps
type ListInferenceRecommendationsJobStepsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceRecommendationsJobStepsPaginator is a paginator for
// ListInferenceRecommendationsJobSteps
type ListInferenceRecommendationsJobStepsPaginator struct {
	options   ListInferenceRecommendationsJobStepsPaginatorOptions
	client    ListInferenceRecommendationsJobStepsAPIClient
	params    *ListInferenceRecommendationsJobStepsInput
	nextToken *string
	firstPage bool
}

// NewListInferenceRecommendationsJobStepsPaginator returns a new
// ListInferenceRecommendationsJobStepsPaginator
func NewListInferenceRecommendationsJobStepsPaginator(client ListInferenceRecommendationsJobStepsAPIClient, params *ListInferenceRecommendationsJobStepsInput, optFns ...func(*ListInferenceRecommendationsJobStepsPaginatorOptions)) *ListInferenceRecommendationsJobStepsPaginator {
	if params == nil {
		params = &ListInferenceRecommendationsJobStepsInput{}
	}

	options := ListInferenceRecommendationsJobStepsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceRecommendationsJobStepsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceRecommendationsJobStepsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceRecommendationsJobSteps page.
func (p *ListInferenceRecommendationsJobStepsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceRecommendationsJobStepsOutput, error) {
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

	result, err := p.client.ListInferenceRecommendationsJobSteps(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInferenceRecommendationsJobSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInferenceRecommendationsJobSteps",
	}
}
