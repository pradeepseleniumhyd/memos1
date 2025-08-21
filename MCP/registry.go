package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_profiling "github.com/input-api/mcp-server/tools/profiling"
	tools_errors "github.com/input-api/mcp-server/tools/errors"
	tools_headers "github.com/input-api/mcp-server/tools/headers"
	tools_proxy "github.com/input-api/mcp-server/tools/proxy"
	tools_explore "github.com/input-api/mcp-server/tools/explore"
	tools_pprof "github.com/input-api/mcp-server/tools/pprof"
	tools_health "github.com/input-api/mcp-server/tools/health"
	tools_metrics "github.com/input-api/mcp-server/tools/metrics"
	tools_rss "github.com/input-api/mcp-server/tools/rss"
	tools_files "github.com/input-api/mcp-server/tools/files"
	tools_oauth2 "github.com/input-api/mcp-server/tools/oauth2"
	tools_memos "github.com/input-api/mcp-server/tools/memos"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_profiling.CreateGetTool(cfg),
		tools_errors.CreateGet_errorTool(cfg),
		tools_headers.CreateHead_content_typeTool(cfg),
		tools_profiling.CreateGet_heapTool(cfg),
		tools_profiling.CreateGet_profileTool(cfg),
		tools_profiling.CreateGet_symbolTool(cfg),
		tools_profiling.CreatePost_symbolTool(cfg),
		tools_profiling.CreateGet_threadcreateTool(cfg),
		tools_profiling.CreateGet_goroutineTool(cfg),
		tools_profiling.CreateGet_traceTool(cfg),
		tools_proxy.CreateGet_api_v1_Tool(cfg),
		tools_explore.CreateGet_explore_rss_xmlTool(cfg),
		tools_pprof.CreateGet_cmdlineTool(cfg),
		tools_errors.CreateGet_errTool(cfg),
		tools_pprof.CreateGet_blockTool(cfg),
		tools_health.CreateGet_healthzTool(cfg),
		tools_metrics.CreateGet_memstatsTool(cfg),
		tools_rss.CreateGet_u_username_rss_xmlTool(cfg),
		tools_files.CreateGet_file_Tool(cfg),
		tools_oauth2.CreateGet_oauth2_userinfoTool(cfg),
		tools_profiling.CreateGet_allocsTool(cfg),
		tools_oauth2.CreateGet_oauth2_tokenTool(cfg),
		tools_memos.CreateGet_memos_api_v1_Tool(cfg),
		tools_profiling.CreateGet_mutexTool(cfg),
	}
}
