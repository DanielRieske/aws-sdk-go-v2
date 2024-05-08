// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or removes permission settings for the specified snapshot. You may add or
// remove specified Amazon Web Services account IDs from a snapshot's list of
// create volume permissions, but you cannot do both in a single operation. If you
// need to both add and remove account IDs for a snapshot, you must use multiple
// operations. You can make up to 500 modifications to a snapshot in a single
// operation.
//
// Encrypted snapshots and snapshots with Amazon Web Services Marketplace product
// codes cannot be made public. Snapshots encrypted with your default KMS key
// cannot be shared with other accounts.
//
// For more information about modifying snapshot permissions, see [Share a snapshot] in the Amazon
// EBS User Guide.
//
// [Share a snapshot]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-modifying-snapshot-permissions.html
func (c *Client) ModifySnapshotAttribute(ctx context.Context, params *ModifySnapshotAttributeInput, optFns ...func(*Options)) (*ModifySnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifySnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifySnapshotAttribute", params, optFns, c.addOperationModifySnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifySnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotAttributeInput struct {

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// The snapshot attribute to modify. Only volume creation permissions can be
	// modified.
	Attribute types.SnapshotAttributeName

	// A JSON representation of the snapshot attribute modification.
	CreateVolumePermission *types.CreateVolumePermissionModifications

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The group to modify for the snapshot.
	GroupNames []string

	// The type of operation to perform to the attribute.
	OperationType types.OperationType

	// The account ID to modify for the snapshot.
	UserIds []string

	noSmithyDocumentSerde
}

type ModifySnapshotAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifySnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifySnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifySnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifySnapshotAttribute"); err != nil {
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
	if err = addOpModifySnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifySnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifySnapshotAttribute",
	}
}
