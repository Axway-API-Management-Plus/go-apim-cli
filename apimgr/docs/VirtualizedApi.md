# VirtualizedApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the virtualized API. | [readonly] 
**OrganizationId** | **string** | The organization that created the virtualized API. | 
**ApiId** | **string** | The identifier of the API that was virtualized.  This is the only attribute that is required to create a virtualized API. | 
**Name** | **string** | The name of this virtualized API. | 
**Version** | **string** | The API version. | 
**ApiRoutingKey** | **string** | The key for routing between two APIs on the same path. | 
**Vhost** | **string** | The virtual host of this virtualized API. | 
**Path** | **string** | The path that services this virtualized API. | 
**DescriptionType** | **string** | Type of descripton to use.  One of: _manual_, _markdown_, _url_, or _original_ (default). | 
**DescriptionManual** | **string** | Markdown string to use as the description of the API. | 
**DescriptionMarkdown** | **string** | Markdown file to use as the description of the API within the API Catalog. | 
**DescriptionUrl** | **string** | External URL to use as the description of the API. | 
**Summary** | **string** | A short summary description of the API. | 
**Retired** | **bool** | Immediately retires the virtualized API. | [default to false]
**Expired** | **bool** | Immediately expires the virtualized API. | [default to false]
**Image** | **string** | URI of the image to be used for this virtualized API. To update the image, please refer to the API. | 
**RetirementDate** | **int64** | Date to retire the virtualized API.  If _retired_ is true, this is the date on which it was retired. | 
**Deprecated** | **bool** | Immediately deprecates the virtualized API.  If deprecated, then _retiredOn_ is optionally used to retire the virtualized API on the specified date. | [default to false]
**State** | **string** | The state of the virtualized API.  One of: unpublished, pending, or published. | 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the virtualized API was created. | [optional] [readonly] 
**CreatedBy** | **string** | The unique identifier for user that created the virtualized API. | [optional] [readonly] 
**CorsProfiles** | [**[]CorsProfile**](CORSProfile.md) | The suite of CORS Profiles for this virtualized API. | 
**SecurityProfiles** | [**[]SecurityProfile**](SecurityProfile.md) | The suite of Security Profiles for this virtualized API. | 
**AuthenticationProfiles** | [**[]AuthenticationProfile**](AuthenticationProfile.md) | The suite of Security Profiles for this virtualized API. | 
**InboundProfiles** | [**map[string]InboundProfiles**](InboundProfiles.md) | The inbound profiles that apply to the virtualized API.  There must always a *\\_default* profile. | 
**OutboundProfiles** | [**map[string]OutboundProfiles**](OutboundProfiles.md) | The outbound profiles that apply to the virtualized API.  There must always a *\\_default* profile. | [optional] 
**ServiceProfiles** | [**map[string]ServiceProfiles**](ServiceProfiles.md) | The inbound profiles that apply to the virtualized API.  There must always a *\\_default* profile. | 
**CaCerts** | [**[]CaCert**](CACert.md) | The outbound profiles that apply to the virtualized API.  There must always a *\\_default* profile. | 
**Tags** | [**map[string][]string**](array.md) | The list of tags associated with this API. Each tag can have multiple values | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


