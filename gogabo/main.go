// Declare this file to be part of the main package so it can be compiled into
// an executable.
package main

// Import all Go packages required for this file.
import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {
	// Discord Authentication Token
	Session.Token = os.Getenv("DG_TOKEN")
}

func main() {

	// Declare any variables needed later.
	var err error

	// Verify a Token was provided
	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	listChannels(Session)

	// log.Printf(Session.State.Guilds[])

	for _, guild := range Session.State.Guilds {
		log.Printf("guild.ID is " + guild.ID)
		log.Printf("Channels for guild are: " + Session.GuildChannels(guild.ID))
	    channels, _ := Session.GuildChannels(guild.ID)
		for _, c := range channels {
            log.Printf("Channel name is " + c.Name)
            log.Printf("Channel ID is " + c.ID)

        }
	}

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}

func listChannels(s *discordgo.Session) {
    // Loop through each guild in the session
    for _, guild := range s.State.Guilds {

        // Get channels for this guild
        channels, _ := s.GuildChannels(guild.ID)

        for _, c := range channels {
            log.Println("Channel name is " + c.Name)
            log.Println("Channel ID is " + c.ID)

        }
    }
}