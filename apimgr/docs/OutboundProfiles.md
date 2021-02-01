# OutboundProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationProfile** | **string** | Outbound authentication credentials if __routeType__ is &#39;proxy&#39;.  Can be null to indicate no credentials. | [optional] 
**RouteType** | **string** | The route type.  Values are either: &#39;proxy&#39; or &#39;policy&#39;.    Can be null and defaults to *proxy*. | [optional] 
**RequestPolicy** | **string** | Request policy that applies to all outbound requests.  Can be null to indicate no policy. | [optional] 
**ResponsePolicy** | **string** | Response policy that applies to all responses from outbound requests.  Can be null to indicate no policy. | [optional] 
**RoutePolicy** | **string** | Route policy if the routeType is &#39;policy&#39;, in which case it must be a valid policy ID.  Can be null if __routeType__ is &#39;proxy&#39;. | [optional] 
**FaultHandlerPolicy** | **string** | Fault handler policy that gets executed in the event of an error.  Can be null to indicate no policy. | [optional] 
**ApiId** | **string** | Route to a different API.  Can be null.  Ignored on the &#39;default&#39; outbound profile. | [optional] 
**ApiMethodId** | **string** | Route to a different API method.  Can be null.  Ignored on the &#39;default&#39; outbound profile. | [optional] 
**Parameters** | [**[]ParamValue**](ParamValue.md) | A list of outbound parameters values - maps from frontend parameters to backend parameters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


