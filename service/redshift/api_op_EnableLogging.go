// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func (c *Client) EnableLogging(ctx context.Context, params *EnableLoggingInput, optFns ...func(*Options)) (*EnableLoggingOutput, error) {
	if params == nil {
		params = &EnableLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableLogging", params, optFns, c.addOperationEnableLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableLoggingInput struct {

	// The identifier of the cluster on which logging is to be started.
	//
	// Example: examplecluster
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of an existing S3 bucket where the log files are to be stored.
	//
	// Constraints:
	//
	//   - Must be in the same region as the cluster
	//
	//   - The cluster must have read bucket and put object permissions
	BucketName *string

	// The log destination type. An enum with possible values of s3 and cloudwatch .
	LogDestinationType types.LogDestinationType

	// The collection of exported log types. Possible values are connectionlog ,
	// useractivitylog , and userlog .
	LogExports []string

	// The prefix applied to the log file names.
	//
	// Constraints:
	//
	//   - Cannot exceed 512 characters
	//
	//   - Cannot contain spaces( ), double quotes ("), single quotes ('), a backslash
	//   (\), or control characters. The hexadecimal codes for invalid characters are:
	//
	//   - x00 to x20
	//
	//   - x22
	//
	//   - x27
	//
	//   - x5c
	//
	//   - x7f or larger
	S3KeyPrefix *string

	noSmithyDocumentSerde
}

// Describes the status of logging for a cluster.
type EnableLoggingOutput struct {

	// The name of the S3 bucket where the log files are stored.
	BucketName *string

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// The log destination type. An enum with possible values of s3 and cloudwatch .
	LogDestinationType types.LogDestinationType

	// The collection of exported log types. Possible values are connectionlog ,
	// useractivitylog , and userlog .
	LogExports []string

	// true if logging is on, false if logging is off.
	LoggingEnabled *bool

	// The prefix applied to the log file names.
	S3KeyPrefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableLogging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableLogging{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableLogging"); err != nil {
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
	if err = addOpEnableLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableLogging(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableLogging",
	}
}
