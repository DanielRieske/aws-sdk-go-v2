// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApiCacheStatus string

// Enum values for ApiCacheStatus
const (
	ApiCacheStatusAvailable ApiCacheStatus = "AVAILABLE"
	ApiCacheStatusCreating  ApiCacheStatus = "CREATING"
	ApiCacheStatusDeleting  ApiCacheStatus = "DELETING"
	ApiCacheStatusModifying ApiCacheStatus = "MODIFYING"
	ApiCacheStatusFailed    ApiCacheStatus = "FAILED"
)

// Values returns all known values for ApiCacheStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApiCacheStatus) Values() []ApiCacheStatus {
	return []ApiCacheStatus{
		"AVAILABLE",
		"CREATING",
		"DELETING",
		"MODIFYING",
		"FAILED",
	}
}

type ApiCacheType string

// Enum values for ApiCacheType
const (
	ApiCacheTypeT2Small   ApiCacheType = "T2_SMALL"
	ApiCacheTypeT2Medium  ApiCacheType = "T2_MEDIUM"
	ApiCacheTypeR4Large   ApiCacheType = "R4_LARGE"
	ApiCacheTypeR4Xlarge  ApiCacheType = "R4_XLARGE"
	ApiCacheTypeR42xlarge ApiCacheType = "R4_2XLARGE"
	ApiCacheTypeR44xlarge ApiCacheType = "R4_4XLARGE"
	ApiCacheTypeR48xlarge ApiCacheType = "R4_8XLARGE"
	ApiCacheTypeSmall     ApiCacheType = "SMALL"
	ApiCacheTypeMedium    ApiCacheType = "MEDIUM"
	ApiCacheTypeLarge     ApiCacheType = "LARGE"
	ApiCacheTypeXlarge    ApiCacheType = "XLARGE"
	ApiCacheTypeLarge2x   ApiCacheType = "LARGE_2X"
	ApiCacheTypeLarge4x   ApiCacheType = "LARGE_4X"
	ApiCacheTypeLarge8x   ApiCacheType = "LARGE_8X"
	ApiCacheTypeLarge12x  ApiCacheType = "LARGE_12X"
)

// Values returns all known values for ApiCacheType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApiCacheType) Values() []ApiCacheType {
	return []ApiCacheType{
		"T2_SMALL",
		"T2_MEDIUM",
		"R4_LARGE",
		"R4_XLARGE",
		"R4_2XLARGE",
		"R4_4XLARGE",
		"R4_8XLARGE",
		"SMALL",
		"MEDIUM",
		"LARGE",
		"XLARGE",
		"LARGE_2X",
		"LARGE_4X",
		"LARGE_8X",
		"LARGE_12X",
	}
}

type ApiCachingBehavior string

// Enum values for ApiCachingBehavior
const (
	ApiCachingBehaviorFullRequestCaching ApiCachingBehavior = "FULL_REQUEST_CACHING"
	ApiCachingBehaviorPerResolverCaching ApiCachingBehavior = "PER_RESOLVER_CACHING"
)

// Values returns all known values for ApiCachingBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApiCachingBehavior) Values() []ApiCachingBehavior {
	return []ApiCachingBehavior{
		"FULL_REQUEST_CACHING",
		"PER_RESOLVER_CACHING",
	}
}

type AssociationStatus string

// Enum values for AssociationStatus
const (
	AssociationStatusProcessing AssociationStatus = "PROCESSING"
	AssociationStatusFailed     AssociationStatus = "FAILED"
	AssociationStatusSuccess    AssociationStatus = "SUCCESS"
)

