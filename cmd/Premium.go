package commands

import (
	"Giveaway/config"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func LoadPremiums() {
	// Finds the absolute path of the file
	path, err := filepath.Abs("./db/premiums.json")
	check(err)
	jsonFile, err := os.Open(path)
	check(err)
	byte, err := ioutil.ReadAll(jsonFile)
	check(err)
	defer jsonFile.Close()
	json.Unmarshal(byte, &config.Premiums)
}

func SavePremiums() {
	file, _ := json.MarshalIndent(config.Premiums, "", " ")
	// Finds the absolute path of the file
	path, err := filepath.Abs("./db/premiums.json")
	check(err)
	_ = ioutil.WriteFile(path, file, 0644)
}

func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

//AddPremium : Adds users from the premium list.
func AddPremium(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != "144472011924570113" && m.Author.ID != "395535610548322326" {
		return
	}
	Message := ""
	Already := ""
	for i := range m.Mentions {
		if isPremium(m.Mentions[i].ID) {
			Already += m.Mentions[i].Username + "\n"
		} else {
			config.Premiums = append(config.Premiums, m.Mentions[i].ID)
			Message += m.Mentions[i].Username + "\n"
		}
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Premium added :",
		Description: Message,
		Color:       0xFFDD00,
	}
	session.ChannelMessageSendEmbed(m.ChannelID, embed)
	if Already != "" {
		embed = &discordgo.MessageEmbed{
			Title:       "Already premium :",
			Description: Already,
			Color:       0xFFDD00,
		}
		session.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
	SavePremiums()
}

//RemovePremium : Removes users from the premium list.
func RemovePremium(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != "144472011924570113" && m.Author.ID != "395535610548322326" {
		return
	}
	Message := ""
	for i := range m.Mentions {
		for k := range config.Premiums {
			if config.Premiums[k] == m.Mentions[i].ID {
				config.Premiums = Remove(config.Premiums, k)
				break
			}
		}
		Message += m.Mentions[i].Username + "\n"
	}
	embed := &discordgo.MessageEmbed{
		Title:       "Premium removed :",
		Description: Message,
		Color:       0xFFDD00,
	}
	_, err := session.ChannelMessageSendEmbed(m.ChannelID, embed)
	check(err)
	SavePremiums()
}

func AmIPremium(session *discordgo.Session, m *discordgo.MessageCreate) {
	if isPremium(m.Author.ID) {
		embed := &discordgo.MessageEmbed{
			Title:       "Premium status :",
			Description: m.Author.Username + "is a premium user!",
			Color:       0xFFDD00,
		}
		_, err := session.ChannelMessageSendEmbed(m.ChannelID, embed)
		check(err)
	} else {
		embed := &discordgo.MessageEmbed{
			Title:       "Premium status :",
			Description: m.Author.Username + " is a not premium user!",
			Color:       0xFFDD00,
		}
		_, err := session.ChannelMessageSendEmbed(m.ChannelID, embed)
		check(err)
	}
}
