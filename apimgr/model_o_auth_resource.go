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
// OAuthResource OAuth protected resource.
type OAuthResource struct {
	// The unique identifier for the oauth protected resource
	Id string `json:"id"`
	// The unique identifier for the application who has access to this resource
	ApplicationId string `json:"applicationId"`
	// The uri prefix which this oauth protected resource relates to
	Uriprefix string `json:"uriprefix,omitempty"`
	// Set of scope strings that have been resolved from querying the OAuth Resource Service at the uri prefix
	Scopes []string `json:"scopes,omitempty"`
	// Flag to indicate if this oauth protected resource is enabled or not
	Enabled bool `json:"enabled,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Scope string `json:"scope,omitempty"`
}
