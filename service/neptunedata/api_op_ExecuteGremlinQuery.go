// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/document"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This commands executes a Gremlin query. Amazon Neptune is compatible with
// Apache TinkerPop3 and Gremlin, so you can use the Gremlin traversal language to
// query the graph, as described under [The Graph]in the Apache TinkerPop3 documentation.
// More details can also be found in [Accessing a Neptune graph with Gremlin].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that enables one of the following IAM actions in that cluster, depending on the
// query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [The Graph]: https://tinkerpop.apache.org/docs/current/reference/#graph
// [Accessing a Neptune graph with Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
func (c *Client) ExecuteGremlinQuery(ctx context.Context, params *ExecuteGremlinQueryInput, optFns ...func(*Options)) (*ExecuteGremlinQueryOutput, error) {
	if params == nil {
		params = &ExecuteGremlinQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteGremlinQuery", params, optFns, c.addOperationExecuteGremlinQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteGremlinQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteGremlinQueryInput struct {

	// Using this API, you can run Gremlin queries in string format much as you can
	// using the HTTP endpoint. The interface is compatible with whatever Gremlin
	// version your DB cluster is using (see the [Tinkerpop client section]to determine which Gremlin releases
	// your engine version supports).
	//
	// [Tinkerpop client section]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin-client.html#best-practices-gremlin-java-latest
	//
	// This member is required.
	GremlinQuery *string

	// If non-null, the query results are returned in a serialized response message in
	// the format specified by this parameter. See the [GraphSON]section in the TinkerPop
	// documentation for a list of the formats that are currently supported.
	//
	// [GraphSON]: https://tinkerpop.apache.org/docs/current/reference/#_graphson
	Serializer *string

	noSmithyDocumentSerde
}

type ExecuteGremlinQueryOutput struct {

	// Metadata about the Gremlin query.
	Meta document.Interface

	// The unique identifier of the Gremlin query.
	RequestId *string

	// The Gremlin query output from the server.
	Result document.Interface

	// The status of the Gremlin query.
	Status *types.GremlinQueryStatusAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteGremlinQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteGremlinQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteGremlinQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExecuteGremlinQuery"); err != nil {
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
	if err = addOpExecuteGremlinQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteGremlinQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteGremlinQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExecuteGremlinQuery",
	}
}
