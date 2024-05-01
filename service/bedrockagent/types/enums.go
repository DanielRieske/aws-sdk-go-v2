// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionGroupSignature string

// Enum values for ActionGroupSignature
const (
	ActionGroupSignatureAmazonUserinput ActionGroupSignature = "AMAZON.UserInput"
)

// Values returns all known values for ActionGroupSignature. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionGroupSignature) Values() []ActionGroupSignature {
	return []ActionGroupSignature{
		"AMAZON.UserInput",
	}
}

type ActionGroupState string

// Enum values for ActionGroupState
const (
	ActionGroupStateEnabled  ActionGroupState = "ENABLED"
	ActionGroupStateDisabled ActionGroupState = "DISABLED"
)

// Values returns all known values for ActionGroupState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionGroupState) Values() []ActionGroupState {
	return []ActionGroupState{
		"ENABLED",
		"DISABLED",
	}
}

type AgentAliasStatus string

// Enum values for AgentAliasStatus
const (
	AgentAliasStatusCreating AgentAliasStatus = "CREATING"
	AgentAliasStatusPrepared AgentAliasStatus = "PREPARED"
	AgentAliasStatusFailed   AgentAliasStatus = "FAILED"
	AgentAliasStatusUpdating AgentAliasStatus = "UPDATING"
	AgentAliasStatusDeleting AgentAliasStatus = "DELETING"
)

// Values returns all known values for AgentAliasStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AgentAliasStatus) Values() []AgentAliasStatus {
	return []AgentAliasStatus{
		"CREATING",
		"PREPARED",
		"FAILED",
		"UPDATING",
		"DELETING",
	}
}

type AgentStatus string

// Enum values for AgentStatus
const (
	AgentStatusCreating    AgentStatus = "CREATING"
	AgentStatusPreparing   AgentStatus = "PREPARING"
	AgentStatusPrepared    AgentStatus = "PREPARED"
	AgentStatusNotPrepared AgentStatus = "NOT_PREPARED"
	AgentStatusDeleting    AgentStatus = "DELETING"
	AgentStatusFailed      AgentStatus = "FAILED"
	AgentStatusVersioning  AgentStatus = "VERSIONING"
	AgentStatusUpdating    AgentStatus = "UPDATING"
)

// Values returns all known values for AgentStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AgentStatus) Values() []AgentStatus {
	return []AgentStatus{
		"CREATING",
		"PREPARING",
		"PREPARED",
		"NOT_PREPARED",
		"DELETING",
		"FAILED",
		"VERSIONING",
		"UPDATING",
	}
}

type ChunkingStrategy string

// Enum values for ChunkingStrategy
const (
	ChunkingStrategyFixedSize ChunkingStrategy = "FIXED_SIZE"
	ChunkingStrategyNone      ChunkingStrategy = "NONE"
)

// Values returns all known values for ChunkingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChunkingStrategy) Values() []ChunkingStrategy {
	return []ChunkingStrategy{
		"FIXED_SIZE",
		"NONE",
	}
}

type CreationMode string

// Enum values for CreationMode
const (
	CreationModeDefault    CreationMode = "DEFAULT"
	CreationModeOverridden CreationMode = "OVERRIDDEN"
)

// Values returns all known values for CreationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CreationMode) Values() []CreationMode {
	return []CreationMode{
		"DEFAULT",
		"OVERRIDDEN",
	}
}

type CustomControlMethod string

// Enum values for CustomControlMethod
const (
	CustomControlMethodReturnControl CustomControlMethod = "RETURN_CONTROL"
)

// Values returns all known values for CustomControlMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CustomControlMethod) Values() []CustomControlMethod {
	return []CustomControlMethod{
		"RETURN_CONTROL",
	}
}

type DataDeletionPolicy string

// Enum values for DataDeletionPolicy
const (
	DataDeletionPolicyRetain DataDeletionPolicy = "RETAIN"
	DataDeletionPolicyDelete DataDeletionPolicy = "DELETE"
)

// Values returns all known values for DataDeletionPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataDeletionPolicy) Values() []DataDeletionPolicy {
	return []DataDeletionPolicy{
		"RETAIN",
		"DELETE",
	}
}

type DataSourceStatus string

// Enum values for DataSourceStatus
const (
	DataSourceStatusAvailable          DataSourceStatus = "AVAILABLE"
	DataSourceStatusDeleting           DataSourceStatus = "DELETING"
	DataSourceStatusDeleteUnsuccessful DataSourceStatus = "DELETE_UNSUCCESSFUL"
)

// Values returns all known values for DataSourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceStatus) Values() []DataSourceStatus {
	return []DataSourceStatus{
		"AVAILABLE",
		"DELETING",
		"DELETE_UNSUCCESSFUL",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeS3 DataSourceType = "S3"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"S3",
	}
}

type IngestionJobFilterAttribute string

// Enum values for IngestionJobFilterAttribute
const (
	IngestionJobFilterAttributeStatus IngestionJobFilterAttribute = "STATUS"
)

// Values returns all known values for IngestionJobFilterAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobFilterAttribute) Values() []IngestionJobFilterAttribute {
	return []IngestionJobFilterAttribute{
		"STATUS",
	}
}

type IngestionJobFilterOperator string

// Enum values for IngestionJobFilterOperator
const (
	IngestionJobFilterOperatorEq IngestionJobFilterOperator = "EQ"
)

