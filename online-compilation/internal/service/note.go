package service

import (
	"context"
	"fmt"
	"online-compilation/internal/biz"

	pb "online-compilation/api/note"
)

type NoteService struct {
	pb.UnimplementedNoteServer
	uc *biz.NoteUsecase
}

func NewNoteService(uc *biz.NoteUsecase) *NoteService {
	return &NoteService{
		uc: uc,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteReply, error) {
	test, err := s.uc.CreateNote(ctx, req)
	fmt.Println(test)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNoteReply{}, nil
}
func (s *NoteService) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteReply, error) {
	return &pb.UpdateNoteReply{}, nil
}
func (s *NoteService) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteReply, error) {
	return &pb.DeleteNoteReply{}, nil
}
func (s *NoteService) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteReply, error) {
	return &pb.GetNoteReply{}, nil
}
func (s *NoteService) ListNote(ctx context.Context, req *pb.ListNoteRequest) (*pb.ListNoteReply, error) {
	return &pb.ListNoteReply{}, nil
}

func (s *NoteService) Code(ctx context.Context, req *pb.CodeInfoRequest) (*pb.CodeInfoReply, error) {
	return &pb.CodeInfoReply{}, nil
}

func (s *NoteService) NoteDownload(ctx context.Context, req *pb.NoteDownloadRequest) (*pb.NoteDownloadReply, error) {
	return &pb.NoteDownloadReply{}, nil
}

func (s *NoteService) CodeCompile(ctx context.Context, req *pb.CodeCompileRequest) (*pb.CodeCompileReply, error) {
	return &pb.CodeCompileReply{}, nil
}
