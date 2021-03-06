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
// VirtualizedApi Represents a virtualized, frontend API.
type VirtualizedApi struct {
	// Unique ID of the virtualized API.
	Id string `json:"id"`
	// The organization that created the virtualized API.
	OrganizationId string `json:"organizationId"`
	// The identifier of the API that was virtualized.  This is the only attribute that is required to create a virtualized API.
	ApiId string `json:"apiId"`
	// The name of this virtualized API.
	Name string `json:"name"`
	// The API version.
	Version string `json:"version"`
	// The key for routing between two APIs on the same path.
	ApiRoutingKey string `json:"apiRoutingKey"`
	// The virtual host of this virtualized API.
	Vhost string `json:"vhost"`
	// The path that services this virtualized API.
	Path string `json:"path"`
	// Type of descripton to use.  One of: _manual_, _markdown_, _url_, or _original_ (default).
	DescriptionType string `json:"descriptionType"`
	// Markdown string to use as the description of the API.
	DescriptionManual string `json:"descriptionManual"`
	// Markdown file to use as the description of the API within the API Catalog.
	DescriptionMarkdown string `json:"descriptionMarkdown"`
	// External URL to use as the description of the API.
	DescriptionUrl string `json:"descriptionUrl"`
	// A short summary description of the API.
	Summary string `json:"summary"`
	// Immediately retires the virtualized API.
	Retired bool `json:"retired"`
	// Immediately expires the virtualized API.
	Expired bool `json:"expired"`
	// URI of the image to be used for this virtualized API. To update the image, please refer to the API.
	Image string `json:"image"`
	// Date to retire the virtualized API.  If _retired_ is true, this is the date on which it was retired.
	RetirementDate int64 `json:"retirementDate"`
	// Immediately deprecates the virtualized API.  If deprecated, then _retiredOn_ is optionally used to retire the virtualized API on the specified date.
	Deprecated bool `json:"deprecated"`
	// The state of the virtualized API.  One of: unpublished, pending, or published.
	State string `json:"state"`
	// Epoch/Unix time stamp when the virtualized API was created.
	CreatedOn int64 `json:"createdOn,omitempty"`
	// The unique identifier for user that created the virtualized API.
	CreatedBy string `json:"createdBy,omitempty"`
	// The suite of CORS Profiles for this virtualized API.
	CorsProfiles []CorsProfile `json:"corsProfiles"`
	// The suite of Security Profiles for this virtualized API.
	SecurityProfiles []SecurityProfile `json:"securityProfiles"`
	// The suite of Security Profiles for this virtualized API.
	AuthenticationProfiles []AuthenticationProfile `json:"authenticationProfiles"`
	// The inbound profiles that apply to the virtualized API.  There must always a *\\_default* profile.
	InboundProfiles map[string]InboundProfiles `json:"inboundProfiles"`
	// The outbound profiles that apply to the virtualized API.  There must always a *\\_default* profile.
	OutboundProfiles map[string]OutboundProfiles `json:"outboundProfiles,omitempty"`
	// The inbound profiles that apply to the virtualized API.  There must always a *\\_default* profile.
	ServiceProfiles map[string]ServiceProfiles `json:"serviceProfiles"`
	// The outbound profiles that apply to the virtualized API.  There must always a *\\_default* profile.
	CaCerts []CaCert `json:"caCerts"`
	// The list of tags associated with this API. Each tag can have multiple values
	Tags map[string][]string `json:"tags"`
}
