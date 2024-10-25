package bot

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"pkg.nit.so/switchboard"

	"github.com/fogo-sh/dunce/database"
	"github.com/fogo-sh/dunce/database/queries"
)

type Config struct {
	DBPath  string `default:"dunce.sqlite" split_words:"true"`
	Token   string
	GuildId string `default:"497544520695808000" split_words:"true"`
	AppId   string `default:"1294448357124870245" split_words:"true"`
}

var config Config
var Bot *discordgo.Session

var db *queries.Queries

func Run(inputConfig Config) error {
	config = inputConfig

	slog.Info("Initializing Dunce...")

	dbInstance, err := database.New(config.DBPath)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	db = dbInstance

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

	slog.Info("Dunce is now running. Press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	slog.Info("Quitting Dunce")

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
