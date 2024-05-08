// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an analysis from Amazon QuickSight. You can optionally include a
// recovery window during which you can restore the analysis. If you don't specify
// a recovery window value, the operation defaults to 30 days. Amazon QuickSight
// attaches a DeletionTime stamp to the response that specifies the end of the
// recovery window. At the end of the recovery window, Amazon QuickSight deletes
// the analysis permanently.
//
// At any time before recovery window ends, you can use the RestoreAnalysis API
// operation to remove the DeletionTime stamp and cancel the deletion of the
// analysis. The analysis remains visible in the API until it's deleted, so you can
// describe it but you can't make a template from it.
//
// An analysis that's scheduled for deletion isn't accessible in the Amazon
// QuickSight console. To access it in the console, restore it. Deleting an
// analysis doesn't delete the dashboards that you publish from it.
func (c *Client) DeleteAnalysis(ctx context.Context, params *DeleteAnalysisInput, optFns ...func(*Options)) (*DeleteAnalysisOutput, error) {
	if params == nil {
		params = &DeleteAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnalysis", params, optFns, c.addOperationDeleteAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAnalysisInput struct {

	// The ID of the analysis that you're deleting.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the Amazon Web Services account where you want to delete an analysis.
	//
	// This member is required.
	AwsAccountId *string

	// This option defaults to the value NoForceDeleteWithoutRecovery . To immediately
	// delete the analysis, add the ForceDeleteWithoutRecovery option. You can't
	// restore an analysis after it's deleted.
	ForceDeleteWithoutRecovery bool

	// A value that specifies the number of days that Amazon QuickSight waits before
	// it deletes the analysis. You can't use this parameter with the
	// ForceDeleteWithoutRecovery option in the same API call. The default value is 30.
	RecoveryWindowInDays *int64

	noSmithyDocumentSerde
}

type DeleteAnalysisOutput struct {

	// The ID of the deleted analysis.
	AnalysisId *string

	// The Amazon Resource Name (ARN) of the deleted analysis.
	Arn *string

	// The date and time that the analysis is scheduled to be deleted.
	DeletionTime *time.Time

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAnalysis"); err != nil {
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
	if err = addOpDeleteAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAnalysis",
	}
}
