package commands

import (
	"Giveaway/config"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func LoadBlackList() {
	// Finds the absolute path of the file
	path, err := filepath.Abs("./db/blacklist.json")
	check(err)
	jsonFile, err := os.Open(path)
	check(err)
	byte, err := ioutil.ReadAll(jsonFile)
	check(err)
	defer jsonFile.Close()
	json.Unmarshal(byte, &config.BlackList)
}

func SaveBlackList() {
	file, _ := json.MarshalIndent(config.BlackList, "", " ")
	// Finds the absolute path of the file
	path, err := filepath.Abs("./db/blacklist.json")
	check(err)
	_ = ioutil.WriteFile(path, file, 0644)
}

//AddBlacklist : Adds users to the black list.
func AddBlacklist(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != "144472011924570113" && m.Author.ID != "395535610548322326" {
		return
	}
	Message := ""
	for i := range m.Mentions {
		config.BlackList = append(config.BlackList, m.Mentions[i].ID)
		Message += m.Mentions[i].Username + "\n"
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Added to blacklist :",
		Description: Message,
		Color:       0xFFDD00,
	}
	session.ChannelMessageSendEmbed(m.ChannelID, embed)
	SaveBlackList()
}

//RemoveBlackList : Removes users from the premium list.
func RemoveBlackList(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != "144472011924570113" && m.Author.ID != "395535610548322326" {
		return
	}
	Message := ""
	for i := range m.Mentions {
		for k := range config.BlackList {
			if config.BlackList[k] == m.Mentions[i].ID {
				config.BlackList = Remove(config.BlackList, k)
				break
			}
		}
		Message += m.Mentions[i].Username + "\n"
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Removed to blacklist :",
		Description: Message,
		Color:       0xFFDD00,
	}
	session.ChannelMessageSendEmbed(m.ChannelID, embed)
	SaveBlackList()
}
