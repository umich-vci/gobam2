# DHCPReverseZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Prefix** | Pointer to [**DHCPReverseZoneAllOfPrefix**](DHCPReverseZoneAllOfPrefix.md) |  | [optional] 

## Methods

### NewDHCPReverseZone

`func NewDHCPReverseZone() *DHCPReverseZone`

NewDHCPReverseZone instantiates a new DHCPReverseZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPReverseZoneWithDefaults

`func NewDHCPReverseZoneWithDefaults() *DHCPReverseZone`

NewDHCPReverseZoneWithDefaults instantiates a new DHCPReverseZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPReverseZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPReverseZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPReverseZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPReverseZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefix

`func (o *DHCPReverseZone) GetPrefix() DHCPReverseZoneAllOfPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DHCPReverseZone) GetPrefixOk() (*DHCPReverseZoneAllOfPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DHCPReverseZone) SetPrefix(v DHCPReverseZoneAllOfPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DHCPReverseZone) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


