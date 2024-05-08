// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The input fails to satisfy the constraints specified by an Amazon Web Services
// service.
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Reason  BadRequestReason
	Details BadRequestDetails

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be processed because of conflict in the current state of
// the resource.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was an internal failure in the AppConfig service.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The configuration size is too large.
type PayloadTooLargeException struct {
	Message *string

	ErrorCodeOverride *string

	Measure BytesMeasure
	Limit   float32
	Size    float32

	noSmithyDocumentSerde
}

func (e *PayloadTooLargeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PayloadTooLargeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PayloadTooLargeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PayloadTooLargeException"
	}
	return *e.ErrorCodeOverride
}
func (e *PayloadTooLargeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested resource could not be found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of one more AppConfig resources exceeds the maximum allowed. Verify
// that your environment doesn't exceed the following service quotas:
//
// Applications: 100 max
//
// Deployment strategies: 20 max
//
// Configuration profiles: 100 max per application
//
// Environments: 20 max per application
//
// To resolve this issue, you can delete one or more resources and try again. Or,
// you can request a quota increase. For more information about quotas and to
// request an increase, see [Service quotas for AppConfig]in the Amazon Web Services General Reference.
//
// [Service quotas for AppConfig]: https://docs.aws.amazon.com/general/latest/gr/appconfig.html#limits_appconfig
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
