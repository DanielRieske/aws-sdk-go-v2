// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Updates capacity settings for a managed EC2 fleet or container fleet. For these
// fleets, you adjust capacity by changing the number of instances in the fleet.
// Fleet capacity determines the number of game sessions and players that the fleet
// can host based on its configuration. For fleets with multiple locations, use
// this operation to manage capacity settings in each location individually.
//
// Use this operation to set these fleet capacity properties:
//
//   - Minimum/maximum size: Set hard limits on the number of Amazon EC2 instances
//     allowed. If Amazon GameLift receives a request--either through manual update or
//     automatic scaling--it won't change the capacity to a value outside of this
//     range.
//
//   - Desired capacity: As an alternative to automatic scaling, manually set the
//     number of Amazon EC2 instances to be maintained. Before changing a fleet's
//     desired capacity, check the maximum capacity of the fleet's Amazon EC2 instance
//     type by calling [DescribeEC2InstanceLimits].
//
// To update capacity for a fleet's home Region, or if the fleet has no remote
// locations, omit the Location parameter. The fleet must be in ACTIVE status.
//
// To update capacity for a fleet's remote location, set the Location parameter to
// the location to update. The location must be in ACTIVE status.
//
// If successful, Amazon GameLift updates the capacity settings and returns the
// identifiers for the updated fleet and/or location. If a requested change to
// desired capacity exceeds the instance type's limit, the LimitExceeded exception
// occurs.
//
// Updates often prompt an immediate change in fleet capacity, such as when
// current capacity is different than the new desired capacity or outside the new
// limits. In this scenario, Amazon GameLift automatically initiates steps to add
// or remove instances in the fleet location. You can track a fleet's current
// capacity by calling [DescribeFleetCapacity]or [DescribeFleetLocationCapacity].
//
// # Learn more
//
// [Scaling fleet capacity]
//
// [Scaling fleet capacity]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-manage-capacity.html
// [DescribeEC2InstanceLimits]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeEC2InstanceLimits.html
// [DescribeFleetLocationCapacity]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html
// [DescribeFleetCapacity]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html
func (c *Client) UpdateFleetCapacity(ctx context.Context, params *UpdateFleetCapacityInput, optFns ...func(*Options)) (*UpdateFleetCapacityOutput, error) {
	if params == nil {
		params = &UpdateFleetCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetCapacity", params, optFns, c.addOperationUpdateFleetCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetCapacityInput struct {

	// A unique identifier for the fleet to update capacity settings for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The number of Amazon EC2 instances you want to maintain in the specified fleet
	// location. This value must fall between the minimum and maximum size limits.
	// Changes in desired instance value can take up to 1 minute to be reflected when
	// viewing the fleet's capacity settings.
	DesiredInstances *int32

	// The name of a remote location to update fleet capacity settings for, in the
	// form of an Amazon Web Services Region code such as us-west-2 .
	Location *string

	// The maximum number of instances that are allowed in the specified fleet
	// location. If this parameter is not set, the default is 1.
	MaxSize *int32

	// The minimum number of instances that are allowed in the specified fleet
	// location. If this parameter is not set, the default is 0.
	MinSize *int32

	noSmithyDocumentSerde
}

type UpdateFleetCapacityOutput struct {

	// The Amazon Resource Name ([ARN] ) that is assigned to a Amazon GameLift fleet
	// resource and uniquely identifies it. ARNs are unique across all Regions. Format
	// is arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	FleetArn *string

	// A unique identifier for the fleet that was updated.
	FleetId *string

	// The remote location being updated, expressed as an Amazon Web Services Region
	// code, such as us-west-2 .
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFleetCapacity"); err != nil {
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
	if err = addOpUpdateFleetCapacityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetCapacity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetCapacity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFleetCapacity",
	}
}
