package utils

import (
	"math"
	"time"

	"gorm.io/gorm"
)

type ForeignKey struct {
	Parent     string
	Children   string
	ForeignKey string
	Action     string
}

var Constraints = []ForeignKey{
	{"images", "users", "avatar_id", "RESTRICT"},
	{"images", "users", "cover_id", "RESTRICT"},
	{"guilds", "users", "guild_id", "SET NULL"},
	{"images", "guilds", "avatar_id", "RESTRICT"},
	{"images", "guilds", "cover_id", "RESTRICT"},
	{"images", "props", "icon_id", "RESTRICT"},
	{"users", "images", "uploader_id", "CASCADE"},
	{"images", "albums", "cover_id", "RESTRICT"},
	{"props", "items", "prop_id", "CASCADE"},
	{"bbs", "forums", "bbs_id", "CASCADE"},
	{"images", "forums", "cover_id", "RESTRICT"},
	{"users", "posts", "author_id", "CASCADE"},
	{"users", "reviews", "author_id", "CASCADE"},
	{"users", "albums", "creator_id", "CASCADE"},
}

var Tables = []any{
	new(Image),    // 图片表
	new(User),     // 用户表
	new(Guild),    // 公会表
	new(Document), // 文档表
	new(Wiki),     // 知识库表
	new(Album),    // 相册表
	new(BBS),      // 论坛组表
	new(Forum),    // 论坛表
	new(Post),     // 帖子表
	new(Review),   // 评论表
	new(Item),     // 道具表
	new(Prop),     // 道具类型表
	new(Online),   // 在线记录表
}

type Content struct {
	Text string  `json:"text"`
	HTML *string `json:"html"`
}

type Owner struct {
	Name     string `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:名称"`
	AvatarID *uint  `json:"-" gorm:"index;comment:头像图片"`
	Avatar   *Image `json:"avatar" gorm:"foreignKey:AvatarID"`
	CoverID  *uint  `json:"-" gorm:"index;comment:封面图片"`
	Cover    *Image `json:"cover" gorm:"foreignKey:CoverID"`
	Exp      uint   `json:"exp" gorm:"not null;comment:经验值"`
	Level    uint   `json:"level" gorm:"-"`
}

func (o *Owner) AfterFind(db *gorm.DB) error {
	result := uint(math.Floor(math.Sqrt(float64(o.Exp)) / 10))
	if result < 6 {
		o.Level = result
	} else {
		o.Level = 6
	}
	return db.
		Preload("Avatar", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("filename")
		}).
		Preload("Cover", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("filename")
		}).Error
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Owner
	Password  string       `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	Admin     bool         `json:"admin,omitempty" gorm:"not null;comment:是否管理员"`
	Voter     bool         `json:"voter,omitempty" gorm:"not null;comment:是否议员"`
	IsMale    *bool        `json:"isMale,omitempty" gorm:"comment:性别"`
	Profile   string       `json:"profile,omitempty" gorm:"type:TEXT;not null;comment:个人介绍"`
	Birthday  *time.Time   `json:"birthday,omitempty" gorm:"comment:生日"`
	Location  *string      `json:"location,omitempty" gorm:"type:VARCHAR(255);comment:地址"`
	DailyCoin uint         `json:"dailyCoin,omitempty" gorm:"not null;comment:签到币"`
	HonorCoin uint         `json:"honorCoin,omitempty" gorm:"not null;comment:贡献币"`
	Checkin   int64        `json:"-" gorm:"not null;comment:签到记录"`
	Email     string       `json:"email,omitempty" gorm:"type:VARCHAR(255);not null;unique;comment:邮箱"`
	QQ        *string      `json:"qq,omitempty" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	MCBEName  *string      `json:"mcbeName,omitempty" gorm:"type:VARCHAR(255);unique;comment:MCBE用户名"`
	MCJEName  *string      `json:"mcjeName,omitempty" gorm:"type:VARCHAR(255);unique;comment:MCJE用户名"`
	GuildID   *uint        `json:"-" gorm:"index;comment:所属公会"`
	Guild     *Guild       `json:"guild,omitempty" gorm:"foreignKey:GuildID"`
	GuildRole *uint        `json:"guildRole,omitempty" gorm:"comment:公会身份"`
	Donation  uint         `json:"donation,omitempty" gorm:"not null;comment:捐赠数额"`
	Setting   *UserSetting `json:"setting,omitempty" gorm:"type:JSON;serializer:json;comment:用户设置"`
	Comments  []Review     `json:"comments,omitempty" gorm:"foreignKey:AuthorID"`
	Reviews   []Review     `json:"reviews,omitempty" gorm:"polymorphic:Refer;polymorphicValue:users"`
	Items     []Item       `json:"items,omitempty" gorm:"polymorphic:Owner;polymorphicValue:users"`
	Albums    []Album      `json:"album,omitempty" gorm:"foreignKey:CreatorID"`
}

