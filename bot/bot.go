package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"pkg.nit.so/switchboard"

	"github.com/fogo-sh/dunce/database"
)

type Config struct {
	DBPath  string `default:"dunce.sqlite" split_words:"true"`
	Token   string
	GuildId string `default:"497544520695808000" split_words:"true"`
	AppId   string `default:"1294448357124870245" split_words:"true"`
}

var config Config
var Bot *discordgo.Session

func Run() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %s\n", err.Error())
	}

	err = envconfig.Process("dunce", &config)
	if err != nil {
		return fmt.Errorf("error loading config: %s", err)
	}

	database.Initialize(config.DBPath)
	database.Migrate()

	Bot, err = discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	if err != nil {
		return fmt.Errorf("error creating Discord session: %w", err)
	}
	Bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	parser := &switchboard.Switchboard{}
	Bot.AddHandler(parser.HandleInteractionCreate)
	RegisterCommands(parser)
	err = parser.SyncCommands(Bot, config.AppId)
	if err != nil {
		return fmt.Errorf("error syncing commands: %w", err)
	}

	if err = Bot.Open(); err != nil {
		return fmt.Errorf("error opening Discord connection: %w", err)
	}

	fmt.Println("Dunce is now running. Press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("Quitting Dunce")

	if err = Bot.Close(); err != nil {
		return fmt.Errorf("error closing Discord connection: %w", err)
	}

	return nil
}

func RegisterCommands(parser *switchboard.Switchboard) {
	_ = parser.AddCommand(&switchboard.Command{
		Name:        "db",
		Description: "Get a copy of the Dunce database.",
		Handler:     DbCommand,
		GuildID:     config.GuildId,
	})

	_ = parser.AddCommand(&switchboard.Command{
		Name:        "test",
		Description: "test command",
		Handler:     TestCommand,
		GuildID:     config.GuildId,
	})
}
