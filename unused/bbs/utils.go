package main

import "github.com/McaxDev/backend/utils"

func CheckPostPerm(u *utils.User, p *utils.Post) bool {

	if u.Admin {
		return true
	}

	if p.UserID != nil && *p.UserID == u.ID {
		return true
	}

	if p.GuildID != nil && u.GuildID != nil {
		if *p.GuildID == *u.GuildID && *u.GuildRole > 3 {
			return true
		}
	}

	return false
}

func CheckCommentPerm(u *utils.User, c *utils.Comment) bool {

	if u.Admin {
		return true
	}

	if c.UserID != nil && u.ID == *c.UserID {
		return true
	}

	return false
}
