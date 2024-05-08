// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the grants distributed for the specified license.
func (c *Client) ListDistributedGrants(ctx context.Context, params *ListDistributedGrantsInput, optFns ...func(*Options)) (*ListDistributedGrantsOutput, error) {
	if params == nil {
		params = &ListDistributedGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributedGrants", params, optFns, c.addOperationListDistributedGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributedGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributedGrantsInput struct {

	// Filters to scope the results. The following filters are supported:
	//
	//   - LicenseArn
	//
	//   - GrantStatus
	//
	//   - GranteePrincipalARN
	//
	//   - ProductSKU
	//
	//   - LicenseIssuerName
	Filters []types.Filter

	// Amazon Resource Names (ARNs) of the grants.
	GrantArns []string

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDistributedGrantsOutput struct {

	// Distributed grant details.
	Grants []types.Grant

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributedGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDistributedGrants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDistributedGrants{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDistributedGrants"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributedGrants(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDistributedGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDistributedGrants",
	}
}
