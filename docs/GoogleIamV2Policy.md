

# GoogleIamV2Policy

Data for an IAM policy.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**annotations** | **Map&lt;String, String&gt;** | A key-value map to store arbitrary metadata for the &#x60;Policy&#x60;. Keys can be up to 63 characters. Values can be up to 255 characters. |  [optional] |
|**createTime** | **String** | Output only. The time when the &#x60;Policy&#x60; was created. |  [optional] [readonly] |
|**deleteTime** | **String** | Output only. The time when the &#x60;Policy&#x60; was deleted. Empty if the policy is not deleted. |  [optional] [readonly] |
|**displayName** | **String** | A user-specified description of the &#x60;Policy&#x60;. This value can be up to 63 characters. |  [optional] |
|**etag** | **String** | An opaque tag that identifies the current version of the &#x60;Policy&#x60;. IAM uses this value to help manage concurrent updates, so they do not cause one update to be overwritten by another. If this field is present in a CreatePolicyRequest, the value is ignored. |  [optional] |
|**kind** | **String** | Output only. The kind of the &#x60;Policy&#x60;. Always contains the value &#x60;DenyPolicy&#x60;. |  [optional] [readonly] |
|**name** | **String** | Immutable. The resource name of the &#x60;Policy&#x60;, which must be unique. Format: &#x60;policies/{attachment_point}/denypolicies/{policy_id}&#x60; The attachment point is identified by its URL-encoded full resource name, which means that the forward-slash character, &#x60;/&#x60;, must be written as &#x60;%2F&#x60;. For example, &#x60;policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-deny-policy&#x60;. For organizations and folders, use the numeric ID in the full resource name. For projects, requests can use the alphanumeric or the numeric ID. Responses always contain the numeric ID. |  [optional] |
|**rules** | [**List&lt;GoogleIamV2PolicyRule&gt;**](GoogleIamV2PolicyRule.md) | A list of rules that specify the behavior of the &#x60;Policy&#x60;. All of the rules should be of the &#x60;kind&#x60; specified in the &#x60;Policy&#x60;. |  [optional] |
|**uid** | **String** | Immutable. The globally unique ID of the &#x60;Policy&#x60;. Assigned automatically when the &#x60;Policy&#x60; is created. |  [optional] |
|**updateTime** | **String** | Output only. The time when the &#x60;Policy&#x60; was last updated. |  [optional] [readonly] |



