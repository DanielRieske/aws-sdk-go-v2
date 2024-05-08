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

// Updates the specified attributes of the Direct Connect gateway association.
//
// Add or remove prefixes from the association.
func (c *Client) UpdateDirectConnectGatewayAssociation(ctx context.Context, params *UpdateDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*UpdateDirectConnectGatewayAssociationOutput, error) {
	if params == nil {
		params = &UpdateDirectConnectGatewayAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDirectConnectGatewayAssociation", params, optFns, c.addOperationUpdateDirectConnectGatewayAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDirectConnectGatewayAssociationInput struct {

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AddAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.
	RemoveAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	noSmithyDocumentSerde
}

type UpdateDirectConnectGatewayAssociationOutput struct {

	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDirectConnectGatewayAssociation"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDirectConnectGatewayAssociation",
	}
}
