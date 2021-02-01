# CorsProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the Profile | [optional] 
**IsDefault** | **bool** | Indicates that this is the default profile.  There can be only one default. | [optional] [default to false]
**Origins** | **[]string** | List of origins for this CORS Profile | [optional] 
**AllowedHeaders** | **[]string** | List of headers... | [optional] 
**ExposedHeaders** | **[]string** | List of headers... | [optional] 
**SupportCredentials** | **bool** | Specifies whether or credentials are supported for APIs/API Methods employing this CORS Profile. | [optional] [default to false]
**MaxAgeSeconds** | **int64** | Specifies the amount of time responses to OPTIONS requests are stored, in seconds, in the preflight result cache | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


