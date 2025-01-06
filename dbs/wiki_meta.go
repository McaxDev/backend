package dbs

type WikiMeta struct {
	Path  string `json:"path" gorm:"comment:'路径'"`
	Title string `json:"title" gorm:"comment:'标题'"`
	Order string `json:"order" gorm:"comment:'次序'"`
}
