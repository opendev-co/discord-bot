package util

import (
	"fmt"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/opendev-co/discord-bot/bot"
	"math"
	"sort"
)

func getDescription(reps map[string][]string, page int) (string, string, int) {
	m := make(map[string]int)
	separator := 10

	for key, value := range reps {
		m[key] = len(value)
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	i := separator * (page - 1)

	message := "Les points de réputations ne représentent rien ⚠️. Ils veulent seulement dire que la personne a aidé plusieurs personnes, ce n'est pas car une personne a beaucoup de points qu'il est de confiance !\n\n"
	maxPage := int(math.Ceil(float64(len(reps)) / float64(separator)))

	for _, key := range keys {
		i++
		message += fmt.Sprintf("- **%v**: <@%s> avec **%d** points\n", i, key, m[key])
	}

	message += "\nCliquez sur le bouton ◀️ pour aller à la page précédente\nCliquez sur le bouton ▶ pour aller à la page suivante\n"
	title := fmt.Sprintf("Classement des points de réputation (%d Points distribués)", getReputationsPointsNumber(reps))

	return message, title, maxPage
}

func GetLeaderboardMessageArguments(reps map[string][]string, page int) (discord.Embed, discord.ContainerComponents, int) {
	description, title, maxPage := getDescription(reps, page)
	user, _ := bot.S.Me()

	embed := discord.Embed{
		Title:       title,
		Description: description,
		Timestamp:   discord.NowTimestamp(),
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf("%d/%v ・ %s", page, maxPage, user.Username),
			Icon: user.AvatarURL(),
		},
		Color: 0x524aff,
	}

	components := discord.ContainerComponents{
		&discord.ActionRowComponent{
			&discord.ButtonComponent{
				CustomID: "before",
				Label:    "◀️",
				Style:    discord.SecondaryButtonStyle(),
			},
			&discord.ButtonComponent{
				CustomID: "next",
				Label:    "▶️",
				Style:    discord.SecondaryButtonStyle(),
			},
		},
	}

	return embed, components, maxPage
}

func getReputationsPointsNumber(reps map[string][]string) int {
	var count int
	for _, value := range reps {
		count += len(value)
	}
	return count
}
