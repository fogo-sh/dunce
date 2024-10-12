package bot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func TestCommand(_ *discordgo.Session, interaction *discordgo.InteractionCreate, args struct{}) {
	users, err := db.GetAllUsers(context.Background())

	if err != nil {
		slog.Error("Failed to get users", "error", err)
		Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to get users",
			},
		})
		return
	}

	countOfUsers := len(users)

	Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("test, %d users", countOfUsers),
		},
	})
}
