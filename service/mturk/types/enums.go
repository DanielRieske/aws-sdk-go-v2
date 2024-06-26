// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssignmentStatus string

// Enum values for AssignmentStatus
const (
	AssignmentStatusSubmitted AssignmentStatus = "Submitted"
	AssignmentStatusApproved  AssignmentStatus = "Approved"
	AssignmentStatusRejected  AssignmentStatus = "Rejected"
)

// Values returns all known values for AssignmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssignmentStatus) Values() []AssignmentStatus {
	return []AssignmentStatus{
		"Submitted",
		"Approved",
		"Rejected",
	}
}

type Comparator string

// Enum values for Comparator
const (
	ComparatorLessThan             Comparator = "LessThan"
	ComparatorLessThanOrEqualTo    Comparator = "LessThanOrEqualTo"
	ComparatorGreaterThan          Comparator = "GreaterThan"
	ComparatorGreaterThanOrEqualTo Comparator = "GreaterThanOrEqualTo"
	ComparatorEqualTo              Comparator = "EqualTo"
	ComparatorNotEqualTo           Comparator = "NotEqualTo"
	ComparatorExists               Comparator = "Exists"
	ComparatorDoesNotExist         Comparator = "DoesNotExist"
	ComparatorIn                   Comparator = "In"
	ComparatorNotIn                Comparator = "NotIn"
)

// Values returns all known values for Comparator. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Comparator) Values() []Comparator {
	return []Comparator{
		"LessThan",
		"LessThanOrEqualTo",
		"GreaterThan",
		"GreaterThanOrEqualTo",
		"EqualTo",
		"NotEqualTo",
		"Exists",
		"DoesNotExist",
		"In",
		"NotIn",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeAssignmentAccepted  EventType = "AssignmentAccepted"
	EventTypeAssignmentAbandoned EventType = "AssignmentAbandoned"
	EventTypeAssignmentReturned  EventType = "AssignmentReturned"
	EventTypeAssignmentSubmitted EventType = "AssignmentSubmitted"
	EventTypeAssignmentRejected  EventType = "AssignmentRejected"
	EventTypeAssignmentApproved  EventType = "AssignmentApproved"
	EventTypeHITCreated          EventType = "HITCreated"
	EventTypeHITExpired          EventType = "HITExpired"
	EventTypeHITReviewable       EventType = "HITReviewable"
	EventTypeHITExtended         EventType = "HITExtended"
	EventTypeHITDisposed         EventType = "HITDisposed"
	EventTypePing                EventType = "Ping"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"AssignmentAccepted",
		"AssignmentAbandoned",
		"AssignmentReturned",
		"AssignmentSubmitted",
		"AssignmentRejected",
		"AssignmentApproved",
		"HITCreated",
		"HITExpired",
		"HITReviewable",
		"HITExtended",
		"HITDisposed",
		"Ping",
	}
}

type HITAccessActions string

// Enum values for HITAccessActions
const (
	HITAccessActionsAccept                   HITAccessActions = "Accept"
	HITAccessActionsPreviewAndAccept         HITAccessActions = "PreviewAndAccept"
	HITAccessActionsDiscoverPreviewAndAccept HITAccessActions = "DiscoverPreviewAndAccept"
)

// Values returns all known values for HITAccessActions. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HITAccessActions) Values() []HITAccessActions {
	return []HITAccessActions{
		"Accept",
		"PreviewAndAccept",
		"DiscoverPreviewAndAccept",
	}
}

type HITReviewStatus string

// Enum values for HITReviewStatus
const (
	HITReviewStatusNotReviewed           HITReviewStatus = "NotReviewed"
	HITReviewStatusMarkedForReview       HITReviewStatus = "MarkedForReview"
	HITReviewStatusReviewedAppropriate   HITReviewStatus = "ReviewedAppropriate"
	HITReviewStatusReviewedInappropriate HITReviewStatus = "ReviewedInappropriate"
)

// Values returns all known values for HITReviewStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HITReviewStatus) Values() []HITReviewStatus {
	return []HITReviewStatus{
		"NotReviewed",
		"MarkedForReview",
		"ReviewedAppropriate",
		"ReviewedInappropriate",
	}
}

