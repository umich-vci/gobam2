# \TaskResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserTask**](TaskResourcesAPI.md#DeleteUserTask) | **Delete** /api/v2/users/{collectionId}/tasks/{id} | Delete a task
[**GetUserTask**](TaskResourcesAPI.md#GetUserTask) | **Get** /api/v2/users/{collectionId}/tasks/{id} | Retrieve a single task
[**GetUserTasks**](TaskResourcesAPI.md#GetUserTasks) | **Get** /api/v2/users/{collectionId}/tasks | Retrieve a collection of tasks
[**PostUserTask**](TaskResourcesAPI.md#PostUserTask) | **Post** /api/v2/users/{collectionId}/tasks | Create a new task
[**PutUserTask**](TaskResourcesAPI.md#PutUserTask) | **Put** /api/v2/users/{collectionId}/tasks/{id} | Update a task



## DeleteUserTask

> DeleteUserTask(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a task



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
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskResourcesAPI.DeleteUserTask(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskResourcesAPI.DeleteUserTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserTaskRequest struct via the builder pattern


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


## GetUserTask

> GetUserTasks200ResponseDataInner GetUserTask(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single task



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
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskResourcesAPI.GetUserTask(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskResourcesAPI.GetUserTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTask`: GetUserTasks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TaskResourcesAPI.GetUserTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetUserTasks200ResponseDataInner**](GetUserTasks200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTasks

> GetUserTasks200Response GetUserTasks(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of tasks



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
	resp, r, err := apiClient.TaskResourcesAPI.GetUserTasks(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskResourcesAPI.GetUserTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTasks`: GetUserTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `TaskResourcesAPI.GetUserTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetUserTasks200Response**](GetUserTasks200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserTask

> GetUserTasks200ResponseDataInner PostUserTask(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).TaskPostRequestBody(taskPostRequestBody).Execute()

Create a new task



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
	taskPostRequestBody := *openapiclient.NewTaskPostRequestBody("task1", "Priority_example", "State_example", int32(123)) // TaskPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskResourcesAPI.PostUserTask(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).TaskPostRequestBody(taskPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskResourcesAPI.PostUserTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserTask`: GetUserTasks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TaskResourcesAPI.PostUserTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUserTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **taskPostRequestBody** | [**TaskPostRequestBody**](TaskPostRequestBody.md) |  | 

### Return type

[**GetUserTasks200ResponseDataInner**](GetUserTasks200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUserTask

> GetUserTasks200ResponseDataInner PutUserTask(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).TaskPutRequestBody(taskPutRequestBody).Execute()

Update a task



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
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	taskPutRequestBody := *openapiclient.NewTaskPutRequestBody("task1", "Priority_example", "State_example", int32(123)) // TaskPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskResourcesAPI.PutUserTask(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).TaskPutRequestBody(taskPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskResourcesAPI.PutUserTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUserTask`: GetUserTasks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `TaskResourcesAPI.PutUserTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **taskPutRequestBody** | [**TaskPutRequestBody**](TaskPutRequestBody.md) |  | 

### Return type

[**GetUserTasks200ResponseDataInner**](GetUserTasks200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

