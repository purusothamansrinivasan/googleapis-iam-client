package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GoogleIamV1BindingDelta represents the GoogleIamV1BindingDelta schema from the OpenAPI specification
type GoogleIamV1BindingDelta struct {
	Member string `json:"member,omitempty"` // A single identity requesting access for a Google Cloud resource. Follows the same format of Binding.members. Required
	Role string `json:"role,omitempty"` // Role that is assigned to `members`. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. Required
	Action string `json:"action,omitempty"` // The action that was performed on a Binding. Required
	Condition GoogleTypeExpr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
}

// GoogleIamV2PolicyOperationMetadata represents the GoogleIamV2PolicyOperationMetadata schema from the OpenAPI specification
type GoogleIamV2PolicyOperationMetadata struct {
	Createtime string `json:"createTime,omitempty"` // Timestamp when the `google.longrunning.Operation` was created.
}

// GoogleIamV2DenyRule represents the GoogleIamV2DenyRule schema from the OpenAPI specification
type GoogleIamV2DenyRule struct {
	Exceptionprincipals []string `json:"exceptionPrincipals,omitempty"` // The identities that are excluded from the deny rule, even if they are listed in the `denied_principals`. For example, you could add a Google group to the `denied_principals`, then exclude specific users who belong to that group. This field can contain the same values as the `denied_principals` field, excluding `principalSet://goog/public:all`, which represents all users on the internet.
	Denialcondition GoogleTypeExpr `json:"denialCondition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Deniedpermissions []string `json:"deniedPermissions,omitempty"` // The permissions that are explicitly denied by this rule. Each permission uses the format `{service_fqdn}/{resource}.{verb}`, where `{service_fqdn}` is the fully qualified domain name for the service. For example, `iam.googleapis.com/roles.list`.
	Deniedprincipals []string `json:"deniedPrincipals,omitempty"` // The identities that are prevented from using one or more permissions on Google Cloud resources. This field can contain the following values: * `principal://goog/subject/{email_id}`: A specific Google Account. Includes Gmail, Cloud Identity, and Google Workspace user accounts. For example, `principal://goog/subject/alice@example.com`. * `principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}`: A Google Cloud service account. For example, `principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com`. * `principalSet://goog/group/{group_id}`: A Google group. For example, `principalSet://goog/group/admins@example.com`. * `principalSet://goog/public:all`: A special identifier that represents any principal that is on the internet, even if they do not have a Google Account or are not logged in. * `principalSet://goog/cloudIdentityCustomerId/{customer_id}`: All of the principals associated with the specified Google Workspace or Cloud Identity customer ID. For example, `principalSet://goog/cloudIdentityCustomerId/C01Abc35`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:principal://goog/subject/{email_id}?uid={uid}`: A specific Google Account that was deleted recently. For example, `deleted:principal://goog/subject/alice@example.com?uid=1234567890`. If the Google Account is recovered, this identifier reverts to the standard identifier for a Google Account. * `deleted:principalSet://goog/group/{group_id}?uid={uid}`: A Google group that was deleted recently. For example, `deleted:principalSet://goog/group/admins@example.com?uid=1234567890`. If the Google group is restored, this identifier reverts to the standard identifier for a Google group. * `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}?uid={uid}`: A Google Cloud service account that was deleted recently. For example, `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com?uid=1234567890`. If the service account is undeleted, this identifier reverts to the standard identifier for a service account. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
	Exceptionpermissions []string `json:"exceptionPermissions,omitempty"` // Specifies the permissions that this rule excludes from the set of denied permissions given by `denied_permissions`. If a permission appears in `denied_permissions` _and_ in `exception_permissions` then it will _not_ be denied. The excluded permissions can be specified using the same syntax as `denied_permissions`.
}

// GoogleIamV2PolicyRule represents the GoogleIamV2PolicyRule schema from the OpenAPI specification
type GoogleIamV2PolicyRule struct {
	Denyrule GoogleIamV2DenyRule `json:"denyRule,omitempty"` // A deny rule in an IAM deny policy.
	Description string `json:"description,omitempty"` // A user-specified description of the rule. This value can be up to 256 characters.
}

// GoogleTypeExpr represents the GoogleTypeExpr schema from the OpenAPI specification
type GoogleTypeExpr struct {
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
}

