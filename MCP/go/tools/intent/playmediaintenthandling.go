package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/sirikit-cloud-media/mcp-server/config"
	"github.com/sirikit-cloud-media/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PlaymediaintenthandlingHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody []PlayMediaIntentHandlingInvocation
		
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
		url := fmt.Sprintf("%s/intent/playMedia", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-applecloudextension-session-id"]; ok {
			req.Header.Set("x-applecloudextension-session-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-applecloudextension-retry-count"]; ok {
			req.Header.Set("x-applecloudextension-retry-count", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Request-Timeout"]; ok {
			req.Header.Set("Request-Timeout", fmt.Sprintf("%v", val))
		}
		if val, ok := args["User-Agent"]; ok {
			req.Header.Set("User-Agent", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Accept-Language"]; ok {
			req.Header.Set("Accept-Language", fmt.Sprintf("%v", val))
		}

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
		var result []interface{}
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

func CreatePlaymediaintenthandlingTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_intent_playMedia",
		mcp.WithDescription("playMedia"),
		mcp.WithString("x-applecloudextension-session-id", mcp.Required(), mcp.Description("")),
		mcp.WithString("x-applecloudextension-retry-count", mcp.Description("")),
		mcp.WithString("Request-Timeout", mcp.Required(), mcp.Description("")),
		mcp.WithString("User-Agent", mcp.Required(), mcp.Description("")),
		mcp.WithString("Accept-Language", mcp.Required(), mcp.Description("")),
		mcp.WithArray("items", mcp.Required(), mcp.Description("Array of objects")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PlaymediaintenthandlingHandler(cfg),
	}
}
