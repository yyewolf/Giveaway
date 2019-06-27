package config

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var Prefix = "$$"

var StartTime = time.Now()

var Menus []Menu
var Categories []Category
var Commands []Command
var HelpEmbeds []*discordgo.MessageEmbed

type Account struct {
	Username string
	Password string
}

//Database : int => Username + Password
var Database []Account

//BlackList : Banned people from the bot
var BlackList []string

//Premiums : People that are premiums
var Premiums []string

//Waiting : People waiting 10 minutes
var Waiting map[string]time.Time

type Menu struct {
	Name        string
	Description string
	Main        bool
}

type Category struct {
	Name string
	Menu Menu
}

type Command struct {
	Category         Category
	Command          string
	ShortDescription string
	LongDescription  string
	Function         func(*discordgo.Session, *discordgo.MessageCreate)
}
