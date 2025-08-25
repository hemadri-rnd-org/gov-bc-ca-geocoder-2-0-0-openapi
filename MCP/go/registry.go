package main

import (
	"github.com/geocoder-rest-api/mcp-server/config"
	"github.com/geocoder-rest-api/mcp-server/models"
	tools_occupants "github.com/geocoder-rest-api/mcp-server/tools/occupants"
	tools_sites "github.com/geocoder-rest-api/mcp-server/tools/sites"
	tools_intersections "github.com/geocoder-rest-api/mcp-server/tools/intersections"
	tools_parcels "github.com/geocoder-rest-api/mcp-server/tools/parcels"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_occupants.CreateGet_occupants_addresses_outputformatTool(cfg),
		tools_sites.CreateGet_sites_siteid_outputformatTool(cfg),
		tools_occupants.CreateGet_occupants_nearest_outputformatTool(cfg),
		tools_sites.CreateGet_sites_within_outputformatTool(cfg),
		tools_sites.CreateGet_sites_near_outputformatTool(cfg),
		tools_sites.CreateGet_sites_nearest_outputformatTool(cfg),
		tools_sites.CreateGet_addresses_outputformatTool(cfg),
		tools_intersections.CreateGet_intersections_intersectionid_outputformatTool(cfg),
		tools_intersections.CreateGet_intersections_within_outputformatTool(cfg),
		tools_parcels.CreateGet_parcels_pids_siteid_outputformatTool(cfg),
		tools_occupants.CreateGet_occupants_near_outputformatTool(cfg),
		tools_occupants.CreateGet_occupants_occupantid_outputformatTool(cfg),
		tools_sites.CreateGet_sites_siteid_subsites_outputformatTool(cfg),
		tools_occupants.CreateGet_occupants_within_outputformatTool(cfg),
		tools_intersections.CreateGet_intersections_near_outputformatTool(cfg),
		tools_intersections.CreateGet_intersections_nearest_outputformatTool(cfg),
	}
}
