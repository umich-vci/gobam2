# GetLocations200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetLocations200ResponseDataInner

`func NewGetLocations200ResponseDataInner() *GetLocations200ResponseDataInner`

NewGetLocations200ResponseDataInner instantiates a new GetLocations200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLocations200ResponseDataInnerWithDefaults

`func NewGetLocations200ResponseDataInnerWithDefaults() *GetLocations200ResponseDataInner`

NewGetLocations200ResponseDataInnerWithDefaults instantiates a new GetLocations200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLocations200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLocations200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLocations200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLocations200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetLocations200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLocations200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLocations200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLocations200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetLocations200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLocations200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLocations200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLocations200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetLocations200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetLocations200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetLocations200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetLocations200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetLocode

`func (o *GetLocations200ResponseDataInner) GetLocode() string`

GetLocode returns the Locode field if non-nil, zero value otherwise.

### GetLocodeOk

`func (o *GetLocations200ResponseDataInner) GetLocodeOk() (*string, bool)`

GetLocodeOk returns a tuple with the Locode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocode

`func (o *GetLocations200ResponseDataInner) SetLocode(v string)`

SetLocode sets Locode field to given value.

### HasLocode

`func (o *GetLocations200ResponseDataInner) HasLocode() bool`

HasLocode returns a boolean if a field has been set.

### GetCode

`func (o *GetLocations200ResponseDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetLocations200ResponseDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetLocations200ResponseDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetLocations200ResponseDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCountry

`func (o *GetLocations200ResponseDataInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetLocations200ResponseDataInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetLocations200ResponseDataInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetLocations200ResponseDataInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDescription

`func (o *GetLocations200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLocations200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLocations200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLocations200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalizedName

`func (o *GetLocations200ResponseDataInner) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *GetLocations200ResponseDataInner) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *GetLocations200ResponseDataInner) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *GetLocations200ResponseDataInner) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetSubdivision

`func (o *GetLocations200ResponseDataInner) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *GetLocations200ResponseDataInner) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *GetLocations200ResponseDataInner) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *GetLocations200ResponseDataInner) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetLongitude

`func (o *GetLocations200ResponseDataInner) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GetLocations200ResponseDataInner) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GetLocations200ResponseDataInner) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GetLocations200ResponseDataInner) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *GetLocations200ResponseDataInner) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GetLocations200ResponseDataInner) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GetLocations200ResponseDataInner) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GetLocations200ResponseDataInner) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLinks

`func (o *GetLocations200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetLocations200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetLocations200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetLocations200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


