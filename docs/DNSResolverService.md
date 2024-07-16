# DNSResolverService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Servers** | Pointer to **[]string** | The list of IP addresses of DNS resolvers. | [optional] 

## Methods

### NewDNSResolverService

`func NewDNSResolverService() *DNSResolverService`

NewDNSResolverService instantiates a new DNSResolverService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSResolverServiceWithDefaults

`func NewDNSResolverServiceWithDefaults() *DNSResolverService`

NewDNSResolverServiceWithDefaults instantiates a new DNSResolverService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSResolverService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSResolverService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSResolverService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSResolverService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServers

`func (o *DNSResolverService) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DNSResolverService) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DNSResolverService) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DNSResolverService) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


