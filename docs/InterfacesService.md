# InterfacesService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**DedicatedManagementEnabled** | Pointer to **bool** |  | [optional] 
**Interfaces** | Pointer to [**[]InterfaceBean**](InterfaceBean.md) |  | [optional] 
**Routes** | Pointer to [**[]RouteBean**](RouteBean.md) |  | [optional] 

## Methods

### NewInterfacesService

`func NewInterfacesService() *InterfacesService`

NewInterfacesService instantiates a new InterfacesService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfacesServiceWithDefaults

`func NewInterfacesServiceWithDefaults() *InterfacesService`

NewInterfacesServiceWithDefaults instantiates a new InterfacesService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InterfacesService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfacesService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfacesService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InterfacesService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDedicatedManagementEnabled

`func (o *InterfacesService) GetDedicatedManagementEnabled() bool`

GetDedicatedManagementEnabled returns the DedicatedManagementEnabled field if non-nil, zero value otherwise.

### GetDedicatedManagementEnabledOk

`func (o *InterfacesService) GetDedicatedManagementEnabledOk() (*bool, bool)`

GetDedicatedManagementEnabledOk returns a tuple with the DedicatedManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedManagementEnabled

`func (o *InterfacesService) SetDedicatedManagementEnabled(v bool)`

SetDedicatedManagementEnabled sets DedicatedManagementEnabled field to given value.

### HasDedicatedManagementEnabled

`func (o *InterfacesService) HasDedicatedManagementEnabled() bool`

HasDedicatedManagementEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *InterfacesService) GetInterfaces() []InterfaceBean`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InterfacesService) GetInterfacesOk() (*[]InterfaceBean, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InterfacesService) SetInterfaces(v []InterfaceBean)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InterfacesService) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetRoutes

`func (o *InterfacesService) GetRoutes() []RouteBean`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *InterfacesService) GetRoutesOk() (*[]RouteBean, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *InterfacesService) SetRoutes(v []RouteBean)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *InterfacesService) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