type HITStatus string

// Enum values for HITStatus
const (
	HITStatusAssignable   HITStatus = "Assignable"
	HITStatusUnassignable HITStatus = "Unassignable"
	HITStatusReviewable   HITStatus = "Reviewable"
	HITStatusReviewing    HITStatus = "Reviewing"
	HITStatusDisposed     HITStatus = "Disposed"
)

// Values returns all known values for HITStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HITStatus) Values() []HITStatus {
	return []HITStatus{
		"Assignable",
		"Unassignable",
		"Reviewable",
		"Reviewing",
		"Disposed",
	}
}

type NotificationTransport string

// Enum values for NotificationTransport
const (
	NotificationTransportEmail NotificationTransport = "Email"
	NotificationTransportSqs   NotificationTransport = "SQS"
	NotificationTransportSns   NotificationTransport = "SNS"
)

// Values returns all known values for NotificationTransport. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotificationTransport) Values() []NotificationTransport {
	return []NotificationTransport{
		"Email",
		"SQS",
		"SNS",
	}
}

type NotifyWorkersFailureCode string

// Enum values for NotifyWorkersFailureCode
const (
	NotifyWorkersFailureCodeSoftFailure NotifyWorkersFailureCode = "SoftFailure"
	NotifyWorkersFailureCodeHardFailure NotifyWorkersFailureCode = "HardFailure"
)

// Values returns all known values for NotifyWorkersFailureCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotifyWorkersFailureCode) Values() []NotifyWorkersFailureCode {
	return []NotifyWorkersFailureCode{
		"SoftFailure",
		"HardFailure",
	}
}

type QualificationStatus string

// Enum values for QualificationStatus
const (
	QualificationStatusGranted QualificationStatus = "Granted"
	QualificationStatusRevoked QualificationStatus = "Revoked"
)

// Values returns all known values for QualificationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QualificationStatus) Values() []QualificationStatus {
	return []QualificationStatus{
		"Granted",
		"Revoked",
	}
}

type QualificationTypeStatus string

// Enum values for QualificationTypeStatus
const (
	QualificationTypeStatusActive   QualificationTypeStatus = "Active"
	QualificationTypeStatusInactive QualificationTypeStatus = "Inactive"
)

// Values returns all known values for QualificationTypeStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QualificationTypeStatus) Values() []QualificationTypeStatus {
	return []QualificationTypeStatus{
		"Active",
		"Inactive",
	}
}

type ReviewableHITStatus string

// Enum values for ReviewableHITStatus
const (
	ReviewableHITStatusReviewable ReviewableHITStatus = "Reviewable"
	ReviewableHITStatusReviewing  ReviewableHITStatus = "Reviewing"
)

// Values returns all known values for ReviewableHITStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReviewableHITStatus) Values() []ReviewableHITStatus {
	return []ReviewableHITStatus{
		"Reviewable",
		"Reviewing",
	}
}

type ReviewActionStatus string

// Enum values for ReviewActionStatus
const (
	ReviewActionStatusIntended  ReviewActionStatus = "Intended"
	ReviewActionStatusSucceeded ReviewActionStatus = "Succeeded"
	ReviewActionStatusFailed    ReviewActionStatus = "Failed"
	ReviewActionStatusCancelled ReviewActionStatus = "Cancelled"
)

// Values returns all known values for ReviewActionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReviewActionStatus) Values() []ReviewActionStatus {
	return []ReviewActionStatus{
		"Intended",
		"Succeeded",
		"Failed",
		"Cancelled",
	}
}

type ReviewPolicyLevel string

// Enum values for ReviewPolicyLevel
const (
	ReviewPolicyLevelAssignment ReviewPolicyLevel = "Assignment"
	ReviewPolicyLevelHit        ReviewPolicyLevel = "HIT"
)

// Values returns all known values for ReviewPolicyLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReviewPolicyLevel) Values() []ReviewPolicyLevel {
	return []ReviewPolicyLevel{
		"Assignment",
		"HIT",
	}
}
