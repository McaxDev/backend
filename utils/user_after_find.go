package utils

import "gorm.io/gorm"

func (user *User) AfterFind(tx *gorm.DB) (err error) {

	/*
		clearPrivacy, ok := tx.Statement.Context.Value("clearPrivacy").(bool)
		if ok && clearPrivacy {
				if !GetSetting[bool](user, "PubEmail") {
					user.Email = nil
				}

				if !GetSetting[bool](user, "PubPhone") {
					user.Phone = nil
				}

				if !GetSetting[bool](user, "PubQQ") {
					user.QQ = nil
				}

				if user.Setting.PubGamename != nil {
					user.JavaName = nil
					user.BedrockName = nil
				}

				if !GetSetting[bool](user, "PubGuild") {
					user.Guild = nil
					user.GuildRole = nil
				}

				if !GetSetting[bool](user, "PubProps") {
					user.Props = nil
				}

				if !GetSetting[bool](user, "PubComments") {
					user.Comments = nil
				}

				if !GetSetting[bool](user, "PubAlbums") {
					user.Albums = nil
				}

				if !GetSetting[bool](user, "PubCoin") {
					user.PermCoin = nil
					user.TempCoin = nil
				}

				if !GetSetting[bool](user, "PubDonation") {
					user.Donation = nil
				}

				if !GetSetting[bool](user, "PubLevel") {
					user.Exp = nil
				} else {
					if user.Exp >= 0 {
						user.Level = 1
					} else if user.Exp >= 100 {
						user.Level = 2
					} else if user.Exp >= 500 {
						user.Level = 3
					} else if user.Exp >= 1000 {
						user.Level = 4
					} else if user.Exp >= 5000 {
						user.Level = 5
					} else if user.Exp >= 10000 {
						user.Level = 6
					}
				}
		}
	*/
	return nil
}
