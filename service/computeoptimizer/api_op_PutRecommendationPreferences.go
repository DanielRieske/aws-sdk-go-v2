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

// Creates a new recommendation preference or updates an existing recommendation
// preference, such as enhanced infrastructure metrics.
//
// For more information, see [Activating enhanced infrastructure metrics] in the Compute Optimizer User Guide.
//
// [Activating enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
func (c *Client) PutRecommendationPreferences(ctx context.Context, params *PutRecommendationPreferencesInput, optFns ...func(*Options)) (*PutRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &PutRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecommendationPreferences", params, optFns, c.addOperationPutRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecommendationPreferencesInput struct {

	// The target resource type of the recommendation preference to create.
	//
	// The Ec2Instance option encompasses standalone instances and instances that are
	// part of Auto Scaling groups. The AutoScalingGroup option encompasses only
	// instances that are part of an Auto Scaling group.
	//
	// The valid values for this parameter are Ec2Instance and AutoScalingGroup .
	//
	// This member is required.
	ResourceType types.ResourceType

	// The status of the enhanced infrastructure metrics recommendation preference to
	// create or update.
	//
	// Specify the Active status to activate the preference, or specify Inactive to
	// deactivate the preference.
	//
	// For more information, see [Enhanced infrastructure metrics] in the Compute Optimizer User Guide.
	//
	// [Enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
	EnhancedInfrastructureMetrics types.EnhancedInfrastructureMetrics

	// The provider of the external metrics recommendation preference to create or
	// update.
	//
	// Specify a valid provider in the source field to activate the preference. To
	// delete this preference, see the DeleteRecommendationPreferencesaction.
	//
	// This preference can only be set for the Ec2Instance resource type.
	//
	// For more information, see [External metrics ingestion] in the Compute Optimizer User Guide.
	//
	// [External metrics ingestion]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/external-metrics-ingestion.html
	ExternalMetricsPreference *types.ExternalMetricsPreference

	// The status of the inferred workload types recommendation preference to create
	// or update.
	//
	// The inferred workload type feature is active by default. To deactivate it,
	// create a recommendation preference.
	//
	// Specify the Inactive status to deactivate the feature, or specify Active to
	// activate it.
	//
	// For more information, see [Inferred workload types] in the Compute Optimizer User Guide.
	//
	// [Inferred workload types]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/inferred-workload-types.html
	InferredWorkloadTypes types.InferredWorkloadTypesPreference

	//  The preference to control the number of days the utilization metrics of the
	// Amazon Web Services resource are analyzed. When this preference isn't specified,
	// we use the default value DAYS_14 .
	//
	// You can only set this preference for the Amazon EC2 instance and Auto Scaling
	// group resource types.
	LookBackPeriod types.LookBackPeriodPreference

	//  The preference to control which resource type values are considered when
	// generating rightsizing recommendations. You can specify this preference as a
	// combination of include and exclude lists. You must specify either an includeList
	// or excludeList . If the preference is an empty set of resource type values, an
	// error occurs.
	//
	// You can only set this preference for the Amazon EC2 instance and Auto Scaling
	// group resource types.
	PreferredResources []types.PreferredResource

	//  The status of the savings estimation mode preference to create or update.
	//
	// Specify the AfterDiscounts status to activate the preference, or specify
	// BeforeDiscounts to deactivate the preference.
	//
	// Only the account manager or delegated administrator of your organization can
	// activate this preference.
	//
	// For more information, see [Savings estimation mode] in the Compute Optimizer User Guide.
	//
	// [Savings estimation mode]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/savings-estimation-mode.html
	SavingsEstimationMode types.SavingsEstimationMode

	// An object that describes the scope of the recommendation preference to create.
	//
	// You can create recommendation preferences at the organization level (for
	// management accounts of an organization only), account level, and resource level.
	// For more information, see [Activating enhanced infrastructure metrics]in the Compute Optimizer User Guide.
	//
	// You cannot create recommendation preferences for Auto Scaling groups at the
	// organization and account levels. You can create recommendation preferences for
	// Auto Scaling groups only at the resource level by specifying a scope name of
	// ResourceArn and a scope value of the Auto Scaling group Amazon Resource Name
	// (ARN). This will configure the preference for all instances that are part of the
	// specified Auto Scaling group. You also cannot create recommendation preferences
	// at the resource level for instances that are part of an Auto Scaling group. You
	// can create recommendation preferences at the resource level only for standalone
	// instances.
	//
	// [Activating enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
	Scope *types.Scope

	//  The preference to control the resource’s CPU utilization threshold, CPU
	// utilization headroom, and memory utilization headroom. When this preference
	// isn't specified, we use the following default values.
	//
	// CPU utilization:
	//
	//   - P99_5 for threshold
	//
	//   - PERCENT_20 for headroom
	//
	// Memory utilization:
	//
	//   - PERCENT_20 for headroom
	//
	//   - You can only set CPU and memory utilization preferences for the Amazon EC2
	//   instance resource type.
	//
	//   - The threshold setting isn’t available for memory utilization.
	UtilizationPreferences []types.UtilizationPreference

	noSmithyDocumentSerde
}

type PutRecommendationPreferencesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRecommendationPreferences"); err != nil {
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
	if err = addOpPutRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRecommendationPreferences",
	}
}
