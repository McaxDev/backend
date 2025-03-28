package utils

import (
	"time"
)

type ForeignKey struct {
	Parent     string
	Children   string
	ForeignKey string
	Action     string // 这是OnDelete操作，而OnUpdate保持默认
}

var Constraints = []ForeignKey{
	{"images", "guilds", "avatar_id", "RESTRICT"},
	{"images", "guilds", "cover_id", "RESTRICT"},
	{"images", "users", "avatar_id", "RESTRICT"},
	{"images", "users", "cover_id", "RESTRICT"},
	{"images", "albums", "cover_id", "RESTRICT"},
	{"images", "forums", "cover_id", "RESTRICT"},
	{"images", "prop_types", "icon_id", "RESTRICT"},

	{"users", "images", "user_id", "CASCADE"},
	{"users", "images", "uploader_id", "SET NULL"},
	{"users", "props", "owner_id", "CASCADE"},
	{"users", "reviews", "author_id", "SET NULL"},
	{"users", "albums", "creator_id", "SET NULL"},
	{"users", "posts", "author_id", "SET NULL"},

	{"guilds", "props", "guild_id", "CASCADE"},
	{"guilds", "posts", "guild_id", "CASCADE"},
	{"guilds", "images", "guild_id", "CASCADE"},
	{"guilds", "users", "guild_id", "SET NULL"},

	{"wiki_groups", "wikis", "wiki_group_id", "CASCADE"},
	{"albums", "images", "album_id", "CASCADE"},
	{"forum_groups", "forums", "forum_group_id", "CASCADE"},
	{"forums", "posts", "forum_id", "CASCADE"},
	{"posts", "reviews", "post_id", "CASCADE"},
	{"reviews", "reviews", "review_id", "CASCADE"},
	{"prop_types", "props", "prop_type_id", "CASCADE"},
}

var Tables = []any{
	new(Image),      // 图片表
	new(User),       // 用户表
	new(Guild),      // 公会表
	new(Wiki),       // 文档表
	new(WikiGroup),  // 文档组表
	new(Album),      // 相册表
	new(ForumGroup), // 论坛组表
	new(Forum),      // 论坛表
	new(Post),       // 帖子表
	new(Review),     // 评论表
	new(Prop),       // 道具表
	new(PropType),   // 道具类型表
	new(Online),     // 在线记录表
}

type Content struct {
	Text string  `json:"text"`
	HTML *string `json:"html"`
}

type User struct {
	ID        uint        `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time   `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"comment:更新时间"`
	Name      string      `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password  string      `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	AvatarID  *uint       `json:"-" gorm:"index;comment:头像图片"`
	Avatar    *Image      `json:"avatar" gorm:"foreignKey:AvatarID"`
	CoverID   *uint       `json:"-" gorm:"index;comment:封面图片"`
	Cover     *Image      `json:"cover" gorm:"foreignKey:CoverID"`
	Admin     bool        `json:"admin" gorm:"not null;comment:是否管理员"`
	Voter     bool        `json:"voter" gorm:"not null;comment:是否议员"`
	Voice     *uint       `json:"voice" gorm:"-"` // 议员投票话语权，计算得到
	IsMale    *bool       `json:"isMale" gorm:"comment:性别"`
	Profile   Content     `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:个人介绍"`
	Birthday  *time.Time  `json:"birthday" gorm:"comment:生日"`
	Location  *string     `json:"location" gorm:"type:VARCHAR(255);comment:地址"`
	Photos    []Image     `json:"photos"`
	DailyCoin uint        `json:"tempCoin" gorm:"not null;comment:签到币"`
	HonorCoin uint        `json:"permCoin" gorm:"not null;comment:贡献币"`
	Checkin   int64       `json:"-" gorm:"not null;comment:签到记录"`
	Email     string      `json:"email" gorm:"type:VARCHAR(255);not null;unique;comment:邮箱"`
	Phone     *string     `json:"phone" gorm:"type:VARCHAR(255);unique;comment:手机号"`
	QQ        *string     `json:"qq" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	MCBEName  *string     `json:"mcbeName" gorm:"type:VARCHAR(255);unique;comment:MCBE用户名"`
	MCJEName  *string     `json:"mcjeName" gorm:"type:VARCHAR(255);unique;comment:MCJE用户名"`
	GuildID   *uint       `json:"-" gorm:"index;comment:所属公会"`
	Guild     *Guild      `json:"guild"`
	GuildRole *uint       `json:"guildRole" gorm:"comment:公会身份角色"`
	Donation  uint        `json:"donation" gorm:"not null;comment:捐赠数额"`
	Exp       uint        `json:"exp" gorm:"not null;comment:经验值"`
	Level     uint        `json:"level" gorm:"-"`
	Setting   UserSetting `json:"setting" gorm:"type:JSON;serializer:json;not null;comment:用户设置"`
	Props     []Prop      `json:"props" gorm:"foreignKey:OwnerID"`
	Reviews   []Review    `json:"reviews" gorm:"foreignKey:AuthorID"`
	Albums    []Album     `json:"albums" gorm:"foreignKey:CreatorID"`
}

type WikiGroup struct {
	ID    uint   `json:"id" gorm:"primarykey;comment:ID"`
	Label string `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:名称"`
	Wikis []Wiki `json:"wikis"`
}

