// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between an approval rule template and one or more
// specified repositories.
func (c *Client) BatchAssociateApprovalRuleTemplateWithRepositories(ctx context.Context, params *BatchAssociateApprovalRuleTemplateWithRepositoriesInput, optFns ...func(*Options)) (*BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error) {
	if params == nil {
		params = &BatchAssociateApprovalRuleTemplateWithRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateApprovalRuleTemplateWithRepositories", params, optFns, c.addOperationBatchAssociateApprovalRuleTemplateWithRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateApprovalRuleTemplateWithRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateApprovalRuleTemplateWithRepositoriesInput struct {

	// The name of the template you want to associate with one or more repositories.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The names of the repositories you want to associate with the template.
	//
	// The length constraint limit is for each string in the array. The array itself
	// can be empty.
	//
	// This member is required.
	RepositoryNames []string

	noSmithyDocumentSerde
}

type BatchAssociateApprovalRuleTemplateWithRepositoriesOutput struct {

	// A list of names of the repositories that have been associated with the template.
	//
	// This member is required.
	AssociatedRepositoryNames []string

	// A list of any errors that might have occurred while attempting to create the
	// association between the template and the repositories.
	//
	// This member is required.
	Errors []types.BatchAssociateApprovalRuleTemplateWithRepositoriesError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAssociateApprovalRuleTemplateWithRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateApprovalRuleTemplateWithRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateApprovalRuleTemplateWithRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchAssociateApprovalRuleTemplateWithRepositories"); err != nil {
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
	if err = addOpBatchAssociateApprovalRuleTemplateWithRepositoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateApprovalRuleTemplateWithRepositories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateApprovalRuleTemplateWithRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchAssociateApprovalRuleTemplateWithRepositories",
	}
}
