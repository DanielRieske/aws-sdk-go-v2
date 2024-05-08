// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates the syntax of a state machine definition.
//
// You can validate that a state machine definition is correct without creating a
// state machine resource. Step Functions will implicitly perform the same syntax
// check when you invoke CreateStateMachine and UpdateStateMachine . State machine
// definitions are specified using a JSON-based, structured language. For more
// information on Amazon States Language see [Amazon States Language](ASL).
//
// Suggested uses for ValidateStateMachineDefinition :
//
//   - Integrate automated checks into your code review or Continuous Integration
//     (CI) process to validate state machine definitions before starting deployments.
//
//   - Run the validation from a Git pre-commit hook to check your state machine
//     definitions before committing them to your source repository.
//
// Errors found in the state machine definition will be returned in the response
// as a list of diagnostic elements, rather than raise an exception.
//
// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
func (c *Client) ValidateStateMachineDefinition(ctx context.Context, params *ValidateStateMachineDefinitionInput, optFns ...func(*Options)) (*ValidateStateMachineDefinitionOutput, error) {
	if params == nil {
		params = &ValidateStateMachineDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateStateMachineDefinition", params, optFns, c.addOperationValidateStateMachineDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateStateMachineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateStateMachineDefinitionInput struct {

	// The Amazon States Language definition of the state machine. For more
	// information, see [Amazon States Language](ASL).
	//
	// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
	//
	// This member is required.
	Definition *string

	// The target type of state machine for this definition. The default is STANDARD .
	Type types.StateMachineType

	noSmithyDocumentSerde
}

type ValidateStateMachineDefinitionOutput struct {

	// If the result is OK , this field will be empty. When there are errors, this
	// field will contain an array of Diagnostic objects to help you troubleshoot.
	//
	// This member is required.
	Diagnostics []types.ValidateStateMachineDefinitionDiagnostic

	// The result value will be OK when no syntax errors are found, or FAIL if the
	// workflow definition does not pass verification.
	//
	// This member is required.
	Result types.ValidateStateMachineDefinitionResultCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateStateMachineDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpValidateStateMachineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpValidateStateMachineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidateStateMachineDefinition"); err != nil {
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
	if err = addOpValidateStateMachineDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateStateMachineDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateStateMachineDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidateStateMachineDefinition",
	}
}
