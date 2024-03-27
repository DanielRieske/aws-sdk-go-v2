// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a knowledge base that contains data sources from which information can
// be queried and used by LLMs. To create a knowledge base, you must first set up
// your data sources and configure a supported vector store. For more information,
// see Set up your data for ingestion (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup.html)
// . If you prefer to let Amazon Bedrock create and manage a vector store for you
// in Amazon OpenSearch Service, use the console. For more information, see Create
// a knowledge base (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create)
// .
//   - Provide the name and an optional description .
//   - Provide the ARN with permissions to create a knowledge base in the roleArn
//     field.
//   - Provide the embedding model to use in the embeddingModelArn field in the
//     knowledgeBaseConfiguration object.
//   - Provide the configuration for your vector store in the storageConfiguration
//     object.
//   - For an Amazon OpenSearch Service database, use the
//     opensearchServerlessConfiguration object. For more information, see Create a
//     vector store in Amazon OpenSearch Service (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-oss.html)
//     .
//   - For an Amazon Aurora database, use the RdsConfiguration object. For more
//     information, see Create a vector store in Amazon Aurora (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-rds.html)
//     .
//   - For a Pinecone database, use the pineconeConfiguration object. For more
//     information, see Create a vector store in Pinecone (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-pinecone.html)
//     .
//   - For a Redis Enterprise Cloud database, use the
//     redisEnterpriseCloudConfiguration object. For more information, see Create a
//     vector store in Redis Enterprise Cloud (https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-redis.html)
//     .
func (c *Client) CreateKnowledgeBase(ctx context.Context, params *CreateKnowledgeBaseInput, optFns ...func(*Options)) (*CreateKnowledgeBaseOutput, error) {
	if params == nil {
		params = &CreateKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKnowledgeBase", params, optFns, c.addOperationCreateKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKnowledgeBaseInput struct {

	// Contains details about the embeddings model used for the knowledge base.
	//
	// This member is required.
	KnowledgeBaseConfiguration *types.KnowledgeBaseConfiguration

	// A name for the knowledge base.
	//
	// This member is required.
	Name *string

	// The ARN of the IAM role with permissions to create the knowledge base.
	//
	// This member is required.
	RoleArn *string

	// Contains details about the configuration of the vector database used for the
	// knowledge base.
	//
	// This member is required.
	StorageConfiguration *types.StorageConfiguration

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see
	// Ensuring idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// A description of the knowledge base.
	Description *string

	// Specify the key-value pairs for the tags that you want to attach to your
	// knowledge base in this object.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKnowledgeBaseOutput struct {

	// Contains details about the knowledge base.
	//
	// This member is required.
	KnowledgeBase *types.KnowledgeBase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKnowledgeBase"); err != nil {
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
	if err = addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKnowledgeBase(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKnowledgeBase struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKnowledgeBase) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKnowledgeBase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKnowledgeBaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKnowledgeBaseInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKnowledgeBase{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKnowledgeBase",
	}
}
