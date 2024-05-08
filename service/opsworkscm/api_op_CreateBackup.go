// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates an application-level backup of a server. While the server is in the
//
// BACKING_UP state, the server cannot be changed, and no additional backup can be
// created.
//
// Backups can be created for servers in RUNNING , HEALTHY , and UNHEALTHY states.
// By default, you can create a maximum of 50 manual backups.
//
// This operation is asynchronous.
//
// A LimitExceededException is thrown when the maximum number of manual backups is
// reached. An InvalidStateException is thrown when the server is not in any of
// the following states: RUNNING, HEALTHY, or UNHEALTHY. A
// ResourceNotFoundException is thrown when the server is not found. A
// ValidationException is thrown when parameters of the request are not valid.
func (c *Client) CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error) {
	if params == nil {
		params = &CreateBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackup", params, optFns, c.addOperationCreateBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupInput struct {

	// The name of the server that you want to back up.
	//
	// This member is required.
	ServerName *string

	//  A user-defined description of the backup.
	Description *string

	// A map that contains tag keys and tag values to attach to an AWS OpsWorks-CM
	// server backup.
	//
	//   - The key cannot be empty.
	//
	//   - The key can be a maximum of 127 characters, and can contain only Unicode
	//   letters, numbers, or separators, or the following special characters: + - = .
	//   _ : /
	//
	//   - The value can be a maximum 255 characters, and contain only Unicode
	//   letters, numbers, or separators, or the following special characters: + - = .
	//   _ : /
	//
	//   - Leading and trailing white spaces are trimmed from both the key and value.
	//
	//   - A maximum of 50 user-applied tags is allowed for tag-supported AWS
	//   OpsWorks-CM resources.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateBackupOutput struct {

	// Backup created by request.
	Backup *types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBackup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBackup"); err != nil {
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
	if err = addOpCreateBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBackup",
	}
}
