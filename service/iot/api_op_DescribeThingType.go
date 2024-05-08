// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the specified thing type.
//
// Requires permission to access the [DescribeThingType] action.
//
// [DescribeThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) DescribeThingType(ctx context.Context, params *DescribeThingTypeInput, optFns ...func(*Options)) (*DescribeThingTypeOutput, error) {
	if params == nil {
		params = &DescribeThingTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeThingType", params, optFns, c.addOperationDescribeThingTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeThingTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeThingType operation.
type DescribeThingTypeInput struct {

	// The name of the thing type.
	//
	// This member is required.
	ThingTypeName *string

	noSmithyDocumentSerde
}

// The output for the DescribeThingType operation.
type DescribeThingTypeOutput struct {

	// The thing type ARN.
	ThingTypeArn *string

	// The thing type ID.
	ThingTypeId *string

	// The ThingTypeMetadata contains additional information about the thing type
	// including: creation date and time, a value indicating whether the thing type is
	// deprecated, and a date and time when it was deprecated.
	ThingTypeMetadata *types.ThingTypeMetadata

	// The name of the thing type.
	ThingTypeName *string

	// The ThingTypeProperties contains information about the thing type including
	// description, and a list of searchable thing attribute names.
	ThingTypeProperties *types.ThingTypeProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeThingTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThingType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThingType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeThingType"); err != nil {
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
	if err = addOpDescribeThingTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThingType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeThingType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeThingType",
	}
}
