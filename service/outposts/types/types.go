// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an address.
type Address struct {

	// The first line of the address.
	//
	// This member is required.
	AddressLine1 *string

	// The city for the address.
	//
	// This member is required.
	City *string

	// The ISO-3166 two-letter country code for the address.
	//
	// This member is required.
	CountryCode *string

	// The postal code for the address.
	//
	// This member is required.
	PostalCode *string

	// The state for the address.
	//
	// This member is required.
	StateOrRegion *string

	// The second line of the address.
	AddressLine2 *string

	// The third line of the address.
	AddressLine3 *string

	// The name of the contact.
	ContactName *string

	// The phone number of the contact.
	ContactPhoneNumber *string

	// The district or county for the address.
	DistrictOrCounty *string

	// The municipality for the address.
	Municipality *string

	noSmithyDocumentSerde
}

// Information about hardware assets.
type AssetInfo struct {

	// The ID of the asset.
	AssetId *string

	// The position of an asset in a rack.
	AssetLocation *AssetLocation

	// The type of the asset.
	AssetType AssetType

	// Information about compute hardware assets.
	ComputeAttributes *ComputeAttributes

	// The rack ID of the asset.
	RackId *string

	noSmithyDocumentSerde
}

// Information about the position of the asset in a rack.
type AssetLocation struct {

	// The position of an asset in a rack measured in rack units.
	RackElevation *float32

	noSmithyDocumentSerde
}

// The capacity tasks that failed.
type CapacityTaskFailure struct {

	// The reason that the specified capacity task failed.
	//
	// This member is required.
	Reason *string

	// The type of failure.
	Type CapacityTaskFailureType

	noSmithyDocumentSerde
}

// The summary of the capacity task.
type CapacityTaskSummary struct {

	// The ID of the specified capacity task.
	CapacityTaskId *string

	// The status of the capacity task.
	CapacityTaskStatus CapacityTaskStatus

	// The date that the specified capacity task successfully ran.
	CompletionDate *time.Time

	// The date that the specified capacity task was created.
	CreationDate *time.Time

	// The date that the specified capacity was last modified.
	LastModifiedDate *time.Time

	// The ID of the Amazon Web Services Outposts order of the host associated with
	// the capacity task.
	OrderId *string

	// The ID of the Outpost associated with the specified capacity task.
	OutpostId *string

	noSmithyDocumentSerde
}

// Information about a catalog item.
type CatalogItem struct {

	// The ID of the catalog item.
	CatalogItemId *string

	// Information about the EC2 capacity of an item.
	EC2Capacities []EC2Capacity

	// The status of a catalog item.
	ItemStatus CatalogItemStatus

	// Information about the power draw of an item.
	PowerKva *float32

	// The supported storage options for the catalog item.
	SupportedStorage []SupportedStorageEnum

	// The uplink speed this catalog item requires for the connection to the Region.
	SupportedUplinkGbps []int32

	// The weight of the item in pounds.
	WeightLbs *int32

	noSmithyDocumentSerde
}

// Information about compute hardware assets.
type ComputeAttributes struct {

	// The host ID of the Dedicated Host on the asset.
	HostId *string

	// A list of the names of instance families that are currently associated with a
	// given asset.
	InstanceFamilies []string

	// The state.
	//   - ACTIVE - The asset is available and can provide capacity for new compute
	//   resources.
	//   - ISOLATED - The asset is undergoing maintenance and can't provide capacity
	//   for new compute resources. Existing compute resources on the asset are not
	//   affected.
	//   - RETIRING - The underlying hardware for the asset is degraded. Capacity for
	//   new compute resources is reduced. Amazon Web Services sends notifications for
	//   resources that must be stopped before the asset can be replaced.
	State ComputeAssetState

	noSmithyDocumentSerde
}

// Information about a connection.
type ConnectionDetails struct {

	// The allowed IP addresses.
	AllowedIps []string

	// The public key of the client.
	ClientPublicKey *string

	// The client tunnel address.
	ClientTunnelAddress *string

	// The endpoint for the server.
	ServerEndpoint *string

	// The public key of the server.
	ServerPublicKey *string

	// The server tunnel address.
	ServerTunnelAddress *string

	noSmithyDocumentSerde
}

// Information about EC2 capacity.
type EC2Capacity struct {

	// The family of the EC2 capacity.
	Family *string

	// The maximum size of the EC2 capacity.
	MaxSize *string

	// The quantity of the EC2 capacity.
	Quantity *string

	noSmithyDocumentSerde
}

// The instance type that you specify determines the combination of CPU, memory,
// storage, and networking capacity.
type InstanceTypeCapacity struct {

	// The number of instances for the specified instance type.
	//
	// This member is required.
	Count int32

	// The instance type of the hosts.
	//
	// This member is required.
	InstanceType *string

	noSmithyDocumentSerde
}

// Information about an instance type.
type InstanceTypeItem struct {

	// The instance type.
	InstanceType *string

	noSmithyDocumentSerde
}

// Information about a line item.
type LineItem struct {

	// Information about assets.
	AssetInformationList []LineItemAssetInformation

	// The ID of the catalog item.
	CatalogItemId *string

	// The ID of the line item.
	LineItemId *string

	// The ID of the previous line item.
	PreviousLineItemId *string

	// The ID of the previous order.
	PreviousOrderId *string

	// The quantity of the line item.
	Quantity *int32

	// Information about a line item shipment.
	ShipmentInformation *ShipmentInformation

	// The status of the line item.
	Status LineItemStatus

	noSmithyDocumentSerde
}

