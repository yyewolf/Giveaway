package commands

import (
	"runtime"
	"strconv"
	"time"

	"Giveaway/config"

	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, m *discordgo.MessageCreate) {
	i := 0
	exist := false
	for i = 0; i < len(config.HelpEmbeds); i++ {
		if config.HelpEmbeds[i].Title == "Help Menu" {
			exist = true
			break
		}
	}
	if exist {
		_, err := session.ChannelMessageSendEmbed(m.ChannelID, config.HelpEmbeds[i])
		check(err)
	}
}

func AdminHelp(session *discordgo.Session, m *discordgo.MessageCreate) {
	i := 0
	exist := false
	for i = 0; i < len(config.HelpEmbeds); i++ {
		if config.HelpEmbeds[i].Title == "Admin Menu" {
			exist = true
			break
		}
	}
	if exist {
		_, err := session.ChannelMessageSendEmbed(m.ChannelID, config.HelpEmbeds[i])
		check(err)
	}
}

func Info(session *discordgo.Session, m *discordgo.MessageCreate) {
	users := 0
	for _, guild := range session.State.Ready.Guilds {
		users += len(guild.Members)
	}
	ServerAmount := len(session.State.Guilds)
	Uptime := time.Since(config.StartTime)

	embed := &discordgo.MessageEmbed{
		Title: "Bot Statistics :",
		Description: `**Servers** : ` + strconv.Itoa(ServerAmount) + `
			**Users** : ` + strconv.Itoa(users) + `
			**Tasks** : ` + strconv.Itoa(runtime.NumGoroutine()) + `
			**Uptime** : ` + strconv.Itoa(int(Uptime.Hours())) + `:` + strconv.Itoa(int(Uptime.Minutes())%60) + `:` + strconv.Itoa(int(Uptime.Seconds())%60),
		Color: 0xFFDD00,
	}
	session.ChannelMessageSendEmbed(m.ChannelID, embed)
}
