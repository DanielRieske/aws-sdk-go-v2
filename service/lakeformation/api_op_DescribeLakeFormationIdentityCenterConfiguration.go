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

// Retrieves the instance ARN and application ARN for the connection.
func (c *Client) DescribeLakeFormationIdentityCenterConfiguration(ctx context.Context, params *DescribeLakeFormationIdentityCenterConfigurationInput, optFns ...func(*Options)) (*DescribeLakeFormationIdentityCenterConfigurationOutput, error) {
	if params == nil {
		params = &DescribeLakeFormationIdentityCenterConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLakeFormationIdentityCenterConfiguration", params, optFns, c.addOperationDescribeLakeFormationIdentityCenterConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLakeFormationIdentityCenterConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLakeFormationIdentityCenterConfigurationInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	noSmithyDocumentSerde
}

type DescribeLakeFormationIdentityCenterConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the Lake Formation application integrated
	// with IAM Identity Center.
	ApplicationArn *string

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// Indicates if external filtering is enabled.
	ExternalFiltering *types.ExternalFilteringConfiguration

	// The Amazon Resource Name (ARN) of the connection.
	InstanceArn *string

	// The Amazon Resource Name (ARN) of the RAM share.
	ResourceShare *string

	// A list of Amazon Web Services account IDs or Amazon Web Services
	// organization/organizational unit ARNs that are allowed to access data managed by
	// Lake Formation. If the ShareRecipients list includes valid values, a resource
	// share is created with the principals you want to have access to the resources as
	// the ShareRecipients . If the ShareRecipients value is null or the list is
	// empty, no resource share is created.
	ShareRecipients []types.DataLakePrincipal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLakeFormationIdentityCenterConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLakeFormationIdentityCenterConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLakeFormationIdentityCenterConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLakeFormationIdentityCenterConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLakeFormationIdentityCenterConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLakeFormationIdentityCenterConfiguration",
	}
}
