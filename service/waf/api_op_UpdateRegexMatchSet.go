// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// Inserts or deletes RegexMatchTuple objects (filters) in a RegexMatchSet. For each RegexMatchSetUpdate
// object, you specify the following values:
//
//   - Whether to insert or delete the object from the array. If you want to
//     change a RegexMatchSetUpdate object, you delete the existing object and add a
//     new one.
//
//   - The part of a web request that you want AWS WAF to inspectupdate, such as a
//     query string or the value of the User-Agent header.
//
//   - The identifier of the pattern (a regular expression) that you want AWS WAF
//     to look for. For more information, see RegexPatternSet.
//
//   - Whether to perform any conversions on the request, such as converting it to
//     lowercase, before inspecting it for the specified string.
//
// For example, you can create a RegexPatternSet that matches any requests with
// User-Agent headers that contain the string B[a@]dB[o0]t . You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet , perform the following steps:
//
//   - Create a RegexMatchSet. For more information, see CreateRegexMatchSet.
//
//   - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
//     an UpdateRegexMatchSet request.
//
//   - Submit an UpdateRegexMatchSet request to specify the part of the request
//     that you want AWS WAF to inspect (for example, the header or the URI) and the
//     identifier of the RegexPatternSet that contain the regular expression patters
//     you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
func (c *Client) UpdateRegexMatchSet(ctx context.Context, params *UpdateRegexMatchSetInput, optFns ...func(*Options)) (*UpdateRegexMatchSetOutput, error) {
	if params == nil {
		params = &UpdateRegexMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRegexMatchSet", params, optFns, c.addOperationUpdateRegexMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRegexMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRegexMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The RegexMatchSetId of the RegexMatchSet that you want to update. RegexMatchSetId is
	// returned by CreateRegexMatchSetand by ListRegexMatchSets.
	//
	// This member is required.
	RegexMatchSetId *string

	// An array of RegexMatchSetUpdate objects that you want to insert into or delete
	// from a RegexMatchSet. For more information, see RegexMatchTuple.
	//
	// This member is required.
	Updates []types.RegexMatchSetUpdate

	noSmithyDocumentSerde
}

type UpdateRegexMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateRegexMatchSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRegexMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRegexMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRegexMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRegexMatchSet"); err != nil {
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
	if err = addOpUpdateRegexMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRegexMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRegexMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRegexMatchSet",
	}
}
