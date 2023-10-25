package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "online-compilation/api/note"
)

// 这个就是注入到service层的
type NoteUsecase struct {
	repo NoteRepo
	log  *log.Helper
}

func (u NoteUsecase) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (string, error) {
	u.repo.CreateNote(ctx)
	return "你好", nil
}

// 这个做初始化的时候使用，程序运行就会先从这个来创建实例
func NewNoteUsecase(repo NoteRepo, logger log.Logger) *NoteUsecase {
	return &NoteUsecase{repo: repo, log: log.NewHelper(logger)}
}

type NoteRepo interface {
	CreateNote(ctx context.Context) //这里是需要数据库层，也就是data层需要实现的函数
}

//下载文件，data层的我没写，大概就是查询数据库拿数据就行
//func (p CommentUsecase) CommentDownload(ctx context.Context, types []string, endTime, startTime, userId, traceId, appId, sessionId string) (commentList []Comment, err error) {
//	commentList = make([]Comment, 0, 10)
//	commentList, err = p.repo.CommentDownloadList(ctx, types, endTime, startTime, userId, traceId, appId, sessionId)
//	if err != nil {
//		return commentList, err
//	}
//
//	return commentList, nil
//}
