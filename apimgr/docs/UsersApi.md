# \UsersApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersForgotpasswordPost**](UsersApi.md#UsersForgotpasswordPost) | **Post** /users/forgotpassword | Allows a user to reset their password.
[**UsersGet**](UsersApi.md#UsersGet) | **Get** /users | Obtains a list of users
[**UsersIdApprovePost**](UsersApi.md#UsersIdApprovePost) | **Post** /users/{id}/approve | Grants approval to a request to create a new user on the system.
[**UsersIdChangepasswordPost**](UsersApi.md#UsersIdChangepasswordPost) | **Post** /users/{id}/changepassword | Updates the password for a given user.
[**UsersIdDelete**](UsersApi.md#UsersIdDelete) | **Delete** /users/{id} | Deletes a user.
[**UsersIdGet**](UsersApi.md#UsersIdGet) | **Get** /users/{id} | Retrieves the details for a given user.
[**UsersIdImageGet**](UsersApi.md#UsersIdImageGet) | **Get** /users/{id}/image | Get the image for a user
[**UsersIdImagePost**](UsersApi.md#UsersIdImagePost) | **Post** /users/{id}/image | Set the image for a user
[**UsersIdPut**](UsersApi.md#UsersIdPut) | **Put** /users/{id} | Updates the details for a given user.
[**UsersIdResetpasswordPut**](UsersApi.md#UsersIdResetpasswordPut) | **Put** /users/{id}/resetpassword | Admin level function to reset the password for a given user.
[**UsersPost**](UsersApi.md#UsersPost) | **Post** /users | Admin function to create a new user on the system
[**UsersRegisterPost**](UsersApi.md#UsersRegisterPost) | **Post** /users/register | Register a new user.
[**UsersResetpasswordGet**](UsersApi.md#UsersResetpasswordGet) | **Get** /users/resetpassword | Validates the user [/forgotpassword](#APIUsersforgotUserPassword) password request.
[**UsersValidateuserGet**](UsersApi.md#UsersValidateuserGet) | **Get** /users/validateuser | Validates the user [/register](#APIUsersregisterUser) request.



## UsersForgotpasswordPost

> UsersForgotpasswordPost(ctx, email, optional)

Allows a user to reset their password.

When this method is invoked, an email is sent to the owner of _email_ to verify that they wish for their password to be reset. The owner of _email_ must click on a link to reset the password. The link should direct the user to [/resetpassword](#APIUsersresetForgottenPassword) with appropriate query paremeters. Redirect URLs may be specified for success and failure conditions. If redirect URLs are specified, they must be a known Static File listener configured in the gateway or the request will be rejected.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| The email address of the user. | 
 **optional** | ***UsersForgotpasswordPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersForgotpasswordPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **success** | **optional.String**| The redirect success location (e.g. /request-forgotten-pw-success) | 
 **failure** | **optional.String**| The redirect failure location (e.g. /request-forgotten-pw-failed) | 

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


## UsersGet

> []User UsersGet(ctx, optional)

Obtains a list of users

Returns a list of users that are visible to the authenticated user.  The list of users can be filtered using the expression: field=__field__&op=__op__&value=__value__.  Optionally, you can add a logical operation for all expressions, using the form: &lop=AND|OR.  By default, the logical operation is AND.  Multiple expression filters can be used, specifying field, op, and value for each filter. The __field__ is one of:  apiid :      Matches the user if the user has explicit access to application(s) that are using the API, specified by ID  appid :      Matches the user if the user has explicit access to the application, specified by ID  description :      The user's description  email :      The user's email address  enabled :      The enabled state of the user, one of: enabled, disabled  createdOn :      The date the user was created on, time in ms, e.g.: 1372755998542  mobile :      The user's mobile phone  name :      The name of the user  loginName :  The login name of the user  orgid :      Matches the user if the user is a member of the organization, specified by ID  phone :      The user's phone  role :      The user's role, one of: user or oadmin  state :      The user's current state, one of: approved, pending  surname :      The surname of the user  The __op__ is an operation and is one of:  eq :      Equal  ne :      Not equal  gt :      Greater than  lt :      Less than  ge :      Greater than or equal  le :      Less than or equal  like :      Like  gete :      Greater than or equal to, and less than or equal to; the __value__ should be a lower-minimum and upper-maximum separated by comma, e.g: value=5,10  The __value__ will be compared against the __field__, according to the supplied __op__. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**optional.Interface of []string**](string.md)| Filter field name. | 
 **op** | [**optional.Interface of []string**](string.md)| Filter operation. | 
 **value** | [**optional.Interface of []string**](string.md)| Filter value | 

### Return type

[**[]User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersIdApprovePost

> User UsersIdApprovePost(ctx, id)

Grants approval to a request to create a new user on the system.

Approving user must be API Manager Administrator or an Organization Administrator of the user's organization with the correct privileges to approve new user requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the user to be approved. | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersIdChangepasswordPost

> UsersIdChangepasswordPost(ctx, id, newPassword)

Updates the password for a given user.

The authenticated user must be API Manager Administrator or an Organization Administrator of the user's organization with the correct privileges to invoke this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the user being updated. | 
**newPassword** | **string**| The new password of the user being updated. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersIdDelete

> UsersIdDelete(ctx, id)

Deletes a user.

Deletes a user with the given user ID.  All the applications and keys associated with the deleted user remain in the organization and can be managed by the Organization Administrator or the API Administrator.  The API Administrator can delete any user.  The Organization Administrators can only delete users belonging to their organizations..

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user ID to delete | 

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


## UsersIdGet

> User UsersIdGet(ctx, id)

Retrieves the details for a given user.

Retrieves user details, given a user ID.  The API Manager Administrator may access all users, otherwise, the user ID must be a member of the authenticated user's own organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the user to be retreived. | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersIdImageGet

> UsersIdImageGet(ctx, id)

Get the image for a user

Returns the jpeg image associated with an user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user ID whos image is to be returned | 

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


## UsersIdImagePost

> UsersIdImagePost(ctx, id, type_, file)

Set the image for a user

Set the jpeg image to be associated with a user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The user ID for which an image is to be updated | 
**type_** | **string**| This value should be unset | 
**file** | ***os.File*****os.File**| The file input data | 

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


## UsersIdPut

> User UsersIdPut(ctx, id, optional)

Updates the details for a given user.

Updates user details, given a user ID.  The API Manager Administrator may update all users, otherwise, the user ID must be a member of the authenticated user's own organization and the authenticated user must be an Organization Administrator.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the user being updated. | 
 **optional** | ***UsersIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersIdResetpasswordPut

> User UsersIdResetpasswordPut(ctx, id)

Admin level function to reset the password for a given user.

The authenticated user must be API Manager Administrator or an Organization Administrator of the user's organization with the correct privileges to invoke this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the user having password reset administratively. | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPost

> User UsersPost(ctx, optional)

Admin function to create a new user on the system

Creates a new user on the system.  Only Organization Administrators and API Manager Administrators may create users. Custom properties can be included on create.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRegisterPost

> UsersRegisterPost(ctx, name, email, password, optional)

Register a new user.

Allows a user to register for an account on the system. A validation email request is sent to the provided email address to confirm ownership. The email should contain a link to [/validateuser](#APIUsersvalidateUser) with appropriate parameters. User properties (including custom properties) may be supplied as form parameters. The method will return JSON, but optionally, redirect URLs may be specified for success and failure conditions. If redirect URLs are specified, they must be a known Static File listener configured in the gateway or the request will be rejected.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the user to register. | 
**email** | **string**| The unique email address of the user to register. | 
**password** | **string**| The password of the user to register. | 
 **optional** | ***UsersRegisterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersRegisterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **token** | **optional.String**| The registration token to use. | 
 **success** | **optional.String**| The redirect success location (e.g. &#39;/registration-success&#39;) | 
 **failure** | **optional.String**| The redirect failure location (e.g. &#39;/registration-failed&#39;) | 

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


## UsersResetpasswordGet

> UsersResetpasswordGet(ctx, email, validator)

Validates the user [/forgotpassword](#APIUsersforgotUserPassword) password request.

User validation code and email address are expected as query string parameters.  When invoked, an email will be sent to the user with their new password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| The email address of the user being validated. | 
**validator** | **string**| Validation string for the user entry. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersValidateuserGet

> UsersValidateuserGet(ctx, email, validator)

Validates the user [/register](#APIUsersregisterUser) request.

User validation code and email address are expected as query parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**| The email address of the user being validated. | 
**validator** | **string**| Validation string for the user entry. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

