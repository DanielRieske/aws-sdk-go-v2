// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the contents of the encrypted fields SecretString or SecretBinary for
// up to 20 secrets. To retrieve a single secret, call GetSecretValue.
//
// To choose which secrets to retrieve, you can specify a list of secrets by name
// or ARN, or you can use filters. If Secrets Manager encounters errors such as
// AccessDeniedException while attempting to retrieve any of the secrets, you can
// see the errors in Errors in the response.
//
// Secrets Manager generates CloudTrail GetSecretValue log entries for each secret
// you request when you call this action. Do not include sensitive information in
// request parameters because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:BatchGetSecretValue , and you must have
// secretsmanager:GetSecretValue for each secret. If you use filters, you must also
// have secretsmanager:ListSecrets . If the secrets are encrypted using
// customer-managed keys instead of the Amazon Web Services managed key
// aws/secretsmanager , then you also need kms:Decrypt permissions for the keys.
// For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func (c *Client) BatchGetSecretValue(ctx context.Context, params *BatchGetSecretValueInput, optFns ...func(*Options)) (*BatchGetSecretValueOutput, error) {
	if params == nil {
		params = &BatchGetSecretValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetSecretValue", params, optFns, c.addOperationBatchGetSecretValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetSecretValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetSecretValueInput struct {

	// The filters to choose which secrets to retrieve. You must include Filters or
	// SecretIdList , but not both.
	Filters []types.Filter

	// The number of results to include in the response.
	//
	// If there are more results available, in the response, Secrets Manager includes
	// NextToken . To get the next results, call BatchGetSecretValue again with the
	// value from NextToken . To use this parameter, you must also use the Filters
	// parameter.
	MaxResults *int32

	// A token that indicates where the output should continue from, if a previous
	// call did not show all results. To get the next results, call BatchGetSecretValue
	// again with this value.
	NextToken *string

	// The ARN or names of the secrets to retrieve. You must include Filters or
	// SecretIdList , but not both.
	SecretIdList []string

	noSmithyDocumentSerde
}

type BatchGetSecretValueOutput struct {

	// A list of errors Secrets Manager encountered while attempting to retrieve
	// individual secrets.
	Errors []types.APIErrorType

	// Secrets Manager includes this value if there's more output available than what
	// is included in the current response. This can occur even when the response
	// includes no values at all, such as when you ask for a filtered view of a long
	// list. To get the next results, call BatchGetSecretValue again with this value.
	NextToken *string

	// A list of secret values.
	SecretValues []types.SecretValueEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetSecretValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetSecretValue"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetSecretValue(options.Region), middleware.Before); err != nil {
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

// BatchGetSecretValueAPIClient is a client that implements the
// BatchGetSecretValue operation.
type BatchGetSecretValueAPIClient interface {
	BatchGetSecretValue(context.Context, *BatchGetSecretValueInput, ...func(*Options)) (*BatchGetSecretValueOutput, error)
}

var _ BatchGetSecretValueAPIClient = (*Client)(nil)

// BatchGetSecretValuePaginatorOptions is the paginator options for
// BatchGetSecretValue
type BatchGetSecretValuePaginatorOptions struct {
	// The number of results to include in the response.
	//
	// If there are more results available, in the response, Secrets Manager includes
	// NextToken . To get the next results, call BatchGetSecretValue again with the
	// value from NextToken . To use this parameter, you must also use the Filters
	// parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// BatchGetSecretValuePaginator is a paginator for BatchGetSecretValue
type BatchGetSecretValuePaginator struct {
	options   BatchGetSecretValuePaginatorOptions
	client    BatchGetSecretValueAPIClient
	params    *BatchGetSecretValueInput
	nextToken *string
	firstPage bool
}

// NewBatchGetSecretValuePaginator returns a new BatchGetSecretValuePaginator
func NewBatchGetSecretValuePaginator(client BatchGetSecretValueAPIClient, params *BatchGetSecretValueInput, optFns ...func(*BatchGetSecretValuePaginatorOptions)) *BatchGetSecretValuePaginator {
	if params == nil {
		params = &BatchGetSecretValueInput{}
	}

	options := BatchGetSecretValuePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &BatchGetSecretValuePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *BatchGetSecretValuePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next BatchGetSecretValue page.
func (p *BatchGetSecretValuePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*BatchGetSecretValueOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.BatchGetSecretValue(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opBatchGetSecretValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetSecretValue",
	}
}
