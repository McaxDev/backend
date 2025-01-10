package dbs

type Question struct {
	Model
	Content string `json:"content" gorm:"not null;type:TEXT;comment:题目内容"`
	Class   string `json:"class" gorm:"index;not null;type:VARCHAR(255);comment:题目分类"`
	Answer  string `json:"answer" gorm:"not null;type:VARCHAR(255);comment:正确答案"`
}
