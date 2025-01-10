package dbs

type Post struct {
	Model
	Title    string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Category string    `json:"category" gorm:"type:VARCHAR(255);comment:分类"`
	Source   string    `json:"source" gorm:"type:TEXT;not null;comment:原内容"`
	Content  string    `json:"content" gorm:"type:TEXT;not null;comment:内容"`
	UserID   *uint     `json:"userId" gorm:"index;comment:作者ID"`
	User     *User     `json:"user" gorm:"constraint:OnDelete:SET NULL"`
	Comments []Comment `json:"comments" gorm:"constraint:OnDelete:CASCADE"`
}