type Wiki struct {
	ID        uint       `json:"id" gorm:"primaryKey;comment:ID"`
	Label     string     `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:名称"`
	Sort      int        `json:"sort" gorm:"not null;comment:排序"`
	Documents []Document `json:"documents" gorm:"polymorphic:Wiki;polymorphicValue:wikis"`
}

type Document struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Path      string    `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;comment:标题"`
	Content   *Content  `json:"content" gorm:"type:JSON;serializer:json;comment:内容"`
	WikiID    uint      `json:"-" gorm:"index:idx_wiki;not null;comment:知识库ID"`
	WikiType  string    `json:"-" gorm:"index:idx_wiki;not null;type:VARCHAR(255);polymorphic:Wiki;comment:知识库类型"`
	Sort      int       `json:"sort" gorm:"not null;comment:排序"`
	Reviews   []Review  `json:"reviews" gorm:"polymorphic:Refer;polymorphicValue:documents"`
}

type Online struct {
	ID     uint      `json:"id" gorm:"primaryKey;comment:ID"`
	Time   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Server string    `json:"server" gorm:"type:VARCHAR(255);not null;index;comment:服务器"`
	Count  *int64    `json:"count" gorm:"comment:在线人数"`
}

type Guild struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Owner
	Count     uint       `json:"count" gorm:"not null;comment:公会人数"`
	Profile   *Content   `json:"profile" gorm:"type:JSON;serializer:json;comment:公会介绍"`
	Money     uint       `json:"money" gorm:"not null;comment:公会资金"`
	Members   []User     `json:"members" gorm:"foreignKey:GuildID"`
	Items     []Item     `json:"items" gorm:"polymorphic:Owner;polymorphicValue:guilds"`
	Photos    []Image    `json:"images" gorm:"polymorphic:Album;polymorphicValue:guilds"`
	Posts     []Post     `json:"posts" gorm:"polymorphic:Forum;polymorphicValue:guilds"`
	Documents []Document `json:"documents" gorm:"polymorphic:Wiki;polymorphicValue:wikis"`
}

type Album struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	CoverID   *uint     `json:"-" gorm:"index;comment:封面ID"`
	Cover     *Image    `json:"cover" gorm:"foreignKey:CoverID"`
	CreatorID *uint     `json:"-" gorm:"index;comment:创建者ID"`
	Creator   *User     `json:"creator" gorm:"foreignKey:CreatorID"`
	Path      string    `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Profile   string    `json:"profile" gorm:"type:TEXT;not null;comment:简介"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	Protected bool      `json:"protected" gorm:"not null;comment:仅创建者修改"`
	Photos    []Image   `json:"photos" gorm:"polymorphic:Album;polymorphicValue:albums"`
	Reviews   []Review  `json:"reviews" gorm:"polymorphic:Refer;polymorphicValue:albums"`
}

type Image struct {
	ID         uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Filename   string    `json:"filename" gorm:"type:VARCHAR(255);not null;unique;comment:文件名"`
	Label      *string   `json:"label,omitempty" gorm:"type:VARCHAR(255);comment:标题"`
	Profile    *string   `json:"profile,omitempty" gorm:"type:VARCHAR(255);comment:简介"`
	Likes      *uint     `json:"likes,omitempty" gorm:"comment:点赞"`
	UploaderID *uint     `json:"-" gorm:"index;comment:上传者用户ID"`
	Uploader   *User     `json:"uploader,omitempty" gorm:"foreignKey:UploaderID"`
	AlbumID    uint      `json:"-" gorm:"index:idx_album;not null;comment:相册ID"`
	AlbumType  string    `json:"-" gorm:"index:idx_album;not null;type:VARCHAR(255);polymorphic:Album;comment:相册类型"`
}

