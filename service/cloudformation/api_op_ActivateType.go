// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates a public third-party extension, making it available for use in stack
// templates. For more information, see [Using public extensions]in the CloudFormation User Guide.
//
// Once you have activated a public third-party extension in your account and
// Region, use [SetTypeConfiguration]to specify configuration properties for the extension. For more
// information, see [Configuring extensions at the account level]in the CloudFormation User Guide.
//
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [Using public extensions]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html
// [Configuring extensions at the account level]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-private.html#registry-set-configuration
func (c *Client) ActivateType(ctx context.Context, params *ActivateTypeInput, optFns ...func(*Options)) (*ActivateTypeOutput, error) {
	if params == nil {
		params = &ActivateTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivateType", params, optFns, c.addOperationActivateTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivateTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ActivateTypeInput struct {

	// Whether to automatically update the extension in this account and Region when a
	// new minor version is published by the extension publisher. Major versions
	// released by the publisher must be manually updated.
	//
	// The default is true .
	AutoUpdate *bool

	// The name of the IAM execution role to use to activate the extension.
	ExecutionRoleArn *string

	// Contains logging configuration information for an extension.
	LoggingConfig *types.LoggingConfig

	// The major version of this extension you want to activate, if multiple major
	// versions are available. The default is the latest major version. CloudFormation
	// uses the latest available minor version of the major version selected.
	//
	// You can specify MajorVersion or VersionBump , but not both.
	MajorVersion *int64

	// The Amazon Resource Name (ARN) of the public extension.
	//
	// Conditional: You must specify PublicTypeArn , or TypeName , Type , and
	// PublisherId .
	PublicTypeArn *string

	// The ID of the extension publisher.
	//
	// Conditional: You must specify PublicTypeArn , or TypeName , Type , and
	// PublisherId .
	PublisherId *string

	// The extension type.
	//
	// Conditional: You must specify PublicTypeArn , or TypeName , Type , and
	// PublisherId .
	Type types.ThirdPartyType

	// The name of the extension.
	//
	// Conditional: You must specify PublicTypeArn , or TypeName , Type , and
	// PublisherId .
	TypeName *string

	// An alias to assign to the public extension, in this account and Region. If you
	// specify an alias for the extension, CloudFormation treats the alias as the
	// extension type name within this account and Region. You must use the alias to
	// refer to the extension in your templates, API calls, and CloudFormation console.
	//
	// An extension alias must be unique within a given account and Region. You can
	// activate the same public resource multiple times in the same account and Region,
	// using different type name aliases.
	TypeNameAlias *string

	// Manually updates a previously-activated type to a new major or minor version,
	// if available. You can also use this parameter to update the value of AutoUpdate .
	//
	//   - MAJOR : CloudFormation updates the extension to the newest major version, if
	//   one is available.
	//
	//   - MINOR : CloudFormation updates the extension to the newest minor version, if
	//   one is available.
	VersionBump types.VersionBump

	noSmithyDocumentSerde
}

type ActivateTypeOutput struct {

	// The Amazon Resource Name (ARN) of the activated extension, in this account and
	// Region.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivateTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpActivateType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpActivateType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ActivateType"); err != nil {
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
	if err = addOpActivateTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivateType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opActivateType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ActivateType",
	}
}
