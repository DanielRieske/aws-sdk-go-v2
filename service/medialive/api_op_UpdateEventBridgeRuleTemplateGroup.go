// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified eventbridge rule template group.
func (c *Client) UpdateEventBridgeRuleTemplateGroup(ctx context.Context, params *UpdateEventBridgeRuleTemplateGroupInput, optFns ...func(*Options)) (*UpdateEventBridgeRuleTemplateGroupOutput, error) {
	if params == nil {
		params = &UpdateEventBridgeRuleTemplateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventBridgeRuleTemplateGroup", params, optFns, c.addOperationUpdateEventBridgeRuleTemplateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventBridgeRuleTemplateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for UpdateEventBridgeRuleTemplateGroupRequest
type UpdateEventBridgeRuleTemplateGroupInput struct {

	// An eventbridge rule template group's identifier. Can be either be its id or
	// current name.
	//
	// This member is required.
	Identifier *string

	// A resource's optional description.
	Description *string

	noSmithyDocumentSerde
}

// Placeholder documentation for UpdateEventBridgeRuleTemplateGroupResponse
type UpdateEventBridgeRuleTemplateGroupOutput struct {

	// An eventbridge rule template group's ARN (Amazon Resource Name)
	Arn *string

	// Placeholder documentation for __timestampIso8601
	CreatedAt *time.Time

	// A resource's optional description.
	Description *string

	// An eventbridge rule template group's id. AWS provided template groups have ids
	// that start with aws-
	Id *string

	// Placeholder documentation for __timestampIso8601
	ModifiedAt *time.Time

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	Name *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventBridgeRuleTemplateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEventBridgeRuleTemplateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEventBridgeRuleTemplateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEventBridgeRuleTemplateGroup"); err != nil {
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
	if err = addOpUpdateEventBridgeRuleTemplateGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventBridgeRuleTemplateGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventBridgeRuleTemplateGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEventBridgeRuleTemplateGroup",
	}
}