// Values returns all known values for AssociationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssociationStatus) Values() []AssociationStatus {
	return []AssociationStatus{
		"PROCESSING",
		"FAILED",
		"SUCCESS",
	}
}

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypeApiKey                 AuthenticationType = "API_KEY"
	AuthenticationTypeAwsIam                 AuthenticationType = "AWS_IAM"
	AuthenticationTypeAmazonCognitoUserPools AuthenticationType = "AMAZON_COGNITO_USER_POOLS"
	AuthenticationTypeOpenidConnect          AuthenticationType = "OPENID_CONNECT"
	AuthenticationTypeAwsLambda              AuthenticationType = "AWS_LAMBDA"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"API_KEY",
		"AWS_IAM",
		"AMAZON_COGNITO_USER_POOLS",
		"OPENID_CONNECT",
		"AWS_LAMBDA",
	}
}

type AuthorizationType string

// Enum values for AuthorizationType
const (
	AuthorizationTypeAwsIam AuthorizationType = "AWS_IAM"
)

// Values returns all known values for AuthorizationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthorizationType) Values() []AuthorizationType {
	return []AuthorizationType{
		"AWS_IAM",
	}
}

type BadRequestReason string

// Enum values for BadRequestReason
const (
	BadRequestReasonCodeError BadRequestReason = "CODE_ERROR"
)

// Values returns all known values for BadRequestReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BadRequestReason) Values() []BadRequestReason {
	return []BadRequestReason{
		"CODE_ERROR",
	}
}

type CacheHealthMetricsConfig string

// Enum values for CacheHealthMetricsConfig
const (
	CacheHealthMetricsConfigEnabled  CacheHealthMetricsConfig = "ENABLED"
	CacheHealthMetricsConfigDisabled CacheHealthMetricsConfig = "DISABLED"
)

// Values returns all known values for CacheHealthMetricsConfig. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CacheHealthMetricsConfig) Values() []CacheHealthMetricsConfig {
	return []CacheHealthMetricsConfig{
		"ENABLED",
		"DISABLED",
	}
}

type ConflictDetectionType string

// Enum values for ConflictDetectionType
const (
	ConflictDetectionTypeVersion ConflictDetectionType = "VERSION"
	ConflictDetectionTypeNone    ConflictDetectionType = "NONE"
)

// Values returns all known values for ConflictDetectionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictDetectionType) Values() []ConflictDetectionType {
	return []ConflictDetectionType{
		"VERSION",
		"NONE",
	}
}

type ConflictHandlerType string

// Enum values for ConflictHandlerType
const (
	ConflictHandlerTypeOptimisticConcurrency ConflictHandlerType = "OPTIMISTIC_CONCURRENCY"
	ConflictHandlerTypeLambda                ConflictHandlerType = "LAMBDA"
	ConflictHandlerTypeAutomerge             ConflictHandlerType = "AUTOMERGE"
	ConflictHandlerTypeNone                  ConflictHandlerType = "NONE"
)

// Values returns all known values for ConflictHandlerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictHandlerType) Values() []ConflictHandlerType {
	return []ConflictHandlerType{
		"OPTIMISTIC_CONCURRENCY",
		"LAMBDA",
		"AUTOMERGE",
		"NONE",
	}
}

type DataSourceIntrospectionStatus string

// Enum values for DataSourceIntrospectionStatus
const (
	DataSourceIntrospectionStatusProcessing DataSourceIntrospectionStatus = "PROCESSING"
	DataSourceIntrospectionStatusFailed     DataSourceIntrospectionStatus = "FAILED"
	DataSourceIntrospectionStatusSuccess    DataSourceIntrospectionStatus = "SUCCESS"
)

// Values returns all known values for DataSourceIntrospectionStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceIntrospectionStatus) Values() []DataSourceIntrospectionStatus {
	return []DataSourceIntrospectionStatus{
		"PROCESSING",
		"FAILED",
		"SUCCESS",
	}
}

type DataSourceLevelMetricsBehavior string

// Enum values for DataSourceLevelMetricsBehavior
const (
	DataSourceLevelMetricsBehaviorFullRequestDataSourceMetrics DataSourceLevelMetricsBehavior = "FULL_REQUEST_DATA_SOURCE_METRICS"
	DataSourceLevelMetricsBehaviorPerDataSourceMetrics         DataSourceLevelMetricsBehavior = "PER_DATA_SOURCE_METRICS"
)

