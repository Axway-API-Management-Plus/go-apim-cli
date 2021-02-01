# RemoteHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the remote host | [optional] [readonly] 
**Name** | **string** | The host name of the remote host | 
**Port** | **int32** | The host port of the remote host | 
**CreatedOn** | **int64** | Epoch/Unix time stamp when the remote host was created | [optional] [readonly] 
**CreatedBy** | **string** | The unique identifier for user that created the remote host | [optional] [readonly] 
**OrganizationId** | **string** | The organization identifier to which the remote host belongs | [optional] 
**MaxConnections** | **int32** | The maximum number of connections to open to a Remote Host. If the maximum number of connections has already been established, the API Gateway instance waits for a connection to drop or become idle before making another request. The default value is -1, meaning there is no limit | [optional] 
**AllowHTTP11** | **bool** | Enables the API Gateway to use HTTP 1.1 when connecting to the remote host. Default value is false, meaning HTTP 1.0 is used | [optional] [default to false]
**IncludeContentLengthRequest** | **bool** | If this option is set, the API Gateway will include the Content-Length HTTP header in all requests to this Remote Host. Default value is false. | [optional] [default to false]
**IncludeContentLengthResponse** | **bool** | If this option is set, if the API Gateway receives a response from this Remote Host that contains a Content-Length HTTP header, it returns this length to the client. Default value is false. | [optional] [default to false]
**OfferTLSServerName** | **bool** | Adds a field to outbound TLS/SSL calls that shows the name that the client used to connect. Default value is false. | [optional] [default to false]
**VerifyServerHostname** | **bool** | Ensures that the certificate presented by the server matches the name of the remote host being connected to. This prevents host spoofing and man-in-the-middle attacks. Default value is false. | [optional] [default to false]
**ConnectionTimeout** | **int64** | If a connection to this remote host is not established within the time set in this field, the connection times out and the connection fails. Default value is 30000 milliseconds (30 seconds). | [optional] 
**ActiveTimeout** | **int64** | The maximum amount of time permitted between reading successive blocks of data. If the Active Timeout is exceeded, the API Gateway closes the connection. This prevents a Remote Host from closing the connection while sending data. Default value is 30000 milliseconds (30 seconds). | [optional] 
**TransactionTimeout** | **int64** | The maximum amount of time permitted to complete the transaction. Default value is 240000 milliseconds (4 minutes). | [optional] 
**IdleTimeout** | **int64** | The maximum amount of time that API Gateway waits after sending a message over a persistent connection to the Remote Host before it closes the connection. Default value is 15000 milliseconds (15 seconds). | [optional] 
**MaxReceiveBytes** | **int64** | The maximum amount of data the API Gateway can receive per transaction. Default value is 20971520 bytes (20MiB). | [optional] 
**MaxSendBytes** | **int64** | The maximum amount of data the API Gateway can transmit per transaction. Default value is 20971520 bytes (20MiB). | [optional] 
**InputBufferSize** | **int64** | The maximum amount of memory allocated to each request. Default value is 8192 bytes. | [optional] 
**OutputBufferSize** | **int64** | The maximum amount of memory allocated to each response. Default value is 8192 bytes. | [optional] 
**AddressCacheTimeout** | **int64** | The period of time to cache addressing information after it has been received from the naming service. Default value is 300000 milliseconds (5 minutes) | [optional] 
**SslSessionCacheSize** | **int32** | Specifies the size of the SSL session cache for connections to the remote host, which controls the number of idle SSL sessions that can be kept in memory. Default value is 32. | [optional] 
**InputEncodings** | **[]string** | Specifies the HTTP content encodings that the API Gateway can accept from peers. Supported encodings: *deflate*, *gzip*. If no encodings are specified the default encoding is applied. | [optional] 
**OutputEncodings** | **[]string** | Specifies the HTTP content encodings that the API Gateway can apply to outgoing messages. Supported encodings: *deflate*, *gzip*. If no encodings are specified the default encoding is applied. | [optional] 
**ExportCorrelationId** | **bool** | Specifies whether to add the X-CorrelationID header to outbound messages | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


