// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApiKeysFormat string

// Enum values for ApiKeysFormat
const (
	ApiKeysFormatCsv ApiKeysFormat = "csv"
)

// Values returns all known values for ApiKeysFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApiKeysFormat) Values() []ApiKeysFormat {
	return []ApiKeysFormat{
		"csv",
	}
}

type ApiKeySourceType string

// Enum values for ApiKeySourceType
const (
	ApiKeySourceTypeHeader     ApiKeySourceType = "HEADER"
	ApiKeySourceTypeAuthorizer ApiKeySourceType = "AUTHORIZER"
)

// Values returns all known values for ApiKeySourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApiKeySourceType) Values() []ApiKeySourceType {
	return []ApiKeySourceType{
		"HEADER",
		"AUTHORIZER",
	}
}

type AuthorizerType string

// Enum values for AuthorizerType
const (
	AuthorizerTypeToken            AuthorizerType = "TOKEN"
	AuthorizerTypeRequest          AuthorizerType = "REQUEST"
	AuthorizerTypeCognitoUserPools AuthorizerType = "COGNITO_USER_POOLS"
)

// Values returns all known values for AuthorizerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthorizerType) Values() []AuthorizerType {
	return []AuthorizerType{
		"TOKEN",
		"REQUEST",
		"COGNITO_USER_POOLS",
	}
}

type CacheClusterSize string

// Enum values for CacheClusterSize
const (
	CacheClusterSizeSize0Point5Gb  CacheClusterSize = "0.5"
	CacheClusterSizeSize1Point6Gb  CacheClusterSize = "1.6"
	CacheClusterSizeSize6Point1Gb  CacheClusterSize = "6.1"
	CacheClusterSizeSize13Point5Gb CacheClusterSize = "13.5"
	CacheClusterSizeSize28Point4Gb CacheClusterSize = "28.4"
	CacheClusterSizeSize58Point2Gb CacheClusterSize = "58.2"
	CacheClusterSizeSize118Gb      CacheClusterSize = "118"
	CacheClusterSizeSize237Gb      CacheClusterSize = "237"
)

// Values returns all known values for CacheClusterSize. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CacheClusterSize) Values() []CacheClusterSize {
	return []CacheClusterSize{
		"0.5",
		"1.6",
		"6.1",
		"13.5",
		"28.4",
		"58.2",
		"118",
		"237",
	}
}

type CacheClusterStatus string

// Enum values for CacheClusterStatus
const (
	CacheClusterStatusCreateInProgress CacheClusterStatus = "CREATE_IN_PROGRESS"
	CacheClusterStatusAvailable        CacheClusterStatus = "AVAILABLE"
	CacheClusterStatusDeleteInProgress CacheClusterStatus = "DELETE_IN_PROGRESS"
	CacheClusterStatusNotAvailable     CacheClusterStatus = "NOT_AVAILABLE"
	CacheClusterStatusFlushInProgress  CacheClusterStatus = "FLUSH_IN_PROGRESS"
)

// Values returns all known values for CacheClusterStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CacheClusterStatus) Values() []CacheClusterStatus {
	return []CacheClusterStatus{
		"CREATE_IN_PROGRESS",
		"AVAILABLE",
		"DELETE_IN_PROGRESS",
		"NOT_AVAILABLE",
		"FLUSH_IN_PROGRESS",
	}
}

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeInternet ConnectionType = "INTERNET"
	ConnectionTypeVpcLink  ConnectionType = "VPC_LINK"
)

// Values returns all known values for ConnectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		"INTERNET",
		"VPC_LINK",
	}
}

type ContentHandlingStrategy string

// Enum values for ContentHandlingStrategy
const (
	ContentHandlingStrategyConvertToBinary ContentHandlingStrategy = "CONVERT_TO_BINARY"
	ContentHandlingStrategyConvertToText   ContentHandlingStrategy = "CONVERT_TO_TEXT"
)

// Values returns all known values for ContentHandlingStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContentHandlingStrategy) Values() []ContentHandlingStrategy {
	return []ContentHandlingStrategy{
		"CONVERT_TO_BINARY",
		"CONVERT_TO_TEXT",
	}
}

type DocumentationPartType string

// Enum values for DocumentationPartType
const (
	DocumentationPartTypeApi            DocumentationPartType = "API"
	DocumentationPartTypeAuthorizer     DocumentationPartType = "AUTHORIZER"
	DocumentationPartTypeModel          DocumentationPartType = "MODEL"
	DocumentationPartTypeResource       DocumentationPartType = "RESOURCE"
	DocumentationPartTypeMethod         DocumentationPartType = "METHOD"
	DocumentationPartTypePathParameter  DocumentationPartType = "PATH_PARAMETER"
	DocumentationPartTypeQueryParameter DocumentationPartType = "QUERY_PARAMETER"
	DocumentationPartTypeRequestHeader  DocumentationPartType = "REQUEST_HEADER"
	DocumentationPartTypeRequestBody    DocumentationPartType = "REQUEST_BODY"
	DocumentationPartTypeResponse       DocumentationPartType = "RESPONSE"
	DocumentationPartTypeResponseHeader DocumentationPartType = "RESPONSE_HEADER"
	DocumentationPartTypeResponseBody   DocumentationPartType = "RESPONSE_BODY"
)

