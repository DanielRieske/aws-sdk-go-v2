// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AllowMessages string

// Enum values for AllowMessages
const (
	AllowMessagesAll  AllowMessages = "ALL"
	AllowMessagesNone AllowMessages = "NONE"
)

// Values returns all known values for AllowMessages. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AllowMessages) Values() []AllowMessages {
	return []AllowMessages{
		"ALL",
		"NONE",
	}
}

type AppInstanceUserEndpointType string

// Enum values for AppInstanceUserEndpointType
const (
	AppInstanceUserEndpointTypeApns        AppInstanceUserEndpointType = "APNS"
	AppInstanceUserEndpointTypeApnsSandbox AppInstanceUserEndpointType = "APNS_SANDBOX"
	AppInstanceUserEndpointTypeGcm         AppInstanceUserEndpointType = "GCM"
)

// Values returns all known values for AppInstanceUserEndpointType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppInstanceUserEndpointType) Values() []AppInstanceUserEndpointType {
	return []AppInstanceUserEndpointType{
		"APNS",
		"APNS_SANDBOX",
		"GCM",
	}
}

type EndpointStatus string

// Enum values for EndpointStatus
const (
	EndpointStatusActive   EndpointStatus = "ACTIVE"
	EndpointStatusInactive EndpointStatus = "INACTIVE"
)

// Values returns all known values for EndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatus) Values() []EndpointStatus {
	return []EndpointStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type EndpointStatusReason string

// Enum values for EndpointStatusReason
const (
	EndpointStatusReasonInvalidDeviceToken EndpointStatusReason = "INVALID_DEVICE_TOKEN"
	EndpointStatusReasonInvalidPinpointArn EndpointStatusReason = "INVALID_PINPOINT_ARN"
)

// Values returns all known values for EndpointStatusReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatusReason) Values() []EndpointStatusReason {
	return []EndpointStatusReason{
		"INVALID_DEVICE_TOKEN",
		"INVALID_PINPOINT_ARN",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotFound                             ErrorCode = "NotFound"
	ErrorCodePreconditionFailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist         ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}

type ExpirationCriterion string

// Enum values for ExpirationCriterion
const (
	ExpirationCriterionCreatedTimestamp ExpirationCriterion = "CREATED_TIMESTAMP"
)

// Values returns all known values for ExpirationCriterion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExpirationCriterion) Values() []ExpirationCriterion {
	return []ExpirationCriterion{
		"CREATED_TIMESTAMP",
	}
}

type RespondsTo string

// Enum values for RespondsTo
const (
	RespondsToStandardMessages RespondsTo = "STANDARD_MESSAGES"
)

// Values returns all known values for RespondsTo. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RespondsTo) Values() []RespondsTo {
	return []RespondsTo{
		"STANDARD_MESSAGES",
	}
}

type StandardMessages string

// Enum values for StandardMessages
const (
	StandardMessagesAuto     StandardMessages = "AUTO"
	StandardMessagesAll      StandardMessages = "ALL"
	StandardMessagesMentions StandardMessages = "MENTIONS"
	StandardMessagesNone     StandardMessages = "NONE"
)

// Values returns all known values for StandardMessages. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StandardMessages) Values() []StandardMessages {
	return []StandardMessages{
		"AUTO",
		"ALL",
		"MENTIONS",
		"NONE",
	}
}

type TargetedMessages string

// Enum values for TargetedMessages
const (
	TargetedMessagesAll  TargetedMessages = "ALL"
	TargetedMessagesNone TargetedMessages = "NONE"
)

// Values returns all known values for TargetedMessages. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetedMessages) Values() []TargetedMessages {
	return []TargetedMessages{
		"ALL",
		"NONE",
	}
}
