// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdminStatus string

// Enum values for AdminStatus
const (
	AdminStatusEnabled           AdminStatus = "ENABLED"
	AdminStatusDisableInProgress AdminStatus = "DISABLE_IN_PROGRESS"
)

// Values returns all known values for AdminStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdminStatus) Values() []AdminStatus {
	return []AdminStatus{
		"ENABLED",
		"DISABLE_IN_PROGRESS",
	}
}

type AssociationStatus string

// Enum values for AssociationStatus
const (
	AssociationStatusEnabled  AssociationStatus = "ENABLED"
	AssociationStatusDisabled AssociationStatus = "DISABLED"
)

// Values returns all known values for AssociationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssociationStatus) Values() []AssociationStatus {
	return []AssociationStatus{
		"ENABLED",
		"DISABLED",
	}
}

type AssociationType string

// Enum values for AssociationType
const (
	AssociationTypeInherited AssociationType = "INHERITED"
	AssociationTypeApplied   AssociationType = "APPLIED"
)

// Values returns all known values for AssociationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssociationType) Values() []AssociationType {
	return []AssociationType{
		"INHERITED",
		"APPLIED",
	}
}

type AutoEnableStandards string

// Enum values for AutoEnableStandards
const (
	AutoEnableStandardsNone    AutoEnableStandards = "NONE"
	AutoEnableStandardsDefault AutoEnableStandards = "DEFAULT"
)

// Values returns all known values for AutoEnableStandards. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutoEnableStandards) Values() []AutoEnableStandards {
	return []AutoEnableStandards{
		"NONE",
		"DEFAULT",
	}
}

type AutomationRulesActionType string

// Enum values for AutomationRulesActionType
const (
	AutomationRulesActionTypeFindingFieldsUpdate AutomationRulesActionType = "FINDING_FIELDS_UPDATE"
)

// Values returns all known values for AutomationRulesActionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutomationRulesActionType) Values() []AutomationRulesActionType {
	return []AutomationRulesActionType{
		"FINDING_FIELDS_UPDATE",
	}
}

type AwsIamAccessKeyStatus string

// Enum values for AwsIamAccessKeyStatus
const (
	AwsIamAccessKeyStatusActive   AwsIamAccessKeyStatus = "Active"
	AwsIamAccessKeyStatusInactive AwsIamAccessKeyStatus = "Inactive"
)

// Values returns all known values for AwsIamAccessKeyStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AwsIamAccessKeyStatus) Values() []AwsIamAccessKeyStatus {
	return []AwsIamAccessKeyStatus{
		"Active",
		"Inactive",
	}
}

type AwsS3BucketNotificationConfigurationS3KeyFilterRuleName string

// Enum values for AwsS3BucketNotificationConfigurationS3KeyFilterRuleName
const (
	AwsS3BucketNotificationConfigurationS3KeyFilterRuleNamePrefix AwsS3BucketNotificationConfigurationS3KeyFilterRuleName = "Prefix"
	AwsS3BucketNotificationConfigurationS3KeyFilterRuleNameSuffix AwsS3BucketNotificationConfigurationS3KeyFilterRuleName = "Suffix"
)

// Values returns all known values for
// AwsS3BucketNotificationConfigurationS3KeyFilterRuleName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AwsS3BucketNotificationConfigurationS3KeyFilterRuleName) Values() []AwsS3BucketNotificationConfigurationS3KeyFilterRuleName {
	return []AwsS3BucketNotificationConfigurationS3KeyFilterRuleName{
		"Prefix",
		"Suffix",
	}
}

type ComplianceStatus string

// Enum values for ComplianceStatus
const (
	ComplianceStatusPassed       ComplianceStatus = "PASSED"
	ComplianceStatusWarning      ComplianceStatus = "WARNING"
	ComplianceStatusFailed       ComplianceStatus = "FAILED"
	ComplianceStatusNotAvailable ComplianceStatus = "NOT_AVAILABLE"
)

// Values returns all known values for ComplianceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ComplianceStatus) Values() []ComplianceStatus {
	return []ComplianceStatus{
		"PASSED",
		"WARNING",
		"FAILED",
		"NOT_AVAILABLE",
	}
}

type ConfigurationPolicyAssociationStatus string

