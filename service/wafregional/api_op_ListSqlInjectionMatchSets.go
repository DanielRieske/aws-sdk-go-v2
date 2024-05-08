// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of SqlInjectionMatchSet objects.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func (c *Client) ListSqlInjectionMatchSets(ctx context.Context, params *ListSqlInjectionMatchSetsInput, optFns ...func(*Options)) (*ListSqlInjectionMatchSetsOutput, error) {
	if params == nil {
		params = &ListSqlInjectionMatchSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSqlInjectionMatchSets", params, optFns, c.addOperationListSqlInjectionMatchSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSqlInjectionMatchSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list the SqlInjectionMatchSet objects created by the current AWS account.
type ListSqlInjectionMatchSetsInput struct {

	// Specifies the number of SqlInjectionMatchSet objects that you want AWS WAF to return for this
	// request. If you have more SqlInjectionMatchSet objects than the number you
	// specify for Limit , the response includes a NextMarker value that you can use
	// to get another batch of Rules .
	Limit int32

	// If you specify a value for Limit and you have more SqlInjectionMatchSet objects than the value of
	// Limit , AWS WAF returns a NextMarker value in the response that allows you to
	// list another group of SqlInjectionMatchSets . For the second and subsequent
	// ListSqlInjectionMatchSets requests, specify the value of NextMarker from the
	// previous response to get information about another batch of
	// SqlInjectionMatchSets .
	NextMarker *string

	noSmithyDocumentSerde
}

// The response to a ListSqlInjectionMatchSets request.
type ListSqlInjectionMatchSetsOutput struct {

	// If you have more SqlInjectionMatchSet objects than the number that you specified for Limit in the
	// request, the response includes a NextMarker value. To list more
	// SqlInjectionMatchSet objects, submit another ListSqlInjectionMatchSets request,
	// and specify the NextMarker value from the response in the NextMarker value in
	// the next request.
	NextMarker *string

	// An array of SqlInjectionMatchSetSummary objects.
	SqlInjectionMatchSets []types.SqlInjectionMatchSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSqlInjectionMatchSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSqlInjectionMatchSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSqlInjectionMatchSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSqlInjectionMatchSets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSqlInjectionMatchSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSqlInjectionMatchSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSqlInjectionMatchSets",
	}
}
