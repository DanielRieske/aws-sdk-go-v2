// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about agents that were instructed to start collecting data.
// Information includes the agent ID, a description of the operation, and whether
// the agent configuration was updated.
type AgentConfigurationStatus struct {

	// The agent ID.
	AgentId *string

	// A description of the operation performed.
	Description *string

	// Information about the status of the StartDataCollection and StopDataCollection
	// operations. The system has recorded the data collection operation. The agent
	// receives this command the next time it polls for a new command.
	OperationSucceeded bool

	noSmithyDocumentSerde
}

// Information about agents associated with the user’s Amazon Web Services
// account. Information includes agent IDs, IP addresses, media access control
// (MAC) addresses, agent or collector status, hostname where the agent resides,
// and agent version for each agent.
type AgentInfo struct {

	// The agent or collector ID.
	AgentId *string

	// Network details about the host where the agent or collector resides.
	AgentNetworkInfoList []AgentNetworkInfo

	// Type of agent.
	AgentType *string

	// Status of the collection process for an agent.
	CollectionStatus *string

	// The ID of the connector.
	ConnectorId *string

	// The health of the agent.
	Health AgentStatus

	// The name of the host where the agent or collector resides. The host can be a
	// server or virtual machine.
	HostName *string

	// Time since agent health was reported.
	LastHealthPingTime *string

	// Agent's first registration timestamp in UTC.
	RegisteredTime *string

	// The agent or collector version.
	Version *string

	noSmithyDocumentSerde
}

// Network details about the host where the agent/collector resides.
type AgentNetworkInfo struct {

	// The IP address for the host where the agent/collector resides.
	IpAddress *string

	// The MAC address for the host where the agent/collector resides.
	MacAddress *string

	noSmithyDocumentSerde
}

