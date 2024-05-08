// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the bundles that you can apply to an Amazon Lightsail instance when you
// create it.
//
// A bundle describes the specifications of an instance, such as the monthly cost,
// amount of memory, the number of vCPUs, amount of storage space, and monthly
// network data transfer quota.
//
// Bundles are referred to as instance plans in the Lightsail console.
func (c *Client) GetBundles(ctx context.Context, params *GetBundlesInput, optFns ...func(*Options)) (*GetBundlesOutput, error) {
	if params == nil {
		params = &GetBundlesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBundles", params, optFns, c.addOperationGetBundlesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBundlesInput struct {

	// Returns a list of bundles that are specific to Lightsail for Research.
	//
	// You must use this parameter to view Lightsail for Research bundles.
	AppCategory types.AppCategory

	// A Boolean value that indicates whether to include inactive (unavailable)
	// bundles in the response of your request.
	IncludeInactive *bool

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetBundles request. If your results are
	// paginated, the response will return a next page token that you can specify as
	// the page token in a subsequent request.
	PageToken *string

	noSmithyDocumentSerde
}

type GetBundlesOutput struct {

	// An array of key-value pairs that contains information about the available
	// bundles.
	Bundles []types.Bundle

	// The token to advance to the next page of results from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetBundles request and specify
	// the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBundlesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBundles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBundles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBundles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBundles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBundles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBundles",
	}
}
