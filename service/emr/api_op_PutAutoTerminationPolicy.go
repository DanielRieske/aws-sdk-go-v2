// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Auto-termination is supported in Amazon EMR releases 5.30.0 and 6.1.0 and
// later. For more information, see [Using an auto-termination policy].
//
// Creates or updates an auto-termination policy for an Amazon EMR cluster. An
// auto-termination policy defines the amount of idle time in seconds after which a
// cluster automatically terminates. For alternative cluster termination options,
// see [Control cluster termination].
//
// [Using an auto-termination policy]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-auto-termination-policy.html
// [Control cluster termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html
func (c *Client) PutAutoTerminationPolicy(ctx context.Context, params *PutAutoTerminationPolicyInput, optFns ...func(*Options)) (*PutAutoTerminationPolicyOutput, error) {
	if params == nil {
		params = &PutAutoTerminationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAutoTerminationPolicy", params, optFns, c.addOperationPutAutoTerminationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAutoTerminationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAutoTerminationPolicyInput struct {

	// Specifies the ID of the Amazon EMR cluster to which the auto-termination policy
	// will be attached.
	//
	// This member is required.
	ClusterId *string

	// Specifies the auto-termination policy to attach to the cluster.
	AutoTerminationPolicy *types.AutoTerminationPolicy

	noSmithyDocumentSerde
}

type PutAutoTerminationPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAutoTerminationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAutoTerminationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAutoTerminationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAutoTerminationPolicy"); err != nil {
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
	if err = addOpPutAutoTerminationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAutoTerminationPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAutoTerminationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAutoTerminationPolicy",
	}
}
