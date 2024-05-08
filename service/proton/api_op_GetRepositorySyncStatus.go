// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the sync status of a repository used for Proton template sync. For more
// information about template sync, see .
//
// A repository sync status isn't tied to the Proton Repository resource (or any
// other Proton resource). Therefore, tags on an Proton Repository resource have no
// effect on this action. Specifically, you can't use these tags to control access
// to this action using Attribute-based access control (ABAC).
//
// For more information about ABAC, see [ABAC] in the Proton User Guide.
//
// [ABAC]: https://docs.aws.amazon.com/proton/latest/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-tags
func (c *Client) GetRepositorySyncStatus(ctx context.Context, params *GetRepositorySyncStatusInput, optFns ...func(*Options)) (*GetRepositorySyncStatusOutput, error) {
	if params == nil {
		params = &GetRepositorySyncStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepositorySyncStatus", params, optFns, c.addOperationGetRepositorySyncStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositorySyncStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRepositorySyncStatusInput struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The repository name.
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The repository sync type.
	//
	// This member is required.
	SyncType types.SyncType

	noSmithyDocumentSerde
}

type GetRepositorySyncStatusOutput struct {

	// The repository sync status detail data that's returned by Proton.
	LatestSync *types.RepositorySyncAttempt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRepositorySyncStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRepositorySyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRepositorySyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRepositorySyncStatus"); err != nil {
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
	if err = addOpGetRepositorySyncStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositorySyncStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRepositorySyncStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRepositorySyncStatus",
	}
}
