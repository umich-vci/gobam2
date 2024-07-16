# BgpPrefixFilterBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** | The type of BGP prefix filter. | [optional] 
**Action** | Pointer to **string** | The action to take on the filtered prefixes. | [optional] 
**Prefix** | Pointer to **string** | The IP address prefix to filter. | [optional] 

## Methods

### NewBgpPrefixFilterBean

`func NewBgpPrefixFilterBean() *BgpPrefixFilterBean`

NewBgpPrefixFilterBean instantiates a new BgpPrefixFilterBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpPrefixFilterBeanWithDefaults

`func NewBgpPrefixFilterBeanWithDefaults() *BgpPrefixFilterBean`

NewBgpPrefixFilterBeanWithDefaults instantiates a new BgpPrefixFilterBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *BgpPrefixFilterBean) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *BgpPrefixFilterBean) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *BgpPrefixFilterBean) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *BgpPrefixFilterBean) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetAction

`func (o *BgpPrefixFilterBean) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BgpPrefixFilterBean) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BgpPrefixFilterBean) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BgpPrefixFilterBean) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrefix

`func (o *BgpPrefixFilterBean) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BgpPrefixFilterBean) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BgpPrefixFilterBean) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BgpPrefixFilterBean) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


