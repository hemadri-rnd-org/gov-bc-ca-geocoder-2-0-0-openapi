package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/geocoder-rest-api/mcp-server/config"
	"github.com/geocoder-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_occupants_near_outputformatHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		outputFormatVal, ok := args["outputFormat"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: outputFormat"), nil
		}
		outputFormat, ok := outputFormatVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: outputFormat"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["point"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("point=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["maxDistance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxDistance=%v", val))
		}
		if val, ok := args["outputSRS"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outputSRS=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["locationDescriptor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locationDescriptor=%v", val))
		}
		if val, ok := args["brief"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("brief=%v", val))
		}
		if val, ok := args["setBack"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("setBack=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/occupants/near.%s%s", cfg.BaseURL, outputFormat, queryString)
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

func CreateGet_occupants_near_outputformatTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_occupants_near_outputFormat",
		mcp.WithDescription("Find occupants of sites near to a geographic point"),
		mcp.WithString("outputFormat", mcp.Required(), mcp.Description("Results format. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#outputFormat target=\"_blank\">outputFormat</a>. \n\nNote: GeoJSON and KML formats only support EPSG:4326 (outputSRS=4326)")),
		mcp.WithString("point", mcp.Required(), mcp.Description("The point (x,y) from which the nearest site will be identified. The coordinates must be specified in the same SRS as given by the 'outputSRS' parameter.")),
		mcp.WithString("tags", mcp.Description("Example: schools;courts;employment<br>A list of tags separated by semicolons.")),
		mcp.WithNumber("maxDistance", mcp.Description("The maximum distance (in metres) to search from the given point.  If not specified, the search distance is unlimited.")),
		mcp.WithNumber("outputSRS", mcp.Description("The EPSG code of the spatial reference system (SRS) to use for output geometries. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#outputSRS target=\"_blank\">outputSRS</a>")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of search results to return.")),
		mcp.WithString("locationDescriptor", mcp.Description("Describes the nature of the address location. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#locationDescriptor target=\"_blank\">locationDescriptor</a>")),
		mcp.WithBoolean("brief", mcp.Description("If true, include only basic match and address details in results. Not supported for shp, csv, and gml formats.")),
		mcp.WithNumber("setBack", mcp.Description("The distance to move the accessPoint away from the curb and towards the inside of the parcel (in metres). Ignored if locationDescriptor not set to accessPoint.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_occupants_near_outputformatHandler(cfg),
	}
}
