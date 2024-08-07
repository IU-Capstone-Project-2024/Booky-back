package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/pkg/storage"
	"booky-back/internal/pkg/validator"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	id := req.GetId()
	v := validator.New()
	if details, err := v.ValidateID(id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "GetCourse: validation error: invalid id")
	} else if len(details) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "GetCourse: validation error: %v", details)
	}

	err := s.Storage.DeleteNote(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "DeleteNote: note not found")
		}
		return nil, status.Errorf(codes.Internal, "DeleteNote: could not delete note: %v", err)
	}

	return &pb.DeleteNoteResponse{}, nil
}
