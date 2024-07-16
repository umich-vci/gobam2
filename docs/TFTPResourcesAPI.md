# \TFTPResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFile**](TFTPResourcesAPI.md#DeleteFile) | **Delete** /api/v2/files/{id} | Delete a file
[**DeleteTftpGroup**](TFTPResourcesAPI.md#DeleteTftpGroup) | **Delete** /api/v2/tftpGroups/{id} | Delete a TFTP group
[**GetCollectionFiles**](TFTPResourcesAPI.md#GetCollectionFiles) | **Get** /api/v2/{collection}/{collectionId}/files | Retrieve a collection of files
[**GetConfigurationTftpGroups**](TFTPResourcesAPI.md#GetConfigurationTftpGroups) | **Get** /api/v2/configurations/{collectionId}/tftpGroups | Retrieve a collection of TFTP groups
[**GetFile**](TFTPResourcesAPI.md#GetFile) | **Get** /api/v2/files/{id} | Retrieve a single file
[**GetFiles**](TFTPResourcesAPI.md#GetFiles) | **Get** /api/v2/files | Retrieve a collection of files
[**GetTftpGroup**](TFTPResourcesAPI.md#GetTftpGroup) | **Get** /api/v2/tftpGroups/{id} | Retrieve a single TFTP group
[**GetTftpGroups**](TFTPResourcesAPI.md#GetTftpGroups) | **Get** /api/v2/tftpGroups | Retrieve a collection of TFTP groups
[**PostCollectionFile**](TFTPResourcesAPI.md#PostCollectionFile) | **Post** /api/v2/{collection}/{collectionId}/files | Create a new file
[**PostConfigurationTftpGroup**](TFTPResourcesAPI.md#PostConfigurationTftpGroup) | **Post** /api/v2/configurations/{collectionId}/tftpGroups | Create a new TFTP group
[**PutFile**](TFTPResourcesAPI.md#PutFile) | **Put** /api/v2/files/{id} | Update a file
[**PutTftpGroup**](TFTPResourcesAPI.md#PutTftpGroup) | **Put** /api/v2/tftpGroups/{id} | Update a TFTP group



## DeleteFile

> DeleteFile(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a file



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
	r, err := apiClient.TFTPResourcesAPI.DeleteFile(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.DeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


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


## DeleteTftpGroup

> DeleteTftpGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a TFTP group



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
	r, err := apiClient.TFTPResourcesAPI.DeleteTftpGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.DeleteTftpGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTftpGroupRequest struct via the builder pattern


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


## GetCollectionFiles

> GetFiles200Response GetCollectionFiles(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of files



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
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.GetCollectionFiles(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetCollectionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionFiles`: GetFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetCollectionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetFiles200Response**](GetFiles200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationTftpGroups

> GetTftpGroups200Response GetConfigurationTftpGroups(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of TFTP groups



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
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.GetConfigurationTftpGroups(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetConfigurationTftpGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationTftpGroups`: GetTftpGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetConfigurationTftpGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationTftpGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTftpGroups200Response**](GetTftpGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> GetFiles200ResponseDataInner GetFile(ctx, id).Fields(fields).Execute()

Retrieve a single file



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
	resp, r, err := apiClient.TFTPResourcesAPI.GetFile(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: GetFiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiles

> GetFiles200Response GetFiles(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of files



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
	resp, r, err := apiClient.TFTPResourcesAPI.GetFiles(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiles`: GetFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetFiles200Response**](GetFiles200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTftpGroup

> GetTftpGroups200ResponseDataInner GetTftpGroup(ctx, id).Fields(fields).Execute()

Retrieve a single TFTP group



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
	resp, r, err := apiClient.TFTPResourcesAPI.GetTftpGroup(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetTftpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTftpGroup`: GetTftpGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetTftpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTftpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetTftpGroups200ResponseDataInner**](GetTftpGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTftpGroups

> GetTftpGroups200Response GetTftpGroups(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of TFTP groups



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
	resp, r, err := apiClient.TFTPResourcesAPI.GetTftpGroups(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.GetTftpGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTftpGroups`: GetTftpGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.GetTftpGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTftpGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTftpGroups200Response**](GetTftpGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionFile

> GetFiles200ResponseDataInner PostCollectionFile(ctx, collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).Id(id).Type_(type_).Name(name).UserDefinedFields(userDefinedFields).Version(version).Description(description).Size(size).Md5Checksum(md5Checksum).Execute()

Create a new file



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
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	id := int64(789) // int64 | The resource identifier. (optional)
	type_ := "type__example" // string | The resource type. (optional)
	name := "name_example" // string | The name of the resource. (optional)
	userDefinedFields := map[string]string{"key": "Inner_example"} // map[string]string | User-defined fields set for the resource. (optional)
	version := "version_example" // string | The version of the TFTP file. (optional)
	description := "description_example" // string | The description of the TFTP file. (optional)
	size := int64(789) // int64 | The size of the TFTP file, in bytes. (optional)
	md5Checksum := "md5Checksum_example" // string | The MD5 checksum of the TFTP file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.PostCollectionFile(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).Id(id).Type_(type_).Name(name).UserDefinedFields(userDefinedFields).Version(version).Description(description).Size(size).Md5Checksum(md5Checksum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.PostCollectionFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionFile`: GetFiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.PostCollectionFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **id** | **int64** | The resource identifier. | 
 **type_** | **string** | The resource type. | 
 **name** | **string** | The name of the resource. | 
 **userDefinedFields** | **map[string]string** | User-defined fields set for the resource. | 
 **version** | **string** | The version of the TFTP file. | 
 **description** | **string** | The description of the TFTP file. | 
 **size** | **int64** | The size of the TFTP file, in bytes. | 
 **md5Checksum** | **string** | The MD5 checksum of the TFTP file. | 

### Return type

[**GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfigurationTftpGroup

> GetTftpGroups200ResponseDataInner PostConfigurationTftpGroup(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).TFTPGroupPostRequestBody(tFTPGroupPostRequestBody).Execute()

Create a new TFTP group



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
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	tFTPGroupPostRequestBody := *openapiclient.NewTFTPGroupPostRequestBody("TFTP GROUP 1") // TFTPGroupPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.PostConfigurationTftpGroup(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).TFTPGroupPostRequestBody(tFTPGroupPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.PostConfigurationTftpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfigurationTftpGroup`: GetTftpGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.PostConfigurationTftpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationTftpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **tFTPGroupPostRequestBody** | [**TFTPGroupPostRequestBody**](TFTPGroupPostRequestBody.md) |  | 

### Return type

[**GetTftpGroups200ResponseDataInner**](GetTftpGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFile

> GetFiles200ResponseDataInner PutFile(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Id2(id2).Type_(type_).Name(name).UserDefinedFields(userDefinedFields).Version(version).Description(description).Size(size).Md5Checksum(md5Checksum).Execute()

Update a file



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
	id2 := int64(789) // int64 | The resource identifier. (optional)
	type_ := "type__example" // string | The resource type. (optional)
	name := "name_example" // string | The name of the resource. (optional)
	userDefinedFields := map[string]string{"key": "Inner_example"} // map[string]string | User-defined fields set for the resource. (optional)
	version := "version_example" // string | The version of the TFTP file. (optional)
	description := "description_example" // string | The description of the TFTP file. (optional)
	size := int64(789) // int64 | The size of the TFTP file, in bytes. (optional)
	md5Checksum := "md5Checksum_example" // string | The MD5 checksum of the TFTP file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.PutFile(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Id2(id2).Type_(type_).Name(name).UserDefinedFields(userDefinedFields).Version(version).Description(description).Size(size).Md5Checksum(md5Checksum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.PutFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFile`: GetFiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.PutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **id2** | **int64** | The resource identifier. | 
 **type_** | **string** | The resource type. | 
 **name** | **string** | The name of the resource. | 
 **userDefinedFields** | **map[string]string** | User-defined fields set for the resource. | 
 **version** | **string** | The version of the TFTP file. | 
 **description** | **string** | The description of the TFTP file. | 
 **size** | **int64** | The size of the TFTP file, in bytes. | 
 **md5Checksum** | **string** | The MD5 checksum of the TFTP file. | 

### Return type

[**GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTftpGroup

> GetTftpGroups200ResponseDataInner PutTftpGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).TFTPGroupPutRequestBody(tFTPGroupPutRequestBody).Execute()

Update a TFTP group



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
	tFTPGroupPutRequestBody := *openapiclient.NewTFTPGroupPutRequestBody() // TFTPGroupPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TFTPResourcesAPI.PutTftpGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).TFTPGroupPutRequestBody(tFTPGroupPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TFTPResourcesAPI.PutTftpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTftpGroup`: GetTftpGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TFTPResourcesAPI.PutTftpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTftpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **tFTPGroupPutRequestBody** | [**TFTPGroupPutRequestBody**](TFTPGroupPutRequestBody.md) |  | 

### Return type

[**GetTftpGroups200ResponseDataInner**](GetTftpGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

