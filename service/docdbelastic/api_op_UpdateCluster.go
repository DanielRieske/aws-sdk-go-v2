// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an elastic cluster. This includes updating admin-username/password,
// upgrading the API version, and setting up a backup window and maintenance window
func (c *Client) UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) {
	if params == nil {
		params = &UpdateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCluster", params, optFns, c.addOperationUpdateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterInput struct {

	// The ARN identifier of the elastic cluster.
	//
	// This member is required.
	ClusterArn *string

	// The password associated with the elastic cluster administrator. This password
	// can contain any printable ASCII character except forward slash (/), double quote
	// ("), or the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	AdminUserPassword *string

	// The authentication type used to determine where to fetch the password used for
	// accessing the elastic cluster. Valid types are PLAIN_TEXT or SECRET_ARN .
	AuthType types.Auth

	// The number of days for which automatic snapshots are retained.
	BackupRetentionPeriod *int32

	// The client token for the elastic cluster.
	ClientToken *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled, as determined by the backupRetentionPeriod .
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Default: a 30-minute window selected at random from an 8-hour block of time for
	// each Amazon Web Services Region, occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The number of vCPUs assigned to each elastic cluster shard. Maximum is 64.
	// Allowed values are 2, 4, 8, 16, 32, 64.
	ShardCapacity *int32

	// The number of shards assigned to the elastic cluster. Maximum is 32.
	ShardCount *int32

	// The number of replica instances applying to all shards in the elastic cluster.
	// A shardInstanceCount value of 1 means there is one writer instance, and any
	// additional instances are replicas that can be used for reads and to improve
	// availability.
	ShardInstanceCount *int32

	// The Amazon EC2 subnet IDs for the elastic cluster.
	SubnetIds []string

	// A list of EC2 VPC security groups to associate with the elastic cluster.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type UpdateClusterOutput struct {

	// Returns information about the updated elastic cluster.
	//
	// This member is required.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCluster"); err != nil {
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
	if err = addIdempotencyToken_opUpdateClusterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCluster(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateClusterInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateClusterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCluster",
	}
}
