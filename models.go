package main

import "gorm.io/gorm"

type GameInstall struct {
	gorm.Model

	Title string
}
