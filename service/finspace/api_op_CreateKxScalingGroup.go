// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new scaling group.
func (c *Client) CreateKxScalingGroup(ctx context.Context, params *CreateKxScalingGroupInput, optFns ...func(*Options)) (*CreateKxScalingGroupOutput, error) {
	if params == nil {
		params = &CreateKxScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxScalingGroup", params, optFns, c.addOperationCreateKxScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxScalingGroupInput struct {

	// The identifier of the availability zones.
	//
	// This member is required.
	AvailabilityZoneId *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	//
	// This member is required.
	ClientToken *string

	// A unique identifier for the kdb environment, where you want to create the
	// scaling group.
	//
	// This member is required.
	EnvironmentId *string

	//  The memory and CPU capabilities of the scaling group host on which FinSpace
	// Managed kdb clusters will be placed.
	//
	// You can add one of the following values:
	//
	//   - kx.sg.4xlarge – The host type with a configuration of 108 GiB memory and 16
	//   vCPUs.
	//
	//   - kx.sg.8xlarge – The host type with a configuration of 216 GiB memory and 32
	//   vCPUs.
	//
	//   - kx.sg.16xlarge – The host type with a configuration of 432 GiB memory and 64
	//   vCPUs.
	//
	//   - kx.sg.32xlarge – The host type with a configuration of 864 GiB memory and
	//   128 vCPUs.
	//
	//   - kx.sg1.16xlarge – The host type with a configuration of 1949 GiB memory and
	//   64 vCPUs.
	//
	//   - kx.sg1.24xlarge – The host type with a configuration of 2948 GiB memory and
	//   96 vCPUs.
	//
	// This member is required.
	HostType *string

	// A unique identifier for the kdb scaling group.
	//
	// This member is required.
	ScalingGroupName *string

	//  A list of key-value pairs to label the scaling group. You can add up to 50
	// tags to a scaling group.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKxScalingGroupOutput struct {

	// The identifier of the availability zones.
	AvailabilityZoneId *string

	//  The timestamp at which the scaling group was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	// A unique identifier for the kdb environment, where you create the scaling
	// group.
	EnvironmentId *string

	//  The memory and CPU capabilities of the scaling group host on which FinSpace
	// Managed kdb clusters will be placed.
	HostType *string

	//  The last time that the scaling group was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// A unique identifier for the kdb scaling group.
	ScalingGroupName *string

	// The status of scaling group.
	//
	//   - CREATING – The scaling group creation is in progress.
	//
	//   - CREATE_FAILED – The scaling group creation has failed.
	//
	//   - ACTIVE – The scaling group is active.
	//
	//   - UPDATING – The scaling group is in the process of being updated.
	//
	//   - UPDATE_FAILED – The update action failed.
	//
	//   - DELETING – The scaling group is in the process of being deleted.
	//
	//   - DELETE_FAILED – The system failed to delete the scaling group.
	//
	//   - DELETED – The scaling group is successfully deleted.
	Status types.KxScalingGroupStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKxScalingGroup"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateKxScalingGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxScalingGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKxScalingGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxScalingGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxScalingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxScalingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxScalingGroupInput ")
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
func addIdempotencyToken_opCreateKxScalingGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxScalingGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKxScalingGroup",
	}
}
