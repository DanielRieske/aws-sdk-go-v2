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

// Gets summary information about a domain configuration.
//
// Requires permission to access the [DescribeDomainConfiguration] action.
//
// [DescribeDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) DescribeDomainConfiguration(ctx context.Context, params *DescribeDomainConfigurationInput, optFns ...func(*Options)) (*DescribeDomainConfigurationOutput, error) {
	if params == nil {
		params = &DescribeDomainConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainConfiguration", params, optFns, c.addOperationDescribeDomainConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainConfigurationInput struct {

	// The name of the domain configuration.
	//
	// This member is required.
	DomainConfigurationName *string

	noSmithyDocumentSerde
}

type DescribeDomainConfigurationOutput struct {

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *types.AuthorizerConfig

	// The ARN of the domain configuration.
	DomainConfigurationArn *string

	// The name of the domain configuration.
	DomainConfigurationName *string

	// A Boolean value that specifies the current state of the domain configuration.
	DomainConfigurationStatus types.DomainConfigurationStatus

	// The name of the domain.
	DomainName *string

	// The type of the domain.
	DomainType types.DomainType

	// The date and time the domain configuration's status was last changed.
	LastStatusChangeDate *time.Time

	// The server certificate configuration.
	ServerCertificateConfig *types.ServerCertificateConfig

	// A list containing summary information about the server certificate included in
	// the domain configuration.
	ServerCertificates []types.ServerCertificateSummary

	// The type of service delivered by the endpoint.
	ServiceType types.ServiceType

	// An object that specifies the TLS configuration for a domain.
	TlsConfig *types.TlsConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDomainConfiguration"); err != nil {
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
	if err = addOpDescribeDomainConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomainConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDomainConfiguration",
	}
}
