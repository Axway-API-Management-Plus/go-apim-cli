# VirtualizedMethodExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the method. | [optional] 
**VirtualizedApiId** | **string** | The identifier of the [VirtualizedAPI](VirtualizedAPI.html). | [optional] 
**Name** | **string** | The virtualized method name.  This defaults to the original [APIDefinition](APIDefinition.html) method name. | [optional] 
**ApiId** | **string** | The reference identifier for the original [APIDefinition](APIDefinition.html) that was virtualized. | [optional] 
**ApiMethodId** | **string** | The reference identifier for the original API [APIDefinition](APIDefinition.html) method that was virtualized. | [optional] 
**Summary** | **string** | A summary of the API Method. | [optional] 
**DescriptionType** | **string** | The source for the method&#39;s description.  One of: *original*, *manual*, *markdown*, or *url*.  Defaults to *original*. | [optional] 
**DescriptionManual** | **string** | Specifies a manual description, which can be markdown text. | [optional] 
**DescriptionMarkdown** | **string** | specifies a markdown file to use for description. | [optional] 
**DescriptionUrl** | **string** | Specifies a URL to use instead of description text. | [optional] 
**Tags** | [**map[string][]string**](array.md) | The list of tags associated with this API method. Each tag can have multiple values. | [optional] 
**Op** | **string** | Internal use only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