// Values returns all known values for DataSourceLevelMetricsBehavior. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceLevelMetricsBehavior) Values() []DataSourceLevelMetricsBehavior {
	return []DataSourceLevelMetricsBehavior{
		"FULL_REQUEST_DATA_SOURCE_METRICS",
		"PER_DATA_SOURCE_METRICS",
	}
}

type DataSourceLevelMetricsConfig string

// Enum values for DataSourceLevelMetricsConfig
const (
	DataSourceLevelMetricsConfigEnabled  DataSourceLevelMetricsConfig = "ENABLED"
	DataSourceLevelMetricsConfigDisabled DataSourceLevelMetricsConfig = "DISABLED"
)

// Values returns all known values for DataSourceLevelMetricsConfig. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceLevelMetricsConfig) Values() []DataSourceLevelMetricsConfig {
	return []DataSourceLevelMetricsConfig{
		"ENABLED",
		"DISABLED",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeAwsLambda               DataSourceType = "AWS_LAMBDA"
	DataSourceTypeAmazonDynamodb          DataSourceType = "AMAZON_DYNAMODB"
	DataSourceTypeAmazonElasticsearch     DataSourceType = "AMAZON_ELASTICSEARCH"
	DataSourceTypeNone                    DataSourceType = "NONE"
	DataSourceTypeHttp                    DataSourceType = "HTTP"
	DataSourceTypeRelationalDatabase      DataSourceType = "RELATIONAL_DATABASE"
	DataSourceTypeAmazonOpensearchService DataSourceType = "AMAZON_OPENSEARCH_SERVICE"
	DataSourceTypeAmazonEventbridge       DataSourceType = "AMAZON_EVENTBRIDGE"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"AWS_LAMBDA",
		"AMAZON_DYNAMODB",
		"AMAZON_ELASTICSEARCH",
		"NONE",
		"HTTP",
		"RELATIONAL_DATABASE",
		"AMAZON_OPENSEARCH_SERVICE",
		"AMAZON_EVENTBRIDGE",
	}
}

type DefaultAction string

// Enum values for DefaultAction
const (
	DefaultActionAllow DefaultAction = "ALLOW"
	DefaultActionDeny  DefaultAction = "DENY"
)

// Values returns all known values for DefaultAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DefaultAction) Values() []DefaultAction {
	return []DefaultAction{
		"ALLOW",
		"DENY",
	}
}

type FieldLogLevel string

// Enum values for FieldLogLevel
const (
	FieldLogLevelNone  FieldLogLevel = "NONE"
	FieldLogLevelError FieldLogLevel = "ERROR"
	FieldLogLevelAll   FieldLogLevel = "ALL"
)

// Values returns all known values for FieldLogLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FieldLogLevel) Values() []FieldLogLevel {
	return []FieldLogLevel{
		"NONE",
		"ERROR",
		"ALL",
	}
}

type GraphQLApiIntrospectionConfig string

// Enum values for GraphQLApiIntrospectionConfig
const (
	GraphQLApiIntrospectionConfigEnabled  GraphQLApiIntrospectionConfig = "ENABLED"
	GraphQLApiIntrospectionConfigDisabled GraphQLApiIntrospectionConfig = "DISABLED"
)

// Values returns all known values for GraphQLApiIntrospectionConfig. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GraphQLApiIntrospectionConfig) Values() []GraphQLApiIntrospectionConfig {
	return []GraphQLApiIntrospectionConfig{
		"ENABLED",
		"DISABLED",
	}
}

type GraphQLApiType string

// Enum values for GraphQLApiType
const (
	GraphQLApiTypeGraphql GraphQLApiType = "GRAPHQL"
	GraphQLApiTypeMerged  GraphQLApiType = "MERGED"
)

