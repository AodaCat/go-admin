package autocode

import "go-admin/global"

// 如果含有time.Time 请自行import time包
type AutoCodeExample struct {
	global.GA_MODEL
	AutoCodeExampleField string `json:"autoCodeExampleField" form:"autoCodeExampleField" gorm:"column:auto_code_example_field;comment:仅作示例条目无实际作用"` // 展示值
}
