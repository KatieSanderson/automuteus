package command

import (
	"strings"

	"github.com/automuteus/utils/pkg/discord"
	"github.com/automuteus/utils/pkg/settings"
	"github.com/bwmarrin/discordgo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type LinkStatus int

const (
	LinkSuccess LinkStatus = iota
	LinkNoPlayer
	LinkNoGameData
)

var Link = discordgo.ApplicationCommand{
	Name:        "link",
	Description: "Link a Discord User to their in-game color (required for >24 colors)",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "User to link",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "color",
			Description: "In-game color",
			Required:    true,
		},
	},
}

func GetLinkParams(s *discordgo.Session, options []*discordgo.ApplicationCommandInteractionDataOption) (string, string) {
	return options[0].UserValue(s).ID, strings.ReplaceAll(strings.ToLower(options[1].StringValue()), " ", "")
}

func LinkResponse(status LinkStatus, userID, color string, sett *settings.GuildSettings) *discordgo.InteractionResponse {
	var content string
	switch status {
	case LinkSuccess:
		content = sett.LocalizeMessage(&i18n.Message{
			ID:    "commands.link.success",
			Other: "Successfully linked {{.UserMention}} to an in-game player with the color: `{{.Color}}`",
		}, map[string]interface{}{
			"UserMention": discord.MentionByUserID(userID),
			"Color":       color,
		})
	case LinkNoPlayer:
		content = sett.LocalizeMessage(&i18n.Message{
			ID:    "commands.link.noplayer",
			Other: "No player in the current game was detected for {{.UserMention}}",
		}, map[string]interface{}{
			"UserMention": discord.MentionByUserID(userID),
		})
	case LinkNoGameData:
		content = sett.LocalizeMessage(&i18n.Message{
			ID:    "commands.link.nogamedata",
			Other: "No game data found for the color `{{.Color}}`",
		}, map[string]interface{}{
			"Color": color,
		})
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: content,
		},
	}
}
