# \MigrateApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrateApplicationsExportCreatePost**](MigrateApi.md#MigrateApplicationsExportCreatePost) | **Post** /migrate/applications/export/create | Creates a set of export options associated with the current http session
[**MigrateApplicationsExportDownloadGet**](MigrateApi.md#MigrateApplicationsExportDownloadGet) | **Get** /migrate/applications/export/download | Exports Application data for migration to other API Gateways
[**MigrateApplicationsExportJsonPost**](MigrateApi.md#MigrateApplicationsExportJsonPost) | **Post** /migrate/applications/export/json | Creates an export of applications based on the export options posted as a JSON object
[**MigrateApplicationsExportPost**](MigrateApi.md#MigrateApplicationsExportPost) | **Post** /migrate/applications/export | Creates an export of applications based on the export options posted in a form data
[**MigrateApplicationsImportPost**](MigrateApi.md#MigrateApplicationsImportPost) | **Post** /migrate/applications/import | Imports applications to the API Gateway



## MigrateApplicationsExportCreatePost

> MigrateApplicationsExportCreatePost(ctx, body)

Creates a set of export options associated with the current http session

Creates a set of export options associated with the current http session. Options include the password used to encrypt the resulting export, export elements: apikeys, oauth & quotas, the filename of the export, and the list of application ids for inclusion in the export. The exported data can be retrieved subsequently with a GET request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ExportOptions**](ExportOptions.md)| The options for creating an application export file | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateApplicationsExportDownloadGet

> MigrateApplicationsExportDownloadGet(ctx, optional)

Exports Application data for migration to other API Gateways

Retrieves the export options associated with the current user HTTP session and creates a stream or returns the exported data in response body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MigrateApplicationsExportDownloadGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MigrateApplicationsExportDownloadGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | **optional.String**| Optional. If present this method will return an octet stream with an file attachment of the same name | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateApplicationsExportJsonPost

> MigrateApplicationsExportJsonPost(ctx, body)

Creates an export of applications based on the export options posted as a JSON object

Creates an export file based on options including the password used to encrypt the resulting export, export elements: apikeys, oauth & quotas, the filename of the export, and the list of application ids for inclusion in the export. The exported data is returned as part of the response body

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ExportOptions**](ExportOptions.md)| Export options | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateApplicationsExportPost

> MigrateApplicationsExportPost(ctx, apikeys, oauth, quota, optional)

Creates an export of applications based on the export options posted in a form data

Creates an export file based on options including the password used to encrypt the resulting export, export elements: apikeys, oauth & quotas, the filename of the export, and the list of application ids for inclusion in the export. The exported data is returned as part of the response body

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikeys** | **string**| True/False. Include/Exclude api keys | 
**oauth** | **string**| True/False. Include/Exclude oauth credentials | 
**quota** | **string**| True/False. Include/Exclude quotas, if available | 
 **optional** | ***MigrateApplicationsExportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MigrateApplicationsExportPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filename** | **optional.String**| The name of the export file | 
 **password** | **optional.String**| The password used to encrypt the exported file | 
 **appIds** | [**optional.Interface of []string**](string.md)| The list of identifiers for the applications to be exported | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateApplicationsImportPost

> MigrateApplicationsImportPost(ctx, type_, file, organizationId, userId, optional)

Imports applications to the API Gateway

Imports a set of applications and assocated API Keys and OAuth credentials, encrypted files require a decryption password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The type of application response required. THe only valid option is iframe - this is for ajax based calls that require an iFrame, i.e. Internet Explorer version prior to version 9. | 
**file** | ***os.File*****os.File**| The file containing application data to be imported | 
**organizationId** | **string**| The Organization to associate the imported applications with. If applicable, for Core OAuth this parameter will be ignored | 
**userId** | **string**| The user to associate the applications with. Default is the API Admin | 
 **optional** | ***MigrateApplicationsImportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MigrateApplicationsImportPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **password** | **optional.String**| Password to be used for decryption | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

