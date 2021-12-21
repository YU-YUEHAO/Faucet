package mysql
import "time"

// 定义数据库的访问模型
type TODO struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	ID      int        `sql:"id" json:"id" form:"id"`                // 唯一标识的
	Title   string     `sql:"title" json:"title" form:"title"`       // 待办的标题
	State   bool       `sql:"state" json:"state" form:"state"`       // 待办的状态，为true表示已完成，为false表示未完成
	Created *time.Time `sql:"created" json:"created" form:"created"` // 创建待办的时间
}