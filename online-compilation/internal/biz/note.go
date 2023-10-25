package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "online-compilation/api/note"
)

type NoteUsecase struct {
	repo NoteRepo
	log  *log.Helper
}

func (u NoteUsecase) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (string, error) {
	u.repo.CreateNote(ctx)
	return "你好", nil
}

func NewNoteUsecase(repo NoteRepo, logger log.Logger) *NoteUsecase {
	return &NoteUsecase{repo: repo, log: log.NewHelper(logger)}
}

type NoteRepo interface {
	CreateNote(ctx context.Context)
}
