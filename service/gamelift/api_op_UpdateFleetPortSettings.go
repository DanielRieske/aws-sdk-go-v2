// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates permissions that allow inbound traffic to connect to game sessions in
// the fleet.
//
// To update settings, specify the fleet ID to be updated and specify the changes
// to be made. List the permissions you want to add in
// InboundPermissionAuthorizations , and permissions you want to remove in
// InboundPermissionRevocations . Permissions to be removed must match existing
// fleet permissions.
//
// For a container fleet, inbound permissions must specify port numbers that are
// defined in the fleet's connection port settings.
//
// If successful, the fleet ID for the updated fleet is returned. For fleets with
// remote locations, port setting updates can take time to propagate across all
// locations. You can check the status of updates in each location by calling
// DescribeFleetPortSettings with a location name.
//
// # Learn more
//
// [Setting up Amazon GameLift fleets]
//
// [Setting up Amazon GameLift fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
func (c *Client) UpdateFleetPortSettings(ctx context.Context, params *UpdateFleetPortSettingsInput, optFns ...func(*Options)) (*UpdateFleetPortSettingsOutput, error) {
	if params == nil {
		params = &UpdateFleetPortSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetPortSettings", params, optFns, c.addOperationUpdateFleetPortSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetPortSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetPortSettingsInput struct {

	// A unique identifier for the fleet to update port settings for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// A collection of port settings to be added to the fleet resource.
	InboundPermissionAuthorizations []types.IpPermission

	// A collection of port settings to be removed from the fleet resource.
	InboundPermissionRevocations []types.IpPermission

	noSmithyDocumentSerde
}

type UpdateFleetPortSettingsOutput struct {

	// The Amazon Resource Name ([ARN] ) that is assigned to a Amazon GameLift fleet
	// resource and uniquely identifies it. ARNs are unique across all Regions. Format
	// is arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	FleetArn *string

	// A unique identifier for the fleet that was updated.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetPortSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFleetPortSettings"); err != nil {
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
	if err = addOpUpdateFleetPortSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetPortSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetPortSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFleetPortSettings",
	}
}
