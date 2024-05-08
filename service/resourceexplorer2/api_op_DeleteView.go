// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified view.
//
// If the specified view is the default view for its Amazon Web Services Region,
// then all Searchoperations in that Region must explicitly specify the view to use
// until you configure a new default by calling the AssociateDefaultViewoperation.
func (c *Client) DeleteView(ctx context.Context, params *DeleteViewInput, optFns ...func(*Options)) (*DeleteViewOutput, error) {
	if params == nil {
		params = &DeleteViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteView", params, optFns, c.addOperationDeleteViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteViewInput struct {

	// The [Amazon resource name (ARN)] of the view that you want to delete.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	ViewArn *string

	noSmithyDocumentSerde
}

type DeleteViewOutput struct {

	// The [Amazon resource name (ARN)] of the view that you successfully deleted.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ViewArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteView{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteView"); err != nil {
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
	if err = addOpDeleteViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteView(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteView",
	}
}
