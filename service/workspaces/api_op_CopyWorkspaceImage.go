// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified image from the specified Region to the current Region. For
// more information about copying images, see [Copy a Custom WorkSpaces Image].
//
// In the China (Ningxia) Region, you can copy images only within the same Region.
//
// In Amazon Web Services GovCloud (US), to copy images to and from other Regions,
// contact Amazon Web Services Support.
//
// Before copying a shared image, be sure to verify that it has been shared from
// the correct Amazon Web Services account. To determine if an image has been
// shared and to see the ID of the Amazon Web Services account that owns an image,
// use the [DescribeWorkSpaceImages]and [DescribeWorkspaceImagePermissions] API operations.
//
// [DescribeWorkspaceImagePermissions]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImagePermissions.html
// [DescribeWorkSpaceImages]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImages.html
// [Copy a Custom WorkSpaces Image]: https://docs.aws.amazon.com/workspaces/latest/adminguide/copy-custom-image.html
func (c *Client) CopyWorkspaceImage(ctx context.Context, params *CopyWorkspaceImageInput, optFns ...func(*Options)) (*CopyWorkspaceImageOutput, error) {
	if params == nil {
		params = &CopyWorkspaceImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyWorkspaceImage", params, optFns, c.addOperationCopyWorkspaceImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyWorkspaceImageInput struct {

	// The name of the image.
	//
	// This member is required.
	Name *string

	// The identifier of the source image.
	//
	// This member is required.
	SourceImageId *string

	// The identifier of the source Region.
	//
	// This member is required.
	SourceRegion *string

	// A description of the image.
	Description *string

	// The tags for the image.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CopyWorkspaceImageOutput struct {

	// The identifier of the image.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyWorkspaceImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopyWorkspaceImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyWorkspaceImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyWorkspaceImage"); err != nil {
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
	if err = addOpCopyWorkspaceImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyWorkspaceImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyWorkspaceImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyWorkspaceImage",
	}
}
