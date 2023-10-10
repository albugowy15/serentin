package common

import (
	"api/internal/db"
	pb "api/proto"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommonServiceServer struct {
	db db.Database
	*pb.UnimplementedCommonServiceServer
}

func NewCommonServer(db *db.Database) *CommonServiceServer {
	return &CommonServiceServer{
		db: *db,
	}
}

func (s *CommonServiceServer) ListPersonalities(context.Context, *pb.ListPersonalitiesRequest) (*pb.ListPersonalitiesResponse, error) {
	type Personalities struct {
		ID          int32
		Personality string
		Description string
	}
	var personalities []Personalities
	err := s.db.Conn.Select(&personalities, "SELECT id, personality, description FROM mbti")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "no personalities found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	var personalitiesResponse []*pb.Personality
	for _, personality := range personalities {
		personalitiesResponse = append(personalitiesResponse, &pb.Personality{
			IdPersonality: personality.ID,
			Personality:   personality.Personality,
			Description:   personality.Description,
		})
	}

	return &pb.ListPersonalitiesResponse{
		Personalities: personalitiesResponse,
	}, nil
}
func (s *CommonServiceServer) ListJobPositions(context.Context, *pb.ListJobPositionsRequest) (*pb.ListJobPositionsResponse, error) {
	type JobPosition struct {
		ID          int32
		Position    string
		Description string
	}
	var jobPositions []JobPosition
	err := s.db.Conn.Select(&jobPositions, "SELECT id, position, description FROM job_positions")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "no positions found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	var jobPositionsResponse []*pb.JobPosition
	for _, jobPosition := range jobPositions {
		jobPositionsResponse = append(jobPositionsResponse, &pb.JobPosition{
			IdJobPosition: jobPosition.ID,
			JobPosition:   jobPosition.Position,
			Description:   jobPosition.Description,
		})
	}

	return &pb.ListJobPositionsResponse{
		JobPositions: jobPositionsResponse,
	}, nil
}