// Information about a line item asset.
type LineItemAssetInformation struct {

	// The ID of the asset.
	AssetId *string

	// The MAC addresses of the asset.
	MacAddressList []string

	noSmithyDocumentSerde
}

// Information about a line item request.
type LineItemRequest struct {

	// The ID of the catalog item.
	CatalogItemId *string

	// The quantity of a line item request.
	Quantity *int32

	noSmithyDocumentSerde
}

// Information about an order.
type Order struct {

	// The line items for the order
	LineItems []LineItem

	// The fulfillment date of the order.
	OrderFulfilledDate *time.Time

	// The ID of the order.
	OrderId *string

	// The submission date for the order.
	OrderSubmissionDate *time.Time

	// The type of order.
	OrderType OrderType

	// The ID of the Outpost in the order.
	OutpostId *string

	// The payment option for the order.
	PaymentOption PaymentOption

	// The payment term.
	PaymentTerm PaymentTerm

	// The status of the order.
	//   - PREPARING - Order is received and being prepared.
	//   - IN_PROGRESS - Order is either being built, shipped, or installed. To get
	//   more details, see the line item status.
	//   - COMPLETED - Order is complete.
	//   - CANCELLED - Order is cancelled.
	//   - ERROR - Customer should contact support.
	// The following status are deprecated: RECEIVED , PENDING , PROCESSING ,
	// INSTALLING , and FULFILLED .
	Status OrderStatus

	noSmithyDocumentSerde
}

// A summary of line items in your order.
type OrderSummary struct {

	// The status of all line items in the order.
	LineItemCountsByStatus map[string]int32

	// The fulfilment date for the order.
	OrderFulfilledDate *time.Time

	// The ID of the order.
	OrderId *string

	// The submission date for the order.
	OrderSubmissionDate *time.Time

	// The type of order.
	OrderType OrderType

	// The ID of the Outpost.
	OutpostId *string

	// The status of the order.
	//   - PREPARING - Order is received and is being prepared.
	//   - IN_PROGRESS - Order is either being built, shipped, or installed. For more
	//   information, see the LineItem status.
	//   - COMPLETED - Order is complete.
	//   - CANCELLED - Order is cancelled.
	//   - ERROR - Customer should contact support.
	// The following statuses are deprecated: RECEIVED , PENDING , PROCESSING ,
	// INSTALLING , and FULFILLED .
	Status OrderStatus

	noSmithyDocumentSerde
}

// Information about an Outpost.
type Outpost struct {

	// The Availability Zone.
	AvailabilityZone *string

	// The ID of the Availability Zone.
	AvailabilityZoneId *string

	// The description of the Outpost.
	Description *string

	// The life cycle status.
	LifeCycleStatus *string

	// The name of the Outpost.
	Name *string

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The ID of the Outpost.
	OutpostId *string

	// The Amazon Web Services account ID of the Outpost owner.
	OwnerId *string

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The hardware type.
	SupportedHardwareType SupportedHardwareType

	// The Outpost tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Information about the physical and logistical details for racks at sites. For
// more information about hardware requirements for racks, see Network readiness
// checklist (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist)
// in the Amazon Web Services Outposts User Guide.
type RackPhysicalProperties struct {

	// The type of fiber used to attach the Outpost to the network.
	FiberOpticCableType FiberOpticCableType

	// The maximum rack weight that this site can support. NO_LIMIT is over 2000 lbs
	// (907 kg).
	MaximumSupportedWeightLbs MaximumSupportedWeightLbs

	// The type of optical standard used to attach the Outpost to the network. This
	// field is dependent on uplink speed, fiber type, and distance to the upstream
	// device. For more information about networking requirements for racks, see
	// Network (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#facility-networking)
	// in the Amazon Web Services Outposts User Guide.
	OpticalStandard OpticalStandard

	// The power connector for the hardware.
	PowerConnector PowerConnector

	// The power draw available at the hardware placement position for the rack.
	PowerDrawKva PowerDrawKva

	// The position of the power feed.
	PowerFeedDrop PowerFeedDrop

	// The power option that you can provide for hardware.
	PowerPhase PowerPhase

	// The number of uplinks each Outpost network device.
	UplinkCount UplinkCount

	// The uplink speed the rack supports for the connection to the Region.
	UplinkGbps UplinkGbps

	noSmithyDocumentSerde
}

// Information about a line item shipment.
type ShipmentInformation struct {

	// The carrier of the shipment.
	ShipmentCarrier ShipmentCarrier

	// The tracking number of the shipment.
	ShipmentTrackingNumber *string

	noSmithyDocumentSerde
}

// Information about a site.
type Site struct {

	// The ID of the Amazon Web Services account.
	AccountId *string

	// The description of the site.
	Description *string

	// The name of the site.
	Name *string

	// Notes about a site.
	Notes *string

	// City where the hardware is installed and powered on.
	OperatingAddressCity *string

	// The ISO-3166 two-letter country code where the hardware is installed and
	// powered on.
	OperatingAddressCountryCode *string

	// State or region where the hardware is installed and powered on.
	OperatingAddressStateOrRegion *string

	// Information about the physical and logistical details for a rack at the site.
	RackPhysicalProperties *RackPhysicalProperties

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The site tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
