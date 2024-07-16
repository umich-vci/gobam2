# LocationPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Locode** | Pointer to **string** | The location code showing the UN/LOCODE country and city location. The code is displayed in the hierarchical structure where the first two characters indicate the standard UN/LOCODE country code and the subsequent three characters indicate the location code. | [optional] [readonly] 
**Code** | Pointer to **string** | The location code showing the UN/LOCODE country and city location, and any appended custom child locations. The code is displayed in the hierarchical structure where the first two characters indicate the standard UN/LOCODE country code and the subsequent three characters indicate the location code. Any custom location codes are appended after the first 5 characters. | [optional] 
**Country** | Pointer to **string** | The two digit country code. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the location. | [optional] 
**LocalizedName** | Pointer to **string** | The localized name of the location. | [optional] 
**Subdivision** | Pointer to **string** | The ISO 1-3 character alphabetic or numeric code for the administrativedivision of the country. | [optional] [readonly] 
**Longitude** | Pointer to **float64** | The geographical longitude coordinate, in decimal degrees format. | [optional] 
**Latitude** | Pointer to **float64** | The geographical latitude coordinate, in decimal degrees format. | [optional] 

## Methods

### NewLocationPutRequestBody

`func NewLocationPutRequestBody() *LocationPutRequestBody`

NewLocationPutRequestBody instantiates a new LocationPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationPutRequestBodyWithDefaults

`func NewLocationPutRequestBodyWithDefaults() *LocationPutRequestBody`

NewLocationPutRequestBodyWithDefaults instantiates a new LocationPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LocationPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LocationPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LocationPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *LocationPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationPutRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationPutRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *LocationPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *LocationPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *LocationPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *LocationPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetLocode

`func (o *LocationPutRequestBody) GetLocode() string`

GetLocode returns the Locode field if non-nil, zero value otherwise.

### GetLocodeOk

`func (o *LocationPutRequestBody) GetLocodeOk() (*string, bool)`

GetLocodeOk returns a tuple with the Locode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocode

`func (o *LocationPutRequestBody) SetLocode(v string)`

SetLocode sets Locode field to given value.

### HasLocode

`func (o *LocationPutRequestBody) HasLocode() bool`

HasLocode returns a boolean if a field has been set.

### GetCode

`func (o *LocationPutRequestBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LocationPutRequestBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LocationPutRequestBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LocationPutRequestBody) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCountry

`func (o *LocationPutRequestBody) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationPutRequestBody) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationPutRequestBody) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LocationPutRequestBody) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *LocationPutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocationPutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocationPutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocationPutRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalizedName

`func (o *LocationPutRequestBody) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *LocationPutRequestBody) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *LocationPutRequestBody) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *LocationPutRequestBody) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetSubdivision

`func (o *LocationPutRequestBody) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *LocationPutRequestBody) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *LocationPutRequestBody) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *LocationPutRequestBody) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetLongitude

`func (o *LocationPutRequestBody) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationPutRequestBody) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationPutRequestBody) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocationPutRequestBody) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationPutRequestBody) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationPutRequestBody) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationPutRequestBody) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocationPutRequestBody) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


