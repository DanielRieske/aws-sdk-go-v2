// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a created decoder manifest.
func (c *Client) GetDecoderManifest(ctx context.Context, params *GetDecoderManifestInput, optFns ...func(*Options)) (*GetDecoderManifestOutput, error) {
	if params == nil {
		params = &GetDecoderManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDecoderManifest", params, optFns, c.addOperationGetDecoderManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDecoderManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDecoderManifestInput struct {

	//  The name of the decoder manifest to retrieve information about.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetDecoderManifestOutput struct {

	//  The Amazon Resource Name (ARN) of the decoder manifest.
	//
	// This member is required.
	Arn *string

	//  The time the decoder manifest was created in seconds since epoch (January 1,
	// 1970 at midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	//  The time the decoder manifest was last updated in seconds since epoch (January
	// 1, 1970 at midnight UTC time).
	//
	// This member is required.
	LastModificationTime *time.Time

	//  The name of the decoder manifest.
	//
	// This member is required.
	Name *string

	//  A brief description of the decoder manifest.
	Description *string

	// The detailed message for the decoder manifest. When a decoder manifest is in an
	// INVALID status, the message contains detailed reason and help information.
	Message *string

	//  The ARN of a vehicle model (model manifest) associated with the decoder
	// manifest.
	ModelManifestArn *string

	//  The state of the decoder manifest. If the status is ACTIVE , the decoder
	// manifest can't be edited. If the status is marked DRAFT , you can edit the
	// decoder manifest.
	Status types.ManifestStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDecoderManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDecoderManifest"); err != nil {
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
	if err = addOpGetDecoderManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDecoderManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDecoderManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDecoderManifest",
	}
}
