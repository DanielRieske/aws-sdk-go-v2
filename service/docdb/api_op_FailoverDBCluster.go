// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Forces a failover for a cluster.
//
// A failover for a cluster promotes one of the Amazon DocumentDB replicas
// (read-only instances) in the cluster to be the primary instance (the cluster
// writer).
//
// If the primary instance fails, Amazon DocumentDB automatically fails over to an
// Amazon DocumentDB replica, if one exists. You can force a failover when you want
// to simulate a failure of a primary instance for testing.
func (c *Client) FailoverDBCluster(ctx context.Context, params *FailoverDBClusterInput, optFns ...func(*Options)) (*FailoverDBClusterOutput, error) {
	if params == nil {
		params = &FailoverDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverDBCluster", params, optFns, c.addOperationFailoverDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to FailoverDBCluster.
type FailoverDBClusterInput struct {

	// A cluster identifier to force a failover for. This parameter is not case
	// sensitive.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DBCluster .
	DBClusterIdentifier *string

	// The name of the instance to promote to the primary instance.
	//
	// You must specify the instance identifier for an Amazon DocumentDB replica in
	// the cluster. For example, mydbcluster-replica1 .
	TargetDBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type FailoverDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFailoverDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "FailoverDBCluster"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFailoverDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "FailoverDBCluster",
	}
}