// GoogleCloudCommonOperationMetadata represents the GoogleCloudCommonOperationMetadata schema from the OpenAPI specification
type GoogleCloudCommonOperationMetadata struct {
	Target string `json:"target,omitempty"` // Output only. Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Output only. Name of the verb executed by the operation.
	Apiversion string `json:"apiVersion,omitempty"` // Output only. API version used to start the operation.
	Cancelrequested bool `json:"cancelRequested,omitempty"` // Output only. Identifies whether the user has requested cancellation of the operation. Operations that have been cancelled successfully have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
	Createtime string `json:"createTime,omitempty"` // Output only. The time the operation was created.
	Endtime string `json:"endTime,omitempty"` // Output only. The time the operation finished running.
	Statusdetail string `json:"statusDetail,omitempty"` // Output only. Human-readable status of the operation, if any.
}

// GoogleIamV1PolicyDelta represents the GoogleIamV1PolicyDelta schema from the OpenAPI specification
type GoogleIamV1PolicyDelta struct {
	Bindingdeltas []GoogleIamV1BindingDelta `json:"bindingDeltas,omitempty"` // The delta for Bindings between two policies.
}

// GoogleIamV1LoggingAuditData represents the GoogleIamV1LoggingAuditData schema from the OpenAPI specification
type GoogleIamV1LoggingAuditData struct {
	Policydelta GoogleIamV1PolicyDelta `json:"policyDelta,omitempty"` // The difference delta between two policies.
}

// GoogleIamV2Policy represents the GoogleIamV2Policy schema from the OpenAPI specification
type GoogleIamV2Policy struct {
	Name string `json:"name,omitempty"` // Immutable. The resource name of the `Policy`, which must be unique. Format: `policies/{attachment_point}/denypolicies/{policy_id}` The attachment point is identified by its URL-encoded full resource name, which means that the forward-slash character, `/`, must be written as `%2F`. For example, `policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-deny-policy`. For organizations and folders, use the numeric ID in the full resource name. For projects, requests can use the alphanumeric or the numeric ID. Responses always contain the numeric ID.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the `Policy` was last updated.
	Etag string `json:"etag,omitempty"` // An opaque tag that identifies the current version of the `Policy`. IAM uses this value to help manage concurrent updates, so they do not cause one update to be overwritten by another. If this field is present in a CreatePolicyRequest, the value is ignored.
	Rules []GoogleIamV2PolicyRule `json:"rules,omitempty"` // A list of rules that specify the behavior of the `Policy`. All of the rules should be of the `kind` specified in the `Policy`.
	Uid string `json:"uid,omitempty"` // Immutable. The globally unique ID of the `Policy`. Assigned automatically when the `Policy` is created.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // A key-value map to store arbitrary metadata for the `Policy`. Keys can be up to 63 characters. Values can be up to 255 characters.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the `Policy` was created.
	Displayname string `json:"displayName,omitempty"` // A user-specified description of the `Policy`. This value can be up to 63 characters.
	Kind string `json:"kind,omitempty"` // Output only. The kind of the `Policy`. Always contains the value `DenyPolicy`.
	Deletetime string `json:"deleteTime,omitempty"` // Output only. The time when the `Policy` was deleted. Empty if the policy is not deleted.
}

// GoogleRpcStatus represents the GoogleRpcStatus schema from the OpenAPI specification
type GoogleRpcStatus struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
}

// GoogleIamAdminV1AuditData represents the GoogleIamAdminV1AuditData schema from the OpenAPI specification
type GoogleIamAdminV1AuditData struct {
	Permissiondelta GoogleIamAdminV1AuditDataPermissionDelta `json:"permissionDelta,omitempty"` // A PermissionDelta message to record the added_permissions and removed_permissions inside a role.
}

// GoogleIamV1betaWorkloadIdentityPoolOperationMetadata represents the GoogleIamV1betaWorkloadIdentityPoolOperationMetadata schema from the OpenAPI specification
type GoogleIamV1betaWorkloadIdentityPoolOperationMetadata struct {
}

// GoogleLongrunningOperation represents the GoogleLongrunningOperation schema from the OpenAPI specification
type GoogleLongrunningOperation struct {
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField GoogleRpcStatus `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
}

// GoogleIamAdminV1AuditDataPermissionDelta represents the GoogleIamAdminV1AuditDataPermissionDelta schema from the OpenAPI specification
type GoogleIamAdminV1AuditDataPermissionDelta struct {
	Addedpermissions []string `json:"addedPermissions,omitempty"` // Added permissions.
	Removedpermissions []string `json:"removedPermissions,omitempty"` // Removed permissions.
}

// GoogleIamV2ListPoliciesResponse represents the GoogleIamV2ListPoliciesResponse schema from the OpenAPI specification
type GoogleIamV2ListPoliciesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A page token that you can use in a ListPoliciesRequest to retrieve the next page. If this field is omitted, there are no additional pages.
	Policies []GoogleIamV2Policy `json:"policies,omitempty"` // Metadata for the policies that are attached to the resource.
}
