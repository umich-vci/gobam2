# KafkaSinkBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**KeyField** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaSinkBean

`func NewKafkaSinkBean() *KafkaSinkBean`

NewKafkaSinkBean instantiates a new KafkaSinkBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaSinkBeanWithDefaults

`func NewKafkaSinkBeanWithDefaults() *KafkaSinkBean`

NewKafkaSinkBeanWithDefaults instantiates a new KafkaSinkBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *KafkaSinkBean) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *KafkaSinkBean) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *KafkaSinkBean) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *KafkaSinkBean) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTopic

`func (o *KafkaSinkBean) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *KafkaSinkBean) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *KafkaSinkBean) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *KafkaSinkBean) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetKeyField

`func (o *KafkaSinkBean) GetKeyField() string`

GetKeyField returns the KeyField field if non-nil, zero value otherwise.

### GetKeyFieldOk

`func (o *KafkaSinkBean) GetKeyFieldOk() (*string, bool)`

GetKeyFieldOk returns a tuple with the KeyField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyField

`func (o *KafkaSinkBean) SetKeyField(v string)`

SetKeyField sets KeyField field to given value.

### HasKeyField

`func (o *KafkaSinkBean) HasKeyField() bool`

HasKeyField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


