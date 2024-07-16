# \ConfigurationResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration**](ConfigurationResourcesAPI.md#DeleteConfiguration) | **Delete** /api/v2/configurations/{id} | Delete a configuration
[**GetConfiguration**](ConfigurationResourcesAPI.md#GetConfiguration) | **Get** /api/v2/configurations/{id} | Retrieve a single configuration
[**GetConfigurations**](ConfigurationResourcesAPI.md#GetConfigurations) | **Get** /api/v2/configurations | Retrieve a collection of configurations
[**PostConfiguration**](ConfigurationResourcesAPI.md#PostConfiguration) | **Post** /api/v2/configurations | Create a new configuration
[**PutConfiguration**](ConfigurationResourcesAPI.md#PutConfiguration) | **Put** /api/v2/configurations/{id} | Update a configuration



## DeleteConfiguration

> DeleteConfiguration(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationResourcesAPI.DeleteConfiguration(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationResourcesAPI.DeleteConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> GetConfigurations200ResponseDataInner GetConfiguration(ctx, id).Fields(fields).Execute()

Retrieve a single configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationResourcesAPI.GetConfiguration(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationResourcesAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: GetConfigurations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationResourcesAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetConfigurations200ResponseDataInner**](GetConfigurations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurations

> GetConfigurations200Response GetConfigurations(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationResourcesAPI.GetConfigurations(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationResourcesAPI.GetConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurations`: GetConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationResourcesAPI.GetConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetConfigurations200Response**](GetConfigurations200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfiguration

> GetConfigurations200ResponseDataInner PostConfiguration(ctx).XBcnChangeControlComment(xBcnChangeControlComment).ConfigurationPostRequestBody(configurationPostRequestBody).Execute()

Create a new configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	configurationPostRequestBody := *openapiclient.NewConfigurationPostRequestBody("name") // ConfigurationPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationResourcesAPI.PostConfiguration(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).ConfigurationPostRequestBody(configurationPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationResourcesAPI.PostConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfiguration`: GetConfigurations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationResourcesAPI.PostConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **configurationPostRequestBody** | [**ConfigurationPostRequestBody**](ConfigurationPostRequestBody.md) |  | 

### Return type

[**GetConfigurations200ResponseDataInner**](GetConfigurations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfiguration

> GetConfigurations200ResponseDataInner PutConfiguration(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).ConfigurationPutRequestBody(configurationPutRequestBody).Execute()

Update a configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	configurationPutRequestBody := *openapiclient.NewConfigurationPutRequestBody("name", false, false, false, false, "CheckIntegrityValidation_example", "CheckMxCnameValidation_example", "CheckMxValidation_example", "CheckNamesValidation_example", "CheckNsValidation_example", "CheckSrvCnameValidation_example", "CheckWildcardValidation_example", false, false, "BlockUsageCalculation_example", false) // ConfigurationPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationResourcesAPI.PutConfiguration(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).ConfigurationPutRequestBody(configurationPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationResourcesAPI.PutConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConfiguration`: GetConfigurations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationResourcesAPI.PutConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **configurationPutRequestBody** | [**ConfigurationPutRequestBody**](ConfigurationPutRequestBody.md) |  | 

### Return type

[**GetConfigurations200ResponseDataInner**](GetConfigurations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