// Values returns all known values for DocumentationPartType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DocumentationPartType) Values() []DocumentationPartType {
	return []DocumentationPartType{
		"API",
		"AUTHORIZER",
		"MODEL",
		"RESOURCE",
		"METHOD",
		"PATH_PARAMETER",
		"QUERY_PARAMETER",
		"REQUEST_HEADER",
		"REQUEST_BODY",
		"RESPONSE",
		"RESPONSE_HEADER",
		"RESPONSE_BODY",
	}
}

type DomainNameStatus string

// Enum values for DomainNameStatus
const (
	DomainNameStatusAvailable                    DomainNameStatus = "AVAILABLE"
	DomainNameStatusUpdating                     DomainNameStatus = "UPDATING"
	DomainNameStatusPending                      DomainNameStatus = "PENDING"
	DomainNameStatusPendingCertificateReimport   DomainNameStatus = "PENDING_CERTIFICATE_REIMPORT"
	DomainNameStatusPendingOwnershipVerification DomainNameStatus = "PENDING_OWNERSHIP_VERIFICATION"
)

// Values returns all known values for DomainNameStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DomainNameStatus) Values() []DomainNameStatus {
	return []DomainNameStatus{
		"AVAILABLE",
		"UPDATING",
		"PENDING",
		"PENDING_CERTIFICATE_REIMPORT",
		"PENDING_OWNERSHIP_VERIFICATION",
	}
}

type EndpointType string

// Enum values for EndpointType
const (
	EndpointTypeRegional EndpointType = "REGIONAL"
	EndpointTypeEdge     EndpointType = "EDGE"
	EndpointTypePrivate  EndpointType = "PRIVATE"
)

// Values returns all known values for EndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EndpointType) Values() []EndpointType {
	return []EndpointType{
		"REGIONAL",
		"EDGE",
		"PRIVATE",
	}
}

type GatewayResponseType string

// Enum values for GatewayResponseType
const (
	GatewayResponseTypeDefault4xx                   GatewayResponseType = "DEFAULT_4XX"
	GatewayResponseTypeDefault5xx                   GatewayResponseType = "DEFAULT_5XX"
	GatewayResponseTypeResourceNotFound             GatewayResponseType = "RESOURCE_NOT_FOUND"
	GatewayResponseTypeUnauthorized                 GatewayResponseType = "UNAUTHORIZED"
	GatewayResponseTypeInvalidApiKey                GatewayResponseType = "INVALID_API_KEY"
	GatewayResponseTypeAccessDenied                 GatewayResponseType = "ACCESS_DENIED"
	GatewayResponseTypeAuthorizerFailure            GatewayResponseType = "AUTHORIZER_FAILURE"
	GatewayResponseTypeAuthorizerConfigurationError GatewayResponseType = "AUTHORIZER_CONFIGURATION_ERROR"
	GatewayResponseTypeInvalidSignature             GatewayResponseType = "INVALID_SIGNATURE"
	GatewayResponseTypeExpiredToken                 GatewayResponseType = "EXPIRED_TOKEN"
	GatewayResponseTypeMissingAuthenticationToken   GatewayResponseType = "MISSING_AUTHENTICATION_TOKEN"
	GatewayResponseTypeIntegrationFailure           GatewayResponseType = "INTEGRATION_FAILURE"
	GatewayResponseTypeIntegrationTimeout           GatewayResponseType = "INTEGRATION_TIMEOUT"
	GatewayResponseTypeApiConfigurationError        GatewayResponseType = "API_CONFIGURATION_ERROR"
	GatewayResponseTypeUnsupportedMediaType         GatewayResponseType = "UNSUPPORTED_MEDIA_TYPE"
	GatewayResponseTypeBadRequestParameters         GatewayResponseType = "BAD_REQUEST_PARAMETERS"
	GatewayResponseTypeBadRequestBody               GatewayResponseType = "BAD_REQUEST_BODY"
	GatewayResponseTypeRequestTooLarge              GatewayResponseType = "REQUEST_TOO_LARGE"
	GatewayResponseTypeThrottled                    GatewayResponseType = "THROTTLED"
	GatewayResponseTypeQuotaExceeded                GatewayResponseType = "QUOTA_EXCEEDED"
	GatewayResponseTypeWafFiltered                  GatewayResponseType = "WAF_FILTERED"
)

// Values returns all known values for GatewayResponseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GatewayResponseType) Values() []GatewayResponseType {
	return []GatewayResponseType{
		"DEFAULT_4XX",
		"DEFAULT_5XX",
		"RESOURCE_NOT_FOUND",
		"UNAUTHORIZED",
		"INVALID_API_KEY",
		"ACCESS_DENIED",
		"AUTHORIZER_FAILURE",
		"AUTHORIZER_CONFIGURATION_ERROR",
		"INVALID_SIGNATURE",
		"EXPIRED_TOKEN",
		"MISSING_AUTHENTICATION_TOKEN",
		"INTEGRATION_FAILURE",
		"INTEGRATION_TIMEOUT",
		"API_CONFIGURATION_ERROR",
		"UNSUPPORTED_MEDIA_TYPE",
		"BAD_REQUEST_PARAMETERS",
		"BAD_REQUEST_BODY",
		"REQUEST_TOO_LARGE",
		"THROTTLED",
		"QUOTA_EXCEEDED",
		"WAF_FILTERED",
	}
}

