// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new table from an existing backup. Any number of users can execute up
// to 50 concurrent restores (any type of restore) in a given account.
//
// You can call RestoreTableFromBackup at a maximum rate of 10 times per second.
//
// You must manually set up the following on the restored table:
//
//   - Auto scaling policies
//
//   - IAM policies
//
//   - Amazon CloudWatch metrics and alarms
//
//   - Tags
//
//   - Stream settings
//
//   - Time to Live (TTL) settings
func (c *Client) RestoreTableFromBackup(ctx context.Context, params *RestoreTableFromBackupInput, optFns ...func(*Options)) (*RestoreTableFromBackupOutput, error) {
	if params == nil {
		params = &RestoreTableFromBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreTableFromBackup", params, optFns, c.addOperationRestoreTableFromBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreTableFromBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableFromBackupInput struct {

	// The Amazon Resource Name (ARN) associated with the backup.
	//
	// This member is required.
	BackupArn *string

	// The name of the new table to which the backup must be restored.
	//
	// This member is required.
	TargetTableName *string

	// The billing mode of the restored table.
	BillingModeOverride types.BillingMode

	// List of global secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	GlobalSecondaryIndexOverride []types.GlobalSecondaryIndex

	// List of local secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	LocalSecondaryIndexOverride []types.LocalSecondaryIndex

	// Sets the maximum number of read and write units for the specified on-demand
	// table. If you use this parameter, you must specify MaxReadRequestUnits ,
	// MaxWriteRequestUnits , or both.
	OnDemandThroughputOverride *types.OnDemandThroughput

	// Provisioned throughput settings for the restored table.
	ProvisionedThroughputOverride *types.ProvisionedThroughput

	// The new server-side encryption settings for the restored table.
	SSESpecificationOverride *types.SSESpecification

	noSmithyDocumentSerde
}

type RestoreTableFromBackupOutput struct {

	// The description of the table created from an existing backup.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreTableFromBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRestoreTableFromBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRestoreTableFromBackup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreTableFromBackup"); err != nil {
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
	if err = addOpRestoreTableFromBackupDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRestoreTableFromBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableFromBackup(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

func addOpRestoreTableFromBackupDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpRestoreTableFromBackupDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpRestoreTableFromBackupDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*RestoreTableFromBackupInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opRestoreTableFromBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreTableFromBackup",
	}
}
