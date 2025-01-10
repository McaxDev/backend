package main

import "github.com/McaxDev/backend/dbs"

func CheckImagePerm(user *dbs.User, image *dbs.Image) bool {

	if user.Admin {
		return true
	}

	if *image.UserID == user.ID {
		return true
	}

	if *image.Album.UserID == user.ID {
		return true
	}

	if *image.Album.GuildID != 0 {
		if user.GuildID == image.Album.GuildID && user.GuildRole > 2 {
			return true
		}
	}

	return false
}

func CheckUploadImage(user *dbs.User, album *dbs.Album) bool {

	if user.Admin {
		return true
	}

	if *album.GuildID == 0 {
		if *album.UserID == user.ID || album.OnlyAdmin == false {
			return true
		}
	} else {
		if user.GuildID == album.GuildID && user.GuildRole > 2 {
			return true
		}
		if album.OnlyAdmin == true {
			return true
		}
	}

	return false
}

func CheckEditAlbum(user *dbs.User, album *dbs.Album) bool {
	if user.Admin {
		return true
	}

	if *album.GuildID == 0 {
		if *album.UserID == user.ID {
			return true
		}
	} else {
		if *album.GuildID != 0 {
			if user.GuildID == album.GuildID && user.GuildRole > 2 {
				return true
			}
			if *album.UserID == user.ID {
				return true
			}
		}
	}

	return false
}
