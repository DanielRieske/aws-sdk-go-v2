// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an IAM Identity Center connection with Lake Formation to allow IAM
// Identity Center users and groups to access Data Catalog resources.
func (c *Client) CreateLakeFormationIdentityCenterConfiguration(ctx context.Context, params *CreateLakeFormationIdentityCenterConfigurationInput, optFns ...func(*Options)) (*CreateLakeFormationIdentityCenterConfigurationOutput, error) {
	if params == nil {
		params = &CreateLakeFormationIdentityCenterConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLakeFormationIdentityCenterConfiguration", params, optFns, c.addOperationCreateLakeFormationIdentityCenterConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLakeFormationIdentityCenterConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLakeFormationIdentityCenterConfigurationInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, view definitions, and other control information to manage
	// your Lake Formation environment.
	CatalogId *string

	// A list of the account IDs of Amazon Web Services accounts of third-party
	// applications that are allowed to access data managed by Lake Formation.
	ExternalFiltering *types.ExternalFilteringConfiguration

	// The ARN of the IAM Identity Center instance for which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// Amazon Web Services Service Namespaces in the Amazon Web Services General
	// Reference.
	InstanceArn *string

	// A list of Amazon Web Services account IDs and/or Amazon Web Services
	// organization/organizational unit ARNs that are allowed to access data managed by
	// Lake Formation. If the ShareRecipients list includes valid values, a resource
	// share is created with the principals you want to have access to the resources.
	// If the ShareRecipients value is null or the list is empty, no resource share is
	// created.
	ShareRecipients []types.DataLakePrincipal

	noSmithyDocumentSerde
}

type CreateLakeFormationIdentityCenterConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the Lake Formation application integrated
	// with IAM Identity Center.
	ApplicationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLakeFormationIdentityCenterConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLakeFormationIdentityCenterConfiguration"); err != nil {
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
	if err = addOpCreateLakeFormationIdentityCenterConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLakeFormationIdentityCenterConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLakeFormationIdentityCenterConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLakeFormationIdentityCenterConfiguration",
	}
}
