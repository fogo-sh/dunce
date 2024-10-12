package bot

import (
	"github.com/bwmarrin/discordgo"
)

func TestCommand(_ *discordgo.Session, interaction *discordgo.InteractionCreate, args struct{}) {
	Bot.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "test",
		},
	})
}
