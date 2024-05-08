// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified scaling plan.
//
// You cannot update a scaling plan if it is in the process of being created,
// updated, or deleted.
func (c *Client) UpdateScalingPlan(ctx context.Context, params *UpdateScalingPlanInput, optFns ...func(*Options)) (*UpdateScalingPlanOutput, error) {
	if params == nil {
		params = &UpdateScalingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScalingPlan", params, optFns, c.addOperationUpdateScalingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScalingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScalingPlanInput struct {

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan. The only valid value is 1 . Currently,
	// you cannot have multiple scaling plan versions.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// A CloudFormation stack or set of tags.
	//
	// For more information, see [ApplicationSource] in the AWS Auto Scaling API Reference.
	//
	// [ApplicationSource]: https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ApplicationSource.html
	ApplicationSource *types.ApplicationSource

	// The scaling instructions.
	//
	// For more information, see [ScalingInstruction] in the AWS Auto Scaling API Reference.
	//
	// [ScalingInstruction]: https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html
	ScalingInstructions []types.ScalingInstruction

	noSmithyDocumentSerde
}

type UpdateScalingPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateScalingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateScalingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateScalingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateScalingPlan"); err != nil {
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
	if err = addOpUpdateScalingPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScalingPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateScalingPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateScalingPlan",
	}
}
