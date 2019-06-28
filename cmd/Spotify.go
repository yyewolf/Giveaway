package commands

import (
	"Giveaway/config"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

//isPremium : gets Premium status of someone.
func isPremium(ID string) bool {
	for i := range config.Premiums {
		if ID == config.Premiums[i] {
			return true
		}
	}
	return false
}

//IsBlackListed : gets BlackListed status of someone.
func IsBlackListed(ID string) bool {
	for i := range config.BlackList {
		if ID == config.BlackList[i] {
			return true
		}
	}
	return false
}

//Spotify : Gives a spotify account to the person that issued the command.
func Spotify(session *discordgo.Session, m *discordgo.MessageCreate) {

	//Will check everything : premium, blacklist and if user is supposed to wait
	if val, ok := config.Waiting[m.Author.ID]; ok {
		delay := val.Sub(time.Now())
		if int(math.Ceil(delay.Seconds())) > 0 {
			if !isPremium(m.Author.ID) {
				embed := &discordgo.MessageEmbed{
					Title:       "Giveaway Bot :",
					Description: `You need to wait : ` + strconv.Itoa(int(math.Floor(delay.Minutes()))) + `:` + strconv.Itoa(int(delay.Seconds()-(60*(math.Floor(delay.Seconds()/60))))) + " .\n\n Contact @:flag_in:|Sunny Singh|™✓:flag_in:#0001 for more info.",
					Footer: &discordgo.MessageEmbedFooter{
						Text: `Made by https://yewolf.ovh`,
					},
					Color: 0xFFDD00,
				}
				_, err := session.ChannelMessageSendEmbed(m.ChannelID, embed)
				check(err)
				return
			}
		} else {
			config.Waiting[m.Author.ID] = time.Now().Add(10 * time.Minute)
		}
	} else {
		config.Waiting[m.Author.ID] = time.Now().Add(10 * time.Minute)
	}

	if IsBlackListed(m.Author.ID) {
		return
	}

	//Seed to make it *almost* random
	rand.Seed(time.Now().UTC().UnixNano())

	channel, err := session.UserChannelCreate(m.Author.ID)
	check(err)
	//Sends a confirmation message.
	embed := &discordgo.MessageEmbed{
		Title:       "Giveaway Bot :",
		Description: `You should have received a DM, if not check that you allowed DMs from people of this server.`,
		Footer: &discordgo.MessageEmbedFooter{
			Text: `Made by https://yewolf.ovh`,
		},
		Color: 0xFFDD00,
	}
	_, err = session.ChannelMessageSendEmbed(m.ChannelID, embed)
	check(err)

	//Choose a *almost* random index.
	index := rand.Intn(len(config.Database))

	//Sends the DM.

	embedDM := &discordgo.MessageEmbed{
		Title:       "Account informations :",
		Description: "Here is your account : \n\nUsername : " + config.Database[index].Username + "\nPassword : " + config.Database[index].Password,
		Footer: &discordgo.MessageEmbedFooter{
			Text: `Made by https://yewolf.ovh`,
		},
		Color: 0xFFDD00,
	}
	_, err = session.ChannelMessageSendEmbed(channel.ID, embedDM)
	check(err)
}
