package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/identity-and-access-management-iam-api/mcp-server/config"
	"github.com/identity-and-access-management-iam-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Iam_policies_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["$.xgafv"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$.xgafv=%v", val))
		}
		if val, ok := args["access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%v", val))
		}
		if val, ok := args["alt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("alt=%v", val))
		}
		if val, ok := args["callback"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("callback=%v", val))
		}
		if val, ok := args["fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields=%v", val))
		}
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		if val, ok := args["oauth_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%v", val))
		}
		if val, ok := args["prettyPrint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("prettyPrint=%v", val))
		}
		if val, ok := args["quotaUser"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("quotaUser=%v", val))
		}
		if val, ok := args["upload_protocol"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("upload_protocol=%v", val))
		}
		if val, ok := args["uploadType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("uploadType=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.GoogleIamV2Policy
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v2/%s%s", cfg.BaseURL, name, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GoogleLongrunningOperation
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateIam_policies_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_v2_name",
		mcp.WithDescription("Updates the specified policy. You can update only the rules and the display name for the policy. To update a policy, you should use a read-modify-write loop: 1. Use GetPolicy to read the current version of the policy. 2. Modify the policy as needed. 3. Use `UpdatePolicy` to write the updated policy. This pattern helps prevent conflicts between concurrent updates."),
		mcp.WithString("$.xgafv", mcp.Description("V1 error format.")),
		mcp.WithString("access_token", mcp.Description("OAuth access token.")),
		mcp.WithString("alt", mcp.Description("Data format for response.")),
		mcp.WithString("callback", mcp.Description("JSONP")),
		mcp.WithString("fields", mcp.Description("Selector specifying which fields to include in a partial response.")),
		mcp.WithString("key", mcp.Description("API key. Your API key identifies your project and provides you with API access, quota, and reports. Required unless you provide an OAuth 2.0 token.")),
		mcp.WithString("oauth_token", mcp.Description("OAuth 2.0 token for the current user.")),
		mcp.WithBoolean("prettyPrint", mcp.Description("Returns response with indentations and line breaks.")),
		mcp.WithString("quotaUser", mcp.Description("Available to use for quota purposes for server-side applications. Can be any arbitrary string assigned to a user, but should not exceed 40 characters.")),
		mcp.WithString("upload_protocol", mcp.Description("Upload protocol for media (e.g. \"raw\", \"multipart\").")),
		mcp.WithString("uploadType", mcp.Description("Legacy upload protocol for media (e.g. \"media\", \"multipart\").")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Immutable. The resource name of the `Policy`, which must be unique. Format: `policies/{attachment_point}/denypolicies/{policy_id}` The attachment point is identified by its URL-encoded full resource name, which means that the forward-slash character, `/`, must be written as `%2F`. For example, `policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-deny-policy`. For organizations and folders, use the numeric ID in the full resource name. For projects, requests can use the alphanumeric or the numeric ID. Responses always contain the numeric ID.")),
		mcp.WithString("etag", mcp.Description("Input parameter: An opaque tag that identifies the current version of the `Policy`. IAM uses this value to help manage concurrent updates, so they do not cause one update to be overwritten by another. If this field is present in a CreatePolicyRequest, the value is ignored.")),
		mcp.WithArray("rules", mcp.Description("Input parameter: A list of rules that specify the behavior of the `Policy`. All of the rules should be of the `kind` specified in the `Policy`.")),
		mcp.WithString("uid", mcp.Description("Input parameter: Immutable. The globally unique ID of the `Policy`. Assigned automatically when the `Policy` is created.")),
		mcp.WithObject("annotations", mcp.Description("Input parameter: A key-value map to store arbitrary metadata for the `Policy`. Keys can be up to 63 characters. Values can be up to 255 characters.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time when the `Policy` was created.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: A user-specified description of the `Policy`. This value can be up to 63 characters.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Output only. The kind of the `Policy`. Always contains the value `DenyPolicy`.")),
		mcp.WithString("deleteTime", mcp.Description("Input parameter: Output only. The time when the `Policy` was deleted. Empty if the policy is not deleted.")),
		mcp.WithString("name", mcp.Description("Input parameter: Immutable. The resource name of the `Policy`, which must be unique. Format: `policies/{attachment_point}/denypolicies/{policy_id}` The attachment point is identified by its URL-encoded full resource name, which means that the forward-slash character, `/`, must be written as `%2F`. For example, `policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-deny-policy`. For organizations and folders, use the numeric ID in the full resource name. For projects, requests can use the alphanumeric or the numeric ID. Responses always contain the numeric ID.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time when the `Policy` was last updated.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Iam_policies_updateHandler(cfg),
	}
}
