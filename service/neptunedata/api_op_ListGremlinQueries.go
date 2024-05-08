// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists active Gremlin queries. See [Gremlin query status API] for details about the output.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [Gremlin query status API]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-api-status.html
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func (c *Client) ListGremlinQueries(ctx context.Context, params *ListGremlinQueriesInput, optFns ...func(*Options)) (*ListGremlinQueriesOutput, error) {
	if params == nil {
		params = &ListGremlinQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGremlinQueries", params, optFns, c.addOperationListGremlinQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGremlinQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGremlinQueriesInput struct {

	// If set to TRUE , the list returned includes waiting queries. The default is
	// FALSE ;
	IncludeWaiting *bool

	noSmithyDocumentSerde
}

type ListGremlinQueriesOutput struct {

	// The number of queries that have been accepted but not yet completed, including
	// queries in the queue.
	AcceptedQueryCount *int32

	// A list of the current queries.
	Queries []types.GremlinQueryStatus

	// The number of Gremlin queries currently running.
	RunningQueryCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGremlinQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGremlinQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGremlinQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGremlinQueries"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGremlinQueries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListGremlinQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGremlinQueries",
	}
}
