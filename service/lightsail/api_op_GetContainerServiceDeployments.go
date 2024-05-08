// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the deployments for your Amazon Lightsail container service
//
// A deployment specifies the settings, such as the ports and launch command, of
// containers that are deployed to your container service.
//
// The deployments are ordered by version in ascending order. The newest version
// is listed at the top of the response.
//
// A set number of deployments are kept before the oldest one is replaced with the
// newest one. For more information, see [Amazon Lightsail endpoints and quotas]in the Amazon Web Services General
// Reference.
//
// [Amazon Lightsail endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/lightsail.html
func (c *Client) GetContainerServiceDeployments(ctx context.Context, params *GetContainerServiceDeploymentsInput, optFns ...func(*Options)) (*GetContainerServiceDeploymentsOutput, error) {
	if params == nil {
		params = &GetContainerServiceDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContainerServiceDeployments", params, optFns, c.addOperationGetContainerServiceDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContainerServiceDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContainerServiceDeploymentsInput struct {

	// The name of the container service for which to return deployments.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type GetContainerServiceDeploymentsOutput struct {

	// An array of objects that describe deployments for a container service.
	Deployments []types.ContainerServiceDeployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContainerServiceDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContainerServiceDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContainerServiceDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetContainerServiceDeployments"); err != nil {
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
	if err = addOpGetContainerServiceDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContainerServiceDeployments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContainerServiceDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetContainerServiceDeployments",
	}
}
