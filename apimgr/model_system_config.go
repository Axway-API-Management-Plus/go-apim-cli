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
// SystemConfig API Manager system configuration
type SystemConfig struct {
	// Returns a list of disabled APIs
	DisabledApis []string `json:"disabledApis,omitempty"`
}
