// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation is not supported by directory buckets.
//
// Deletes the access point policy for the specified access point.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to DeleteAccessPointPolicy :
//
// [PutAccessPointPolicy]
//
// [GetAccessPointPolicy]
//
// [GetAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples
func (c *Client) DeleteAccessPointPolicy(ctx context.Context, params *DeleteAccessPointPolicyInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyOutput, error) {
	if params == nil {
		params = &DeleteAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessPointPolicy", params, optFns, c.addOperationDeleteAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPointPolicyInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy you want to delete.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/ . For example, to access the access
	// point reports-ap through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap
	// . The value must be URL encoded.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

func (in *DeleteAccessPointPolicyInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.AccessPointName = in.Name
	p.RequiresAccountId = ptr.Bool(true)
}

type DeleteAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAccessPointPolicy"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteAccessPointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessPointPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteAccessPointPolicyUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (m *DeleteAccessPointPolicyInput) GetARNMember() (*string, bool) {
	if m.Name == nil {
		return nil, false
	}
	return m.Name, true
}

func (m *DeleteAccessPointPolicyInput) SetARNMember(v string) error {
	m.Name = &v
	return nil
}

type endpointPrefix_opDeleteAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessPointPolicyMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*DeleteAccessPointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDeleteAccessPointPolicyMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAccessPointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAccessPointPolicy",
	}
}

func copyDeleteAccessPointPolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteAccessPointPolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteAccessPointPolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *DeleteAccessPointPolicyInput) copy() interface{} {
	v := *in
	return &v
}
func getDeleteAccessPointPolicyARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteAccessPointPolicyInput)
	if in.Name == nil {
		return nil, false
	}
	return in.Name, true
}
func setDeleteAccessPointPolicyARNMember(input interface{}, v string) error {
	in := input.(*DeleteAccessPointPolicyInput)
	in.Name = &v
	return nil
}
func backFillDeleteAccessPointPolicyAccountID(input interface{}, v string) error {
	in := input.(*DeleteAccessPointPolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteAccessPointPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteAccessPointPolicyARNMember,
			BackfillAccountID: backFillDeleteAccessPointPolicyAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteAccessPointPolicyARNMember,
			CopyInput:         copyDeleteAccessPointPolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
