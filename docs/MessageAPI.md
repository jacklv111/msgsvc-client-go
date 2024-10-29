# \MessageAPI

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](MessageAPI.md#CreateMessage) | **Post** /message | Create message
[**DownloadData**](MessageAPI.md#DownloadData) | **Get** /download | Download message data
[**UploadData**](MessageAPI.md#UploadData) | **Post** /upload | Upload message data if necessary



## CreateMessage

> CreateMessage201Response CreateMessage(ctx).Authorization(authorization).CreateMessage(createMessage).Execute()

Create message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/msgsvc-client-go"
)

func main() {
	authorization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Access token
	createMessage := *openapiclient.NewCreateMessage() // CreateMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.CreateMessage(context.Background()).Authorization(authorization).CreateMessage(createMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.CreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessage`: CreateMessage201Response
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.CreateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token | 
 **createMessage** | [**CreateMessage**](CreateMessage.md) |  | 

### Return type

[**CreateMessage201Response**](CreateMessage201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadData

> *os.File DownloadData(ctx).Authorization(authorization).MessageId(messageId).Execute()

Download message data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/msgsvc-client-go"
)

func main() {
	authorization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Access token
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.DownloadData(context.Background()).Authorization(authorization).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.DownloadData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadData`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.DownloadData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token | 
 **messageId** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadData

> UploadData200Response UploadData(ctx).Authorization(authorization).Body(body).Execute()

Upload message data if necessary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/jacklv111/msgsvc-client-go"
)

func main() {
	authorization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Access token
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.UploadData(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.UploadData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadData`: UploadData200Response
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.UploadData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token | 
 **body** | ***os.File** |  | 

### Return type

[**UploadData200Response**](UploadData200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

