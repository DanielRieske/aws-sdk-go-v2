// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified provisioning artifact (also known as a version) for the
// specified product.
//
// You cannot update a provisioning artifact for a product that was shared with
// you.
func (c *Client) UpdateProvisioningArtifact(ctx context.Context, params *UpdateProvisioningArtifactInput, optFns ...func(*Options)) (*UpdateProvisioningArtifactOutput, error) {
	if params == nil {
		params = &UpdateProvisioningArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisioningArtifact", params, optFns, c.addOperationUpdateProvisioningArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisioningArtifactInput struct {

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The identifier of the provisioning artifact.
	//
	// This member is required.
	ProvisioningArtifactId *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// Indicates whether the product version is active.
	//
	// Inactive provisioning artifacts are invisible to end users. End users cannot
	// launch or update a provisioned product from an inactive provisioning artifact.
	Active *bool

	// The updated description of the provisioning artifact.
	Description *string

	// Information set by the administrator to provide guidance to end users about
	// which provisioning artifacts to use.
	//
	// The DEFAULT value indicates that the product version is active.
	//
	// The administrator can set the guidance to DEPRECATED to inform users that the
	// product version is deprecated. Users are able to make updates to a provisioned
	// product of a deprecated version but cannot launch new provisioned products using
	// a deprecated version.
	Guidance types.ProvisioningArtifactGuidance

	// The updated name of the provisioning artifact.
	Name *string

	noSmithyDocumentSerde
}

type UpdateProvisioningArtifactOutput struct {

	// The URL of the CloudFormation template in Amazon S3 or GitHub in JSON format.
	Info map[string]string

	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *types.ProvisioningArtifactDetail

	// The status of the current request.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProvisioningArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProvisioningArtifact"); err != nil {
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
	if err = addOpUpdateProvisioningArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisioningArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProvisioningArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProvisioningArtifact",
	}
}
