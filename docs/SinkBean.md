# SinkBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SinkType** | Pointer to **string** | The output sink type. | [optional] 
**HealthCheckEnabled** | Pointer to **bool** | Indicates whether health check service is enabled on the downstream sink server. | [optional] 
**CaCertificate** | Pointer to **string** | The CA certificate used to authenticate with the remote host. | [optional] 
**CertificateVerificationEnabled** | Pointer to **bool** | Indicates whether the server attempts a TLS handshake using the uploaded CA certificate with the remote host&#39;s TLS server certificate. | [optional] 
**HostnameVerificationEnabled** | Pointer to **bool** | Indicates whether the server validates the hostname part of the URI against the CN (Common Name) or SAN (Subject Alternative Name) of the server certificate during the TLS handshake. | [optional] 

## Methods

### NewSinkBean

`func NewSinkBean() *SinkBean`

NewSinkBean instantiates a new SinkBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinkBeanWithDefaults

`func NewSinkBeanWithDefaults() *SinkBean`

NewSinkBeanWithDefaults instantiates a new SinkBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSinkType

`func (o *SinkBean) GetSinkType() string`

GetSinkType returns the SinkType field if non-nil, zero value otherwise.

### GetSinkTypeOk

`func (o *SinkBean) GetSinkTypeOk() (*string, bool)`

GetSinkTypeOk returns a tuple with the SinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkType

`func (o *SinkBean) SetSinkType(v string)`

SetSinkType sets SinkType field to given value.

### HasSinkType

`func (o *SinkBean) HasSinkType() bool`

HasSinkType returns a boolean if a field has been set.

### GetHealthCheckEnabled

`func (o *SinkBean) GetHealthCheckEnabled() bool`

GetHealthCheckEnabled returns the HealthCheckEnabled field if non-nil, zero value otherwise.

### GetHealthCheckEnabledOk

`func (o *SinkBean) GetHealthCheckEnabledOk() (*bool, bool)`

GetHealthCheckEnabledOk returns a tuple with the HealthCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEnabled

`func (o *SinkBean) SetHealthCheckEnabled(v bool)`

SetHealthCheckEnabled sets HealthCheckEnabled field to given value.

### HasHealthCheckEnabled

`func (o *SinkBean) HasHealthCheckEnabled() bool`

HasHealthCheckEnabled returns a boolean if a field has been set.

### GetCaCertificate

`func (o *SinkBean) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *SinkBean) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *SinkBean) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *SinkBean) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificateVerificationEnabled

`func (o *SinkBean) GetCertificateVerificationEnabled() bool`

GetCertificateVerificationEnabled returns the CertificateVerificationEnabled field if non-nil, zero value otherwise.

### GetCertificateVerificationEnabledOk

`func (o *SinkBean) GetCertificateVerificationEnabledOk() (*bool, bool)`

GetCertificateVerificationEnabledOk returns a tuple with the CertificateVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVerificationEnabled

`func (o *SinkBean) SetCertificateVerificationEnabled(v bool)`

SetCertificateVerificationEnabled sets CertificateVerificationEnabled field to given value.

### HasCertificateVerificationEnabled

`func (o *SinkBean) HasCertificateVerificationEnabled() bool`

HasCertificateVerificationEnabled returns a boolean if a field has been set.

### GetHostnameVerificationEnabled

`func (o *SinkBean) GetHostnameVerificationEnabled() bool`

GetHostnameVerificationEnabled returns the HostnameVerificationEnabled field if non-nil, zero value otherwise.

### GetHostnameVerificationEnabledOk

`func (o *SinkBean) GetHostnameVerificationEnabledOk() (*bool, bool)`

GetHostnameVerificationEnabledOk returns a tuple with the HostnameVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationEnabled

`func (o *SinkBean) SetHostnameVerificationEnabled(v bool)`

SetHostnameVerificationEnabled sets HostnameVerificationEnabled field to given value.

### HasHostnameVerificationEnabled

`func (o *SinkBean) HasHostnameVerificationEnabled() bool`

HasHostnameVerificationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


