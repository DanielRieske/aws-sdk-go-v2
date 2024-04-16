// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the status, metrics, and errors (if there are any) that are associated
// with a job.
func (c *Client) GetIdMappingJob(ctx context.Context, params *GetIdMappingJobInput, optFns ...func(*Options)) (*GetIdMappingJobOutput, error) {
	if params == nil {
		params = &GetIdMappingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdMappingJob", params, optFns, c.addOperationGetIdMappingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdMappingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdMappingJobInput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

type GetIdMappingJobOutput struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The time at which the job was started.
	//
	// This member is required.
	StartTime *time.Time

	// The current status of the job.
	//
	// This member is required.
	Status types.JobStatus

	// The time at which the job has finished.
	EndTime *time.Time

	// An object containing an error message, if there was an error.
	ErrorDetails *types.ErrorDetails

	// Metrics associated with the execution, specifically total records processed,
	// unique IDs generated, and records the execution skipped.
	Metrics *types.IdMappingJobMetrics

	// A list of OutputSource objects.
	OutputSourceConfig []types.IdMappingJobOutputSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdMappingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIdMappingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIdMappingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdMappingJob"); err != nil {
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
	if err = addOpGetIdMappingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdMappingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdMappingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdMappingJob",
	}
}
