# \MetricsApi

All URIs are relative to *http://localhost/api/portal/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsFieldsGet**](MetricsApi.md#MetricsFieldsGet) | **Get** /metrics/fields | Gets a list of metric field names available for summary and timeline queries.
[**MetricsReportsTypeSummaryLevelGet**](MetricsApi.md#MetricsReportsTypeSummaryLevelGet) | **Get** /metrics/reports/{type}/summary/{level} | Gets a summary report for application usage
[**MetricsReportsTypeTimelineLevelMetricTypeGet**](MetricsApi.md#MetricsReportsTypeTimelineLevelMetricTypeGet) | **Get** /metrics/reports/{type}/timeline/{level}/{metricType} | Gets a timeline report for application usage



## MetricsFieldsGet

> []MetricField MetricsFieldsGet(ctx, )

Gets a list of metric field names available for summary and timeline queries.

Retrieves a set of metric fields that may be used when querying or interpreting the summary and timeline reports.  The __metricType__ is the metric name.  The __aggreggateName__ is the metric name for the aggregated __metricType__.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MetricField**](MetricField.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsReportsTypeSummaryLevelGet

> MetricsReportsTypeSummaryLevelGet(ctx, type_, level, from, to, optional)

Gets a summary report for application usage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The report type, either &#39;app&#39; or &#39;api&#39; | 
**level** | **int32**| The report level (0 or 1 for drill-through) | 
**from** | **string**| The starting date/time for the report. | 
**to** | **string**| The end date/time for the report. | 
 **optional** | ***MetricsReportsTypeSummaryLevelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MetricsReportsTypeSummaryLevelGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **client** | [**optional.Interface of []string**](string.md)| Filter a specific client ID (multiple permitted). | 
 **service** | [**optional.Interface of []string**](string.md)| Filter a specific service name (multiple permitted). | 
 **method** | **optional.String**| Filter a specific method. | 
 **organization** | **optional.String**| Filter a specific organziation. | 
 **reportsubtype** | **optional.String**| Define the report subtype. Allowed values are : &#39;original&#39;, &#39;trafficAll&#39; or &#39;trafficSubset&#39;. Defaults to &#39;original&#39; if absent | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsReportsTypeTimelineLevelMetricTypeGet

> MetricTimeline MetricsReportsTypeTimelineLevelMetricTypeGet(ctx, type_, level, from, to, metricType, optional)

Gets a timeline report for application usage

Produces a timeline report for a __metricType__ over a specified time range.  The __from__ and __two__ parameters should be a URL encoded ISO-8601 combined date and time format (e.g. 2013-03-13T00%3A00%3A00Z).  The __metricType__ name is one of the types returned from [fields](#APIMetricsgetMetricFields).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| The report type, either &#39;app&#39; or &#39;api&#39; | 
**level** | **int32**| The report level (0 or 1 for drill-through) | 
**from** | **string**| The starting date/time for the report. | 
**to** | **string**| The end date/time for the report. | 
**metricType** | **string**| The metric type to query. | 
 **optional** | ***MetricsReportsTypeTimelineLevelMetricTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MetricsReportsTypeTimelineLevelMetricTypeGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **client** | [**optional.Interface of []string**](string.md)| Filter a specific client ID (multiple permitted). | 
 **service** | [**optional.Interface of []string**](string.md)| Filter a specific service name (multiple permitted). | 
 **method** | **optional.String**| Filter a specific method. | 
 **organization** | **optional.String**| Filter a specific organziation. | 
 **reportsubtype** | **optional.String**| Define the report subtype. Allowed values are : &#39;original&#39;, &#39;trafficAll&#39; or &#39;trafficSubset&#39;. Defaults to &#39;original&#39; if absent | 

### Return type

[**MetricTimeline**](MetricTimeline.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