// Enum values for ConfigurationPolicyAssociationStatus
const (
	ConfigurationPolicyAssociationStatusPending ConfigurationPolicyAssociationStatus = "PENDING"
	ConfigurationPolicyAssociationStatusSuccess ConfigurationPolicyAssociationStatus = "SUCCESS"
	ConfigurationPolicyAssociationStatusFailed  ConfigurationPolicyAssociationStatus = "FAILED"
)

// Values returns all known values for ConfigurationPolicyAssociationStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationPolicyAssociationStatus) Values() []ConfigurationPolicyAssociationStatus {
	return []ConfigurationPolicyAssociationStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
	}
}

type ControlFindingGenerator string

// Enum values for ControlFindingGenerator
const (
	ControlFindingGeneratorStandardControl ControlFindingGenerator = "STANDARD_CONTROL"
	ControlFindingGeneratorSecurityControl ControlFindingGenerator = "SECURITY_CONTROL"
)

// Values returns all known values for ControlFindingGenerator. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ControlFindingGenerator) Values() []ControlFindingGenerator {
	return []ControlFindingGenerator{
		"STANDARD_CONTROL",
		"SECURITY_CONTROL",
	}
}

type ControlStatus string

// Enum values for ControlStatus
const (
	ControlStatusEnabled  ControlStatus = "ENABLED"
	ControlStatusDisabled ControlStatus = "DISABLED"
)

// Values returns all known values for ControlStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ControlStatus) Values() []ControlStatus {
	return []ControlStatus{
		"ENABLED",
		"DISABLED",
	}
}

type DateRangeUnit string

// Enum values for DateRangeUnit
const (
	DateRangeUnitDays DateRangeUnit = "DAYS"
)

// Values returns all known values for DateRangeUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DateRangeUnit) Values() []DateRangeUnit {
	return []DateRangeUnit{
		"DAYS",
	}
}

type FindingHistoryUpdateSourceType string

// Enum values for FindingHistoryUpdateSourceType
const (
	FindingHistoryUpdateSourceTypeBatchUpdateFindings FindingHistoryUpdateSourceType = "BATCH_UPDATE_FINDINGS"
	FindingHistoryUpdateSourceTypeBatchImportFindings FindingHistoryUpdateSourceType = "BATCH_IMPORT_FINDINGS"
)

// Values returns all known values for FindingHistoryUpdateSourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingHistoryUpdateSourceType) Values() []FindingHistoryUpdateSourceType {
	return []FindingHistoryUpdateSourceType{
		"BATCH_UPDATE_FINDINGS",
		"BATCH_IMPORT_FINDINGS",
	}
}

type IntegrationType string

// Enum values for IntegrationType
const (
	IntegrationTypeSendFindingsToSecurityHub      IntegrationType = "SEND_FINDINGS_TO_SECURITY_HUB"
	IntegrationTypeReceiveFindingsFromSecurityHub IntegrationType = "RECEIVE_FINDINGS_FROM_SECURITY_HUB"
	IntegrationTypeUpdateFindingsInSecurityHub    IntegrationType = "UPDATE_FINDINGS_IN_SECURITY_HUB"
)

// Values returns all known values for IntegrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationType) Values() []IntegrationType {
	return []IntegrationType{
		"SEND_FINDINGS_TO_SECURITY_HUB",
		"RECEIVE_FINDINGS_FROM_SECURITY_HUB",
		"UPDATE_FINDINGS_IN_SECURITY_HUB",
	}
}

type MalwareState string

// Enum values for MalwareState
const (
	MalwareStateObserved      MalwareState = "OBSERVED"
	MalwareStateRemovalFailed MalwareState = "REMOVAL_FAILED"
	MalwareStateRemoved       MalwareState = "REMOVED"
)

// Values returns all known values for MalwareState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MalwareState) Values() []MalwareState {
	return []MalwareState{
		"OBSERVED",
		"REMOVAL_FAILED",
		"REMOVED",
	}
}

type MalwareType string

