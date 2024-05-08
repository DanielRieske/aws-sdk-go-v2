// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the association between the specified Direct Connect gateway and
// virtual private gateway.
//
// We recommend that you specify the associationID to delete the association.
// Alternatively, if you own virtual gateway and a Direct Connect gateway
// association, you can specify the virtualGatewayId and directConnectGatewayId to
// delete an association.
func (c *Client) DeleteDirectConnectGatewayAssociation(ctx context.Context, params *DeleteDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayAssociationOutput, error) {
	if params == nil {
		params = &DeleteDirectConnectGatewayAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDirectConnectGatewayAssociation", params, optFns, c.addOperationDeleteDirectConnectGatewayAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayAssociationInput struct {

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The ID of the virtual private gateway.
	VirtualGatewayId *string

	noSmithyDocumentSerde
}

type DeleteDirectConnectGatewayAssociationOutput struct {

	// Information about the deleted association.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDirectConnectGatewayAssociation"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDirectConnectGatewayAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDirectConnectGatewayAssociation",
	}
}
