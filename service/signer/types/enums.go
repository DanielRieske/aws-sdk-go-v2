// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Category string

// Enum values for Category
const (
	CategoryAWSIoT Category = "AWSIoT"
)

// Values returns all known values for Category. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Category) Values() []Category {
	return []Category{
		"AWSIoT",
	}
}

type EncryptionAlgorithm string

// Enum values for EncryptionAlgorithm
const (
	EncryptionAlgorithmRsa   EncryptionAlgorithm = "RSA"
	EncryptionAlgorithmEcdsa EncryptionAlgorithm = "ECDSA"
)

// Values returns all known values for EncryptionAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionAlgorithm) Values() []EncryptionAlgorithm {
	return []EncryptionAlgorithm{
		"RSA",
		"ECDSA",
	}
}

type HashAlgorithm string

// Enum values for HashAlgorithm
const (
	HashAlgorithmSha1   HashAlgorithm = "SHA1"
	HashAlgorithmSha256 HashAlgorithm = "SHA256"
)

// Values returns all known values for HashAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HashAlgorithm) Values() []HashAlgorithm {
	return []HashAlgorithm{
		"SHA1",
		"SHA256",
	}
}

type ImageFormat string

// Enum values for ImageFormat
const (
	ImageFormatJson         ImageFormat = "JSON"
	ImageFormatJSONEmbedded ImageFormat = "JSONEmbedded"
	ImageFormatJSONDetached ImageFormat = "JSONDetached"
)

// Values returns all known values for ImageFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageFormat) Values() []ImageFormat {
	return []ImageFormat{
		"JSON",
		"JSONEmbedded",
		"JSONDetached",
	}
}

type SigningProfileStatus string

// Enum values for SigningProfileStatus
const (
	SigningProfileStatusActive   SigningProfileStatus = "Active"
	SigningProfileStatusCanceled SigningProfileStatus = "Canceled"
	SigningProfileStatusRevoked  SigningProfileStatus = "Revoked"
)

// Values returns all known values for SigningProfileStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SigningProfileStatus) Values() []SigningProfileStatus {
	return []SigningProfileStatus{
		"Active",
		"Canceled",
		"Revoked",
	}
}

type SigningStatus string

// Enum values for SigningStatus
const (
	SigningStatusInProgress SigningStatus = "InProgress"
	SigningStatusFailed     SigningStatus = "Failed"
	SigningStatusSucceeded  SigningStatus = "Succeeded"
)

// Values returns all known values for SigningStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SigningStatus) Values() []SigningStatus {
	return []SigningStatus{
		"InProgress",
		"Failed",
		"Succeeded",
	}
}

type ValidityType string

// Enum values for ValidityType
const (
	ValidityTypeDays   ValidityType = "DAYS"
	ValidityTypeMonths ValidityType = "MONTHS"
	ValidityTypeYears  ValidityType = "YEARS"
)

// Values returns all known values for ValidityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidityType) Values() []ValidityType {
	return []ValidityType{
		"DAYS",
		"MONTHS",
		"YEARS",
	}
}
