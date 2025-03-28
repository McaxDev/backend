package utils

type Setting struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Type    string `json:"type"`
	Default any    `json:"default"`
}

type UserSetting struct {
	PubEmail    *bool   `json:"pubEmail,omitempty" label:"公开邮箱" default:"false" group:"privacy"`
	PubPhone    *bool   `json:"pubPhone,omitempty" label:"公开手机号" default:"false" group:"privacy"`
	PubQQ       *bool   `json:"pubQQ,omitempty" label:"公开QQ号" default:"false" group:"privacy"`
	PubGamename *bool   `json:"pubGamename,omitempty" label:"公开游戏名称" default:"false" group:"privacy"`
	PubGuild    *bool   `label:"公开公会" default:"false" group:"privacy"`
	PubProps    *bool   `label:"公开道具" default:"false" group:"privacy"`
	PubComments *bool   `label:"公开评论" default:"false" group:"privacy"`
	PubAlbums   *bool   `label:"公开相册" default:"false" group:"privacy"`
	PubCoin     *bool   `label:"公开金币" default:"false" group:"privacy"`
	PubGamedata *bool   `label:"公开游戏数据" default:"false" group:"privacy"`
	PubDonation *bool   `label:"公开捐赠数额" default:"false" group:"privacy"`
	PubLevel    *bool   `label:"公开等级" default:"false" group:"privacy"`
	UseMFA      *bool   `label:"登录启用MFA验证" default:"false" group:"privacy"`
	UseCaptcha  *bool   `label:"登录启用Captcha验证" default:"false" group:"privacy"`
	BEChatFmt   *string `label:"基岩聊天格式" default:"false" group:"privacy"`
	JEChatFmt   *string `label:"MCJE聊天格式" default:"false" group:"privacy"`
}

/*
	var SettingMap = map[string]Setting{
		"PubEmail":    {"公开邮箱", "boolean", false},
		"PubPhone":    {"公开手机号", "boolean", false},
		"PubQQ":       {"公开QQ号", "boolean", false},
		"PubGameName": {"公开游戏名称", "boolean", true},
		"PubGuild":    {"公开公会", "boolean", true},
		"PubProps":    {"公开道具", "boolean", false},
		"PubComments": {"公开评论", "boolean", true},
		"PubAlbums":   {"公开相册", "boolean", true},
		"PubCoin":     {"公开金币", "boolean", true},
		"PubGameData": {"公开游戏数据", "boolean", true},
		"PubDonation": {"公开捐赠数额", "boolean", true},
		"PubLevel":    {"公开等级", "boolean", true},
		"UseMFA":      {"登录启用MFA验证", "boolean", false},
		"UseCaptcha":  {"登录启用Captcha验证", "boolean", false},
		"BEChatFmt":   {"基岩版聊天格式", "string", ""},
		"JEChatFmt":   {"Java版聊天格式", "string", ""},
	}
*/

var SettingOrder = []struct {
	Label    string
	Settings []string
}{
	{
		Label: "隐私设置",
		Settings: []string{
			"PubEmail",
			"PubPhone",
			"PubQQ",
			"PubGameName",
			"PubGuild",
			"PubProps",
			"PubComments",
			"PubAlbums",
			"PubCoin",
			"PubGameData",
			"PubDonation",
			"PubLevel",
		},
	},
	{
		Label: "安全设置",
		Settings: []string{
			"UseMFA",
			"UseCaptcha",
		},
	},
	{
		Label: "个性化设置",
		Settings: []string{
			"BEChatFmt",
			"JEChatFmt",
		},
	},
}
