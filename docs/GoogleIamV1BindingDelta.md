

# GoogleIamV1BindingDelta

One delta entry for Binding. Each individual change (only one member in each entry) to a binding will be a separate entry.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**action** | [**ActionEnum**](#ActionEnum) | The action that was performed on a Binding. Required |  [optional] |
|**condition** | [**GoogleTypeExpr**](GoogleTypeExpr.md) |  |  [optional] |
|**member** | **String** | A single identity requesting access for a Google Cloud resource. Follows the same format of Binding.members. Required |  [optional] |
|**role** | **String** | Role that is assigned to &#x60;members&#x60;. For example, &#x60;roles/viewer&#x60;, &#x60;roles/editor&#x60;, or &#x60;roles/owner&#x60;. Required |  [optional] |



## Enum: ActionEnum

| Name | Value |
|---- | -----|
| ACTION_UNSPECIFIED | &quot;ACTION_UNSPECIFIED&quot; |
| ADD | &quot;ADD&quot; |
| REMOVE | &quot;REMOVE&quot; |



