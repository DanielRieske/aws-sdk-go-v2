// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Promotes an AppInstanceUser or AppInstanceBot to an AppInstanceAdmin . The
// promoted entity can perform the following actions.
//
//   - ChannelModerator actions across all channels in the AppInstance .
//
//   - DeleteChannelMessage actions.
//
// Only an AppInstanceUser and AppInstanceBot can be promoted to an
// AppInstanceAdmin role.
func (c *Client) CreateAppInstanceAdmin(ctx context.Context, params *CreateAppInstanceAdminInput, optFns ...func(*Options)) (*CreateAppInstanceAdminOutput, error) {
	if params == nil {
		params = &CreateAppInstanceAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppInstanceAdmin", params, optFns, c.addOperationCreateAppInstanceAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppInstanceAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInstanceAdminInput struct {

	// The ARN of the administrator of the current AppInstance .
	//
	// This member is required.
	AppInstanceAdminArn *string

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	noSmithyDocumentSerde
}

type CreateAppInstanceAdminOutput struct {

	// The ARN and name of the administrator, the ARN of the AppInstance , and the
	// created and last-updated timestamps. All timestamps use epoch milliseconds.
	AppInstanceAdmin *types.Identity

	// The ARN of the of the admin for the AppInstance .
	AppInstanceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppInstanceAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAppInstanceAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAppInstanceAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppInstanceAdmin"); err != nil {
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
	if err = addOpCreateAppInstanceAdminValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppInstanceAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppInstanceAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppInstanceAdmin",
	}
}
