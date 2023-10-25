package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"online-compilation/internal/biz"
)

type noteRepo struct {
	Data *Data
	Log  *log.Helper
}

func (n noteRepo) CreateNote(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewNoteRepo(data *Data, logger log.Logger) biz.NoteRepo {
	return &noteRepo{
		Data: data,
		Log:  log.NewHelper(logger),
	}
}
