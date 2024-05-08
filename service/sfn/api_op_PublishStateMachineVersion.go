// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a [version] from the current revision of a state machine. Use versions to create
// immutable snapshots of your state machine. You can start executions from
// versions either directly or with an alias. To create an alias, use CreateStateMachineAlias.
//
// You can publish up to 1000 versions for each state machine. You must manually
// delete unused versions using the DeleteStateMachineVersionAPI action.
//
// PublishStateMachineVersion is an idempotent API. It doesn't create a duplicate
// state machine version if it already exists for the current revision. Step
// Functions bases PublishStateMachineVersion 's idempotency check on the
// stateMachineArn , name , and revisionId parameters. Requests with the same
// parameters return a successful idempotent response. If you don't specify a
// revisionId , Step Functions checks for a previously published version of the
// state machine's current revision.
//
// Related operations:
//
// # DeleteStateMachineVersion
//
// # ListStateMachineVersions
//
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func (c *Client) PublishStateMachineVersion(ctx context.Context, params *PublishStateMachineVersionInput, optFns ...func(*Options)) (*PublishStateMachineVersionOutput, error) {
	if params == nil {
		params = &PublishStateMachineVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishStateMachineVersion", params, optFns, c.addOperationPublishStateMachineVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishStateMachineVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishStateMachineVersionInput struct {

	// The Amazon Resource Name (ARN) of the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// An optional description of the state machine version.
	Description *string

	// Only publish the state machine version if the current state machine's revision
	// ID matches the specified ID.
	//
	// Use this option to avoid publishing a version if the state machine changed
	// since you last updated it. If the specified revision ID doesn't match the state
	// machine's current revision ID, the API returns ConflictException .
	//
	// To specify an initial revision ID for a state machine with no revision ID
	// assigned, specify the string INITIAL for the revisionId parameter. For example,
	// you can specify a revisionID of INITIAL when you create a state machine using
	// the CreateStateMachineAPI action.
	RevisionId *string

	noSmithyDocumentSerde
}

type PublishStateMachineVersionOutput struct {

	// The date the version was created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon Resource Name (ARN) (ARN) that identifies the state machine version.
	//
	// This member is required.
	StateMachineVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishStateMachineVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPublishStateMachineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPublishStateMachineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PublishStateMachineVersion"); err != nil {
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
	if err = addOpPublishStateMachineVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishStateMachineVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishStateMachineVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PublishStateMachineVersion",
	}
}
