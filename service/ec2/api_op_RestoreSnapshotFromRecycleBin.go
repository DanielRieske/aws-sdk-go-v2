// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a snapshot from the Recycle Bin. For more information, see [Restore snapshots from the Recycle Bin] in the
// Amazon EBS User Guide.
//
// [Restore snapshots from the Recycle Bin]: https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin-working-with-snaps.html#recycle-bin-restore-snaps
func (c *Client) RestoreSnapshotFromRecycleBin(ctx context.Context, params *RestoreSnapshotFromRecycleBinInput, optFns ...func(*Options)) (*RestoreSnapshotFromRecycleBinOutput, error) {
	if params == nil {
		params = &RestoreSnapshotFromRecycleBinInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreSnapshotFromRecycleBin", params, optFns, c.addOperationRestoreSnapshotFromRecycleBinMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreSnapshotFromRecycleBinOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreSnapshotFromRecycleBinInput struct {

	// The ID of the snapshot to restore.
	//
	// This member is required.
	SnapshotId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type RestoreSnapshotFromRecycleBinOutput struct {

	// The description for the snapshot.
	Description *string

	// Indicates whether the snapshot is encrypted.
	Encrypted *bool

	// The ARN of the Outpost on which the snapshot is stored. For more information,
	// see [Amazon EBS local snapshots on Outposts]in the Amazon EBS User Guide.
	//
	// [Amazon EBS local snapshots on Outposts]: https://docs.aws.amazon.com/ebs/latest/userguide/snapshots-outposts.html
	OutpostArn *string

	// The ID of the Amazon Web Services account that owns the EBS snapshot.
	OwnerId *string

	// The progress of the snapshot, as a percentage.
	Progress *string

	// The ID of the snapshot.
	SnapshotId *string

	// Reserved for future use.
	SseType types.SSEType

	// The time stamp when the snapshot was initiated.
	StartTime *time.Time

	// The state of the snapshot.
	State types.SnapshotState

	// The ID of the volume that was used to create the snapshot.
	VolumeId *string

	// The size of the volume, in GiB.
	VolumeSize *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreSnapshotFromRecycleBinMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpRestoreSnapshotFromRecycleBin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRestoreSnapshotFromRecycleBin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreSnapshotFromRecycleBin"); err != nil {
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
	if err = addOpRestoreSnapshotFromRecycleBinValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreSnapshotFromRecycleBin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreSnapshotFromRecycleBin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreSnapshotFromRecycleBin",
	}
}