// Enum values for MalwareType
const (
	MalwareTypeAdware              MalwareType = "ADWARE"
	MalwareTypeBlendedThreat       MalwareType = "BLENDED_THREAT"
	MalwareTypeBotnetAgent         MalwareType = "BOTNET_AGENT"
	MalwareTypeCoinMiner           MalwareType = "COIN_MINER"
	MalwareTypeExploitKit          MalwareType = "EXPLOIT_KIT"
	MalwareTypeKeylogger           MalwareType = "KEYLOGGER"
	MalwareTypeMacro               MalwareType = "MACRO"
	MalwareTypePotentiallyUnwanted MalwareType = "POTENTIALLY_UNWANTED"
	MalwareTypeSpyware             MalwareType = "SPYWARE"
	MalwareTypeRansomware          MalwareType = "RANSOMWARE"
	MalwareTypeRemoteAccess        MalwareType = "REMOTE_ACCESS"
	MalwareTypeRootkit             MalwareType = "ROOTKIT"
	MalwareTypeTrojan              MalwareType = "TROJAN"
	MalwareTypeVirus               MalwareType = "VIRUS"
	MalwareTypeWorm                MalwareType = "WORM"
)

// Values returns all known values for MalwareType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MalwareType) Values() []MalwareType {
	return []MalwareType{
		"ADWARE",
		"BLENDED_THREAT",
		"BOTNET_AGENT",
		"COIN_MINER",
		"EXPLOIT_KIT",
		"KEYLOGGER",
		"MACRO",
		"POTENTIALLY_UNWANTED",
		"SPYWARE",
		"RANSOMWARE",
		"REMOTE_ACCESS",
		"ROOTKIT",
		"TROJAN",
		"VIRUS",
		"WORM",
	}
}

type MapFilterComparison string

// Enum values for MapFilterComparison
const (
	MapFilterComparisonEquals      MapFilterComparison = "EQUALS"
	MapFilterComparisonNotEquals   MapFilterComparison = "NOT_EQUALS"
	MapFilterComparisonContains    MapFilterComparison = "CONTAINS"
	MapFilterComparisonNotContains MapFilterComparison = "NOT_CONTAINS"
)

// Values returns all known values for MapFilterComparison. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MapFilterComparison) Values() []MapFilterComparison {
	return []MapFilterComparison{
		"EQUALS",
		"NOT_EQUALS",
		"CONTAINS",
		"NOT_CONTAINS",
	}
}

type NetworkDirection string

// Enum values for NetworkDirection
const (
	NetworkDirectionIn  NetworkDirection = "IN"
	NetworkDirectionOut NetworkDirection = "OUT"
)

// Values returns all known values for NetworkDirection. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NetworkDirection) Values() []NetworkDirection {
	return []NetworkDirection{
		"IN",
		"OUT",
	}
}

type OrganizationConfigurationConfigurationType string

// Enum values for OrganizationConfigurationConfigurationType
const (
	OrganizationConfigurationConfigurationTypeCentral OrganizationConfigurationConfigurationType = "CENTRAL"
	OrganizationConfigurationConfigurationTypeLocal   OrganizationConfigurationConfigurationType = "LOCAL"
)

// Values returns all known values for OrganizationConfigurationConfigurationType.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationConfigurationConfigurationType) Values() []OrganizationConfigurationConfigurationType {
	return []OrganizationConfigurationConfigurationType{
		"CENTRAL",
		"LOCAL",
	}
}

type OrganizationConfigurationStatus string

// Enum values for OrganizationConfigurationStatus
const (
	OrganizationConfigurationStatusPending OrganizationConfigurationStatus = "PENDING"
	OrganizationConfigurationStatusEnabled OrganizationConfigurationStatus = "ENABLED"
	OrganizationConfigurationStatusFailed  OrganizationConfigurationStatus = "FAILED"
)

// Values returns all known values for OrganizationConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationConfigurationStatus) Values() []OrganizationConfigurationStatus {
	return []OrganizationConfigurationStatus{
		"PENDING",
		"ENABLED",
		"FAILED",
	}
}

type ParameterValueType string

// Enum values for ParameterValueType
const (
	ParameterValueTypeDefault ParameterValueType = "DEFAULT"
	ParameterValueTypeCustom  ParameterValueType = "CUSTOM"
)

// Values returns all known values for ParameterValueType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParameterValueType) Values() []ParameterValueType {
	return []ParameterValueType{
		"DEFAULT",
		"CUSTOM",
	}
}

type Partition string

