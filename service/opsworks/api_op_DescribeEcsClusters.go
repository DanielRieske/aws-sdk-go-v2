// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes Amazon ECS clusters that are registered with a stack. If you specify
// only a stack ID, you can use the MaxResults and NextToken parameters to
// paginate the response. However, AWS OpsWorks Stacks currently supports only one
// cluster per layer, so the result set has a maximum of one element.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack or an attached policy that explicitly
// grants permission. For more information about user permissions, see [Managing User Permissions].
//
// This call accepts only one resource-identifying parameter.
//
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) DescribeEcsClusters(ctx context.Context, params *DescribeEcsClustersInput, optFns ...func(*Options)) (*DescribeEcsClustersOutput, error) {
	if params == nil {
		params = &DescribeEcsClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEcsClusters", params, optFns, c.addOperationDescribeEcsClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEcsClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEcsClustersInput struct {

	// A list of ARNs, one for each cluster to be described.
	EcsClusterArns []string

	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call DescribeEcsClusters again and assign that token
	// to the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null .
	NextToken *string

	// A stack ID. DescribeEcsClusters returns a description of the cluster that is
	// registered with the stack.
	StackId *string

	noSmithyDocumentSerde
}

// Contains the response to a DescribeEcsClusters request.
type DescribeEcsClustersOutput struct {

	// A list of EcsCluster objects containing the cluster descriptions.
	EcsClusters []types.EcsCluster

	// If a paginated request does not return all of the remaining results, this
	// parameter is set to a token that you can assign to the request object's
	// NextToken parameter to retrieve the next set of results. If the previous
	// paginated request returned all of the remaining results, this parameter is set
	// to null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEcsClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEcsClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEcsClusters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEcsClusters"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEcsClusters(options.Region), middleware.Before); err != nil {
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

// DescribeEcsClustersAPIClient is a client that implements the
// DescribeEcsClusters operation.
type DescribeEcsClustersAPIClient interface {
	DescribeEcsClusters(context.Context, *DescribeEcsClustersInput, ...func(*Options)) (*DescribeEcsClustersOutput, error)
}

var _ DescribeEcsClustersAPIClient = (*Client)(nil)

// DescribeEcsClustersPaginatorOptions is the paginator options for
// DescribeEcsClusters
type DescribeEcsClustersPaginatorOptions struct {
	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEcsClustersPaginator is a paginator for DescribeEcsClusters
type DescribeEcsClustersPaginator struct {
	options   DescribeEcsClustersPaginatorOptions
	client    DescribeEcsClustersAPIClient
	params    *DescribeEcsClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeEcsClustersPaginator returns a new DescribeEcsClustersPaginator
func NewDescribeEcsClustersPaginator(client DescribeEcsClustersAPIClient, params *DescribeEcsClustersInput, optFns ...func(*DescribeEcsClustersPaginatorOptions)) *DescribeEcsClustersPaginator {
	if params == nil {
		params = &DescribeEcsClustersInput{}
	}

	options := DescribeEcsClustersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEcsClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEcsClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEcsClusters page.
func (p *DescribeEcsClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEcsClustersOutput, error) {
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

	result, err := p.client.DescribeEcsClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEcsClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEcsClusters",
	}
}
