package handler

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/opendev-co/discord-bot/bot"
	"github.com/opendev-co/discord-bot/discord"
)

var MessageCreate = func(m *gateway.MessageCreateEvent) {
	// Regex from https://github.com/HimbeersaftLP/MagicalHourglass/blob/main/src/commands/github.js#L7

	match := discord.GithubLinkMatch(m.Content)

	if !discord.ValidGithubLink(match) {
		return
	}

	message, ok := discord.FormatGithubLines(match)

	if !ok {
		return
	}

	_, _ = bot.S.SendMessageReply(m.ChannelID, message, m.ID)
	_ = discord.RemoveMessagesIntegrations(bot.S.Token, m.Message)
}
