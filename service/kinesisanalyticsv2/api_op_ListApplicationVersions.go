// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the versions for the specified application, including versions that
// were rolled back. The response also includes a summary of the configuration
// associated with each version.
//
// To get the complete description of a specific application version, invoke the DescribeApplicationVersion
// operation.
//
// This operation is supported only for Managed Service for Apache Flink.
func (c *Client) ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) {
	if params == nil {
		params = &ListApplicationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationVersions", params, optFns, c.addOperationListApplicationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationVersionsInput struct {

	// The name of the application for which you want to list all versions.
	//
	// This member is required.
	ApplicationName *string

	// The maximum number of versions to list in this invocation of the operation.
	Limit *int32

	// If a previous invocation of this operation returned a pagination token, pass it
	// into this value to retrieve the next set of results. For more information about
	// pagination, see [Using the Amazon Command Line Interface's Pagination Options].
	//
	// [Using the Amazon Command Line Interface's Pagination Options]: https://docs.aws.amazon.com/cli/latest/userguide/pagination.html
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationVersionsOutput struct {

	// A list of the application versions and the associated configuration summaries.
	// The list includes application versions that were rolled back.
	//
	// To get the complete description of a specific application version, invoke the DescribeApplicationVersion
	// operation.
	ApplicationVersionSummaries []types.ApplicationVersionSummary

	// The pagination token for the next set of results, or null if there are no
	// additional results. To retrieve the next set of items, pass this token into a
	// subsequent invocation of this operation. For more information about pagination,
	// see [Using the Amazon Command Line Interface's Pagination Options].
	//
	// [Using the Amazon Command Line Interface's Pagination Options]: https://docs.aws.amazon.com/cli/latest/userguide/pagination.html
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationVersions"); err != nil {
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
	if err = addOpListApplicationVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListApplicationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationVersions",
	}
}
