package handlers

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/services/orchestrator"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (_ *Server) SendCrossCheckReport(_ context.Context, request *orchestrator.SendCrossCheckReportRequest) (_ *emptypb.Empty, _ error) {
	var err error

	var adoption1 = &db.Adoption{
		WorkID: request.FirstWorkID,
	}

	var adoption2 = &db.Adoption{
		WorkID: request.SecondWorkID,
	}

	for i, _ := range request.Match {
		adoption1.Path = &request.Match[i].FirstWorkPath
		adoption1.PartOffset = &request.Match[i].FirstWorkStart
		adoption1.PartSize = &request.Match[i].FirstWorkSize

		adoption2.Path = &request.Match[i].SecondWorkPath
		adoption2.PartOffset = &request.Match[i].SecondWorkStart
		adoption2.PartSize = &request.Match[i].SecondWorkSize

		if err = db.CreateAdoption(adoption1); err != nil {
			log.Println("Error saving cross-check report:", err.Error())
			continue
		}

		adoption2.RefersTo = &adoption1.ID
		if err = db.CreateAdoption(adoption2); err != nil {
			log.Println("Error saving cross-check report:", err.Error())
			continue
		}
	}

	return
}

func (_ *Server) SendDefaultReport(_ context.Context, request *orchestrator.SendDefaultReportRequest) (_ *emptypb.Empty, _ error) {
	var err error

	var adoption = &db.Adoption{
		WorkID: request.WorkID,
	}

	for i, _ := range request.Segment {
		adoption.Path = &request.Segment[i].WorkPath
		adoption.PartOffset = &request.Segment[i].WorkStart
		adoption.PartSize = &request.Segment[i].WorkSize
		adoption.SimilarityScore = request.Segment[i].Accuracy

		if err = db.CreateAdoption(adoption); err != nil {
			log.Println("Error saving default report:", err.Error())
			continue
		}
	}

	return
}
