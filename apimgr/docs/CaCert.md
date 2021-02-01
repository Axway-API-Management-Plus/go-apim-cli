# CaCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBlob** | **string** | Raw certificate data | [optional] 
**Name** | **string** | Name of the certificate | [optional] 
**Alias** | **string** | Certificate alias | [optional] 
**Subject** | **string** | Certificate subject | [optional] 
**Issuer** | **string** | Certificate issuer | [optional] 
**Version** | **int32** | Version of the certificate | [optional] 
**NotValidBefore** | **int64** | Start date of the certificate | [optional] 
**NotValidAfter** | **int64** | Expiry date of the certificate | [optional] 
**SignatureAlgorithm** | **string** | The algorithm used to sign the certificate. | [optional] 
**Sha1Fingerprint** | **string** | SHA1 fingerprint. | [optional] 
**Md5Fingerprint** | **string** | MD5 fingerprint. | [optional] 
**Expired** | **bool** | Flag indicating whether or not the certificate is expired. | [optional] [default to false]
**NotYetValid** | **bool** | Flag indicating whether or not the certificate is valid yet. | [optional] [default to false]
**Inbound** | **bool** | Flag indicating whether this certificate is used for inbound SSL connections when invoking the virtualized API. | [optional] [default to false]
**Outbound** | **bool** | Flag indicating whether this certificate is used for outbound SSL connections when invoking the virtualized API. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