type IntegrationType string

// Enum values for IntegrationType
const (
	IntegrationTypeHttp      IntegrationType = "HTTP"
	IntegrationTypeAws       IntegrationType = "AWS"
	IntegrationTypeMock      IntegrationType = "MOCK"
	IntegrationTypeHttpProxy IntegrationType = "HTTP_PROXY"
	IntegrationTypeAwsProxy  IntegrationType = "AWS_PROXY"
)

// Values returns all known values for IntegrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationType) Values() []IntegrationType {
	return []IntegrationType{
		"HTTP",
		"AWS",
		"MOCK",
		"HTTP_PROXY",
		"AWS_PROXY",
	}
}

type LocationStatusType string

// Enum values for LocationStatusType
const (
	LocationStatusTypeDocumented   LocationStatusType = "DOCUMENTED"
	LocationStatusTypeUndocumented LocationStatusType = "UNDOCUMENTED"
)

// Values returns all known values for LocationStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LocationStatusType) Values() []LocationStatusType {
	return []LocationStatusType{
		"DOCUMENTED",
		"UNDOCUMENTED",
	}
}

type Op string

// Enum values for Op
const (
	OpAdd     Op = "add"
	OpRemove  Op = "remove"
	OpReplace Op = "replace"
	OpMove    Op = "move"
	OpCopy    Op = "copy"
	OpTest    Op = "test"
)

// Values returns all known values for Op. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Op) Values() []Op {
	return []Op{
		"add",
		"remove",
		"replace",
		"move",
		"copy",
		"test",
	}
}

type PutMode string

// Enum values for PutMode
const (
	PutModeMerge     PutMode = "merge"
	PutModeOverwrite PutMode = "overwrite"
)

// Values returns all known values for PutMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PutMode) Values() []PutMode {
	return []PutMode{
		"merge",
		"overwrite",
	}
}

type QuotaPeriodType string

// Enum values for QuotaPeriodType
const (
	QuotaPeriodTypeDay   QuotaPeriodType = "DAY"
	QuotaPeriodTypeWeek  QuotaPeriodType = "WEEK"
	QuotaPeriodTypeMonth QuotaPeriodType = "MONTH"
)

// Values returns all known values for QuotaPeriodType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QuotaPeriodType) Values() []QuotaPeriodType {
	return []QuotaPeriodType{
		"DAY",
		"WEEK",
		"MONTH",
	}
}

type SecurityPolicy string

// Enum values for SecurityPolicy
const (
	SecurityPolicyTls10 SecurityPolicy = "TLS_1_0"
	SecurityPolicyTls12 SecurityPolicy = "TLS_1_2"
)

// Values returns all known values for SecurityPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SecurityPolicy) Values() []SecurityPolicy {
	return []SecurityPolicy{
		"TLS_1_0",
		"TLS_1_2",
	}
}

type UnauthorizedCacheControlHeaderStrategy string

// Enum values for UnauthorizedCacheControlHeaderStrategy
const (
	UnauthorizedCacheControlHeaderStrategyFailWith403                  UnauthorizedCacheControlHeaderStrategy = "FAIL_WITH_403"
	UnauthorizedCacheControlHeaderStrategySucceedWithResponseHeader    UnauthorizedCacheControlHeaderStrategy = "SUCCEED_WITH_RESPONSE_HEADER"
	UnauthorizedCacheControlHeaderStrategySucceedWithoutResponseHeader UnauthorizedCacheControlHeaderStrategy = "SUCCEED_WITHOUT_RESPONSE_HEADER"
)

// Values returns all known values for UnauthorizedCacheControlHeaderStrategy.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UnauthorizedCacheControlHeaderStrategy) Values() []UnauthorizedCacheControlHeaderStrategy {
	return []UnauthorizedCacheControlHeaderStrategy{
		"FAIL_WITH_403",
		"SUCCEED_WITH_RESPONSE_HEADER",
		"SUCCEED_WITHOUT_RESPONSE_HEADER",
	}
}

type VpcLinkStatus string

// Enum values for VpcLinkStatus
const (
	VpcLinkStatusAvailable VpcLinkStatus = "AVAILABLE"
	VpcLinkStatusPending   VpcLinkStatus = "PENDING"
	VpcLinkStatusDeleting  VpcLinkStatus = "DELETING"
	VpcLinkStatusFailed    VpcLinkStatus = "FAILED"
)

// Values returns all known values for VpcLinkStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VpcLinkStatus) Values() []VpcLinkStatus {
	return []VpcLinkStatus{
		"AVAILABLE",
		"PENDING",
		"DELETING",
		"FAILED",
	}
}
