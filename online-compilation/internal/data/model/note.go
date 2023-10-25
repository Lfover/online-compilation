package model

import "time"

type NoteDetail struct {
	ID        int       `gorm:"column:id" db:"id" json:"id" form:"id"`                                 //  主键id
	Title     string    `gorm:"column:title" db:"title" json:"title" form:"title"`                     //  笔记标题
	Note      string    `gorm:"column:note" db:"note" json:"note" form:"note"`                         //  笔记内容
	Type      string    `gorm:"column:type" db:"type" json:"type" form:"type"`                         //  语言类型
	Code      string    `gorm:"column:code" db:"code" json:"code" form:"code"`                         //  代码
	Result    string    `gorm:"column:result" db:"result" json:"result" form:"result"`                 // 运行结果
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"` //  创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"` //  更新时间
}

func (*NoteDetail) TableName() string {
	return "note_detail"
}
