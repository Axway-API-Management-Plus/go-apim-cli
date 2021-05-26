/*
 * API Manager API v1.3
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.0
 * Contact: support@axway.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apimgr
// InboundProfiles Inbound profiles
type InboundProfiles struct {
	// The name of the Security Profile to be used for this Virtualized API.
	SecurityProfile string `json:"securityProfile,omitempty"`
	// The name of the CORS Profile to be used for this Virtualized API.
	CorsProfile string `json:"corsProfile,omitempty"`
	// Enables metrics monitoring for the API
	MonitorAPI bool `json:"monitorAPI,omitempty"`
	// Identifies the client for metrics monitoring
	MonitorSubject string `json:"monitorSubject,omitempty"`
}