// Enum values for Partition
const (
	PartitionAws      Partition = "aws"
	PartitionAwsCn    Partition = "aws-cn"
	PartitionAwsUsGov Partition = "aws-us-gov"
)

// Values returns all known values for Partition. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Partition) Values() []Partition {
	return []Partition{
		"aws",
		"aws-cn",
		"aws-us-gov",
	}
}

type RecordState string

// Enum values for RecordState
const (
	RecordStateActive   RecordState = "ACTIVE"
	RecordStateArchived RecordState = "ARCHIVED"
)

// Values returns all known values for RecordState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordState) Values() []RecordState {
	return []RecordState{
		"ACTIVE",
		"ARCHIVED",
	}
}

type RegionAvailabilityStatus string

// Enum values for RegionAvailabilityStatus
const (
	RegionAvailabilityStatusAvailable   RegionAvailabilityStatus = "AVAILABLE"
	RegionAvailabilityStatusUnavailable RegionAvailabilityStatus = "UNAVAILABLE"
)

// Values returns all known values for RegionAvailabilityStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RegionAvailabilityStatus) Values() []RegionAvailabilityStatus {
	return []RegionAvailabilityStatus{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

type RuleStatus string

// Enum values for RuleStatus
const (
	RuleStatusEnabled  RuleStatus = "ENABLED"
	RuleStatusDisabled RuleStatus = "DISABLED"
)

// Values returns all known values for RuleStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RuleStatus) Values() []RuleStatus {
	return []RuleStatus{
		"ENABLED",
		"DISABLED",
	}
}

type SecurityControlProperty string

// Enum values for SecurityControlProperty
const (
	SecurityControlPropertyParameters SecurityControlProperty = "Parameters"
)

// Values returns all known values for SecurityControlProperty. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SecurityControlProperty) Values() []SecurityControlProperty {
	return []SecurityControlProperty{
		"Parameters",
	}
}

type SeverityLabel string

// Enum values for SeverityLabel
const (
	SeverityLabelInformational SeverityLabel = "INFORMATIONAL"
	SeverityLabelLow           SeverityLabel = "LOW"
	SeverityLabelMedium        SeverityLabel = "MEDIUM"
	SeverityLabelHigh          SeverityLabel = "HIGH"
	SeverityLabelCritical      SeverityLabel = "CRITICAL"
)

// Values returns all known values for SeverityLabel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SeverityLabel) Values() []SeverityLabel {
	return []SeverityLabel{
		"INFORMATIONAL",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

type SeverityRating string

// Enum values for SeverityRating
const (
	SeverityRatingLow      SeverityRating = "LOW"
	SeverityRatingMedium   SeverityRating = "MEDIUM"
	SeverityRatingHigh     SeverityRating = "HIGH"
	SeverityRatingCritical SeverityRating = "CRITICAL"
)

// Values returns all known values for SeverityRating. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SeverityRating) Values() []SeverityRating {
	return []SeverityRating{
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"asc",
		"desc",
	}
}

type StandardsStatus string

// Enum values for StandardsStatus
const (
	StandardsStatusPending    StandardsStatus = "PENDING"
	StandardsStatusReady      StandardsStatus = "READY"
	StandardsStatusFailed     StandardsStatus = "FAILED"
	StandardsStatusDeleting   StandardsStatus = "DELETING"
	StandardsStatusIncomplete StandardsStatus = "INCOMPLETE"
)

// Values returns all known values for StandardsStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StandardsStatus) Values() []StandardsStatus {
	return []StandardsStatus{
		"PENDING",
		"READY",
		"FAILED",
		"DELETING",
		"INCOMPLETE",
	}
}

type StatusReasonCode string

// Enum values for StatusReasonCode
const (
	StatusReasonCodeNoAvailableConfigurationRecorder StatusReasonCode = "NO_AVAILABLE_CONFIGURATION_RECORDER"
	StatusReasonCodeInternalError                    StatusReasonCode = "INTERNAL_ERROR"
)

// Values returns all known values for StatusReasonCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StatusReasonCode) Values() []StatusReasonCode {
	return []StatusReasonCode{
		"NO_AVAILABLE_CONFIGURATION_RECORDER",
		"INTERNAL_ERROR",
	}
}

type StringFilterComparison string

