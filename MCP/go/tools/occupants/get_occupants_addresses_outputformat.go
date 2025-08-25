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

func Get_occupants_addresses_outputformatHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["addressString"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addressString=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["locationDescriptor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locationDescriptor=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["interpolation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("interpolation=%v", val))
		}
		if val, ok := args["echo"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("echo=%v", val))
		}
		if val, ok := args["brief"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("brief=%v", val))
		}
		if val, ok := args["autoComplete"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("autoComplete=%v", val))
		}
		if val, ok := args["setBack"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("setBack=%v", val))
		}
		if val, ok := args["outputSRS"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outputSRS=%v", val))
		}
		if val, ok := args["minScore"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minScore=%v", val))
		}
		if val, ok := args["matchPrecision"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("matchPrecision=%v", val))
		}
		if val, ok := args["matchPrecisionNot"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("matchPrecisionNot=%v", val))
		}
		if val, ok := args["siteName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteName=%v", val))
		}
		if val, ok := args["unitDesignator"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unitDesignator=%v", val))
		}
		if val, ok := args["unitNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unitNumber=%v", val))
		}
		if val, ok := args["unitNumberSuffix"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("unitNumberSuffix=%v", val))
		}
		if val, ok := args["civicNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("civicNumber=%v", val))
		}
		if val, ok := args["civicNumberSuffix"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("civicNumberSuffix=%v", val))
		}
		if val, ok := args["streetName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("streetName=%v", val))
		}
		if val, ok := args["streetType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("streetType=%v", val))
		}
		if val, ok := args["streetDirection"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("streetDirection=%v", val))
		}
		if val, ok := args["streetQualifier"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("streetQualifier=%v", val))
		}
		if val, ok := args["localityName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("localityName=%v", val))
		}
		if val, ok := args["provinceCode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("provinceCode=%v", val))
		}
		if val, ok := args["localities"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("localities=%v", val))
		}
		if val, ok := args["notLocalities"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("notLocalities=%v", val))
		}
		if val, ok := args["bbox"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bbox=%v", val))
		}
		if val, ok := args["centre"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("centre=%v", val))
		}
		if val, ok := args["maxDistance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxDistance=%v", val))
		}
		if val, ok := args["extrapolate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("extrapolate=%v", val))
		}
		if val, ok := args["parcelPoint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parcelPoint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/occupants/addresses.%s%s", cfg.BaseURL, outputFormat, queryString)
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

