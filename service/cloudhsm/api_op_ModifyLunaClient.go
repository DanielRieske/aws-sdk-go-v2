// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Modifies the certificate used by the client.
//
// This action can potentially start a workflow to install the new certificate on
// the client's HSMs.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
func (c *Client) ModifyLunaClient(ctx context.Context, params *ModifyLunaClientInput, optFns ...func(*Options)) (*ModifyLunaClientOutput, error) {
	if params == nil {
		params = &ModifyLunaClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyLunaClient", params, optFns, c.addOperationModifyLunaClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyLunaClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyLunaClientInput struct {

	// The new certificate for the client.
	//
	// This member is required.
	Certificate *string

	// The ARN of the client.
	//
	// This member is required.
	ClientArn *string

	noSmithyDocumentSerde
}

type ModifyLunaClientOutput struct {

	// The ARN of the client.
	ClientArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyLunaClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyLunaClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyLunaClient{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyLunaClient"); err != nil {
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
	if err = addOpModifyLunaClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyLunaClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyLunaClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyLunaClient",
	}
}
