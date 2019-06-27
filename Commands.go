package main

import "Giveaway/config"
import commands "Giveaway/cmd"

func AllCommands() {
	mainMenu := config.Menu{
		Name:        "Help Menu",
		Description: "The prefix of the bot is currently : **$$**",
		Main:        true,
	}

	adminMenu := config.Menu{
		Name:        "Admin Menu",
		Description: "The prefix of the bot is currently : **$$**",
		Main:        false,
	}

	adminCategory := config.Category{
		Name: "Administration :",
		Menu: adminMenu,
	}

	accountsCategory := config.Category{
		Name: "Account types :",
		Menu: mainMenu,
	}

	spotify := config.Command{
		Category:         accountsCategory,
		Command:          config.Prefix + "spotify",
		ShortDescription: "**spotify** : Will give you a spotify premium account right into your DMs.",
		Function:         commands.Spotify,
	}

	help := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "help",
		ShortDescription: "**help** : The bot will send you the default help menu.",
		Function:         commands.Help,
	}
	
	helpAdmin := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "help 2",
		ShortDescription: "**help 2** : The bot will send you this menu.",
		Function:         commands.AdminHelp,
	}

	info := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "info",
		ShortDescription: "**info** : The bot will send it's info.",
		Function:         commands.Info,
	}

	//ADMIN ONLY COMMANDS :

	addpremium := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "addpremium",
		ShortDescription: "**addpremium** : Adds someone to the premium list.",
		Function:         commands.AddPremium,
	}

	removepremium := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "unpremium",
		ShortDescription: "**unpremium** : Removes someone from the premium list.",
		Function:         commands.RemovePremium,
	}

	addblacklist := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "blacklist",
		ShortDescription: "**blacklist** : Adds someone to the blacklist.",
		Function:         commands.AddBlacklist,
	}

	removeblacklist := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "whitelist",
		ShortDescription: "**whitelist** : Removes someone from the black list.",
		Function:         commands.RemoveBlackList,
	}

	premiuminfo := config.Command{
		Category:         adminCategory,
		Command:          config.Prefix + "premiuminfo",
		ShortDescription: "**premiuminfo** : Will tell you if you are premium on this bot.",
		Function:         commands.AmIPremium,
	}
	
	config.AddCommand(helpAdmin)
	config.AddCommand(help)
	config.AddCommand(info)
	config.AddCommand(spotify)

	config.AddCommand(addpremium)
	config.AddCommand(removepremium)
	config.AddCommand(addblacklist)
	config.AddCommand(removeblacklist)

	config.AddCommand(premiuminfo)

	config.CreateEmbeds()
}
