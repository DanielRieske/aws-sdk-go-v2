// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports optimization recommendations for Amazon EC2 instances.
//
// Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (JSON) (.json) file, to an existing
// Amazon Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Amazon EC2 instance export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func (c *Client) ExportEC2InstanceRecommendations(ctx context.Context, params *ExportEC2InstanceRecommendationsInput, optFns ...func(*Options)) (*ExportEC2InstanceRecommendationsOutput, error) {
	if params == nil {
		params = &ExportEC2InstanceRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportEC2InstanceRecommendations", params, optFns, c.addOperationExportEC2InstanceRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportEC2InstanceRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportEC2InstanceRecommendationsInput struct {

	// An object to specify the destination Amazon Simple Storage Service (Amazon S3)
	// bucket name and key prefix for the export job.
	//
	// You must create the destination Amazon S3 bucket for your recommendations
	// export before you create the export job. Compute Optimizer does not create the
	// S3 bucket for you. After you create the S3 bucket, ensure that it has the
	// required permissions policy to allow Compute Optimizer to write the export file
	// to it. If you plan to specify an object prefix when you create the export job,
	// you must include the object prefix in the policy that you add to the S3 bucket.
	// For more information, see [Amazon S3 Bucket Policy for Compute Optimizer]in the Compute Optimizer User Guide.
	//
	// [Amazon S3 Bucket Policy for Compute Optimizer]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html
	//
	// This member is required.
	S3DestinationConfig *types.S3DestinationConfig

	// The IDs of the Amazon Web Services accounts for which to export instance
	// recommendations.
	//
	// If your account is the management account of an organization, use this
	// parameter to specify the member account for which you want to export
	// recommendations.
	//
	// This parameter cannot be specified together with the include member accounts
	// parameter. The parameters are mutually exclusive.
	//
	// Recommendations for member accounts are not included in the export if this
	// parameter, or the include member accounts parameter, is omitted.
	//
	// You can specify multiple account IDs per request.
	AccountIds []string

	// The recommendations data to include in the export file. For more information
	// about the fields that can be exported, see [Exported files]in the Compute Optimizer User Guide.
	//
	// [Exported files]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html#exported-files
	FieldsToExport []types.ExportableInstanceField

	// The format of the export file.
	//
	// The only export file format currently supported is Csv .
	FileFormat types.FileFormat

	// An array of objects to specify a filter that exports a more specific set of
	// instance recommendations.
	Filters []types.Filter

	// Indicates whether to include recommendations for resources in all member
	// accounts of the organization if your account is the management account of an
	// organization.
	//
	// The member accounts must also be opted in to Compute Optimizer, and trusted
	// access for Compute Optimizer must be enabled in the organization account. For
	// more information, see [Compute Optimizer and Amazon Web Services Organizations trusted access]in the Compute Optimizer User Guide.
	//
	// Recommendations for member accounts of the organization are not included in the
	// export file if this parameter is omitted.
	//
	// Recommendations for member accounts are not included in the export if this
	// parameter, or the account IDs parameter, is omitted.
	//
	// [Compute Optimizer and Amazon Web Services Organizations trusted access]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html#trusted-service-access
	IncludeMemberAccounts bool

	// An object to specify the preferences for the Amazon EC2 instance
	// recommendations to export.
	RecommendationPreferences *types.RecommendationPreferences

	noSmithyDocumentSerde
}

type ExportEC2InstanceRecommendationsOutput struct {

	// The identification number of the export job.
	//
	// Use the DescribeRecommendationExportJobs action, and specify the job ID to view the status of an export job.
	JobId *string

	// An object that describes the destination Amazon S3 bucket of a recommendations
	// export file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportEC2InstanceRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportEC2InstanceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportEC2InstanceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportEC2InstanceRecommendations"); err != nil {
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
	if err = addOpExportEC2InstanceRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportEC2InstanceRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportEC2InstanceRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportEC2InstanceRecommendations",
	}
}
