// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationTagStatus string

// Enum values for ApplicationTagStatus
const (
	ApplicationTagStatusInProgress ApplicationTagStatus = "IN_PROGRESS"
	ApplicationTagStatusSuccess    ApplicationTagStatus = "SUCCESS"
	ApplicationTagStatusFailure    ApplicationTagStatus = "FAILURE"
)

// Values returns all known values for ApplicationTagStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationTagStatus) Values() []ApplicationTagStatus {
	return []ApplicationTagStatus{
		"IN_PROGRESS",
		"SUCCESS",
		"FAILURE",
	}
}

type AssociationOption string

// Enum values for AssociationOption
const (
	AssociationOptionApplyApplicationTag AssociationOption = "APPLY_APPLICATION_TAG"
	AssociationOptionSkipApplicationTag  AssociationOption = "SKIP_APPLICATION_TAG"
)

// Values returns all known values for AssociationOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssociationOption) Values() []AssociationOption {
	return []AssociationOption{
		"APPLY_APPLICATION_TAG",
		"SKIP_APPLICATION_TAG",
	}
}

type ResourceGroupState string

// Enum values for ResourceGroupState
const (
	ResourceGroupStateCreating       ResourceGroupState = "CREATING"
	ResourceGroupStateCreateComplete ResourceGroupState = "CREATE_COMPLETE"
	ResourceGroupStateCreateFailed   ResourceGroupState = "CREATE_FAILED"
	ResourceGroupStateUpdating       ResourceGroupState = "UPDATING"
	ResourceGroupStateUpdateComplete ResourceGroupState = "UPDATE_COMPLETE"
	ResourceGroupStateUpdateFailed   ResourceGroupState = "UPDATE_FAILED"
)

// Values returns all known values for ResourceGroupState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceGroupState) Values() []ResourceGroupState {
	return []ResourceGroupState{
		"CREATING",
		"CREATE_COMPLETE",
		"CREATE_FAILED",
		"UPDATING",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}

type ResourceItemStatus string

// Enum values for ResourceItemStatus
const (
	ResourceItemStatusSuccess    ResourceItemStatus = "SUCCESS"
	ResourceItemStatusFailed     ResourceItemStatus = "FAILED"
	ResourceItemStatusInProgress ResourceItemStatus = "IN_PROGRESS"
	ResourceItemStatusSkipped    ResourceItemStatus = "SKIPPED"
)

// Values returns all known values for ResourceItemStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceItemStatus) Values() []ResourceItemStatus {
	return []ResourceItemStatus{
		"SUCCESS",
		"FAILED",
		"IN_PROGRESS",
		"SKIPPED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeCfnStack         ResourceType = "CFN_STACK"
	ResourceTypeResourceTagValue ResourceType = "RESOURCE_TAG_VALUE"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"CFN_STACK",
		"RESOURCE_TAG_VALUE",
	}
}

type SyncAction string

// Enum values for SyncAction
const (
	SyncActionStartSync SyncAction = "START_SYNC"
	SyncActionNoAction  SyncAction = "NO_ACTION"
)

// Values returns all known values for SyncAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SyncAction) Values() []SyncAction {
	return []SyncAction{
		"START_SYNC",
		"NO_ACTION",
	}
}
