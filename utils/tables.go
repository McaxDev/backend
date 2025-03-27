package utils

import (
	"bytes"
	"time"

	"github.com/yuin/goldmark"
	"gorm.io/gorm"
)

type Content struct {
	Text string  `json:"text"`
	HTML *string `json:"html"`
}

func (c *Content) Render(MD goldmark.Markdown) error {
	var buffer bytes.Buffer
	if err := MD.Convert([]byte(c.Text), &buffer); err != nil {
		return err
	}
	html := buffer.String()
	c.HTML = &html
	return nil
}

// not null的字段如果是指针，代表如果用户设置保密此项，就查询后设为nil
type User struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
	Name      string         `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password  string         `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	AvatarID  *uint          `json:"-" gorm:"index;comment:头像图片ID"`
	Avatar    *Image         `json:"avatar" gorm:"foreignKey:AvatarID"`
	CoverID   *uint          `json:"-" gorm:"index;comment:背景图片ID"`
	Cover     *Image         `json:"cover" gorm:"foreignKey:CoverID"`
	Admin     bool           `json:"admin" gorm:"not null;comment:是否管理员"`
	Voter     bool           `json:"voter" gorm:"not null;comment:是否议员"`
	Voice     *uint          `json:"voice" gorm:"-"` // 议员投票话语权，计算得到
	IsMale    *bool          `json:"isMale" gorm:"comment:性别"`
	Profile   Content        `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:个人介绍"`
	Birthday  *time.Time     `json:"birthday" gorm:"comment:生日"`
	Photos    []Image        `json:"photos" gorm:"constraint:OnDelete:CASCADE"`
	DailyCoin *uint          `json:"tempCoin" gorm:"not null;comment:签到币"` // 空指针代表应用层保密
	HonorCoin *uint          `json:"permCoin" gorm:"not null;comment:贡献币"` // 空指针代表应用层保密
	Checkin   int64          `json:"-" gorm:"not null;comment:签到记录"`
	Email     string         `json:"email" gorm:"type:VARCHAR(255);not null;unique;comment:邮箱"`
	Phone     *string        `json:"phone" gorm:"type:VARCHAR(255);unique;comment:手机号"`
	QQ        *string        `json:"qq" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	MCBEName  *string        `json:"mcbeName" gorm:"type:VARCHAR(255);unique;comment:MCBE用户名"`
	MCJEName  *string        `json:"mcjeName" gorm:"type:VARCHAR(255);unique;comment:MCJE用户名"`
	GuildID   *uint          `json:"-" gorm:"index;comment:公会ID"`
	Guild     *Guild         `json:"guild" gorm:"constraint:OnDelete:SET NULL"`
	GuildRole *uint          `json:"guildRole" gorm:"comment:公会身份角色"`
	Donation  uint           `json:"donation" gorm:"not null;comment:捐赠数额"`
	Exp       uint           `json:"exp" gorm:"not null;comment:经验值"`
	Level     uint           `json:"level" gorm:"-"`
	Setting   UserSetting    `json:"setting" gorm:"type:JSON;serializer:json;comment:用户设置"`
	Props     []Property     `json:"props" gorm:"constraint:OnDelete:CASCADE"`
	Reviews   []Review       `json:"reviews" gorm:"constraint:OnDelete:SET NULL"`
	Albums    []Album        `json:"albums" gorm:"constraint:OnDelete:SET NULL"`
}

type WikiGroup struct {
	ID    uint   `json:"id" gorm:"comment:ID"`
	Label string `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:名称"`
	Wikis []Wiki `json:"wikis" gorm:"constraint:OnDelete:CASCADE"`
}

type Wiki struct {
	ID          uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Path        string    `json:"path" gorm:"not null;type:VARCHAR(255);comment:路径"`
	Title       string    `json:"title" gorm:"not null;type:VARCHAR(255);comment:标题"`
	Content     Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:内容"`
	WikiGroupID uint      `json:"-" gorm:"index;not null"`
}

type Online struct {
	ID     uint      `json:"id" gorm:"primarykey;comment:ID"`
	Time   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Server string    `json:"server" gorm:"type:VARCHAR(255);not null;index;comment:服务器"`
	Count  *int64    `json:"count" gorm:"comment:在线人数"`
}

type Guild struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
	Name      string         `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:公会名"`
	Count     uint           `json:"count" gorm:"not null;comment:公会人数"`
	AvatarID  *uint          `json:"-" gorm:"index;comment:LOGO图片ID"`
	Avatar    *Image         `json:"avatar" gorm:"foreignKey:AvatarID"`
	CoverID   *uint          `json:"-" gorm:"index;comment:背景图片ID"`
	Cover     *Image         `json:"cover" gorm:"foreignKey:CoverID"`
	Profile   Content        `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:公会介绍"`
	Notice    Content        `json:"notice" gorm:"type:JSON;serializer:json;not null;comment:公会公告"`
	Money     uint           `json:"money" gorm:"not null;comment:公会资金"`
	Level     uint           `json:"level" gorm:"not null;comment:公会等级"`
	Posts     []Post         `json:"posts" gorm:"constraint:OnDelete:CASCADE"`
	Photos    []Image        `json:"photos" gorm:"constraint:OnDelete:CASCADE"`
	Members   []User         `json:"users" gorm:"constraint:OnDelete:SET NULL"`
}

