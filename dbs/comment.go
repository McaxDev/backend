package dbs

type Comment struct {
	Model
	Source    string    `json:"source" gorm:"type:TEXT;not null;comment:源内容"`
	Content   string    `json:"content" gorm:"type:TEXT;not null;comment:内容"`
	Attitude  int       `json:"attitude" gorm:"type:TINYINT;not null;comment:态度"`
	UserID    *uint     `json:"userId,omitempty" gorm:"index;comment:作者ID"`
	User      User      `gorm:"constraint:OnDelete:SET NULL"`
	Comments  []Comment `gorm:"constraint:OnDelete:CASCADE"`
	ImageID   *uint     `json:"imageId,omitempty" gorm:"index;comment:图片ID"`
	PostID    *uint     `json:"postId,omitempty" gorm:"index;comment:帖子ID"`
	CommentID *uint     `json:"commentId,omitempty" gorm:"index;comment:评论ID"`
}
