# OspfRoutingProtocolBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationEnabled** | Pointer to **bool** | Indicates whether authentication is enabled. | [optional] 
**AuthenticationPassword** | Pointer to **string** | The password for authentication. | [optional] 
**AuthenticationKey** | Pointer to **string** | The authentication key used for OSPF. The authentication key is used to compute the MD5 hash that is used to authenticate OSPF packets. | [optional] 
**DeadInterval** | Pointer to **string** | The length of time that the peer/neighbor router will maintain a route to the primary router in the absence of hello messages. | [optional] 
**HelloInterval** | Pointer to **string** | The length of time that the primary router contacts its peer/neighbor to indicate that it&#39;s still active. | [optional] 
**AreaId** | Pointer to **string** | The OSPF area ID. | [optional] 
**StubArea** | Pointer to **bool** | Indicates whether an OSPF subnet is used. | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewOspfRoutingProtocolBean

`func NewOspfRoutingProtocolBean() *OspfRoutingProtocolBean`

NewOspfRoutingProtocolBean instantiates a new OspfRoutingProtocolBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfRoutingProtocolBeanWithDefaults

`func NewOspfRoutingProtocolBeanWithDefaults() *OspfRoutingProtocolBean`

NewOspfRoutingProtocolBeanWithDefaults instantiates a new OspfRoutingProtocolBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationEnabled

`func (o *OspfRoutingProtocolBean) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *OspfRoutingProtocolBean) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *OspfRoutingProtocolBean) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *OspfRoutingProtocolBean) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *OspfRoutingProtocolBean) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *OspfRoutingProtocolBean) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *OspfRoutingProtocolBean) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *OspfRoutingProtocolBean) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *OspfRoutingProtocolBean) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *OspfRoutingProtocolBean) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *OspfRoutingProtocolBean) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *OspfRoutingProtocolBean) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetDeadInterval

`func (o *OspfRoutingProtocolBean) GetDeadInterval() string`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *OspfRoutingProtocolBean) GetDeadIntervalOk() (*string, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *OspfRoutingProtocolBean) SetDeadInterval(v string)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *OspfRoutingProtocolBean) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *OspfRoutingProtocolBean) GetHelloInterval() string`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *OspfRoutingProtocolBean) GetHelloIntervalOk() (*string, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *OspfRoutingProtocolBean) SetHelloInterval(v string)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *OspfRoutingProtocolBean) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetAreaId

`func (o *OspfRoutingProtocolBean) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *OspfRoutingProtocolBean) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *OspfRoutingProtocolBean) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *OspfRoutingProtocolBean) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetStubArea

`func (o *OspfRoutingProtocolBean) GetStubArea() bool`

GetStubArea returns the StubArea field if non-nil, zero value otherwise.

### GetStubAreaOk

`func (o *OspfRoutingProtocolBean) GetStubAreaOk() (*bool, bool)`

GetStubAreaOk returns a tuple with the StubArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubArea

`func (o *OspfRoutingProtocolBean) SetStubArea(v bool)`

SetStubArea sets StubArea field to given value.

### HasStubArea

`func (o *OspfRoutingProtocolBean) HasStubArea() bool`

HasStubArea returns a boolean if a field has been set.

### GetPassword

`func (o *OspfRoutingProtocolBean) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OspfRoutingProtocolBean) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OspfRoutingProtocolBean) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OspfRoutingProtocolBean) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


