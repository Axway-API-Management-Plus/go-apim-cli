# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortalName** | **string** | The name of the API Manager. | [optional] 
**PortalHostname** | **string** | The network hostname or IP Address of the API Manager which will be used in email links. | [optional] 
**ApiPortalName** | **string** | The name of the API Portal | [optional] 
**ApiPortalHostname** | **string** | The network hostname or IP Address of the API Portal which will be used in email links. | [optional] 
**IsApiPortalConfigured** | **bool** | Indicates if the API Portal is configured. | [optional] [default to false]
**RegistrationEnabled** | **bool** | Enables/disables user registration for the API Manager | [optional] [default to false]
**ResetPasswordEnabled** | **bool** | Enables/disables spport for resetting user passwords for the API Manager | [optional] [default to false]
**MinimumPasswordLength** | **int32** | The minimum password length. | [optional] 
**AutoApproveUserRegistration** | **bool** | Enables/disables auto-approve for user registration whereby API Administrator or Organization Administrator approval is not required. | [optional] [default to false]
**SystemOAuthScopesEnabled** | **bool** | Enables/disables the ability to add System scopes to an Application. These scopes represent Gateway OAuth resources that are not covered by APIs. | [optional] [default to false]
**AutoApproveApplications** | **bool** | Enables/disables auto-application approval whereby users do not need API Administrator or Organization Administrator approval. | [optional] [default to false]
**DelegateUserAdministration** | **bool** | Enables/disables user administration to the Organization Administrators. | [optional] [default to false]
**DelegateApplicationAdministration** | **bool** | Enables/disables application administration to the Organization Administrators. | [optional] [default to false]
**ApiDefaultVirtualHost** | **string** | The network host and port that serves as the default virtual host from which API Manager registered API will be accessible through. | [optional] 
**ApiRoutingKeyEnabled** | **bool** | Enable routing to APIs on the same base path. | [optional] [default to false]
**ApiRoutingKeyLocation** | **string** | An additional routing key is required to support multiple APIs registered on the same base path. This indicates where to look for the value. | [optional] 
**EmailFrom** | **string** | The &#39;from&#39; address used in emails. | [optional] 
**EmailBounceAddress** | **string** | An email address where undeliverable emails will be bounced to. | [optional] 
**PromoteApiViaPolicy** | **bool** | Enables/disables API promotion via policy. | [optional] [default to false]
**GlobalPoliciesEnabled** | **bool** | Enables/disables Global policies. | [optional] [default to false]
**GlobalRequestPolicy** | **string** | The Global Request Policy to be executed for all Frontend API calls. Must be a valid policy ID. Can be null to indicate no policy | [optional] 
**GlobalResponsePolicy** | **string** | The Global Response Policy to be executed for all Frontend API calls. Must be a valid policy ID. Can be null to indicate no policy | [optional] 
**FaultHandlersEnabled** | **bool** | Enables/disables API Manager fault handlers. | [optional] [default to false]
**GlobalFaultHandlerPolicy** | **string** | The Global Fault Handler Policy to be used by all Frontend APIs in the event of an error. Must be a valid policy ID. Can be null to indicate no policy | [optional] 
**BaseOAuth** | **bool** |  | [optional] [default to false]
**ExternalUserName** | **string** | External user name | [optional] 
**ExternalUserDescription** | **string** | External user description | [optional] 
**ExternalUserPhone** | **string** | External user phone | [optional] 
**ExternalUserEmail** | **string** | External user email | [optional] 
**ExternalUserOrganization** | **string** | External user organization name | [optional] 
**ExternalUserRole** | **string** | External user role | [optional] 
**ExternalUserEnabled** | **string** | External user enabled | [optional] 
**SessionIdleTimeout** | **int64** | Idle session timeout in milliseconds | 
**IsTrial** | **bool** | Is trial enabled | [optional] [default to false]
**DefaultTrialDuration** | **int32** | Default trial duration in days | [optional] 
**LoginNameRegex** | **string** | Login name validation regex | [optional] 
**ProductVersion** | **string** | The Version information of API Manager. | [optional] 
**Os** | **string** | The operating system on which API Manager server is running. | [optional] 
**Architecture** | **string** | The architecture of the operating system on which the API Manager server is running. Supported values: [ win-x86-32, linux-x86-64 ] | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


