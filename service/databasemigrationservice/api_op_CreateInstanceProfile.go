// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the instance profile using the specified parameters.
func (c *Client) CreateInstanceProfile(ctx context.Context, params *CreateInstanceProfileInput, optFns ...func(*Options)) (*CreateInstanceProfileOutput, error) {
	if params == nil {
		params = &CreateInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstanceProfile", params, optFns, c.addOperationCreateInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstanceProfileInput struct {

	// The Availability Zone where the instance profile will be created. The default
	// value is a random, system-chosen Availability Zone in the Amazon Web Services
	// Region where your data provider is created, for examplem us-east-1d .
	AvailabilityZone *string

	// A user-friendly description of the instance profile.
	Description *string

	// A user-friendly name for the instance profile.
	InstanceProfileName *string

	// The Amazon Resource Name (ARN) of the KMS key that is used to encrypt the
	// connection parameters for the instance profile.
	//
	// If you don't specify a value for the KmsKeyArn parameter, then DMS uses your
	// default encryption key.
	//
	// KMS creates the default encryption key for your Amazon Web Services account.
	// Your Amazon Web Services account has a different default encryption key for each
	// Amazon Web Services Region.
	KmsKeyArn *string

	// Specifies the network type for the instance profile. A value of IPV4 represents
	// an instance profile with IPv4 network type and only supports IPv4 addressing. A
	// value of IPV6 represents an instance profile with IPv6 network type and only
	// supports IPv6 addressing. A value of DUAL represents an instance profile with
	// dual network type that supports IPv4 and IPv6 addressing.
	NetworkType *string

	// Specifies the accessibility options for the instance profile. A value of true
	// represents an instance profile with a public IP address. A value of false
	// represents an instance profile with a private IP address. The default value is
	// true .
	PubliclyAccessible *bool

	// A subnet group to associate with the instance profile.
	SubnetGroupIdentifier *string

	// One or more tags to be assigned to the instance profile.
	Tags []types.Tag

	// Specifies the VPC security group names to be used with the instance profile.
	// The VPC security group must work with the VPC containing the instance profile.
	VpcSecurityGroups []string

	noSmithyDocumentSerde
}

type CreateInstanceProfileOutput struct {

	// The instance profile that was created.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateInstanceProfile"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstanceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateInstanceProfile",
	}
}
