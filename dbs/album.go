package dbs

type Album struct {
	Model
	Cover       string  `json:"cover" gorm:"type:VARCHAR(255);comment:封面文件名"`
	Path        string  `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title       string  `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Description string  `json:"description" gorm:"type:TEXT;not null;comment:简介"`
	OnlyAdmin   bool    `json:"only_admin" gorm:"not null;comment:仅允许管理员"`
	Pinned      bool    `json:"pinned" gorm:"not null;comment:是否置顶"`
	GuildID     *uint   `json:"guild_id" gorm:"index;comment:公会ID"`
	Guild       *Guild  `json:"guild" gorm:"constraint:OnDelete:SET NULL;"`
	UserID      *uint   `json:"userId" gorm:"index;comment:创建者"`
	User        *User   `json:"user" gorm:"constraint:OnDelete:SET NULL;"`
	Images      []Image `json:"images" gorm:"constraint:OnDelete:CASCADE"`
}
