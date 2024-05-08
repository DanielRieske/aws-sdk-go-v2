// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about a state machine [alias].
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # ListStateMachineAliases
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func (c *Client) DescribeStateMachineAlias(ctx context.Context, params *DescribeStateMachineAliasInput, optFns ...func(*Options)) (*DescribeStateMachineAliasOutput, error) {
	if params == nil {
		params = &DescribeStateMachineAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachineAlias", params, optFns, c.addOperationDescribeStateMachineAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineAliasInput struct {

	// The Amazon Resource Name (ARN) of the state machine alias.
	//
	// This member is required.
	StateMachineAliasArn *string

	noSmithyDocumentSerde
}

type DescribeStateMachineAliasOutput struct {

	// The date the state machine alias was created.
	CreationDate *time.Time

	// A description of the alias.
	Description *string

	// The name of the state machine alias.
	Name *string

	// The routing configuration of the alias.
	RoutingConfiguration []types.RoutingConfigurationListItem

	// The Amazon Resource Name (ARN) of the state machine alias.
	StateMachineAliasArn *string

	// The date the state machine alias was last updated.
	//
	// For a newly created state machine, this is the same as the creation date.
	UpdateDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStateMachineAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachineAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachineAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStateMachineAlias"); err != nil {
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
	if err = addOpDescribeStateMachineAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachineAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachineAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStateMachineAlias",
	}
}