func CreateGet_occupants_addresses_outputformatTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_occupants_addresses_outputFormat",
		mcp.WithDescription("Geocode an address and identify site occupants"),
		mcp.WithString("outputFormat", mcp.Required(), mcp.Description("Results format. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#outputFormat target=\"_blank\">outputFormat</a>. \n\nNote: GeoJSON and KML formats only support EPSG:4326 (outputSRS=4326)")),
		mcp.WithString("addressString", mcp.Description("Occupant name OR Occupant name ** address")),
		mcp.WithString("tags", mcp.Description("Example: schools;courts;employment<br>A list of tags separated by semicolons.")),
		mcp.WithString("locationDescriptor", mcp.Description("Describes the nature of the address location. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#locationDescriptor target=\"_blank\">locationDescriptor</a>")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of search results to return.")),
		mcp.WithString("interpolation", mcp.Description("accessPoint interpolation method. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#interpolation target=\"_blank\">interpolation</a>")),
		mcp.WithBoolean("echo", mcp.Description("If true, include unmatched address details such as site name in results.")),
		mcp.WithBoolean("brief", mcp.Description("If true, include only basic match and address details in results. Not supported for shp, csv, and gml formats.")),
		mcp.WithBoolean("autoComplete", mcp.Description("If true, addressString is expected to contain a partial address that requires completion. Not supported for shp, csv, gml formats.")),
		mcp.WithNumber("setBack", mcp.Description("The distance to move the accessPoint away from the curb and towards the inside of the parcel (in metres). Ignored if locationDescriptor not set to accessPoint.")),
		mcp.WithNumber("outputSRS", mcp.Description("The EPSG code of the spatial reference system (SRS) to use for output geometries. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#outputSRS target=\"_blank\">outputSRS</a>")),
		mcp.WithNumber("minScore", mcp.Description("The minimum score required for a match to be returned. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#minScore target=\"_blank\">minScore</a>")),
		mcp.WithString("matchPrecision", mcp.Description("Example: street,locality.  A comma separated list of individual match precision levels to include in results. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#matchPrecision target=\"_blank\">matchPrecision</a>")),
		mcp.WithString("matchPrecisionNot", mcp.Description("Example: street,locality.  A comma separated list of individual match precision levels to exclude from results. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#matchPrecisionNot target=\"_blank\">matchPrecisionNot</a>")),
		mcp.WithString("siteName", mcp.Description("A string containing the name of the building, facility, or institution (e.g., Duck Building, Casa Del Mar, Crystal Garden, Bluebird House).See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#siteName target=\"_blank\">siteName</a>")),
		mcp.WithString("unitDesignator", mcp.Description("The type of unit within a house or building. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#unitDesignator target=\"_blank\">unitDesignator</a>")),
		mcp.WithString("unitNumber", mcp.Description("The number of the unit, suite, or apartment within a house or building.")),
		mcp.WithString("unitNumberSuffix", mcp.Description("A letter that follows the unit number as in Unit 1A or Suite 302B.")),
		mcp.WithString("civicNumber", mcp.Description("The official number assigned to a site by an address authority. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#civicNumber target=\"_blank\">civicNumber</a>")),
		mcp.WithString("civicNumberSuffix", mcp.Description("A letter or fraction that follows the civic number (e.g., the A in 1039A Bledsoe St). See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#civicNumberSuffix target=\"_blank\">civicNumberSuffix</a>")),
		mcp.WithString("streetName", mcp.Description("The official name of the street as assigned by an address authority (e.g., the Douglas in 1175 Douglas Street). See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#streetName target=\"_blank\">streetName</a>")),
		mcp.WithString("streetType", mcp.Description("The type of street as assigned by a municipality (e.g., the ST in 1175 DOUGLAS St). See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#streetType target=\"_blank\">streetType</a>")),
		mcp.WithString("streetDirection", mcp.Description("The abbreviated compass direction as defined by Canada Post and B.C. civic addressing authorities. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#streetDirection target=\"_blank\">streetDirection</a>")),
		mcp.WithString("streetQualifier", mcp.Description("The qualifier of a street name (e.g., the Bridge in Johnson St Bridge)")),
		mcp.WithString("localityName", mcp.Description("The name of the locality assigned to a given site by an address authority. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#streetDirection target=\"_blank\">streetDirection</a>")),
		mcp.WithString("provinceCode", mcp.Description("The ISO 3166-2 Sub-Country Code. The code for British Columbia is BC.")),
		mcp.WithString("localities", mcp.Description("A comma separated list of locality names that matched addresses must belong to. For example, setting localities to Nanaimo only returns addresses in Nanaimo")),
		mcp.WithString("notLocalities", mcp.Description("A comma-separated list of localities to exclude from the search.")),
		mcp.WithString("bbox", mcp.Description("Example: -126.07929,49.7628,-126.0163,49.7907.  A bounding box (xmin,ymin,xmax,ymax) that limits the search area. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#bbox target=\"_blank\">bbox</a>")),
		mcp.WithString("centre", mcp.Description("Example: -124.0165926,49.2296251 .  The coordinates of a centre point (x,y) used to define a bounding circle that will limit the search area. This parameter must be specified together with 'maxDistance'. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#centre target='_blank'>centre</a>")),
		mcp.WithString("maxDistance", mcp.Description("The maximum distance (in metres) to search from the given point.  If not specified, the search distance is unlimited.")),
		mcp.WithBoolean("extrapolate", mcp.Description("If true, uses supplied parcelPoint to derive an appropriate accessPoint.           See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#extrapolate target=\"_blank\">extrapolate</a>")),
		mcp.WithString("parcelPoint", mcp.Description("The coordinates of a point (x,y) known to be inside the parcel containing a given address. See <a href=https://github.com/bcgov/ols-geocoder/blob/gh-pages/glossary.md#parcelPoint target=\"_blank\">parcelPoint</a>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_occupants_addresses_outputformatHandler(cfg),
	}
}