type Wiki struct {
	ID          uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Path        string    `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title       string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Content     Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:内容"`
	WikiGroupID uint      `json:"-" gorm:"index;comment:文档组"`
}

type Online struct {
	ID     uint      `json:"id" gorm:"primarykey;comment:ID"`
	Time   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Server string    `json:"server" gorm:"type:VARCHAR(255);not null;index;comment:服务器"`
	Count  *int64    `json:"count" gorm:"comment:在线人数"`
}

type Guild struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Name      string    `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:公会名"`
	Count     uint      `json:"count" gorm:"not null;comment:公会人数"`
	AvatarID  *uint     `json:"-" gorm:"index;comment:LOGO图片ID"`
	Avatar    *Image    `json:"avatar" gorm:"foreignKey:AvatarID"`
	CoverID   *uint     `json:"-" gorm:"index;comment:背景图片ID"`
	Cover     *Image    `json:"cover" gorm:"foreignKey:CoverID"`
	Profile   Content   `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:公会介绍"`
	Notice    Content   `json:"notice" gorm:"type:JSON;serializer:json;not null;comment:公会公告"`
	Money     uint      `json:"money" gorm:"not null;comment:公会资金"`
	Level     uint      `json:"level" gorm:"not null;comment:公会等级"`
	Props     []Prop    `json:"props"`
	Posts     []Post    `json:"posts"`
	Photos    []Image   `json:"photos"`
	Members   []User    `json:"members"`
}

type Album struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	CoverID   *uint     `json:"-" gorm:"index;comment:封面ID"`
	Cover     *Image    `json:"cover" gorm:"foreignKey:CoverID"`
	Path      string    `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Profile   Content   `json:"profile" gorm:"type:JSON;serializer:json;not null;comment:简介"`
	Admin     bool      `json:"admin" gorm:"not null;comment:仅允许管理员"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	CreatorID *uint     `json:"-" gorm:"index;comment:创建者"`
	Creator   *User     `json:"creator" gorm:"foreignKey:CreatorID"`
	Photos    []Image   `json:"photos"`
}

// 图片可以属于公共相册或公会相册或个人精选照片（此时具备点赞标题描述上传者）
// 图片也可以也可以不属于上述（即用于头像或封面等，不具备前述，为空指针）
type Image struct {
	ID         uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Filename   string    `json:"filename" gorm:"type:VARCHAR(255);not null;unique;comment:文件名"`
	Label      *string   `json:"label" gorm:"type:VARCHAR(255);comment:标题"`
	Profile    *string   `json:"profile" gorm:"type:VARCHAR(255);comment:简介"`
	Likes      *uint     `json:"likes" gorm:"comment:点赞"`
	UploaderID *uint     `json:"-" gorm:"index;comment:上传者用户ID"`
	Uploader   *User     `json:"uploader" gorm:"foreignKey:UploaderID"`
	UserID     *uint     `json:"-" gorm:"index;comment:用户照片ID"`
	User       *User     `json:"user"`
	AlbumID    *uint     `json:"-" gorm:"index;comment:相册ID"`
	Album      *Album    `json:"album"`
	GuildID    *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild      *Guild    `json:"guild"`
}

