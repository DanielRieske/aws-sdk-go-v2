// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A what-if forecast is a forecast that is created from a modified version of the
// baseline forecast. Each what-if forecast incorporates either a replacement
// dataset or a set of transformations to the original dataset.
func (c *Client) CreateWhatIfForecast(ctx context.Context, params *CreateWhatIfForecastInput, optFns ...func(*Options)) (*CreateWhatIfForecastOutput, error) {
	if params == nil {
		params = &CreateWhatIfForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWhatIfForecast", params, optFns, c.addOperationCreateWhatIfForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWhatIfForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWhatIfForecastInput struct {

	// The Amazon Resource Name (ARN) of the what-if analysis.
	//
	// This member is required.
	WhatIfAnalysisArn *string

	// The name of the what-if forecast. Names must be unique within each what-if
	// analysis.
	//
	// This member is required.
	WhatIfForecastName *string

	// A list of [tags] to apply to the what if forecast.
	//
	// [tags]: https://docs.aws.amazon.com/forecast/latest/dg/tagging-forecast-resources.html
	Tags []types.Tag

	// The replacement time series dataset, which contains the rows that you want to
	// change in the related time series dataset. A replacement time series does not
	// need to contain all rows that are in the baseline related time series. Include
	// only the rows (measure-dimension combinations) that you want to include in the
	// what-if forecast.
	//
	// This dataset is merged with the original time series to create a transformed
	// dataset that is used for the what-if analysis.
	//
	// This dataset should contain the items to modify (such as item_id or
	// workforce_type), any relevant dimensions, the timestamp column, and at least one
	// of the related time series columns. This file should not contain duplicate
	// timestamps for the same time series.
	//
	// Timestamps and item_ids not included in this dataset are not included in the
	// what-if analysis.
	TimeSeriesReplacementsDataSource *types.TimeSeriesReplacementsDataSource

	// The transformations that are applied to the baseline time series. Each
	// transformation contains an action and a set of conditions. An action is applied
	// only when all conditions are met. If no conditions are provided, the action is
	// applied to all items.
	TimeSeriesTransformations []types.TimeSeriesTransformation

	noSmithyDocumentSerde
}

type CreateWhatIfForecastOutput struct {

	// The Amazon Resource Name (ARN) of the what-if forecast.
	WhatIfForecastArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWhatIfForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWhatIfForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWhatIfForecast{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWhatIfForecast"); err != nil {
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
	if err = addOpCreateWhatIfForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWhatIfForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWhatIfForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWhatIfForecast",
	}
}
