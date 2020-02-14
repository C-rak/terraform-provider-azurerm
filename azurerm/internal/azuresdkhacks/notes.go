package azuresdkhacks

// There's a design flaw that exists between the Azure SDK for Go and Azure Resource Manager API
// where to perform a delta update unchanged fields are omited from the response; however the API
// also allows for passing `null` to disable the value of a field.
//
// Ultimately the Azure SDK for Go has opted to serialise structs with `json:"name,omitempty"` which
// means that this value will be omited if nil to allow for delta updates - however this means there's
// no means of removing/resetting a value of a nested object once provided since a `nil` object will be
// reset
//
// As such this set of well intentioned hacks is intended to force this behaviour where necessary.
//
// It's worth noting that these hacks are a last resort and the Swagger/API/SDK should almost always be
// fixed instead.
