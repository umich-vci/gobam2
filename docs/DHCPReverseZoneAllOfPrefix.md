# DHCPReverseZoneAllOfPrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Range** | **string** | The address range for the IP block or network. | 

## Methods

### NewDHCPReverseZoneAllOfPrefix

`func NewDHCPReverseZoneAllOfPrefix(id int64, range_ string, ) *DHCPReverseZoneAllOfPrefix`

NewDHCPReverseZoneAllOfPrefix instantiates a new DHCPReverseZoneAllOfPrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPReverseZoneAllOfPrefixWithDefaults

`func NewDHCPReverseZoneAllOfPrefixWithDefaults() *DHCPReverseZoneAllOfPrefix`

NewDHCPReverseZoneAllOfPrefixWithDefaults instantiates a new DHCPReverseZoneAllOfPrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DHCPReverseZoneAllOfPrefix) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DHCPReverseZoneAllOfPrefix) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DHCPReverseZoneAllOfPrefix) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *DHCPReverseZoneAllOfPrefix) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPReverseZoneAllOfPrefix) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPReverseZoneAllOfPrefix) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPReverseZoneAllOfPrefix) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRange

`func (o *DHCPReverseZoneAllOfPrefix) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *DHCPReverseZoneAllOfPrefix) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *DHCPReverseZoneAllOfPrefix) SetRange(v string)`

SetRange sets Range field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


