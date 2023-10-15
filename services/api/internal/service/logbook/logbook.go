package logbook

import (
	"api/internal/db"
	"api/pkg/session"
	"api/pkg/validator"
	pb "api/proto"
	"context"
	"database/sql"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogbookServiceServer struct {
	db db.Database
	*pb.UnimplementedLogbookServiceServer
}

func NewLogbookServer(db *db.Database) *LogbookServiceServer {
	return &LogbookServiceServer{
		db: *db,
	}
}

func (s *LogbookServiceServer) Create(ctx context.Context, req *pb.CreateLogbookRequest) (*pb.LogbookResponseMessage, error) {
	session := session.GetSession(ctx)
	var userIdDb int64
	err := s.db.Conn.Get(&userIdDb, "SELECT id FROM users WHERE id=?", session.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	if err := validator.ValidateLogbook(req.GetLogs()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid logbook: %v", err)
	}

	// create logbook
	res, err := s.db.Conn.Exec("INSERT INTO logbook (id_user, logs, log_date) VALUES (?, ?, ?, ?)",
		session.ID, req.GetLogs(), req.GetLogDate())
	if err != nil {
		log.Printf("Error insert new logbook: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	logbookId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error get last insert id: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	// return message response
	return &pb.LogbookResponseMessage{
		IdLogbook: logbookId,
		Message:   "success create logbook",
	}, nil
}
func (s *LogbookServiceServer) List(ctx context.Context, req *pb.ListLogbookRequest) (*pb.ListLogbookResponse, error) {
	// check id user exists
	session := session.GetSession(ctx)
	var userIdDb int64
	err := s.db.Conn.Get(&userIdDb, "SELECT id FROM users WHERE id=?", session.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("error get user id: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	type Logbook struct {
		ID      int64
		Logs    string
		LogDate string `db:"log_date"`
	}
	logbooks := []Logbook{}
	parsed_start_date, err := time.Parse("2006-01-02", req.GetStartDate())
	if err != nil {
		log.Printf("Error parse start date: %v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "start_date error: %v\n", err)
	}
	parsed_end_date, err := time.Parse("2006-01-02", req.GetEndDate())
	if err != nil {
		log.Printf("Error parse end date: %v\n", err)
		return nil, status.Errorf(codes.InvalidArgument, "end_date error: %v\n", err)
	}
	if err := s.db.Conn.Select(&logbooks,
		"SELECT id, logs, log_date FROM logbook WHERE id_user=? AND log_date between ? AND ?",
		userIdDb,
		parsed_start_date,
		parsed_end_date,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "logbook not found")
		}
		log.Printf("Error get logbook: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	// return list of logbooks
	var logbooksResponse []*pb.LogbookResponse
	for _, logbook := range logbooks {
		logbooksResponse = append(logbooksResponse, &pb.LogbookResponse{
			IdLogbook: logbook.ID,
			Logs:      logbook.Logs,
			LogDate:   logbook.LogDate,
		})
	}
	return &pb.ListLogbookResponse{
		StartDate:    req.GetStartDate(),
		EndDate:      req.GetEndDate(),
		TotalRecords: int32(len(logbooksResponse)),
		Logbooks:     logbooksResponse,
	}, nil
}
func (s *LogbookServiceServer) Update(ctx context.Context, req *pb.UpdateLogbookRequest) (*pb.LogbookResponseMessage, error) {
	// some validation
	// id logbook exists
	// logbook min length 10, and max length 100
	logbookId := req.GetIdLogbook()
	var logbookIdDb int64
	err := s.db.Conn.Get(&logbookIdDb, "SELECT id FROM logbook WHERE id=?", logbookId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "logbook not found")
		}
		log.Printf("Error get logbook with id %d: %v\n", logbookId, err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	if err := validator.ValidateLogbook(req.GetLogs()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid logbook: %v", err)
	}

	// do update
	_, err = s.db.Conn.Exec("UPDATE logbook SET logs=?, log_date=? WHERE id=?", req.GetLogs(), req.GetLogDate(), req.GetIdLogbook())
	if err != nil {
		log.Printf("Error update logbook: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	// return message response
	return &pb.LogbookResponseMessage{
		IdLogbook: logbookIdDb,
		Message:   "success update logbook",
	}, nil
}
func (s *LogbookServiceServer) Delete(ctx context.Context, req *pb.DeleteLogbookRequest) (*pb.LogbookResponseMessage, error) {
	// id logbook exists
	logbookId := req.GetIdLogbook()
	var logbookIdDb int64
	err := s.db.Conn.Get(&logbookIdDb, "SELECT id FROM logbook WHERE id=?", logbookId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "logbook not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	// delete logbook
	_, err = s.db.Conn.Exec("DELETE FROM logbook WHERE id=?", logbookId)
	if err != nil {
		log.Printf("Error delete logbook: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	// return message response
	return &pb.LogbookResponseMessage{
		IdLogbook: logbookIdDb,
		Message:   "success delete logbook",
	}, nil
}
