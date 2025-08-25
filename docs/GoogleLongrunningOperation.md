

# GoogleLongrunningOperation

This resource represents a long-running operation that is the result of a network API call.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**done** | **Boolean** | If the value is &#x60;false&#x60;, it means the operation is still in progress. If &#x60;true&#x60;, the operation is completed, and either &#x60;error&#x60; or &#x60;response&#x60; is available. |  [optional] |
|**error** | [**GoogleRpcStatus**](GoogleRpcStatus.md) |  |  [optional] |
|**metadata** | **Map&lt;String, Object&gt;** | Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any. |  [optional] |
|**name** | **String** | The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the &#x60;name&#x60; should be a resource name ending with &#x60;operations/{unique_id}&#x60;. |  [optional] |
|**response** | **Map&lt;String, Object&gt;** | The normal, successful response of the operation. If the original method returns no data on success, such as &#x60;Delete&#x60;, the response is &#x60;google.protobuf.Empty&#x60;. If the original method is standard &#x60;Get&#x60;/&#x60;Create&#x60;/&#x60;Update&#x60;, the response should be the resource. For other methods, the response should have the type &#x60;XxxResponse&#x60;, where &#x60;Xxx&#x60; is the original method name. For example, if the original method name is &#x60;TakeSnapshot()&#x60;, the inferred response type is &#x60;TakeSnapshotResponse&#x60;. |  [optional] |



