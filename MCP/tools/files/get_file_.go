package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_file_Handler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		pathVal, ok := args["path"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: path"), nil
		}
		path, ok := pathVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: path"), nil
		}
		filepathVal, ok := args["filepath"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: filepath"), nil
		}
		filepath, ok := filepathVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: filepath"), nil
		}
		url := fmt.Sprintf("%s/file/*", cfg.BaseURL, path, filepath)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
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
		var result interface{}
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

func CreateGet_file_Tool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_file_*",
		mcp.WithDescription("Serve file content"),
		mcp.WithString("path", mcp.Required(), mcp.Description("File path")),
		mcp.WithString("filepath", mcp.Required(), mcp.Description("Path to the file")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_file_Handler(cfg),
	}
}