// Enum values for StringFilterComparison
const (
	StringFilterComparisonEquals          StringFilterComparison = "EQUALS"
	StringFilterComparisonPrefix          StringFilterComparison = "PREFIX"
	StringFilterComparisonNotEquals       StringFilterComparison = "NOT_EQUALS"
	StringFilterComparisonPrefixNotEquals StringFilterComparison = "PREFIX_NOT_EQUALS"
	StringFilterComparisonContains        StringFilterComparison = "CONTAINS"
	StringFilterComparisonNotContains     StringFilterComparison = "NOT_CONTAINS"
)

// Values returns all known values for StringFilterComparison. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StringFilterComparison) Values() []StringFilterComparison {
	return []StringFilterComparison{
		"EQUALS",
		"PREFIX",
		"NOT_EQUALS",
		"PREFIX_NOT_EQUALS",
		"CONTAINS",
		"NOT_CONTAINS",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeAccount            TargetType = "ACCOUNT"
	TargetTypeOrganizationalUnit TargetType = "ORGANIZATIONAL_UNIT"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"ACCOUNT",
		"ORGANIZATIONAL_UNIT",
	}
}

type ThreatIntelIndicatorCategory string

// Enum values for ThreatIntelIndicatorCategory
const (
	ThreatIntelIndicatorCategoryBackdoor          ThreatIntelIndicatorCategory = "BACKDOOR"
	ThreatIntelIndicatorCategoryCardStealer       ThreatIntelIndicatorCategory = "CARD_STEALER"
	ThreatIntelIndicatorCategoryCommandAndControl ThreatIntelIndicatorCategory = "COMMAND_AND_CONTROL"
	ThreatIntelIndicatorCategoryDropSite          ThreatIntelIndicatorCategory = "DROP_SITE"
	ThreatIntelIndicatorCategoryExploitSite       ThreatIntelIndicatorCategory = "EXPLOIT_SITE"
	ThreatIntelIndicatorCategoryKeylogger         ThreatIntelIndicatorCategory = "KEYLOGGER"
)

// Values returns all known values for ThreatIntelIndicatorCategory. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelIndicatorCategory) Values() []ThreatIntelIndicatorCategory {
	return []ThreatIntelIndicatorCategory{
		"BACKDOOR",
		"CARD_STEALER",
		"COMMAND_AND_CONTROL",
		"DROP_SITE",
		"EXPLOIT_SITE",
		"KEYLOGGER",
	}
}

type ThreatIntelIndicatorType string

// Enum values for ThreatIntelIndicatorType
const (
	ThreatIntelIndicatorTypeDomain       ThreatIntelIndicatorType = "DOMAIN"
	ThreatIntelIndicatorTypeEmailAddress ThreatIntelIndicatorType = "EMAIL_ADDRESS"
	ThreatIntelIndicatorTypeHashMd5      ThreatIntelIndicatorType = "HASH_MD5"
	ThreatIntelIndicatorTypeHashSha1     ThreatIntelIndicatorType = "HASH_SHA1"
	ThreatIntelIndicatorTypeHashSha256   ThreatIntelIndicatorType = "HASH_SHA256"
	ThreatIntelIndicatorTypeHashSha512   ThreatIntelIndicatorType = "HASH_SHA512"
	ThreatIntelIndicatorTypeIpv4Address  ThreatIntelIndicatorType = "IPV4_ADDRESS"
	ThreatIntelIndicatorTypeIpv6Address  ThreatIntelIndicatorType = "IPV6_ADDRESS"
	ThreatIntelIndicatorTypeMutex        ThreatIntelIndicatorType = "MUTEX"
	ThreatIntelIndicatorTypeProcess      ThreatIntelIndicatorType = "PROCESS"
	ThreatIntelIndicatorTypeUrl          ThreatIntelIndicatorType = "URL"
)

// Values returns all known values for ThreatIntelIndicatorType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelIndicatorType) Values() []ThreatIntelIndicatorType {
	return []ThreatIntelIndicatorType{
		"DOMAIN",
		"EMAIL_ADDRESS",
		"HASH_MD5",
		"HASH_SHA1",
		"HASH_SHA256",
		"HASH_SHA512",
		"IPV4_ADDRESS",
		"IPV6_ADDRESS",
		"MUTEX",
		"PROCESS",
		"URL",
	}
}

