# Swagger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An identifier | [optional] [readonly] 
**Title** | **string** | Schema title | [optional] 
**Description** | **string** | Description of the Schema | [optional] 
**Format** | **string** | The format ex: int32, int64, float, double, byte, binary, date, date-time or password | [optional] 
**Required** | **[]string** | Specifies if the parameter is required | [optional] 
**Properties** | [**map[string]SchemaObject**](SchemaObject.md) | Not used beacause our model does not support inline nested types | [optional] 
**Items** | [**SchemaObject**](SchemaObject.md) |  | [optional] 
**Example** | [**map[string]interface{}**](.md) | if the schema is an array specifies the items type | [optional] 
**Discriminator** | **string** |  | [optional] 
**ApiVersion** | **string** | The API version | [optional] 
**SwaggerVersion** | **string** | The Swagger version | [optional] 
**BasePath** | **string** | The base path | [optional] 
**ResourcePath** | **string** | The resource path hosted | [optional] 
**Models** | **map[string]map[string]interface{}** | The models/schema for the API | [optional] 
**Consumes** | **[]string** | The content types that the API consumes | [optional] 
**Produces** | **[]string** | The content types that the API produces | [optional] 
**Authorizations** | [**map[string]Authorization**](Authorization.md) | The Authorization schemes provided for this API | [optional] 
**Name** | **string** | The name of the API | [optional] 
**Summary** | **string** | Brief summary of the API. | [optional] 
**SecurityProfile** | [**SwaggerSecurityProfile**](SwaggerSecurityProfile.md) |  | [optional] 
**BasePaths** | **[]string** | Array of basePaths supported for this API/service, based on the configured ports | [optional] 
**Image** | **string** | API image URL | [optional] 
**State** | **string** | The state of the API.  Possible values: &#39;pending&#39;, &#39;unpublished&#39;, or &#39;published&#39;. | [optional] 
**Cors** | **bool** | Indicates that the API is CORS enabled | [optional] [default to false]
**Expired** | **bool** | Indicates that the API is expired. | [optional] [default to false]
**Deprecated** | **bool** | Indicates that the API is deprecated.  If &#39;true&#39;, then the API may have a &#39;retirementDate&#39;. | [optional] [default to false]
**RetirementDate** | **int64** | Indicates that the API is deprecated and will be retired on the supplied date (in milliseconds). | [optional] 
**Tags** | [**map[string][]string**](array.md) | The list of tags associated with this API. Each tag can have multiple values | [optional] 
**DocumentationUrl** | **string** | The documentation URL for the operation | [optional] 
**AvailableApiDefinitions** | **map[string]string** | The schema definitions that this API supports and links to those definitions. | [optional] 
**AvailableSDK** | **map[string]string** | The SDK downloads that this API supports and links to those downloads. | [optional] 
**Apis** | [**[]Api**](API.md) | The API resources | [optional] 
**Ref** | **string** | A Reference to a definition on definitions object | [optional] 
**Default** | [**map[string]interface{}**](.md) | Default value for this schema if it is applicable | [optional] 
**Type** | **string** | The resource type. Possible values: &#39;rest&#39;, &#39;wsdl&#39; | [optional] 
**Enum** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


