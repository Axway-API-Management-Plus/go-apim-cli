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
// VirtualizedApiMethod Represents a virtualized API method as part of a frontend API.
type VirtualizedApiMethod struct {
	// Unique ID of the method.
	Id string `json:"id,omitempty"`
	// The identifier of the [VirtualizedAPI](VirtualizedAPI.html).
	VirtualizedApiId string `json:"virtualizedApiId,omitempty"`
	// The virtualized method name.  This defaults to the original [APIDefinition](APIDefinition.html) method name.
	Name string `json:"name,omitempty"`
	// The reference identifier for the original [APIDefinition](APIDefinition.html) that was virtualized.
	ApiId string `json:"apiId,omitempty"`
	// The reference identifier for the original API [APIDefinition](APIDefinition.html) method that was virtualized.
	ApiMethodId string `json:"apiMethodId,omitempty"`
	// A summary of the API Method.
	Summary string `json:"summary,omitempty"`
	// The source for the method's description.  One of: *original*, *manual*, *markdown*, or *url*.  Defaults to *original*.
	DescriptionType string `json:"descriptionType,omitempty"`
	// Specifies a manual description, which can be markdown text.
	DescriptionManual string `json:"descriptionManual,omitempty"`
	// specifies a markdown file to use for description.
	DescriptionMarkdown string `json:"descriptionMarkdown,omitempty"`
	// Specifies a URL to use instead of description text.
	DescriptionUrl string `json:"descriptionUrl,omitempty"`
	// The list of tags associated with this API method. Each tag can have multiple values.
	Tags map[string][]string `json:"tags,omitempty"`
}
