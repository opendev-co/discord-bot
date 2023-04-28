package handler

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/opendev-co/discord-bot/bot"
	"github.com/opendev-co/discord-bot/util"
)

var MessageCreate = func(m *gateway.MessageCreateEvent) {
	// Regex from https://github.com/HimbeersaftLP/MagicalHourglass/blob/main/src/commands/github.js#L7

	if m.Author.Bot {
		return
	}

	match := util.GithubLinkMatch(m.Content)
	if !util.ValidGithubLink(match) {
		return
	}

	message, ok := util.FormatGithubLines(match)
	if !ok {
		return
	}

	_, _ = bot.S.SendMessageReply(m.ChannelID, message, m.ID)
	_ = util.RemoveMessagesIntegrations(bot.S.Token, m.Message)
}
