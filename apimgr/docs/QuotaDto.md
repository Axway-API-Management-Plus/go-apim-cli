# QuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The quota identifier | 
**Type** | **string** | The quota type, either API or APPLICATION | 
**Name** | **string** | The name of the quota | 
**Description** | **string** | The quota for MyApplication the overrides default Application quota | [optional] 
**Restrictions** | [**[]QuotaApiConstraintDto**](QuotaApiConstraintDTO.md) | An array of restrictions imposed on the quota | [optional] 
**System** | **bool** | Indicates if the quota is system (protected) | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