type ForumGroup struct {
	ID     uint    `json:"id" gorm:"primarykey;comment:ID"`
	Label  string  `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Forums []Forum `json:"forums"`
}

type Forum struct {
	ID           uint   `json:"id" gorm:"primarykey;comment:ID"`
	ForumGroupID uint   `json:"-" gorm:"index;not null;comment:论坛组ID"`
	Path         string `json:"path" gorm:"type:VARCHAR(255);not null;unique;comment:路径"`
	Title        string `json:"title" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	Profile      string `json:"profile" gorm:"type:VARCHAR(255);not null;comment:简介"`
	CoverID      *uint  `json:"-" gorm:"index;comment:封面ID"`
	Cover        *Image `json:"cover" gorm:"foreignKey:CoverID"`
	Posts        []Post `json:"posts"`
}

// 帖子可以属于公共论坛，也可以属于公会论坛
type Post struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Pinned    bool      `json:"pinned" gorm:"not null;comment:是否置顶"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;comment:标题"`
	ForumID   *uint     `json:"-" gorm:"index;comment:论坛ID"`
	Forum     *Forum    `json:"forum"`
	GuildID   *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild     *Guild    `json:"guild"`
	Content   Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:原内容"`
	AuthorID  *uint     `json:"-" gorm:"index;comment:作者ID"`
	Author    *User     `json:"author" gorm:"foreignKey:AuthorID"`
	Reviews   []Review  `json:"reviews"`
}

// 评论，可以属于相册、帖子或其他评论
type Review struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	Content   Content   `json:"content" gorm:"type:JSON;serializer:json;not null;comment:源内容"`
	Attitude  *bool     `json:"attitude" gorm:"comment:态度"`
	AuthorID  *uint     `json:"-" gorm:"index;comment:作者ID"`
	Author    *User     `json:"author" gorm:"foreignKey:AuthorID"`
	AlbumID   *uint     `json:"-" gorm:"index;comment:相册ID"`
	PostID    *uint     `json:"-" gorm:"index;comment:帖子ID"`
	ReviewID  *uint     `json:"-" gorm:"index;comment:评论ID"`
	Reviews   []Review  `json:"reviews"`
}

type Prop struct {
	ID         uint      `json:"id" gorm:"primarykey;comment:ID"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	OwnerID    *uint     `json:"-" gorm:"index;comment:用户ID"`
	Owner      *User     `json:"owner" gorm:"foreignKey:OwnerID"`
	GuildID    *uint     `json:"-" gorm:"index;comment:公会ID"`
	Guild      *Guild    `json:"guild"`
	PropTypeID uint      `json:"-" gorm:"index;not null;comment:道具类型ID"`
	PropType   PropType  `json:"propType"`
	Count      uint      `json:"count" gorm:"not null;comment:数量"`
}

type PropType struct {
	ID       uint   `json:"id" gorm:"primarykey;comment:ID"`
	Label    string `json:"label" gorm:"type:VARCHAR(255);not null;unique;comment:道具名称"`
	Profile  string `json:"profile" gorm:"type:TEXT;not null;comment:道具简介"`
	IconID   *uint  `json:"-" gorm:"index;comment:图标ID"`
	Icon     *Image `json:"icon" gorm:"foreignKey:IconID"`
	Function string `json:"function" gorm:"type:VARCHAR(255);not null;comment:功能"`
	Params   []byte `json:"params" gorm:"type:JSON;serializer:json;not null;comment:功能参数"`
}
