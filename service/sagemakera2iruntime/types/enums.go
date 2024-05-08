// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContentClassifier string

// Enum values for ContentClassifier
const (
	ContentClassifierFreeOfPersonallyIdentifiableInformation ContentClassifier = "FreeOfPersonallyIdentifiableInformation"
	ContentClassifierFreeOfAdultContent                      ContentClassifier = "FreeOfAdultContent"
)

// Values returns all known values for ContentClassifier. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContentClassifier) Values() []ContentClassifier {
	return []ContentClassifier{
		"FreeOfPersonallyIdentifiableInformation",
		"FreeOfAdultContent",
	}
}

type HumanLoopStatus string

// Enum values for HumanLoopStatus
const (
	HumanLoopStatusInProgress HumanLoopStatus = "InProgress"
	HumanLoopStatusFailed     HumanLoopStatus = "Failed"
	HumanLoopStatusCompleted  HumanLoopStatus = "Completed"
	HumanLoopStatusStopped    HumanLoopStatus = "Stopped"
	HumanLoopStatusStopping   HumanLoopStatus = "Stopping"
)

// Values returns all known values for HumanLoopStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HumanLoopStatus) Values() []HumanLoopStatus {
	return []HumanLoopStatus{
		"InProgress",
		"Failed",
		"Completed",
		"Stopped",
		"Stopping",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "Ascending"
	SortOrderDescending SortOrder = "Descending"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"Ascending",
		"Descending",
	}
}