type Album struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	CoverID   *uint     `json:"-" gorm:"index;comment:封面ID"`
	Cover     *Image    `json:"cover"`
	Path      string    `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Profile   Content   `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:简介"`
	Admin     bool      `json:"admin" gorm:"not null;comment:仅允许管理员"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	GuildID   *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild     *Guild    `json:"guild" gorm:"constraint:OnDelete:SET NULL;"`
	UserID    *uint     `json:"-" gorm:"index;comment:创建者"`
	User      *User     `json:"user" gorm:"constraint:OnDelete:SET NULL;"`
	Photos    []Image   `json:"photos" gorm:"constraint:OnDelete:CASCADE"`
}

// 图片可以属于公共相册或公会相册或个人精选照片（此时具备点赞标题描述上传者）
// 图片也可以也可以不属于上述（即用于头像或封面等，不具备前述，为空指针）
type Image struct {
	ID         uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Filename   string    `json:"filename" gorm:"type:VARCHAR(255);not null;unique;comment:文件名"`
	Label      *string   `json:"label" gorm:"type:VARCHAR(255);unique;comment:标题"`
	Profile    *string   `json:"profile" gorm:"type:TEXT;comment:简介"`
	Likes      *uint     `json:"likes" gorm:"comment:点赞"`
	UploaderID *uint     `json:"-" gorm:"index;comment:上传者用户ID"`
	Uploader   *User     `json:"uploader" gorm:"constraint:OnDelete:SET NULL"`
	UserID     *uint     `json:"-" gorm:"index;comment:所属用户照片ID"`
	User       *User     `json:"user" gorm:"constraint:OnDelete:CASCADE"`
	AlbumID    *uint     `json:"-" gorm:"index;comment:所属相册ID"`
	Album      *Album    `json:"album" gorm:"constraint:OnDelete:CASCADE"`
	GuildID    *uint     `json:"-" gorm:"index;comment:所属公会ID"`
	Guild      *Guild    `json:"guild" gorm:"constraint:OnDelete:CASCADE"`
}

type ForumGroup struct {
	ID     uint    `json:"id" gorm:"primarykey;comment:ID"`
	Label  string  `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Forums []Forum `json:"forums" gorm:"constraint:OnDelete:CASCADE"`
}

type Forum struct {
	ID           uint    `json:"id" gorm:"primarykey;comment:ID"`
	ForumGroupID uint    `json:"-" gorm:"index;not null;comment:论坛组ID"`
	Path         string  `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title        string  `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	SubTitle     string  `json:"subTitle" gorm:"type:VARCHAR(255);not null;comment:副标题"`
	Profile      Content `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:介绍"`
	CoverID      *uint   `json:"-" gorm:"index;comment:封面ID"`
	Cover        *Image  `json:"cover"`
	Posts        []Post  `json:"posts" gorm:"constraint:OnDelete:CASCADE"`
}

// 帖子可以属于公共论坛，也可以属于公会论坛
type Post struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;comment:标题"`
	ForumID   *uint     `json:"-" gorm:"index;comment:论坛ID"`
	Forum     *Forum    `json:"forum" gorm:"constraint:OnDelete:CASCADE"`
	GuildID   *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild     *Guild    `json:"guild" gorm:"constraint:OnDelete:CASCADE"`
	Content   Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:原内容"`
	UserID    *uint     `json:"-" gorm:"index;comment:作者ID"`
	User      *User     `json:"user" gorm:"constraint:OnDelete:SET NULL"`
	Reviews   []Review  `json:"comments" gorm:"constraint:OnDelete:CASCADE"`
}

// 评论，可以属于相册、帖子或其他评论
type Review struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Content   Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:源内容"`
	Attitude  *bool     `json:"attitude" gorm:"comment:态度"`
	UserID    *uint     `json:"userId" gorm:"index;comment:作者ID"`
	User      *User     `gorm:"constraint:OnDelete:SET NULL"`
	AlbumID   *uint     `json:"-" gorm:"index;comment:相册ID"`
	PostID    *uint     `json:"-" gorm:"index;comment:帖子ID"`
	ReviewID  *uint     `json:"-" gorm:"index;comment:评论ID"`
	Reviews   []Review  `gorm:"constraint:OnDelete:CASCADE"`
}

type Property struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	UserID    uint      `json:"-" gorm:"index;not null;comment:用户ID"`
	Property  string    `json:"property" gorm:"not null;comment:道具ID"`
	Count     uint      `json:"count" gorm:"not null;comment:数量"`
}
