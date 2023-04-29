package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/opendev-co/discord-bot/bot"
	"github.com/opendev-co/discord-bot/util"
)

type Clear struct {
	Amount int                    `description:"Nombre de messages à supprimer"`
	Member cmd.Optional[cmd.User] `description:"Utilisateur à clear"`
}

func (c Clear) Run(interaction *cmd.Interaction) {
	_, err := bot.S.Channel(interaction.ChannelID())
	if err != nil {
		return
	}

	if c.Amount <= 0 {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Embeds:    []discord.Embed{util.EmbedError("Le nombre de messages à supprimer doit être supérieur à 0")},
			Ephemeral: true,
		})
		return
	}

	if !util.HasPermission(interaction.GuildID(), *interaction.Member(), discord.PermissionManageMessages) {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Embeds:    []discord.Embed{util.EmbedError("Vous n'avez pas la permission de supprimer des messages")},
			Ephemeral: true,
		})

		return
	}

	member := c.Member.GetOrFallback(cmd.User(0))
	var messagesID []discord.MessageID

	amount := 0

	if member != 0 {
		messages, _ := bot.S.Messages(interaction.ChannelID(), 9999)

		for _, message := range messages {
			if message.Author.ID == discord.UserID(member) {
				amount++
				messagesID = append(messagesID, message.ID)
			}
		}
	} else {
		messages, _ := bot.S.Messages(interaction.ChannelID(), uint(c.Amount))

		for _, message := range messages {
			amount++
			messagesID = append(messagesID, message.ID)
		}
	}

	_ = bot.S.DeleteMessages(interaction.ChannelID(), messagesID, api.AuditLogReason("Clear by "+interaction.User().Tag()))

	if member != 0 {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Embeds: []discord.Embed{util.EmbedSuccess(fmt.Sprintf("%v messages de <@%v> ont été supprimés", amount, member))},
		})

		return
	} else {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Embeds: []discord.Embed{util.EmbedSuccess(fmt.Sprintf("%v messages ont été supprimés", amount))},
		})
	}
}
