# DnsActivityFilterBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of DNS events to exclude from logging | [optional] 

## Methods

### NewDnsActivityFilterBean

`func NewDnsActivityFilterBean() *DnsActivityFilterBean`

NewDnsActivityFilterBean instantiates a new DnsActivityFilterBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsActivityFilterBeanWithDefaults

`func NewDnsActivityFilterBeanWithDefaults() *DnsActivityFilterBean`

NewDnsActivityFilterBeanWithDefaults instantiates a new DnsActivityFilterBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *DnsActivityFilterBean) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DnsActivityFilterBean) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DnsActivityFilterBean) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DnsActivityFilterBean) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


