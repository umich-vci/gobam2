/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
)

// checks if the GetReconciliationPolicyDiscoveredDevices200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReconciliationPolicyDiscoveredDevices200ResponseDataInner{}

// GetReconciliationPolicyDiscoveredDevices200ResponseDataInner struct for GetReconciliationPolicyDiscoveredDevices200ResponseDataInner
type GetReconciliationPolicyDiscoveredDevices200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	Device *DiscoveredDevice `json:"device,omitempty"`
	Address *string `json:"address,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	Interface *DiscoveredInterface `json:"interface,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	InterfaceIndex *int32 `json:"interfaceIndex,omitempty"`
	Speed *int32 `json:"speed,omitempty"`
	Connector *string `json:"connector,omitempty"`
	Alias *string `json:"alias,omitempty"`
	PortMode *string `json:"portMode,omitempty"`
	Range *string `json:"range,omitempty"`
	Location *string `json:"location,omitempty"`
	VlanId *int32 `json:"vlanId,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInner instantiates a new GetReconciliationPolicyDiscoveredDevices200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInner() *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner {
	this := GetReconciliationPolicyDiscoveredDevices200ResponseDataInner{}
	return &this
}

// NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInnerWithDefaults instantiates a new GetReconciliationPolicyDiscoveredDevices200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReconciliationPolicyDiscoveredDevices200ResponseDataInnerWithDefaults() *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner {
	this := GetReconciliationPolicyDiscoveredDevices200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDevice() DiscoveredDevice {
	if o == nil || IsNil(o.Device) {
		var ret DiscoveredDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDeviceOk() (*DiscoveredDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DiscoveredDevice and assigns it to the Device field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetDevice(v DiscoveredDevice) {
	o.Device = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetAddress(v string) {
	o.Address = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterface() DiscoveredInterface {
	if o == nil || IsNil(o.Interface) {
		var ret DiscoveredInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceOk() (*DiscoveredInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given DiscoveredInterface and assigns it to the Interface field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetInterface(v DiscoveredInterface) {
	o.Interface = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetInterfaceIndex returns the InterfaceIndex field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceIndex() int32 {
	if o == nil || IsNil(o.InterfaceIndex) {
		var ret int32
		return ret
	}
	return *o.InterfaceIndex
}

// GetInterfaceIndexOk returns a tuple with the InterfaceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetInterfaceIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.InterfaceIndex) {
		return nil, false
	}
	return o.InterfaceIndex, true
}

// HasInterfaceIndex returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasInterfaceIndex() bool {
	if o != nil && !IsNil(o.InterfaceIndex) {
		return true
	}

	return false
}

// SetInterfaceIndex gets a reference to the given int32 and assigns it to the InterfaceIndex field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetInterfaceIndex(v int32) {
	o.InterfaceIndex = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetSpeed(v int32) {
	o.Speed = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetConnector() string {
	if o == nil || IsNil(o.Connector) {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetConnectorOk() (*string, bool) {
	if o == nil || IsNil(o.Connector) {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasConnector() bool {
	if o != nil && !IsNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetConnector(v string) {
	o.Connector = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetAlias(v string) {
	o.Alias = &v
}

// GetPortMode returns the PortMode field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetPortMode() string {
	if o == nil || IsNil(o.PortMode) {
		var ret string
		return ret
	}
	return *o.PortMode
}

// GetPortModeOk returns a tuple with the PortMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetPortModeOk() (*string, bool) {
	if o == nil || IsNil(o.PortMode) {
		return nil, false
	}
	return o.PortMode, true
}

// HasPortMode returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasPortMode() bool {
	if o != nil && !IsNil(o.PortMode) {
		return true
	}

	return false
}

// SetPortMode gets a reference to the given string and assigns it to the PortMode field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetPortMode(v string) {
	o.PortMode = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetRange(v string) {
	o.Range = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetLocation(v string) {
	o.Location = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InterfaceIndex) {
		toSerialize["interfaceIndex"] = o.InterfaceIndex
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.PortMode) {
		toSerialize["portMode"] = o.PortMode
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner struct {
	value *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner
	isSet bool
}

func (v NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) Get() *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner {
	return v.value
}

func (v *NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) Set(val *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner(val *GetReconciliationPolicyDiscoveredDevices200ResponseDataInner) *NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner {
	return &NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReconciliationPolicyDiscoveredDevices200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


