// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an existing workload.
func (c *Client) UpdateWorkload(ctx context.Context, params *UpdateWorkloadInput, optFns ...func(*Options)) (*UpdateWorkloadOutput, error) {
	if params == nil {
		params = &UpdateWorkloadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkload", params, optFns, c.addOperationUpdateWorkloadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to update a workload.
type UpdateWorkloadInput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The list of Amazon Web Services account IDs associated with the workload.
	AccountIds []string

	// List of AppRegistry application ARNs to associate to the workload.
	Applications []string

	// The URL of the architectural design for the workload.
	ArchitecturalDesign *string

	// The list of Amazon Web Services Regions associated with the workload, for
	// example, us-east-2 , or ca-central-1 .
	AwsRegions []string

	// The description for the workload.
	Description *string

	// Well-Architected discovery configuration settings to associate to the workload.
	DiscoveryConfig *types.WorkloadDiscoveryConfig

	// The environment for the workload.
	Environment types.WorkloadEnvironment

	// The improvement status for a workload.
	ImprovementStatus types.WorkloadImprovementStatus

	// The industry for the workload.
	Industry *string

	// The industry type for the workload. If specified, must be one of the following:
	//   - Agriculture
	//   - Automobile
	//   - Defense
	//   - Design and Engineering
	//   - Digital Advertising
	//   - Education
	//   - Environmental Protection
	//   - Financial Services
	//   - Gaming
	//   - General Public Services
	//   - Healthcare
	//   - Hospitality
	//   - InfoTech
	//   - Justice and Public Safety
	//   - Life Sciences
	//   - Manufacturing
	//   - Media & Entertainment
	//   - Mining & Resources
	//   - Oil & Gas
	//   - Power & Utilities
	//   - Professional Services
	//   - Real Estate & Construction
	//   - Retail & Wholesale
	//   - Social Protection
	//   - Telecommunications
	//   - Travel, Transportation & Logistics
	//   - Other
	IndustryType *string

	// Flag indicating whether the workload owner has acknowledged that the Review
	// owner field is required. If a Review owner is not added to the workload within
	// 60 days of acknowledgement, access to the workload is restricted until an owner
	// is added.
	IsReviewOwnerUpdateAcknowledged *bool

	// Configuration of the Jira integration.
	JiraConfiguration *types.WorkloadJiraConfigurationInput

	// The list of non-Amazon Web Services Regions associated with the workload.
	NonAwsRegions []string

	// The notes associated with the workload. For a review template, these are the
	// notes that will be associated with the workload when the template is applied.
	Notes *string

	// The priorities of the pillars, which are used to order items in the improvement
	// plan. Each pillar is represented by its PillarReviewSummary$PillarId .
	PillarPriorities []string

	// The review owner of the workload. The name, email address, or identifier for
	// the primary group or individual that owns the workload review process.
	ReviewOwner *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// Output of an update workload call.
type UpdateWorkloadOutput struct {

	// A workload return object.
	Workload *types.Workload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWorkload"); err != nil {
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
	if err = addOpUpdateWorkloadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWorkload",
	}
}
