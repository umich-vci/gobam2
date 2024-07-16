# DHCPForwardZoneAllOfZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**AbsoluteName** | **string** | The fully qualified name of the zone. | 

## Methods

### NewDHCPForwardZoneAllOfZone

`func NewDHCPForwardZoneAllOfZone(id int64, absoluteName string, ) *DHCPForwardZoneAllOfZone`

NewDHCPForwardZoneAllOfZone instantiates a new DHCPForwardZoneAllOfZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPForwardZoneAllOfZoneWithDefaults

`func NewDHCPForwardZoneAllOfZoneWithDefaults() *DHCPForwardZoneAllOfZone`

NewDHCPForwardZoneAllOfZoneWithDefaults instantiates a new DHCPForwardZoneAllOfZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DHCPForwardZoneAllOfZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DHCPForwardZoneAllOfZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DHCPForwardZoneAllOfZone) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *DHCPForwardZoneAllOfZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPForwardZoneAllOfZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPForwardZoneAllOfZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPForwardZoneAllOfZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DHCPForwardZoneAllOfZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DHCPForwardZoneAllOfZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DHCPForwardZoneAllOfZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DHCPForwardZoneAllOfZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *DHCPForwardZoneAllOfZone) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *DHCPForwardZoneAllOfZone) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *DHCPForwardZoneAllOfZone) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


