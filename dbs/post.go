package dbs

type Post struct {
	Model
	Title    string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Forum    string    `json:"forum,omitempty" gorm:"type:VARCHAR(255);comment:论坛"`
	Source   string    `json:"source,omitempty" gorm:"type:TEXT;not null;comment:原内容"`
	Content  string    `json:"content,omitempty" gorm:"type:TEXT;not null;comment:内容"`
	GuildID  *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild    *Guild    `json:"guild" gorm:"constraint:OnDelete:CASCADE"`
	UserID   *uint     `json:"-" gorm:"index;comment:作者ID"`
	User     *User     `json:"user" gorm:"constraint:OnDelete:SET NULL"`
	Comments []Comment `json:"comments" gorm:"constraint:OnDelete:CASCADE"`
}
