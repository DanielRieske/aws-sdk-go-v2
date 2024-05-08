// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Migrates a WorkSpace from one operating system or bundle type to another, while
// retaining the data on the user volume.
//
// The migration process recreates the WorkSpace by using a new root volume from
// the target bundle image and the user volume from the last available snapshot of
// the original WorkSpace. During migration, the original D:\Users\%USERNAME% user
// profile folder is renamed to D:\Users\%USERNAME%MMddyyTHHmmss%.NotMigrated . A
// new D:\Users\%USERNAME%\ folder is generated by the new OS. Certain files in
// the old user profile are moved to the new user profile.
//
// For available migration scenarios, details about what happens during migration,
// and best practices, see [Migrate a WorkSpace].
//
// [Migrate a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/migrate-workspaces.html
func (c *Client) MigrateWorkspace(ctx context.Context, params *MigrateWorkspaceInput, optFns ...func(*Options)) (*MigrateWorkspaceOutput, error) {
	if params == nil {
		params = &MigrateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MigrateWorkspace", params, optFns, c.addOperationMigrateWorkspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MigrateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MigrateWorkspaceInput struct {

	// The identifier of the target bundle type to migrate the WorkSpace to.
	//
	// This member is required.
	BundleId *string

	// The identifier of the WorkSpace to migrate from.
	//
	// This member is required.
	SourceWorkspaceId *string

	noSmithyDocumentSerde
}

type MigrateWorkspaceOutput struct {

	// The original identifier of the WorkSpace that is being migrated.
	SourceWorkspaceId *string

	// The new identifier of the WorkSpace that is being migrated. If the migration
	// does not succeed, the target WorkSpace ID will not be used, and the WorkSpace
	// will still have the original WorkSpace ID.
	TargetWorkspaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMigrateWorkspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMigrateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMigrateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "MigrateWorkspace"); err != nil {
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
	if err = addOpMigrateWorkspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMigrateWorkspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMigrateWorkspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "MigrateWorkspace",
	}
}
