package booky

import (
	pb "booky-back/api/booky"
	"booky-back/internal/models"
	"booky-back/internal/storage"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "UpdateNote: note id is required")
	}

	note, err := s.Storage.GetNote(req.Id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "UpdateNote: note not found")
		}
		return nil, status.Errorf(codes.Internal, "UpdateNote: could not get note: %v", err)
	}

	err = note.BindUpdateNote(req.GetData())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "UpdateNote: could not bind note: %v", err)
	}

	user, err := s.Storage.GetUser(note.Publisher.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "CreateNote: could not get user: %v", err)
	}

	note.Publisher = *user

	updatedNote, err := s.Storage.UpdateNote(note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "UpdateNote: could not update note: %v", err)
	}

	grpcNote, err := models.BindNoteToGRPC(updatedNote)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "UpdateNote: could not bind note to grpc: %v", err)
	}

	return &pb.UpdateNoteResponse{
		Note: grpcNote,
	}, nil
}
