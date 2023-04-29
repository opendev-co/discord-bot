package util

import "github.com/diamondburned/arikawa/v3/discord"

func EmbedColor(msg string, title string, color discord.Color) discord.Embed {
	e := discord.Embed{
		Title:       title,
		Description: msg,
		Color:       color,
	}
	return e
}

func EmbedSuccess(msg string) discord.Embed {
	return EmbedColor(msg, "", 0x524aff)
}

func EmbedSuccessTitle(msg string, title string) discord.Embed {
	return EmbedColor(msg, title, 0x524aff)
}

func EmbedError(msg string) discord.Embed {
	return EmbedColor(msg, "", 0xd34a52)
}
