// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Start import task for provisioning Sidewalk devices in bulk using an S3 CSV
// file.
func (c *Client) StartWirelessDeviceImportTask(ctx context.Context, params *StartWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*StartWirelessDeviceImportTaskOutput, error) {
	if params == nil {
		params = &StartWirelessDeviceImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartWirelessDeviceImportTask", params, optFns, c.addOperationStartWirelessDeviceImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartWirelessDeviceImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWirelessDeviceImportTaskInput struct {

	// The name of the Sidewalk destination that describes the IoT rule to route
	// messages from the devices in the import task that are onboarded to AWS IoT
	// Wireless.
	//
	// This member is required.
	DestinationName *string

	// The Sidewalk-related parameters for importing wireless devices that need to be
	// provisioned in bulk.
	//
	// This member is required.
	Sidewalk *types.SidewalkStartImportInfo

	// Each resource must have a unique client request token. The client token is used
	// to implement idempotency. It ensures that the request completes no more than one
	// time. If you retry a request with the same token and the same parameters, the
	// request will complete successfully. However, if you try to create a new resource
	// using the same token but different parameters, an HTTP 409 conflict occurs. If
	// you omit this value, AWS SDKs will automatically generate a unique client
	// request. For more information about idempotency, see Ensuring idempotency in
	// Amazon EC2 API requests (https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html)
	// .
	ClientRequestToken *string

	// The tag to attach to the specified resource. Tags are metadata that you can use
	// to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartWirelessDeviceImportTaskOutput struct {

	// The ARN (Amazon Resource Name) of the import task.
	Arn *string

	// The import task ID.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartWirelessDeviceImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartWirelessDeviceImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartWirelessDeviceImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartWirelessDeviceImportTask"); err != nil {
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
	if err = addIdempotencyToken_opStartWirelessDeviceImportTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartWirelessDeviceImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartWirelessDeviceImportTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartWirelessDeviceImportTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartWirelessDeviceImportTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartWirelessDeviceImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartWirelessDeviceImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartWirelessDeviceImportTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartWirelessDeviceImportTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartWirelessDeviceImportTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartWirelessDeviceImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartWirelessDeviceImportTask",
	}
}
