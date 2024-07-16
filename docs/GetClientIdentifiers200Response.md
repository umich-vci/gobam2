# GetClientIdentifiers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetClientIdentifiers200ResponseDataInner**](GetClientIdentifiers200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetClientIdentifiers200Response

`func NewGetClientIdentifiers200Response() *GetClientIdentifiers200Response`

NewGetClientIdentifiers200Response instantiates a new GetClientIdentifiers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientIdentifiers200ResponseWithDefaults

`func NewGetClientIdentifiers200ResponseWithDefaults() *GetClientIdentifiers200Response`

NewGetClientIdentifiers200ResponseWithDefaults instantiates a new GetClientIdentifiers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetClientIdentifiers200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetClientIdentifiers200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetClientIdentifiers200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetClientIdentifiers200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetClientIdentifiers200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetClientIdentifiers200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetClientIdentifiers200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetClientIdentifiers200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetClientIdentifiers200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetClientIdentifiers200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetClientIdentifiers200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetClientIdentifiers200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetClientIdentifiers200Response) GetData() []GetClientIdentifiers200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetClientIdentifiers200Response) GetDataOk() (*[]GetClientIdentifiers200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetClientIdentifiers200Response) SetData(v []GetClientIdentifiers200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetClientIdentifiers200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


