package main

import (
	"github.com/identity-and-access-management-iam-api/mcp-server/config"
	"github.com/identity-and-access-management-iam-api/mcp-server/models"
	tools_policies "github.com/identity-and-access-management-iam-api/mcp-server/tools/policies"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_policies.CreateIam_policies_listpoliciesTool(cfg),
		tools_policies.CreateIam_policies_createpolicyTool(cfg),
		tools_policies.CreateIam_policies_operations_getTool(cfg),
		tools_policies.CreateIam_policies_updateTool(cfg),
		tools_policies.CreateIam_policies_deleteTool(cfg),
	}
}