type UnprocessedErrorCode string

// Enum values for UnprocessedErrorCode
const (
	UnprocessedErrorCodeInvalidInput  UnprocessedErrorCode = "INVALID_INPUT"
	UnprocessedErrorCodeAccessDenied  UnprocessedErrorCode = "ACCESS_DENIED"
	UnprocessedErrorCodeNotFound      UnprocessedErrorCode = "NOT_FOUND"
	UnprocessedErrorCodeLimitExceeded UnprocessedErrorCode = "LIMIT_EXCEEDED"
)

// Values returns all known values for UnprocessedErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UnprocessedErrorCode) Values() []UnprocessedErrorCode {
	return []UnprocessedErrorCode{
		"INVALID_INPUT",
		"ACCESS_DENIED",
		"NOT_FOUND",
		"LIMIT_EXCEEDED",
	}
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusReady    UpdateStatus = "READY"
	UpdateStatusUpdating UpdateStatus = "UPDATING"
)

// Values returns all known values for UpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateStatus) Values() []UpdateStatus {
	return []UpdateStatus{
		"READY",
		"UPDATING",
	}
}

type VerificationState string

// Enum values for VerificationState
const (
	VerificationStateUnknown        VerificationState = "UNKNOWN"
	VerificationStateTruePositive   VerificationState = "TRUE_POSITIVE"
	VerificationStateFalsePositive  VerificationState = "FALSE_POSITIVE"
	VerificationStateBenignPositive VerificationState = "BENIGN_POSITIVE"
)

// Values returns all known values for VerificationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VerificationState) Values() []VerificationState {
	return []VerificationState{
		"UNKNOWN",
		"TRUE_POSITIVE",
		"FALSE_POSITIVE",
		"BENIGN_POSITIVE",
	}
}

type VulnerabilityExploitAvailable string

// Enum values for VulnerabilityExploitAvailable
const (
	VulnerabilityExploitAvailableYes VulnerabilityExploitAvailable = "YES"
	VulnerabilityExploitAvailableNo  VulnerabilityExploitAvailable = "NO"
)

// Values returns all known values for VulnerabilityExploitAvailable. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VulnerabilityExploitAvailable) Values() []VulnerabilityExploitAvailable {
	return []VulnerabilityExploitAvailable{
		"YES",
		"NO",
	}
}

type VulnerabilityFixAvailable string

// Enum values for VulnerabilityFixAvailable
const (
	VulnerabilityFixAvailableYes     VulnerabilityFixAvailable = "YES"
	VulnerabilityFixAvailableNo      VulnerabilityFixAvailable = "NO"
	VulnerabilityFixAvailablePartial VulnerabilityFixAvailable = "PARTIAL"
)

// Values returns all known values for VulnerabilityFixAvailable. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VulnerabilityFixAvailable) Values() []VulnerabilityFixAvailable {
	return []VulnerabilityFixAvailable{
		"YES",
		"NO",
		"PARTIAL",
	}
}

type WorkflowState string

// Enum values for WorkflowState
const (
	WorkflowStateNew        WorkflowState = "NEW"
	WorkflowStateAssigned   WorkflowState = "ASSIGNED"
	WorkflowStateInProgress WorkflowState = "IN_PROGRESS"
	WorkflowStateDeferred   WorkflowState = "DEFERRED"
	WorkflowStateResolved   WorkflowState = "RESOLVED"
)

// Values returns all known values for WorkflowState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowState) Values() []WorkflowState {
	return []WorkflowState{
		"NEW",
		"ASSIGNED",
		"IN_PROGRESS",
		"DEFERRED",
		"RESOLVED",
	}
}

type WorkflowStatus string

// Enum values for WorkflowStatus
const (
	WorkflowStatusNew        WorkflowStatus = "NEW"
	WorkflowStatusNotified   WorkflowStatus = "NOTIFIED"
	WorkflowStatusResolved   WorkflowStatus = "RESOLVED"
	WorkflowStatusSuppressed WorkflowStatus = "SUPPRESSED"
)

// Values returns all known values for WorkflowStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowStatus) Values() []WorkflowStatus {
	return []WorkflowStatus{
		"NEW",
		"NOTIFIED",
		"RESOLVED",
		"SUPPRESSED",
	}
}
