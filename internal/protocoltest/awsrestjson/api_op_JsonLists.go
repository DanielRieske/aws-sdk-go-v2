// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This test case serializes JSON lists for the following cases for both input and
// output:
//
//   - Normal JSON lists.
//   - Normal JSON sets.
//   - JSON lists of lists.
//   - Lists of structures.
func (c *Client) JsonLists(ctx context.Context, params *JsonListsInput, optFns ...func(*Options)) (*JsonListsOutput, error) {
	if params == nil {
		params = &JsonListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonLists", params, optFns, c.addOperationJsonListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonListsInput struct {
	BooleanList []bool

	EnumList []types.FooEnum

	IntEnumList []types.IntegerEnum

	IntegerList []int32

	// A list of lists of strings.
	NestedStringList [][]string

	StringList []string

	StringSet []string

	StructureList []types.StructureListMember

	TimestampList []time.Time

	noSmithyDocumentSerde
}

type JsonListsOutput struct {
	BooleanList []bool

	EnumList []types.FooEnum

	IntEnumList []types.IntegerEnum

	IntegerList []int32

	// A list of lists of strings.
	NestedStringList [][]string

	StringList []string

	StringSet []string

	StructureList []types.StructureListMember

	TimestampList []time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJsonListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJsonLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "JsonLists"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJsonLists(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJsonLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonLists",
	}
}
