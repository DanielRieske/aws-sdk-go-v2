// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes one or more remediation exceptions mentioned in the resource keys.
//
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
func (c *Client) DeleteRemediationExceptions(ctx context.Context, params *DeleteRemediationExceptionsInput, optFns ...func(*Options)) (*DeleteRemediationExceptionsOutput, error) {
	if params == nil {
		params = &DeleteRemediationExceptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRemediationExceptions", params, optFns, c.addOperationDeleteRemediationExceptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRemediationExceptionsInput struct {

	// The name of the Config rule for which you want to delete remediation exception
	// configuration.
	//
	// This member is required.
	ConfigRuleName *string

	// An exception list of resource exception keys to be processed with the current
	// request. Config adds exception for each resource key. For example, Config adds 3
	// exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []types.RemediationExceptionResourceKey

	noSmithyDocumentSerde
}

type DeleteRemediationExceptionsOutput struct {

	// Returns a list of failed delete remediation exceptions batch objects. Each
	// object in the batch consists of a list of failed items and failure messages.
	FailedBatches []types.FailedDeleteRemediationExceptionsBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRemediationExceptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteRemediationExceptions"); err != nil {
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
	if err = addOpDeleteRemediationExceptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRemediationExceptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRemediationExceptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteRemediationExceptions",
	}
}
