// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the public key and an import token you need to import or reimport key
// material for a KMS key.
//
// By default, KMS keys are created with key material that KMS generates. This
// operation supports [Importing key material], an advanced feature that lets you generate and import the
// cryptographic key material for a KMS key. For more information about importing
// key material into KMS, see [Importing key material]in the Key Management Service Developer Guide.
//
// Before calling GetParametersForImport , use the CreateKey operation with an Origin value
// of EXTERNAL to create a KMS key with no key material. You can import key
// material for a symmetric encryption KMS key, HMAC KMS key, asymmetric encryption
// KMS key, or asymmetric signing KMS key. You can also import key material into a [multi-Region key]
// of any supported type. However, you can't import key material into a KMS key in
// a [custom key store]. You can also use GetParametersForImport to get a public key and import
// token to [reimport the original key material]into a KMS key whose key material expired or was deleted.
//
// GetParametersForImport returns the items that you need to import your key
// material.
//
//   - The public key (or "wrapping key") of an RSA key pair that KMS generates.
//
// You will use this public key to encrypt ("wrap") your key material while it's
//
//	in transit to KMS.
//
//	- A import token that ensures that KMS can decrypt your key material and
//	associate it with the correct KMS key.
//
// The public key and its import token are permanently linked and must be used
// together. Each public key and import token set is valid for 24 hours. The
// expiration date and time appear in the ParametersValidTo field in the
// GetParametersForImport response. You cannot use an expired public key or import
// token in an ImportKeyMaterialrequest. If your key and token expire, send another
// GetParametersForImport request.
//
// GetParametersForImport requires the following information:
//
//   - The key ID of the KMS key for which you are importing the key material.
//
//   - The key spec of the public key ("wrapping key") that you will use to
//     encrypt your key material during import.
//
//   - The wrapping algorithm that you will use with the public key to encrypt
//     your key material.
//
// You can use the same or a different public key spec and wrapping algorithm each
// time you import or reimport the same key material.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:GetParametersForImport] (key policy)
//
// Related operations:
//
// # ImportKeyMaterial
//
// # DeleteImportedKeyMaterial
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [kms:GetParametersForImport]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [reimport the original key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [multi-Region key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html
func (c *Client) GetParametersForImport(ctx context.Context, params *GetParametersForImportInput, optFns ...func(*Options)) (*GetParametersForImportOutput, error) {
	if params == nil {
		params = &GetParametersForImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParametersForImport", params, optFns, c.addOperationGetParametersForImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParametersForImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersForImportInput struct {

	// The identifier of the KMS key that will be associated with the imported key
	// material. The Origin of the KMS key must be EXTERNAL .
	//
	// All KMS key types are supported, including multi-Region keys. However, you
	// cannot import key material into a KMS key in a custom key store.
	//
	// Specify the key ID or key ARN of the KMS key.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// The algorithm you will use with the RSA public key ( PublicKey ) in the response
	// to protect your key material during import. For more information, see Select a wrapping algorithmin the
	// Key Management Service Developer Guide.
	//
	// For RSA_AES wrapping algorithms, you encrypt your key material with an AES key
	// that you generate, then encrypt your AES key with the RSA public key from KMS.
	// For RSAES wrapping algorithms, you encrypt your key material directly with the
	// RSA public key from KMS.
	//
	// The wrapping algorithms that you can use depend on the type of key material
	// that you are importing. To import an RSA private key, you must use an RSA_AES
	// wrapping algorithm.
	//
	//   - RSA_AES_KEY_WRAP_SHA_256 — Supported for wrapping RSA and ECC key material.
	//
	//   - RSA_AES_KEY_WRAP_SHA_1 — Supported for wrapping RSA and ECC key material.
	//
	//   - RSAES_OAEP_SHA_256 — Supported for all types of key material, except RSA
	//   key material (private key).
	//
	// You cannot use the RSAES_OAEP_SHA_256 wrapping algorithm with the RSA_2048
	//   wrapping key spec to wrap ECC_NIST_P521 key material.
	//
	//   - RSAES_OAEP_SHA_1 — Supported for all types of key material, except RSA key
	//   material (private key).
	//
	// You cannot use the RSAES_OAEP_SHA_1 wrapping algorithm with the RSA_2048
	//   wrapping key spec to wrap ECC_NIST_P521 key material.
	//
	//   - RSAES_PKCS1_V1_5 (Deprecated) — As of October 10, 2023, KMS does not
	//   support the RSAES_PKCS1_V1_5 wrapping algorithm.
	//
	// This member is required.
	WrappingAlgorithm types.AlgorithmSpec

	// The type of RSA public key to return in the response. You will use this
	// wrapping key with the specified wrapping algorithm to protect your key material
	// during import.
	//
	// Use the longest RSA wrapping key that is practical.
	//
	// You cannot use an RSA_2048 public key to directly wrap an ECC_NIST_P521 private
	// key. Instead, use an RSA_AES wrapping algorithm or choose a longer RSA public
	// key.
	//
	// This member is required.
	WrappingKeySpec types.WrappingKeySpec

	noSmithyDocumentSerde
}

type GetParametersForImportOutput struct {

	// The import token to send in a subsequent ImportKeyMaterial request.
	ImportToken []byte

	// The Amazon Resource Name ([key ARN] ) of the KMS key to use in a subsequent ImportKeyMaterial request.
	// This is the same KMS key specified in the GetParametersForImport request.
	//
	// [key ARN]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN
	KeyId *string

	// The time at which the import token and public key are no longer valid. After
	// this time, you cannot use them to make an ImportKeyMaterialrequest and you must send another
	// GetParametersForImport request to get new ones.
	ParametersValidTo *time.Time

	// The public key to use to encrypt the key material before importing it with ImportKeyMaterial.
	PublicKey []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParametersForImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetParametersForImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParametersForImport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetParametersForImport"); err != nil {
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
	if err = addOpGetParametersForImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersForImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetParametersForImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetParametersForImport",
	}
}
