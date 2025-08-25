package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/geocoder-rest-api/mcp-server/config"
	"github.com/geocoder-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_parcels_pids_siteid_outputformatHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		siteIDVal, ok := args["siteID"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: siteID"), nil
		}
		siteID, ok := siteIDVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: siteID"), nil
		}
		outputFormatVal, ok := args["outputFormat"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: outputFormat"), nil
		}
		outputFormat, ok := outputFormatVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: outputFormat"), nil
		}
		url := fmt.Sprintf("%s/parcels/pids/%s.%s", cfg.BaseURL, siteID, outputFormat)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("apikey", cfg.APIKey)
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
		var result map[string]interface{}
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

func CreateGet_parcels_pids_siteid_outputformatTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_parcels_pids_siteID_outputFormat",
		mcp.WithDescription("Get a comma-separated string of all pids for a given site"),
		mcp.WithString("siteID", mcp.Required(), mcp.Description("A unique, but not immutable, site identifier.")),
		mcp.WithString("outputFormat", mcp.Required(), mcp.Description("Results format. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#outputFormat target=\"_blank\">outputFormat</a>. \n\nNote: GeoJSON and KML formats only support EPSG:4326 (outputSRS=4326)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_parcels_pids_siteid_outputformatHandler(cfg),
	}
}
