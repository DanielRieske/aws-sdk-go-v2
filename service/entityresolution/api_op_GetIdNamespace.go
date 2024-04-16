// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the IdNamespace with a given name, if it exists.
func (c *Client) GetIdNamespace(ctx context.Context, params *GetIdNamespaceInput, optFns ...func(*Options)) (*GetIdNamespaceOutput, error) {
	if params == nil {
		params = &GetIdNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdNamespace", params, optFns, c.addOperationGetIdNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdNamespaceInput struct {

	// The name of the ID namespace.
	//
	// This member is required.
	IdNamespaceName *string

	noSmithyDocumentSerde
}

type GetIdNamespaceOutput struct {

	// The timestamp of when the ID namespace was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the ID namespace.
	//
	// This member is required.
	IdNamespaceArn *string

	// The name of the ID namespace.
	//
	// This member is required.
	IdNamespaceName *string

	// The type of ID namespace. There are two types: SOURCE and TARGET . The SOURCE
	// contains configurations for sourceId data that will be processed in an ID
	// mapping workflow. The TARGET contains a configuration of targetId to which all
	// sourceIds will resolve to.
	//
	// This member is required.
	Type types.IdNamespaceType

	// The timestamp of when the ID namespace was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The description of the ID namespace.
	Description *string

	// Determines the properties of IdMappingWorkflow where this IdNamespace can be
	// used as a Source or a Target .
	IdMappingWorkflowProperties []types.IdNamespaceIdMappingWorkflowProperties

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	InputSourceConfig []types.IdNamespaceInputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access the resources defined in this IdNamespace on your behalf as part
	// of a workflow run.
	RoleArn *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIdNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIdNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdNamespace"); err != nil {
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
	if err = addOpGetIdNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdNamespace",
	}
}
