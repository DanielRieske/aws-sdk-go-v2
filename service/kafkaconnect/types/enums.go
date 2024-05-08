// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ConnectorState string

// Enum values for ConnectorState
const (
	ConnectorStateRunning  ConnectorState = "RUNNING"
	ConnectorStateCreating ConnectorState = "CREATING"
	ConnectorStateUpdating ConnectorState = "UPDATING"
	ConnectorStateDeleting ConnectorState = "DELETING"
	ConnectorStateFailed   ConnectorState = "FAILED"
)

// Values returns all known values for ConnectorState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConnectorState) Values() []ConnectorState {
	return []ConnectorState{
		"RUNNING",
		"CREATING",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type CustomPluginContentType string

// Enum values for CustomPluginContentType
const (
	CustomPluginContentTypeJar CustomPluginContentType = "JAR"
	CustomPluginContentTypeZip CustomPluginContentType = "ZIP"
)

// Values returns all known values for CustomPluginContentType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomPluginContentType) Values() []CustomPluginContentType {
	return []CustomPluginContentType{
		"JAR",
		"ZIP",
	}
}

type CustomPluginState string

// Enum values for CustomPluginState
const (
	CustomPluginStateCreating     CustomPluginState = "CREATING"
	CustomPluginStateCreateFailed CustomPluginState = "CREATE_FAILED"
	CustomPluginStateActive       CustomPluginState = "ACTIVE"
	CustomPluginStateUpdating     CustomPluginState = "UPDATING"
	CustomPluginStateUpdateFailed CustomPluginState = "UPDATE_FAILED"
	CustomPluginStateDeleting     CustomPluginState = "DELETING"
)

// Values returns all known values for CustomPluginState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomPluginState) Values() []CustomPluginState {
	return []CustomPluginState{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
		"UPDATING",
		"UPDATE_FAILED",
		"DELETING",
	}
}

type KafkaClusterClientAuthenticationType string

// Enum values for KafkaClusterClientAuthenticationType
const (
	KafkaClusterClientAuthenticationTypeNone KafkaClusterClientAuthenticationType = "NONE"
	KafkaClusterClientAuthenticationTypeIam  KafkaClusterClientAuthenticationType = "IAM"
)

// Values returns all known values for KafkaClusterClientAuthenticationType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KafkaClusterClientAuthenticationType) Values() []KafkaClusterClientAuthenticationType {
	return []KafkaClusterClientAuthenticationType{
		"NONE",
		"IAM",
	}
}

type KafkaClusterEncryptionInTransitType string

// Enum values for KafkaClusterEncryptionInTransitType
const (
	KafkaClusterEncryptionInTransitTypePlaintext KafkaClusterEncryptionInTransitType = "PLAINTEXT"
	KafkaClusterEncryptionInTransitTypeTls       KafkaClusterEncryptionInTransitType = "TLS"
)

// Values returns all known values for KafkaClusterEncryptionInTransitType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KafkaClusterEncryptionInTransitType) Values() []KafkaClusterEncryptionInTransitType {
	return []KafkaClusterEncryptionInTransitType{
		"PLAINTEXT",
		"TLS",
	}
}

type WorkerConfigurationState string

// Enum values for WorkerConfigurationState
const (
	WorkerConfigurationStateActive   WorkerConfigurationState = "ACTIVE"
	WorkerConfigurationStateDeleting WorkerConfigurationState = "DELETING"
)

// Values returns all known values for WorkerConfigurationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkerConfigurationState) Values() []WorkerConfigurationState {
	return []WorkerConfigurationState{
		"ACTIVE",
		"DELETING",
	}
}