// Values returns all known values for GraphQLApiType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GraphQLApiType) Values() []GraphQLApiType {
	return []GraphQLApiType{
		"GRAPHQL",
		"MERGED",
	}
}

type GraphQLApiVisibility string

// Enum values for GraphQLApiVisibility
const (
	GraphQLApiVisibilityGlobal  GraphQLApiVisibility = "GLOBAL"
	GraphQLApiVisibilityPrivate GraphQLApiVisibility = "PRIVATE"
)

// Values returns all known values for GraphQLApiVisibility. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GraphQLApiVisibility) Values() []GraphQLApiVisibility {
	return []GraphQLApiVisibility{
		"GLOBAL",
		"PRIVATE",
	}
}

type MergeType string

// Enum values for MergeType
const (
	MergeTypeManualMerge MergeType = "MANUAL_MERGE"
	MergeTypeAutoMerge   MergeType = "AUTO_MERGE"
)

// Values returns all known values for MergeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MergeType) Values() []MergeType {
	return []MergeType{
		"MANUAL_MERGE",
		"AUTO_MERGE",
	}
}

type OperationLevelMetricsConfig string

// Enum values for OperationLevelMetricsConfig
const (
	OperationLevelMetricsConfigEnabled  OperationLevelMetricsConfig = "ENABLED"
	OperationLevelMetricsConfigDisabled OperationLevelMetricsConfig = "DISABLED"
)

// Values returns all known values for OperationLevelMetricsConfig. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OperationLevelMetricsConfig) Values() []OperationLevelMetricsConfig {
	return []OperationLevelMetricsConfig{
		"ENABLED",
		"DISABLED",
	}
}

type OutputType string

// Enum values for OutputType
const (
	OutputTypeSdl  OutputType = "SDL"
	OutputTypeJson OutputType = "JSON"
)

// Values returns all known values for OutputType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutputType) Values() []OutputType {
	return []OutputType{
		"SDL",
		"JSON",
	}
}

type Ownership string

// Enum values for Ownership
const (
	OwnershipCurrentAccount Ownership = "CURRENT_ACCOUNT"
	OwnershipOtherAccounts  Ownership = "OTHER_ACCOUNTS"
)

// Values returns all known values for Ownership. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Ownership) Values() []Ownership {
	return []Ownership{
		"CURRENT_ACCOUNT",
		"OTHER_ACCOUNTS",
	}
}

type RelationalDatabaseSourceType string

// Enum values for RelationalDatabaseSourceType
const (
	RelationalDatabaseSourceTypeRdsHttpEndpoint RelationalDatabaseSourceType = "RDS_HTTP_ENDPOINT"
)

// Values returns all known values for RelationalDatabaseSourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RelationalDatabaseSourceType) Values() []RelationalDatabaseSourceType {
	return []RelationalDatabaseSourceType{
		"RDS_HTTP_ENDPOINT",
	}
}

type ResolverKind string

// Enum values for ResolverKind
const (
	ResolverKindUnit     ResolverKind = "UNIT"
	ResolverKindPipeline ResolverKind = "PIPELINE"
)

// Values returns all known values for ResolverKind. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverKind) Values() []ResolverKind {
	return []ResolverKind{
		"UNIT",
		"PIPELINE",
	}
}

type ResolverLevelMetricsBehavior string

// Enum values for ResolverLevelMetricsBehavior
const (
	ResolverLevelMetricsBehaviorFullRequestResolverMetrics ResolverLevelMetricsBehavior = "FULL_REQUEST_RESOLVER_METRICS"
	ResolverLevelMetricsBehaviorPerResolverMetrics         ResolverLevelMetricsBehavior = "PER_RESOLVER_METRICS"
)

