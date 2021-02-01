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
// CaCert Trusted CA certificate
type CaCert struct {
	// Raw certificate data
	CertBlob string `json:"certBlob,omitempty"`
	// Name of the certificate
	Name string `json:"name,omitempty"`
	// Certificate alias
	Alias string `json:"alias,omitempty"`
	// Certificate subject
	Subject string `json:"subject,omitempty"`
	// Certificate issuer
	Issuer string `json:"issuer,omitempty"`
	// Version of the certificate
	Version int32 `json:"version,omitempty"`
	// Start date of the certificate
	NotValidBefore int64 `json:"notValidBefore,omitempty"`
	// Expiry date of the certificate
	NotValidAfter int64 `json:"notValidAfter,omitempty"`
	// The algorithm used to sign the certificate.
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`
	// SHA1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// Flag indicating whether or not the certificate is expired.
	Expired bool `json:"expired,omitempty"`
	// Flag indicating whether or not the certificate is valid yet.
	NotYetValid bool `json:"notYetValid,omitempty"`
	// Flag indicating whether this certificate is used for inbound SSL connections when invoking the virtualized API.
	Inbound bool `json:"inbound,omitempty"`
	// Flag indicating whether this certificate is used for outbound SSL connections when invoking the virtualized API.
	Outbound bool `json:"outbound,omitempty"`
}
