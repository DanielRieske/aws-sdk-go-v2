// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stop an already running Bot Recommendation request.
func (c *Client) StopBotRecommendation(ctx context.Context, params *StopBotRecommendationInput, optFns ...func(*Options)) (*StopBotRecommendationOutput, error) {
	if params == nil {
		params = &StopBotRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopBotRecommendation", params, optFns, c.addOperationStopBotRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopBotRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopBotRecommendationInput struct {

	// The unique identifier of the bot containing the bot recommendation to be
	// stopped.
	//
	// This member is required.
	BotId *string

	// The unique identifier of the bot recommendation to be stopped.
	//
	// This member is required.
	BotRecommendationId *string

	// The version of the bot containing the bot recommendation.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the bot recommendation to stop.
	// The string must match one of the supported locales. For more information, see [Supported languages]
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type StopBotRecommendationOutput struct {

	// The unique identifier of the bot containing the bot recommendation that is
	// being stopped.
	BotId *string

	// The unique identifier of the bot recommendation that is being stopped.
	BotRecommendationId *string

	// The status of the bot recommendation. If the status is Failed, then the reasons
	// for the failure are listed in the failureReasons field.
	BotRecommendationStatus types.BotRecommendationStatus

	// The version of the bot containing the recommendation that is being stopped.
	BotVersion *string

	// The identifier of the language and locale of the bot response to stop. The
	// string must match one of the supported locales. For more information, see [Supported languages]
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	LocaleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopBotRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopBotRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopBotRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopBotRecommendation"); err != nil {
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
	if err = addOpStopBotRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopBotRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopBotRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopBotRecommendation",
	}
}
