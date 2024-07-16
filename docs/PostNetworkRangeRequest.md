# PostNetworkRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Range** | **string** | The address range for the DHCP range. | 
**Usage** | Pointer to **map[string]int64** | Values indicating the current amount of DHCP allocated, DHCP abandoned, DHCP reserved, and unassigned IP addresses within the range. | [optional] [readonly] 
**InheritedFields** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewPostNetworkRangeRequest

`func NewPostNetworkRangeRequest(type_ string, range_ string, ) *PostNetworkRangeRequest`

NewPostNetworkRangeRequest instantiates a new PostNetworkRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostNetworkRangeRequestWithDefaults

`func NewPostNetworkRangeRequestWithDefaults() *PostNetworkRangeRequest`

NewPostNetworkRangeRequestWithDefaults instantiates a new PostNetworkRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostNetworkRangeRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostNetworkRangeRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostNetworkRangeRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostNetworkRangeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostNetworkRangeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostNetworkRangeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostNetworkRangeRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostNetworkRangeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostNetworkRangeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostNetworkRangeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostNetworkRangeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostNetworkRangeRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostNetworkRangeRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostNetworkRangeRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostNetworkRangeRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostNetworkRangeRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostNetworkRangeRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostNetworkRangeRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostNetworkRangeRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRange

`func (o *PostNetworkRangeRequest) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *PostNetworkRangeRequest) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *PostNetworkRangeRequest) SetRange(v string)`

SetRange sets Range field to given value.


### GetUsage

`func (o *PostNetworkRangeRequest) GetUsage() map[string]int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PostNetworkRangeRequest) GetUsageOk() (*map[string]int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PostNetworkRangeRequest) SetUsage(v map[string]int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PostNetworkRangeRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetInheritedFields

`func (o *PostNetworkRangeRequest) GetInheritedFields() []string`

GetInheritedFields returns the InheritedFields field if non-nil, zero value otherwise.

### GetInheritedFieldsOk

`func (o *PostNetworkRangeRequest) GetInheritedFieldsOk() (*[]string, bool)`

GetInheritedFieldsOk returns a tuple with the InheritedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFields

`func (o *PostNetworkRangeRequest) SetInheritedFields(v []string)`

SetInheritedFields sets InheritedFields field to given value.

### HasInheritedFields

`func (o *PostNetworkRangeRequest) HasInheritedFields() bool`

HasInheritedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


