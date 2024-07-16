# Location

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

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Location) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Location) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Location) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Location) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *Location) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *Location) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *Location) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *Location) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetLocode

`func (o *Location) GetLocode() string`

GetLocode returns the Locode field if non-nil, zero value otherwise.

### GetLocodeOk

`func (o *Location) GetLocodeOk() (*string, bool)`

GetLocodeOk returns a tuple with the Locode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocode

`func (o *Location) SetLocode(v string)`

SetLocode sets Locode field to given value.

### HasLocode

`func (o *Location) HasLocode() bool`

HasLocode returns a boolean if a field has been set.

### GetCode

`func (o *Location) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Location) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Location) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Location) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCountry

`func (o *Location) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Location) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *Location) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Location) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Location) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Location) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalizedName

`func (o *Location) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *Location) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *Location) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *Location) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetSubdivision

`func (o *Location) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *Location) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *Location) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *Location) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetLongitude

`func (o *Location) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *Location) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


