// QSFDKSJhflkjSDHLKFJHqklfKQSHQSFDKSJhflkjSDHLKFJHqklfKQSHQSFDKSJhflkjSDHLKFJHqklfKQSHQSFDKSJhflkjSDHLKFJHqklfKQSH

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	commands "Giveaway/cmd"
	"Giveaway/config"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func LoadDatabase() {
	// Finds the absolute path of the file
	path, err := filepath.Abs("./db/accounts.txt")
	check(err)
	//Read the file
	file, err := ioutil.ReadFile(path)
	check(err)
	//Convert files to string
	str := string(file)
	Accounts := strings.Split(str, "\n")
	// Formats the database here :
	for i := range Accounts {
		Account := strings.Split(Accounts[i], ":")
		ToAdd := config.Account{
			Username: Account[0],
			Password: Account[len(Account)-1],
		}
		config.Database = append(config.Database, ToAdd)
	}
	//Database has been formatted !
	color.Green("Database successfully taken.")
}

func main() {
	config.Waiting = make(map[string]time.Time)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	LoadDatabase()
	commands.LoadPremiums()
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(botReady)
	AllCommands()
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func botReady(session *discordgo.Session, evt *discordgo.Ready) {
	color.Green("Bot is now running.  Press CTRL-C to exit.")
	session.UpdateStatus(0, "$$help for help")
}
