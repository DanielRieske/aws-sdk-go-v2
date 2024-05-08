// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate a lens to a workload.
//
// Up to 10 lenses can be associated with a workload in a single API operation. A
// maximum of 20 lenses can be associated with a workload.
//
// # Disclaimer
//
// By accessing and/or applying custom lenses created by another Amazon Web
// Services user or account, you acknowledge that custom lenses created by other
// users and shared with you are Third Party Content as defined in the Amazon Web
// Services Customer Agreement.
func (c *Client) AssociateLenses(ctx context.Context, params *AssociateLensesInput, optFns ...func(*Options)) (*AssociateLensesOutput, error) {
	if params == nil {
		params = &AssociateLensesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateLenses", params, optFns, c.addOperationAssociateLensesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateLensesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to associate lens reviews.
type AssociateLensesInput struct {

	// List of lens aliases to associate or disassociate with a workload. Up to 10
	// lenses can be specified.
	//
	// Identify a lens using its LensSummary$LensAlias.
	//
	// This member is required.
	LensAliases []string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	noSmithyDocumentSerde
}

type AssociateLensesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateLensesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateLenses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateLenses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateLenses"); err != nil {
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
	if err = addOpAssociateLensesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateLenses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateLenses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateLenses",
	}
}
