// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an existing Amazon Q Business web experience.
func (c *Client) GetWebExperience(ctx context.Context, params *GetWebExperienceInput, optFns ...func(*Options)) (*GetWebExperienceOutput, error) {
	if params == nil {
		params = &GetWebExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWebExperience", params, optFns, c.addOperationGetWebExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWebExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWebExperienceInput struct {

	// The identifier of the Amazon Q Business application linked to the web
	// experience.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the Amazon Q Business web experience.
	//
	// This member is required.
	WebExperienceId *string

	noSmithyDocumentSerde
}

type GetWebExperienceOutput struct {

	// The identifier of the Amazon Q Business application linked to the web
	// experience.
	ApplicationId *string

	// The authentication configuration information for your Amazon Q Business web
	// experience.
	//
	// Deprecated: Property associated with legacy SAML IdP flow. Deprecated in favor
	// of using AWS IAM Identity Center for user management.
	AuthenticationConfiguration types.WebExperienceAuthConfiguration

	// The Unix timestamp when the Amazon Q Business web experience was last created.
	CreatedAt *time.Time

	// The endpoint of your Amazon Q Business web experience.
	DefaultEndpoint *string

	// When the Status field value is FAILED , the ErrorMessage field contains a
	// description of the error that caused the data source connector to fail.
	Error *types.ErrorDetail

	//  The Amazon Resource Name (ARN) of the service role attached to your web
	// experience.
	RoleArn *string

	// Determines whether sample prompts are enabled in the web experience for an end
	// user.
	SamplePromptsControlMode types.WebExperienceSamplePromptsControlMode

	// The current status of the Amazon Q Business web experience. When the Status
	// field value is FAILED , the ErrorMessage field contains a description of the
	// error that caused the data source connector to fail.
	Status types.WebExperienceStatus

	// The subtitle for your Amazon Q Business web experience.
	Subtitle *string

	// The title for your Amazon Q Business web experience.
	Title *string

	// The Unix timestamp when the Amazon Q Business web experience was last updated.
	UpdatedAt *time.Time

	// The Amazon Resource Name (ARN) of the role with the permission to access the
	// Amazon Q Business web experience and required resources.
	WebExperienceArn *string

	// The identifier of the Amazon Q Business web experience.
	WebExperienceId *string

	// The customized welcome message for end users of an Amazon Q Business web
	// experience.
	WelcomeMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWebExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWebExperience"); err != nil {
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
	if err = addOpGetWebExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWebExperience(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWebExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWebExperience",
	}
}
