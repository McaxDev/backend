package dbs

type Wiki struct {
	Model
	Path     string `json:"path" gorm:"not null;type:VARCHAR(255);comment:路径"`
	Title    string `json:"title" gorm:"not null;type:VARCHAR(255);comment:标题"`
	Markdown string `json:"markdown,omitempty" gorm:"not null;type:TEXT;comment:内容"`
	HTML     string `json:"html,omitempty" gorm:"not null;type:TEXT;comment:HTML内容"`
	Category string `json:"category" gorm:"type:VARCHAR(255);not null;comment:分类ID"`
}
