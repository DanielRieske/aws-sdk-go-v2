// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type PayloadFormatIndicator string

// Enum values for PayloadFormatIndicator
const (
	PayloadFormatIndicatorUnspecifiedBytes PayloadFormatIndicator = "UNSPECIFIED_BYTES"
	PayloadFormatIndicatorUtf8Data         PayloadFormatIndicator = "UTF8_DATA"
)

// Values returns all known values for PayloadFormatIndicator. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PayloadFormatIndicator) Values() []PayloadFormatIndicator {
	return []PayloadFormatIndicator{
		"UNSPECIFIED_BYTES",
		"UTF8_DATA",
	}
}