// Values returns all known values for ResolverLevelMetricsBehavior. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverLevelMetricsBehavior) Values() []ResolverLevelMetricsBehavior {
	return []ResolverLevelMetricsBehavior{
		"FULL_REQUEST_RESOLVER_METRICS",
		"PER_RESOLVER_METRICS",
	}
}

type ResolverLevelMetricsConfig string

// Enum values for ResolverLevelMetricsConfig
const (
	ResolverLevelMetricsConfigEnabled  ResolverLevelMetricsConfig = "ENABLED"
	ResolverLevelMetricsConfigDisabled ResolverLevelMetricsConfig = "DISABLED"
)

// Values returns all known values for ResolverLevelMetricsConfig. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverLevelMetricsConfig) Values() []ResolverLevelMetricsConfig {
	return []ResolverLevelMetricsConfig{
		"ENABLED",
		"DISABLED",
	}
}

type RuntimeName string

// Enum values for RuntimeName
const (
	RuntimeNameAppsyncJs RuntimeName = "APPSYNC_JS"
)

// Values returns all known values for RuntimeName. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RuntimeName) Values() []RuntimeName {
	return []RuntimeName{
		"APPSYNC_JS",
	}
}

type SchemaStatus string

// Enum values for SchemaStatus
const (
	SchemaStatusProcessing    SchemaStatus = "PROCESSING"
	SchemaStatusActive        SchemaStatus = "ACTIVE"
	SchemaStatusDeleting      SchemaStatus = "DELETING"
	SchemaStatusFailed        SchemaStatus = "FAILED"
	SchemaStatusSuccess       SchemaStatus = "SUCCESS"
	SchemaStatusNotApplicable SchemaStatus = "NOT_APPLICABLE"
)

// Values returns all known values for SchemaStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaStatus) Values() []SchemaStatus {
	return []SchemaStatus{
		"PROCESSING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"SUCCESS",
		"NOT_APPLICABLE",
	}
}

type SourceApiAssociationStatus string

// Enum values for SourceApiAssociationStatus
const (
	SourceApiAssociationStatusMergeScheduled          SourceApiAssociationStatus = "MERGE_SCHEDULED"
	SourceApiAssociationStatusMergeFailed             SourceApiAssociationStatus = "MERGE_FAILED"
	SourceApiAssociationStatusMergeSuccess            SourceApiAssociationStatus = "MERGE_SUCCESS"
	SourceApiAssociationStatusMergeInProgress         SourceApiAssociationStatus = "MERGE_IN_PROGRESS"
	SourceApiAssociationStatusAutoMergeScheduleFailed SourceApiAssociationStatus = "AUTO_MERGE_SCHEDULE_FAILED"
	SourceApiAssociationStatusDeletionScheduled       SourceApiAssociationStatus = "DELETION_SCHEDULED"
	SourceApiAssociationStatusDeletionInProgress      SourceApiAssociationStatus = "DELETION_IN_PROGRESS"
	SourceApiAssociationStatusDeletionFailed          SourceApiAssociationStatus = "DELETION_FAILED"
)

// Values returns all known values for SourceApiAssociationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SourceApiAssociationStatus) Values() []SourceApiAssociationStatus {
	return []SourceApiAssociationStatus{
		"MERGE_SCHEDULED",
		"MERGE_FAILED",
		"MERGE_SUCCESS",
		"MERGE_IN_PROGRESS",
		"AUTO_MERGE_SCHEDULE_FAILED",
		"DELETION_SCHEDULED",
		"DELETION_IN_PROGRESS",
		"DELETION_FAILED",
	}
}

type TypeDefinitionFormat string

// Enum values for TypeDefinitionFormat
const (
	TypeDefinitionFormatSdl  TypeDefinitionFormat = "SDL"
	TypeDefinitionFormatJson TypeDefinitionFormat = "JSON"
)

// Values returns all known values for TypeDefinitionFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TypeDefinitionFormat) Values() []TypeDefinitionFormat {
	return []TypeDefinitionFormat{
		"SDL",
		"JSON",
	}
}
