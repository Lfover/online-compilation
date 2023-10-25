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

//这是一个老师写的从网页上下载文件的方法，你可以参考然后你写你的笔记导出
//func (s *CommentService) CommentDownload(ctx http.Context) error {
//	rawURL := ctx.Request().URL
//	u, err := url.Parse(rawURL.String())
//	if err != nil {
//		return err
//	}
//	traceId := u.Query().Get("trace_id")
//	userId := u.Query().Get("user_id")
//	commentType := u.Query().Get("type")
//	appId := u.Query().Get("app_id")
//	startTime := u.Query().Get("start_time")
//	endTime := u.Query().Get("end_time")
//	sessionId := u.Query().Get("session_id")
//	var types []string
//	if err = json.Unmarshal([]byte(commentType), &types); err != nil {
//		return err
//	}
//	err = checkParam(startTime, endTime)
//	if err != nil {
//		return err
//	}
//	commentList, err := s.commentUsecase.CommentDownload(ctx, types, endTime, startTime, userId, traceId, appId, sessionId)
//	if err != nil {
//		return err
//	}
//	apps, err := s.appRepo.QueryList(ctx, 0)
//	if err != nil {
//		s.log.Errorf("QueryList err: %v", err)
//		return err
//	}
//	appMap := make(map[string]string)
//	for _, app := range apps {
//		appMap[app.AppId] = app.AppName
//	}
//	f := download_excel.NewExcel()
//	sheetName := "Sheet1"
//	fileName := "comments.xlsx"
//	f.File.NewSheet(sheetName)
//	// 设置表头
//	headers := []string{"应用名称", "trace_id", "session_id", "user_id", "类型", "时间"}
//	f.SetHeader(sheetName, headers)
//	for i, comment := range commentList {
//		row := i + 2
//		appName, ok := appMap[comment.AppId]
//		if !ok {
//			appName = comment.AppId
//		}
//		f.File.SetCellValue(sheetName, "A"+strconv.Itoa(row), appName)
//		f.File.SetCellValue(sheetName, "B"+strconv.Itoa(row), comment.TraceId)
//		f.File.SetCellValue(sheetName, "C"+strconv.Itoa(row), comment.SessionId)
//		f.File.SetCellValue(sheetName, "D"+strconv.Itoa(row), comment.UserId)
//		f.File.SetCellValue(sheetName, "E"+strconv.Itoa(row), comment.TypeChinese)
//		f.File.SetCellValue(sheetName, "F"+strconv.Itoa(row), comment.CreatedAt)
//	}
//	buf := new(bytes.Buffer)
//	err = f.File.Write(buf)
//	if err != nil {
//		return err
//	}
//	if err = f.File.SaveAs(fileName); err != nil {
//		return err
//	}
//	// 写入http响应体
//	if err = f.Write(ctx, buf.Bytes(), fileName); err != nil {
//		return err
//	}
//	return nil
//}
//
//func checkParam(startTime, endTime string) error {
//	if startTime == "" || endTime == "" {
//		return errors.New("param time can not be empty")
//	}
//	// 解析时间字符串
//	sTime, err := time.Parse("2006-01-02 15:04:05", startTime)
//	if err != nil {
//		return err
//	}
//	eTime, err := time.Parse("2006-01-02 15:04:05", endTime)
//	if err != nil {
//		return err
//	}
//	if eTime.Unix()-sTime.Unix() > 86400*7 {
//		return errors.New("param time can not be more than 7 days")
//	}
//	return nil
//}
