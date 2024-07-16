# GetZoneDeclarations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetZoneDeclarations200ResponseDataInner**](GetZoneDeclarations200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetZoneDeclarations200Response

`func NewGetZoneDeclarations200Response() *GetZoneDeclarations200Response`

NewGetZoneDeclarations200Response instantiates a new GetZoneDeclarations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneDeclarations200ResponseWithDefaults

`func NewGetZoneDeclarations200ResponseWithDefaults() *GetZoneDeclarations200Response`

NewGetZoneDeclarations200ResponseWithDefaults instantiates a new GetZoneDeclarations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetZoneDeclarations200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetZoneDeclarations200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetZoneDeclarations200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetZoneDeclarations200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetZoneDeclarations200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetZoneDeclarations200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetZoneDeclarations200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetZoneDeclarations200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetZoneDeclarations200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetZoneDeclarations200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetZoneDeclarations200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetZoneDeclarations200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetZoneDeclarations200Response) GetData() []GetZoneDeclarations200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetZoneDeclarations200Response) GetDataOk() (*[]GetZoneDeclarations200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetZoneDeclarations200Response) SetData(v []GetZoneDeclarations200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetZoneDeclarations200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