//	An object representing the agent or data collector that failed to delete, each
//
// containing agentId, errorMessage, and errorCode.
type BatchDeleteAgentError struct {

	//  The ID of the agent or data collector to delete.
	//
	// This member is required.
	AgentId *string

	//  The type of error that occurred for the delete failed agent. Valid status are:
	// AGENT_IN_USE | NOT_FOUND | INTERNAL_SERVER_ERROR.
	//
	// This member is required.
	ErrorCode DeleteAgentErrorCode

	//  The description of the error that occurred for the delete failed agent.
	//
	// This member is required.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// A metadata object that represents the deletion task being executed.
type BatchDeleteConfigurationTask struct {

	//  The type of configuration item to delete. Supported types are: SERVER.
	ConfigurationType DeletionConfigurationItemType

	//  The list of configuration IDs that were successfully deleted by the deletion
	// task.
	DeletedConfigurations []string

	//  A list of configuration IDs that produced warnings regarding their deletion,
	// paired with a warning message.
	DeletionWarnings []DeletionWarning

	//  An epoch seconds timestamp (UTC) of when the deletion task was completed or
	// failed.
	EndTime *time.Time

	//  A list of configuration IDs that failed to delete during the deletion task,
	// each paired with an error message.
	FailedConfigurations []FailedConfiguration

	//  The list of configuration IDs that were originally requested to be deleted by
	// the deletion task.
	RequestedConfigurations []string

	//  An epoch seconds timestamp (UTC) of when the deletion task was started.
	StartTime *time.Time

	//  The current execution status of the deletion task. Valid status are:
	// INITIALIZING | VALIDATING | DELETING | COMPLETED | FAILED.
	Status BatchDeleteConfigurationTaskStatus

	//  The deletion task's unique identifier.
	TaskId *string

	noSmithyDocumentSerde
}

// Error messages returned for each import task that you deleted as a response for
// this command.
type BatchDeleteImportDataError struct {

	// The type of error that occurred for a specific import task.
	ErrorCode BatchDeleteImportDataErrorCode

	// The description of the error that occurred for a specific import task.
	ErrorDescription *string

	// The unique import ID associated with the error that occurred.
	ImportTaskId *string

	noSmithyDocumentSerde
}

// Tags for a configuration item. Tags are metadata that help you categorize IT
// assets.
type ConfigurationTag struct {

	// The configuration ID for the item to tag. You can specify a list of keys and
	// values.
	ConfigurationId *string

	// A type of IT asset to tag.
	ConfigurationType ConfigurationItemType

	// A type of tag on which to filter. For example, serverType.
	Key *string

	// The time the configuration tag was created in Coordinated Universal Time (UTC).
	TimeOfCreation *time.Time

	// A value on which to filter. For example key = serverType and value = web server.
	Value *string

	noSmithyDocumentSerde
}

// A list of continuous export descriptions.
type ContinuousExportDescription struct {

	// The type of data collector used to gather this data (currently only offered for
	// AGENT).
	DataSource DataSource

	// The unique ID assigned to this export.
	ExportId *string

	// The name of the s3 bucket where the export data parquet files are stored.
	S3Bucket *string

	// An object which describes how the data is stored.
	//
	//   - databaseName - the name of the Glue database used to store the schema.
	SchemaStorageConfig map[string]string

	// The timestamp representing when the continuous export was started.
	StartTime *time.Time

	// Describes the status of the export. Can be one of the following values:
	//
	//   - START_IN_PROGRESS - setting up resources to start continuous export.
	//
	//   - START_FAILED - an error occurred setting up continuous export. To recover,
	//   call start-continuous-export again.
	//
	//   - ACTIVE - data is being exported to the customer bucket.
	//
	//   - ERROR - an error occurred during export. To fix the issue, call
	//   stop-continuous-export and start-continuous-export.
	//
	//   - STOP_IN_PROGRESS - stopping the export.
	//
	//   - STOP_FAILED - an error occurred stopping the export. To recover, call
	//   stop-continuous-export again.
	//
	//   - INACTIVE - the continuous export has been stopped. Data is no longer being
	//   exported to the customer bucket.
	Status ContinuousExportStatus

	// Contains information about any errors that have occurred. This data type can
	// have the following values:
	//
	//   - ACCESS_DENIED - You don’t have permission to start Data Exploration in
	//   Amazon Athena. Contact your Amazon Web Services administrator for help. For more
	//   information, see [Setting Up Amazon Web Services Application Discovery Service]in the Application Discovery Service User Guide.
	//
	//   - DELIVERY_STREAM_LIMIT_FAILURE - You reached the limit for Amazon Kinesis
	//   Data Firehose delivery streams. Reduce the number of streams or request a limit
	//   increase and try again. For more information, see [Kinesis Data Streams Limits]in the Amazon Kinesis Data
	//   Streams Developer Guide.
	//
	//   - FIREHOSE_ROLE_MISSING - The Data Exploration feature is in an error state
	//   because your user is missing the Amazon Web
	//   ServicesApplicationDiscoveryServiceFirehose role. Turn on Data Exploration in
	//   Amazon Athena and try again. For more information, see [Creating the Amazon Web ServicesApplicationDiscoveryServiceFirehose Role]in the Application
	//   Discovery Service User Guide.
	//
	//   - FIREHOSE_STREAM_DOES_NOT_EXIST - The Data Exploration feature is in an
	//   error state because your user is missing one or more of the Kinesis data
	//   delivery streams.
	//
	//   - INTERNAL_FAILURE - The Data Exploration feature is in an error state
	//   because of an internal failure. Try again later. If this problem persists,
	//   contact Amazon Web Services Support.
	//
	//   - LAKE_FORMATION_ACCESS_DENIED - You don't have sufficient lake formation
	//   permissions to start continuous export. For more information, see [Upgrading Amazon Web Services Glue Data Permissions to the Amazon Web Services Lake Formation Model]in the
	//   Amazon Web Services Lake Formation Developer Guide.
	//
	// You can use one of the following two ways to resolve this issue.
	//
	//   - If you don’t want to use the Lake Formation permission model, you can
	//   change the default Data Catalog settings to use only Amazon Web Services
	//   Identity and Access Management (IAM) access control for new databases. For more
	//   information, see [Change Data Catalog Settings]in the Lake Formation Developer Guide.
	//
	//   - You can give the service-linked IAM roles
	//   AWSServiceRoleForApplicationDiscoveryServiceContinuousExport and
	//   AWSApplicationDiscoveryServiceFirehose the required Lake Formation permissions.
	//   For more information, see [Granting Database Permissions]in the Lake Formation Developer Guide.
	//
	//   - AWSServiceRoleForApplicationDiscoveryServiceContinuousExport - Grant
	//   database creator permissions, which gives the role database creation ability and
	//   implicit permissions for any created tables. For more information, see [Implicit Lake Formation Permissions]in the
	//   Lake Formation Developer Guide.
	//
	//   - AWSApplicationDiscoveryServiceFirehose - Grant describe permissions for all
	//   tables in the database.
	//
	//   - S3_BUCKET_LIMIT_FAILURE - You reached the limit for Amazon S3 buckets.
	//   Reduce the number of S3 buckets or request a limit increase and try again. For
	//   more information, see [Bucket Restrictions and Limitations]in the Amazon Simple Storage Service Developer Guide.
	//
	//   - S3_NOT_SIGNED_UP - Your account is not signed up for the Amazon S3 service.
	//   You must sign up before you can use Amazon S3. You can sign up at the following
	//   URL: [https://aws.amazon.com/s3].
	//
	// [Creating the Amazon Web ServicesApplicationDiscoveryServiceFirehose Role]: https://docs.aws.amazon.com/application-discovery/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-create-firehose-role
	// [Implicit Lake Formation Permissions]: https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html
	// [Granting Database Permissions]: https://docs.aws.amazon.com/lake-formation/latest/dg/granting-database-permissions.html
	// [Kinesis Data Streams Limits]: http://docs.aws.amazon.com/streams/latest/dev/service-sizes-and-limits.html
	// [Setting Up Amazon Web Services Application Discovery Service]: http://docs.aws.amazon.com/application-discovery/latest/userguide/setting-up.html
	// [Change Data Catalog Settings]: https://docs.aws.amazon.com/lake-formation/latest/dg/getting-started-setup.html#setup-change-cat-settings
	// [https://aws.amazon.com/s3]: https://aws.amazon.com/s3
	// [Upgrading Amazon Web Services Glue Data Permissions to the Amazon Web Services Lake Formation Model]: http://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html
	// [Bucket Restrictions and Limitations]: http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html
	StatusDetail *string

	// The timestamp that represents when this continuous export was stopped.
	StopTime *time.Time

	noSmithyDocumentSerde
}

// Inventory data for installed discovery agents.
type CustomerAgentInfo struct {

	// Number of active discovery agents.
	//
	// This member is required.
	ActiveAgents int32

	// Number of blacklisted discovery agents.
	//
	// This member is required.
	BlackListedAgents int32

	// Number of healthy discovery agents
	//
	// This member is required.
	HealthyAgents int32

	// Number of discovery agents with status SHUTDOWN.
	//
	// This member is required.
	ShutdownAgents int32

	// Total number of discovery agents.
	//
	// This member is required.
	TotalAgents int32

	// Number of unhealthy discovery agents.
	//
	// This member is required.
	UnhealthyAgents int32

	// Number of unknown discovery agents.
	//
	// This member is required.
	UnknownAgents int32

	noSmithyDocumentSerde
}

// The inventory data for installed Agentless Collector collectors.
type CustomerAgentlessCollectorInfo struct {

	// The number of active Agentless Collector collectors.
	//
	// This member is required.
	ActiveAgentlessCollectors int32

	// The number of deny-listed Agentless Collector collectors.
	//
	// This member is required.
	DenyListedAgentlessCollectors int32

	// The number of healthy Agentless Collector collectors.
	//
	// This member is required.
	HealthyAgentlessCollectors int32

	// The number of Agentless Collector collectors with SHUTDOWN status.
	//
	// This member is required.
	ShutdownAgentlessCollectors int32

	//  The total number of Agentless Collector collectors.
	//
	// This member is required.
	TotalAgentlessCollectors int32

	//  The number of unhealthy Agentless Collector collectors.
	//
	// This member is required.
	UnhealthyAgentlessCollectors int32

	//  The number of unknown Agentless Collector collectors.
	//
	// This member is required.
	UnknownAgentlessCollectors int32

	noSmithyDocumentSerde
}

// Inventory data for installed discovery connectors.
type CustomerConnectorInfo struct {

	// Number of active discovery connectors.
	//
	// This member is required.
	ActiveConnectors int32

	// Number of blacklisted discovery connectors.
	//
	// This member is required.
	BlackListedConnectors int32

	// Number of healthy discovery connectors.
	//
	// This member is required.
	HealthyConnectors int32

	// Number of discovery connectors with status SHUTDOWN,
	//
	// This member is required.
	ShutdownConnectors int32

	// Total number of discovery connectors.
	//
	// This member is required.
	TotalConnectors int32

	// Number of unhealthy discovery connectors.
	//
	// This member is required.
	UnhealthyConnectors int32

	// Number of unknown discovery connectors.
	//
	// This member is required.
	UnknownConnectors int32

	noSmithyDocumentSerde
}

// The inventory data for installed Migration Evaluator collectors.
type CustomerMeCollectorInfo struct {

	//  The number of active Migration Evaluator collectors.
	//
	// This member is required.
	ActiveMeCollectors int32

	//  The number of deny-listed Migration Evaluator collectors.
	//
	// This member is required.
	DenyListedMeCollectors int32

	//  The number of healthy Migration Evaluator collectors.
	//
	// This member is required.
	HealthyMeCollectors int32

	//  The number of Migration Evaluator collectors with SHUTDOWN status.
	//
	// This member is required.
	ShutdownMeCollectors int32

	//  The total number of Migration Evaluator collectors.
	//
	// This member is required.
	TotalMeCollectors int32

	//  The number of unhealthy Migration Evaluator collectors.
	//
	// This member is required.
	UnhealthyMeCollectors int32

	//  The number of unknown Migration Evaluator collectors.
	//
	// This member is required.
	UnknownMeCollectors int32

	noSmithyDocumentSerde
}

//	An object representing the agent or data collector to be deleted along with
//
// the optional configurations for error handling.
type DeleteAgent struct {

	//  The ID of the agent or data collector to delete.
	//
	// This member is required.
	AgentId *string

	//  Optional flag used to force delete an agent or data collector. It is needed to
	// delete any agent in HEALTHY/UNHEALTHY/RUNNING status. Note that deleting an
	// agent that is actively reporting health causes it to be re-registered with a
	// different agent ID after data collector re-connects with Amazon Web Services.
	Force bool

	noSmithyDocumentSerde
}

// A configuration ID paired with a warning message.
type DeletionWarning struct {

	//  The unique identifier of the configuration that produced a warning.
	ConfigurationId *string

	//  The integer warning code associated with the warning message.
	WarningCode int32

	//  A descriptive message of the warning the associated configuration ID produced.
	WarningText *string

	noSmithyDocumentSerde
}

//	Indicates that the exported data must include EC2 instance type matches for
//
// on-premises servers that are discovered through Amazon Web Services Application
// Discovery Service.
type Ec2RecommendationsExportPreferences struct {

	//  The recommended EC2 instance type that matches the CPU usage metric of server
	// performance data.
	CpuPerformanceMetricBasis *UsageMetricBasis

	//  If set to true, the export [preferences] is set to Ec2RecommendationsExportPreferences .
	//
	// [preferences]: https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartExportTask.html#API_StartExportTask_RequestSyntax
	Enabled bool

	//  An array of instance types to exclude from recommendations.
	ExcludedInstanceTypes []string

	//  The target Amazon Web Services Region for the recommendations. You can use any
	// of the Region codes available for the chosen service, as listed in [Amazon Web Services service endpoints]in the
	// Amazon Web Services General Reference.
	//
	// [Amazon Web Services service endpoints]: https://docs.aws.amazon.com/general/latest/gr/rande.html
	PreferredRegion *string

	//  The recommended EC2 instance type that matches the Memory usage metric of
	// server performance data.
	RamPerformanceMetricBasis *UsageMetricBasis

	//  The contract type for a reserved instance. If blank, we assume an On-Demand
	// instance is preferred.
	ReservedInstanceOptions *ReservedInstanceOptions

	//  The target tenancy to use for your recommended EC2 instances.
	Tenancy Tenancy

	noSmithyDocumentSerde
}

// Used to select which agent's data is to be exported. A single agent ID may be
// selected for export using the [StartExportTask]action.
//
// [StartExportTask]: http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartExportTask.html
type ExportFilter struct {

	// Supported condition: EQUALS
	//
	// This member is required.
	Condition *string

	// A single ExportFilter name. Supported filters: agentIds .
	//
	// This member is required.
	Name *string

	// A single agent ID for a Discovery Agent. An agent ID can be found using the [DescribeAgents]
	// action. Typically an ADS agent ID is in the form o-0123456789abcdef0 .
	//
	// [DescribeAgents]: http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeAgents.html
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Information regarding the export status of discovered data. The value is an
// array of objects.
type ExportInfo struct {

	// A unique identifier used to query an export.
	//
	// This member is required.
	ExportId *string

	// The time that the data export was initiated.
	//
	// This member is required.
	ExportRequestTime *time.Time

	// The status of the data export job.
	//
	// This member is required.
	ExportStatus ExportStatus

	// A status message provided for API callers.
	//
	// This member is required.
	StatusMessage *string

	// A URL for an Amazon S3 bucket where you can review the exported data. The URL
	// is displayed only if the export succeeded.
	ConfigurationsDownloadUrl *string

	// If true, the export of agent information exceeded the size limit for a single
	// export and the exported data is incomplete for the requested time range. To
	// address this, select a smaller time range for the export by using startDate and
	// endDate .
	IsTruncated bool

	// The endTime used in the StartExportTask request. If no endTime was requested,
	// this result does not appear in ExportInfo .
	RequestedEndTime *time.Time

	// The value of startTime parameter in the StartExportTask request. If no startTime
	// was requested, this result does not appear in ExportInfo .
	RequestedStartTime *time.Time

	noSmithyDocumentSerde
}

//	Indicates the type of data that is being exported. Only one ExportPreferences
//
// can be enabled for a [StartExportTask]action.
//
// The following types satisfy this interface:
//
//	ExportPreferencesMemberEc2RecommendationsPreferences
//
// [StartExportTask]: https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartExportTask.html
type ExportPreferences interface {
	isExportPreferences()
}

//	If enabled, exported data includes EC2 instance type matches for on-premises
//
// servers discovered through Amazon Web Services Application Discovery Service.
type ExportPreferencesMemberEc2RecommendationsPreferences struct {
	Value Ec2RecommendationsExportPreferences

	noSmithyDocumentSerde
}

func (*ExportPreferencesMemberEc2RecommendationsPreferences) isExportPreferences() {}

// A configuration ID paired with an error message.
type FailedConfiguration struct {

	//  The unique identifier of the configuration the failed to delete.
	ConfigurationId *string

	//  A descriptive message indicating why the associated configuration failed to
	// delete.
	ErrorMessage *string

	//  The integer error code associated with the error message.
	ErrorStatusCode int32

	noSmithyDocumentSerde
}

// A filter that can use conditional operators.
//
// For more information about filters, see [Querying Discovered Configuration Items] in the Amazon Web Services Application
// Discovery Service User Guide.
//
// [Querying Discovered Configuration Items]: https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html
type Filter struct {

	// A conditional operator. The following operators are valid: EQUALS, NOT_EQUALS,
	// CONTAINS, NOT_CONTAINS. If you specify multiple filters, the system utilizes all
	// filters as though concatenated by AND. If you specify multiple values for a
	// particular filter, the system differentiates the values using OR. Calling either
	// DescribeConfigurations or ListConfigurations returns attributes of matching
	// configuration items.
	//
	// This member is required.
	Condition *string

	// The name of the filter.
	//
	// This member is required.
	Name *string

	// A string value on which to filter. For example, if you choose the
	// destinationServer.osVersion filter name, you could specify Ubuntu for the value.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// An array of information related to the import task request that includes status
// information, times, IDs, the Amazon S3 Object URL for the import file, and more.
type ImportTask struct {

	// The total number of application records in the import file that failed to be
	// imported.
	ApplicationImportFailure int32

	// The total number of application records in the import file that were
	// successfully imported.
	ApplicationImportSuccess int32

	// A unique token used to prevent the same import request from occurring more than
	// once. If you didn't provide a token, a token was automatically generated when
	// the import task request was sent.
	ClientRequestToken *string

	// A link to a compressed archive folder (in the ZIP format) that contains an
	// error log and a file of failed records. You can use these two files to quickly
	// identify records that failed, why they failed, and correct those records.
	// Afterward, you can upload the corrected file to your Amazon S3 bucket and create
	// another import task request.
	//
	// This field also includes authorization information so you can confirm the
	// authenticity of the compressed archive before you download it.
	//
	// If some records failed to be imported we recommend that you correct the records
	// in the failed entries file and then imports that failed entries file. This
	// prevents you from having to correct and update the larger original file and
	// attempt importing it again.
	ErrorsAndFailedEntriesZip *string

	// The time that the import task request finished, presented in the Unix time
	// stamp format.
	ImportCompletionTime *time.Time

	// The time that the import task request was deleted, presented in the Unix time
	// stamp format.
	ImportDeletedTime *time.Time

	// The time that the import task request was made, presented in the Unix time
	// stamp format.
	ImportRequestTime *time.Time

	// The unique ID for a specific import task. These IDs aren't globally unique, but
	// they are unique within an Amazon Web Services account.
	ImportTaskId *string

	// The URL for your import file that you've uploaded to Amazon S3.
	ImportUrl *string

	// A descriptive name for an import task. You can use this name to filter future
	// requests related to this import task, such as identifying applications and
	// servers that were included in this import task. We recommend that you use a
	// meaningful name for each import task.
	Name *string

	// The total number of server records in the import file that failed to be
	// imported.
	ServerImportFailure int32

	// The total number of server records in the import file that were successfully
	// imported.
	ServerImportSuccess int32

	// The status of the import task. An import can have the status of IMPORT_COMPLETE
	// and still have some records fail to import from the overall request. More
	// information can be found in the downloadable archive defined in the
	// errorsAndFailedEntriesZip field, or in the Migration Hub management console.
	Status ImportStatus

	noSmithyDocumentSerde
}

// A name-values pair of elements you can use to filter the results when querying
// your import tasks. Currently, wildcards are not supported for filters.
//
// When filtering by import status, all other filter values are ignored.
type ImportTaskFilter struct {

	// The name, status, or import task ID for a specific import task.
	Name ImportTaskFilterName

	// An array of strings that you can provide to match against a specific name,
	// status, or import task ID to filter the results for your import task queries.
	Values []string

	noSmithyDocumentSerde
}

// Details about neighboring servers.
type NeighborConnectionDetail struct {

	// The number of open network connections with the neighboring server.
	//
	// This member is required.
	ConnectionsCount int64

	// The ID of the server that accepted the network connection.
	//
	// This member is required.
	DestinationServerId *string

	// The ID of the server that opened the network connection.
	//
	// This member is required.
	SourceServerId *string

	// The destination network port for the connection.
	DestinationPort *int32

	// The network protocol used for the connection.
	TransportProtocol *string

	noSmithyDocumentSerde
}

// A field and direction for ordered output.
type OrderByElement struct {

	// The field on which to order.
	//
	// This member is required.
	FieldName *string

	// Ordering direction.
	SortOrder OrderString

	noSmithyDocumentSerde
}

// Used to provide Reserved Instance preferences for the recommendation.
type ReservedInstanceOptions struct {

	//  The flexibility to change the instance types needed for your Reserved
	// Instance.
	//
	// This member is required.
	OfferingClass OfferingClass

	//  The payment plan to use for your Reserved Instance.
	//
	// This member is required.
	PurchasingOption PurchasingOption

	//  The preferred duration of the Reserved Instance term.
	//
	// This member is required.
	TermLength TermLength

	noSmithyDocumentSerde
}

// Metadata that help you categorize IT assets.
//
// Do not store sensitive information (like personal data) in tags.
type Tag struct {

	// The type of tag on which to filter.
	//
	// This member is required.
	Key *string

	// A value for a tag key on which to filter.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The tag filter. Valid names are: tagKey , tagValue , configurationId .
type TagFilter struct {

	// A name of the tag filter.
	//
	// This member is required.
	Name *string

	// Values for the tag filter.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

//	Specifies the performance metrics to use for the server that is used for
//
// recommendations.
type UsageMetricBasis struct {

	//  A utilization metric that is used by the recommendations.
	Name *string

	//  Specifies the percentage of the specified utilization metric that is used by
	// the recommendations.
	PercentageAdjust *float64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isExportPreferences() {}
