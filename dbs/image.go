package dbs

type Image struct {
	Model
	Filename    string    `json:"filename" gorm:"not null;type:VARCHAR(255);comment:文件名"`
	Title       string    `json:"title" gorm:"not null;type:VARCHAR(255);unique;comment:标题"`
	Description string    `json:"description" gorm:"not null;type:TEXT;comment:简介"`
	Likes       uint      `json:"likes" gorm:"not null;comment:点赞"`
	UserID      *uint     `json:"userId" gorm:"index;comment:上传者"`
	User        *User     `json:"user" gorm:"constraint:OnDelete:SET NULL"`
	AlbumID     *uint     `json:"albumId" gorm:"index;not null;comment:相册ID"`
	Album       *Album    `json:"album" gorm:"constraint:OnDelete:CASCADE"`
	Comments    []Comment `gorm:"constraint:OnDelete:CASCADE"`
}
