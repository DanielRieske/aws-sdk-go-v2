// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing report plan identified by its ReportPlanName with the input
// document in JSON format.
func (c *Client) UpdateReportPlan(ctx context.Context, params *UpdateReportPlanInput, optFns ...func(*Options)) (*UpdateReportPlanOutput, error) {
	if params == nil {
		params = &UpdateReportPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReportPlan", params, optFns, c.addOperationUpdateReportPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReportPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReportPlanInput struct {

	// The unique name of the report plan. This name is between 1 and 256 characters,
	// starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and
	// underscores (_).
	//
	// This member is required.
	ReportPlanName *string

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to UpdateReportPlanInput . Retrying a successful request with
	// the same idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	// A structure that contains information about where to deliver your reports,
	// specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your
	// reports.
	ReportDeliveryChannel *types.ReportDeliveryChannel

	// An optional description of the report plan with a maximum 1,024 characters.
	ReportPlanDescription *string

	// Identifies the report template for the report. Reports are built using a report
	// template. The report templates are:
	//
	//     RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT |
	//     COPY_JOB_REPORT | RESTORE_JOB_REPORT
	//
	// If the report template is RESOURCE_COMPLIANCE_REPORT or
	// CONTROL_COMPLIANCE_REPORT , this API resource also describes the report coverage
	// by Amazon Web Services Regions and frameworks.
	ReportSetting *types.ReportSetting

	noSmithyDocumentSerde
}

type UpdateReportPlanOutput struct {

	// The date and time that a report plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationTime is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationTime *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	ReportPlanArn *string

	// The unique name of the report plan.
	ReportPlanName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReportPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReportPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReportPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateReportPlan"); err != nil {
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
	if err = addIdempotencyToken_opUpdateReportPlanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateReportPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReportPlan(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateReportPlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateReportPlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateReportPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateReportPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateReportPlanInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateReportPlanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateReportPlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateReportPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateReportPlan",
	}
}
