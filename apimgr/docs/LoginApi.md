# \LoginApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginDelete**](LoginApi.md#LoginDelete) | **Delete** /login | Logout from API Manager
[**LoginPost**](LoginApi.md#LoginPost) | **Post** /login | Login to API Manager



## LoginDelete

> LoginDelete(ctx, )

Logout from API Manager

Destroys the caller session with the API Manager.

### Required Parameters

This endpoint does not need any parameter.

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


## LoginPost

> LoginPost(ctx, username, password, optional)

Login to API Manager

Logs a user, identified by _username_ and _password_, into the API Manager by creating unique a session cookie.  The _success_ parameter defaults to a named URL identified by \"/home\" and will return a redirect to the portal after login.  The _failure_ parameter is optional.  If _failure_ is not specified, and the login attempt fails, this method returns a JSON error response.  If _failure_ is specified, and the login attempt fails, then this method will redirect to a named URL, identified by \"/login-failed\".

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| The login name of the authenticating user | 
**password** | **string**| The password of the authenticating user | 
 **optional** | ***LoginPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **success** | **optional.String**| The redirect success location (defaults to: /home) | 
 **failure** | **optional.String**| Optional redirect failure location (e.g. /login-failed | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

