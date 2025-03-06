package handlers

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/pkg/s3storage"
	"SparkGuardBackend/services/orchestrator"
	"context"
	"fmt"
)

func (_ *Server) GetWorksOfEvent(_ context.Context, request *orchestrator.GetWorksOfEventRequest) (*orchestrator.GetWorksOfEventResponse, error) {
	works, err := db.GetWorksIdOfEvent(request.EventID)

	if err != nil {
		return nil, err
	}

	response := &orchestrator.GetWorksOfEventResponse{
		WorkID: works,
	}

	return response, nil
}

func (_ *Server) GetWorksDownloadLinks(_ context.Context, request *orchestrator.GetWorksDownloadLinksRequest) (response *orchestrator.GetWorksDownloadLinksResponse, _ error) {
	response = &orchestrator.GetWorksDownloadLinksResponse{
		Item: make([]*orchestrator.GetWorksDownloadLinksResponseItem, len(request.WorkID)),
	}

	for ind, id := range request.WorkID {
		link, _ := s3storage.ShareFile(fmt.Sprintf("%d.zip", id))

		response.Item[ind] = &orchestrator.GetWorksDownloadLinksResponseItem{
			WorkID:       id,
			DownloadLink: link,
		}
	}

	return response, nil
}
