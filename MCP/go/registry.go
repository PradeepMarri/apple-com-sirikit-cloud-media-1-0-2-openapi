package main

import (
	"github.com/sirikit-cloud-media/mcp-server/config"
	"github.com/sirikit-cloud-media/mcp-server/models"
	tools_intent "github.com/sirikit-cloud-media/mcp-server/tools/intent"
	tools_queues "github.com/sirikit-cloud-media/mcp-server/tools/queues"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_intent.CreateAddmediaintenthandlingTool(cfg),
		tools_intent.CreatePlaymediaintenthandlingTool(cfg),
		tools_intent.CreateUpdatemediaaffinityintenthandlingTool(cfg),
		tools_queues.CreatePlaymediaonqueueTool(cfg),
		tools_queues.CreateUpdateactivityonqueueTool(cfg),
	}
}
