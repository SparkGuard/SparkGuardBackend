package handlers

import "SparkGuardBackend/services/orchestrator"

type Server struct {
	orchestrator.UnimplementedOrchestratorServer
}