type BBS struct {
	ID     uint    `json:"id" gorm:"primaryKey;comment:ID"`
	Label  string  `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Sort   int     `json:"sort" gorm:"not null;comment:排序"`
	Forums []Forum `json:"forums"`
}

type Forum struct {
	ID       uint   `json:"id" gorm:"primaryKey;comment:ID"`
	BBSID    uint   `json:"-" gorm:"index;not null;comment:论坛组ID"`
	Path     string `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title    string `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	SubTitle string `json:"subTitle" gorm:"type:VARCHAR(255);not null;comment:副标题"`
	Profile  string `json:"profile" gorm:"type:TEXT;not null;comment:简介"`
	CoverID  *uint  `json:"-" gorm:"index;comment:封面ID"`
	Cover    *Image `json:"cover" gorm:"foreignKey:CoverID"`
	Sort     int    `json:"sort" gorm:"not null;comment:排序"`
	Posts    []Post `json:"posts" gorm:"polymorphic:Forum;polymorphicValue:forums"`
}

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;comment:标题"`
	ForumID   uint      `json:"-" gorm:"index:idx_forum;not null;comment:论坛ID"`
	ForumType string    `json:"-" gorm:"index:idx_forum;not null;type:VARCHAR(255);polymorphic:Forum;comment:论坛类型"`
	Content   *Content  `json:"content" gorm:"type:JSON;serializer:json;comment:原内容"`
	AuthorID  uint      `json:"-" gorm:"index;not null;comment:作者ID"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`
	Reviews   []Review  `json:"reviews" gorm:"polymorphic:Refer;polymorphicValue:posts"`
}

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Content   Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:源内容"`
	Attitude  *bool     `json:"attitude" gorm:"comment:态度"`
	AuthorID  uint      `json:"-" gorm:"index;not null;comment:作者ID"`
	Author    *User     `json:"author" gorm:"foreignKey:AuthorID"`
	ReferID   uint      `json:"-" gorm:"index:idx_refer;not null;comment:指向ID"`
	ReferType string    `json:"-" gorm:"index:idx_refer;not null;type:VARCHAR(255);polymorphic:Refer;comment:指向类型"`
	Refer     string    `json:"refer" gorm:"-"`
	Reviews   []Review  `json:"reviews" gorm:"polymorphic:Refer;polymorphicValue:reviews"`
}

var ReferMap = map[string]string{
	"users":     "name",
	"albums":    "title",
	"posts":     "title",
	"reviews":   "content",
	"documents": "title",
}

func (r *Review) AfterFind(db *gorm.DB) error {
	return db.Table(r.ReferType).Select(ReferMap[r.ReferType]).First(&r.Refer, "id = ?", r.ReferID).Error
}

type Item struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:ID"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	PropID    uint      `json:"-" gorm:"index;not null;comment:道具类型ID"`
	Prop      Prop      `json:"prop"`
	OwnerID   uint      `json:"-" gorm:"index:idx_owner;not null;comment:拥有者ID"`
	OwnerType string    `json:"-" gorm:"index:idx_owner;not null;type:VARCHAR(255);polymorphic:Owner;comment:拥有者类型"`
	Owner     *Owner    `json:"owner" gorm:"-"`
	Count     uint      `json:"count" gorm:"not null;comment:数量"`
}

func (i *Item) AfterFind(db *gorm.DB) error {
	return db.Table(i.OwnerType).First(i.Owner, "id = ?", i.OwnerID).Error
}

type Prop struct {
	ID       uint   `json:"id" gorm:"primaryKey;comment:ID"`
	Label    string `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:道具名称"`
	Profile  string `json:"profile" gorm:"type:TEXT;not null;comment:道具简介"`
	IconID   *uint  `json:"-" gorm:"index;comment:图标ID"`
	Icon     *Image `json:"icon" gorm:"foreignKey:IconID"`
	Function string `json:"function" gorm:"type:VARCHAR(255);not null;comment:功能"`
	Params   []byte `json:"params" gorm:"type:JSON;serializer:json;not null;comment:功能参数"`
}
