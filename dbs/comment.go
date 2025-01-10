package dbs

type Comment struct {
	Model
	Source    string    `json:"source" gorm:"type:TEXT;not null;comment:源内容"`
	Content   string    `json:"content" gorm:"type:TEXT;not null;comment:内容"`
	Attitude  int       `json:"attitude" gorm:"type:TINYINT;not null;comment:态度"`
	UserID    *uint     `json:"userId" gorm:"index;comment:作者ID"`
	User      User      `gorm:"constraint:OnDelete:SET NULL"`
	Comments  []Comment `gorm:"constraint:OnDelete:CASCADE"`
	ImageID   *uint     `json:"imageId" gorm:"index;comment:图片ID"`
	PostID    *uint     `json:"postId" gorm:"index;comment:帖子ID"`
	CommentID *uint     `json:"commentId" gorm:"index;comment:评论ID"`
}
