package physicaldata

import (
	"api/internal/db"
	pb "api/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecordPhysicalDataServiceServer struct {
	db db.Database
	*pb.UnimplementedRecordPhysicalDataServiceServer
}

func NewRecordPhysicalDataServer(db *db.Database) *RecordPhysicalDataServiceServer {
	return &RecordPhysicalDataServiceServer{
		db: *db,
	}
}

func (s *RecordPhysicalDataServiceServer) Save(ctx context.Context, req *pb.SavePhysicalDataRequest) (*pb.SavePhysicalDataResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (s *RecordPhysicalDataServiceServer) List(ctx context.Context, req *pb.ListPhysicalDataRequest) (*pb.ListPhysicalDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
