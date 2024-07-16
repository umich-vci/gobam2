# DHCPForwardZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Zone** | Pointer to [**DHCPForwardZoneAllOfZone**](DHCPForwardZoneAllOfZone.md) |  | [optional] 

## Methods

### NewDHCPForwardZone

`func NewDHCPForwardZone() *DHCPForwardZone`

NewDHCPForwardZone instantiates a new DHCPForwardZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPForwardZoneWithDefaults

`func NewDHCPForwardZoneWithDefaults() *DHCPForwardZone`

NewDHCPForwardZoneWithDefaults instantiates a new DHCPForwardZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPForwardZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPForwardZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPForwardZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPForwardZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *DHCPForwardZone) GetZone() DHCPForwardZoneAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DHCPForwardZone) GetZoneOk() (*DHCPForwardZoneAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DHCPForwardZone) SetZone(v DHCPForwardZoneAllOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *DHCPForwardZone) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


