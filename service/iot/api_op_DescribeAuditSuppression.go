// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Defender audit suppression.
func (c *Client) DescribeAuditSuppression(ctx context.Context, params *DescribeAuditSuppressionInput, optFns ...func(*Options)) (*DescribeAuditSuppressionOutput, error) {
	if params == nil {
		params = &DescribeAuditSuppressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAuditSuppression", params, optFns, c.addOperationDescribeAuditSuppressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAuditSuppressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAuditSuppressionInput struct {

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	//
	// This member is required.
	CheckName *string

	// Information that identifies the noncompliant resource.
	//
	// This member is required.
	ResourceIdentifier *types.ResourceIdentifier

	noSmithyDocumentSerde
}

type DescribeAuditSuppressionOutput struct {

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	CheckName *string

	//  The description of the audit suppression.
	Description *string

	//  The epoch timestamp in seconds at which this suppression expires.
	ExpirationDate *time.Time

	// Information that identifies the noncompliant resource.
	ResourceIdentifier *types.ResourceIdentifier

	//  Indicates whether a suppression should exist indefinitely or not.
	SuppressIndefinitely *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAuditSuppressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAuditSuppression"); err != nil {
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
	if err = addOpDescribeAuditSuppressionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAuditSuppression(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAuditSuppression(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAuditSuppression",
	}
}
