// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of extension pack associations for the specified
// migration project. An extension pack is an add-on module that emulates functions
// present in a source database that are required when converting objects to the
// target database.
func (c *Client) DescribeExtensionPackAssociations(ctx context.Context, params *DescribeExtensionPackAssociationsInput, optFns ...func(*Options)) (*DescribeExtensionPackAssociationsOutput, error) {
	if params == nil {
		params = &DescribeExtensionPackAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExtensionPackAssociations", params, optFns, c.addOperationDescribeExtensionPackAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExtensionPackAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExtensionPackAssociationsInput struct {

	// The name or Amazon Resource Name (ARN) for the migration project.
	//
	// This member is required.
	MigrationProjectIdentifier *string

	// Filters applied to the extension pack associations described in the form of
	// key-value pairs.
	Filters []types.Filter

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeExtensionPackAssociationsOutput struct {

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// A paginated list of extension pack associations for the specified migration
	// project.
	Requests []types.SchemaConversionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExtensionPackAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExtensionPackAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExtensionPackAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExtensionPackAssociations"); err != nil {
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
	if err = addOpDescribeExtensionPackAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExtensionPackAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeExtensionPackAssociationsAPIClient is a client that implements the
// DescribeExtensionPackAssociations operation.
type DescribeExtensionPackAssociationsAPIClient interface {
	DescribeExtensionPackAssociations(context.Context, *DescribeExtensionPackAssociationsInput, ...func(*Options)) (*DescribeExtensionPackAssociationsOutput, error)
}

var _ DescribeExtensionPackAssociationsAPIClient = (*Client)(nil)

// DescribeExtensionPackAssociationsPaginatorOptions is the paginator options for
// DescribeExtensionPackAssociations
type DescribeExtensionPackAssociationsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeExtensionPackAssociationsPaginator is a paginator for
// DescribeExtensionPackAssociations
type DescribeExtensionPackAssociationsPaginator struct {
	options   DescribeExtensionPackAssociationsPaginatorOptions
	client    DescribeExtensionPackAssociationsAPIClient
	params    *DescribeExtensionPackAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeExtensionPackAssociationsPaginator returns a new
// DescribeExtensionPackAssociationsPaginator
func NewDescribeExtensionPackAssociationsPaginator(client DescribeExtensionPackAssociationsAPIClient, params *DescribeExtensionPackAssociationsInput, optFns ...func(*DescribeExtensionPackAssociationsPaginatorOptions)) *DescribeExtensionPackAssociationsPaginator {
	if params == nil {
		params = &DescribeExtensionPackAssociationsInput{}
	}

	options := DescribeExtensionPackAssociationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeExtensionPackAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeExtensionPackAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeExtensionPackAssociations page.
func (p *DescribeExtensionPackAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeExtensionPackAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeExtensionPackAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeExtensionPackAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExtensionPackAssociations",
	}
}