// Values returns all known values for IngestionJobFilterOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobFilterOperator) Values() []IngestionJobFilterOperator {
	return []IngestionJobFilterOperator{
		"EQ",
	}
}

type IngestionJobSortByAttribute string

// Enum values for IngestionJobSortByAttribute
const (
	IngestionJobSortByAttributeStatus    IngestionJobSortByAttribute = "STATUS"
	IngestionJobSortByAttributeStartedAt IngestionJobSortByAttribute = "STARTED_AT"
)

// Values returns all known values for IngestionJobSortByAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobSortByAttribute) Values() []IngestionJobSortByAttribute {
	return []IngestionJobSortByAttribute{
		"STATUS",
		"STARTED_AT",
	}
}

type IngestionJobStatus string

// Enum values for IngestionJobStatus
const (
	IngestionJobStatusStarting   IngestionJobStatus = "STARTING"
	IngestionJobStatusInProgress IngestionJobStatus = "IN_PROGRESS"
	IngestionJobStatusComplete   IngestionJobStatus = "COMPLETE"
	IngestionJobStatusFailed     IngestionJobStatus = "FAILED"
)

// Values returns all known values for IngestionJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobStatus) Values() []IngestionJobStatus {
	return []IngestionJobStatus{
		"STARTING",
		"IN_PROGRESS",
		"COMPLETE",
		"FAILED",
	}
}

type KnowledgeBaseState string

// Enum values for KnowledgeBaseState
const (
	KnowledgeBaseStateEnabled  KnowledgeBaseState = "ENABLED"
	KnowledgeBaseStateDisabled KnowledgeBaseState = "DISABLED"
)

// Values returns all known values for KnowledgeBaseState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseState) Values() []KnowledgeBaseState {
	return []KnowledgeBaseState{
		"ENABLED",
		"DISABLED",
	}
}

type KnowledgeBaseStatus string

// Enum values for KnowledgeBaseStatus
const (
	KnowledgeBaseStatusCreating           KnowledgeBaseStatus = "CREATING"
	KnowledgeBaseStatusActive             KnowledgeBaseStatus = "ACTIVE"
	KnowledgeBaseStatusDeleting           KnowledgeBaseStatus = "DELETING"
	KnowledgeBaseStatusUpdating           KnowledgeBaseStatus = "UPDATING"
	KnowledgeBaseStatusFailed             KnowledgeBaseStatus = "FAILED"
	KnowledgeBaseStatusDeleteUnsuccessful KnowledgeBaseStatus = "DELETE_UNSUCCESSFUL"
)

// Values returns all known values for KnowledgeBaseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStatus) Values() []KnowledgeBaseStatus {
	return []KnowledgeBaseStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"UPDATING",
		"FAILED",
		"DELETE_UNSUCCESSFUL",
	}
}

type KnowledgeBaseStorageType string

// Enum values for KnowledgeBaseStorageType
const (
	KnowledgeBaseStorageTypeOpensearchServerless KnowledgeBaseStorageType = "OPENSEARCH_SERVERLESS"
	KnowledgeBaseStorageTypePinecone             KnowledgeBaseStorageType = "PINECONE"
	KnowledgeBaseStorageTypeRedisEnterpriseCloud KnowledgeBaseStorageType = "REDIS_ENTERPRISE_CLOUD"
	KnowledgeBaseStorageTypeRds                  KnowledgeBaseStorageType = "RDS"
	KnowledgeBaseStorageTypeMongoDbAtlas         KnowledgeBaseStorageType = "MONGO_DB_ATLAS"
)

// Values returns all known values for KnowledgeBaseStorageType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStorageType) Values() []KnowledgeBaseStorageType {
	return []KnowledgeBaseStorageType{
		"OPENSEARCH_SERVERLESS",
		"PINECONE",
		"REDIS_ENTERPRISE_CLOUD",
		"RDS",
		"MONGO_DB_ATLAS",
	}
}

type KnowledgeBaseType string

// Enum values for KnowledgeBaseType
const (
	KnowledgeBaseTypeVector KnowledgeBaseType = "VECTOR"
)

// Values returns all known values for KnowledgeBaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseType) Values() []KnowledgeBaseType {
	return []KnowledgeBaseType{
		"VECTOR",
	}
}

type PromptState string

// Enum values for PromptState
const (
	PromptStateEnabled  PromptState = "ENABLED"
	PromptStateDisabled PromptState = "DISABLED"
)

// Values returns all known values for PromptState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PromptState) Values() []PromptState {
	return []PromptState{
		"ENABLED",
		"DISABLED",
	}
}

type PromptType string

// Enum values for PromptType
const (
	PromptTypePreProcessing                   PromptType = "PRE_PROCESSING"
	PromptTypeOrchestration                   PromptType = "ORCHESTRATION"
	PromptTypePostProcessing                  PromptType = "POST_PROCESSING"
	PromptTypeKnowledgeBaseResponseGeneration PromptType = "KNOWLEDGE_BASE_RESPONSE_GENERATION"
)

// Values returns all known values for PromptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PromptType) Values() []PromptType {
	return []PromptType{
		"PRE_PROCESSING",
		"ORCHESTRATION",
		"POST_PROCESSING",
		"KNOWLEDGE_BASE_RESPONSE_GENERATION",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type Type string

// Enum values for Type
const (
	TypeString  Type = "string"
	TypeNumber  Type = "number"
	TypeInteger Type = "integer"
	TypeBoolean Type = "boolean"
	TypeArray   Type = "array"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"string",
		"number",
		"integer",
		"boolean",
		"array",
	}
}
