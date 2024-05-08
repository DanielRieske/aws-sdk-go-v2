// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables Health to work with Organizations. You can use the organizational view
// feature to aggregate events from all Amazon Web Services accounts in your
// organization in a centralized location.
//
// This operation also creates a service-linked role for the management account in
// the organization.
//
// To call this operation, you must meet the following requirements:
//
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan
//     from [Amazon Web Services Support]to use the Health API. If you call the Health API from an Amazon Web
//     Services account that doesn't have a Business, Enterprise On-Ramp, or Enterprise
//     Support plan, you receive a SubscriptionRequiredException error.
//
//   - You must have permission to call this operation from the organization's
//     management account. For example IAM policies, see [Health identity-based policy examples].
//
// If you don't have the required support plan, you can instead use the Health
// console to enable the organizational view feature. For more information, see [Aggregating Health events]in
// the Health User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Aggregating Health events]: https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html
// [Health identity-based policy examples]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html
func (c *Client) EnableHealthServiceAccessForOrganization(ctx context.Context, params *EnableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*EnableHealthServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &EnableHealthServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableHealthServiceAccessForOrganization", params, optFns, c.addOperationEnableHealthServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableHealthServiceAccessForOrganizationInput struct {
	noSmithyDocumentSerde
}

type EnableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableHealthServiceAccessForOrganization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableHealthServiceAccessForOrganization",
	}
}
