package command

import (
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/opendev-co/discord-bot/util"
)

type Github struct {
	Link string `description:"Lien github"`
}

func (g Github) Run(interaction *cmd.Interaction) {
	match := util.GithubLinkMatch(g.Link)

	if !util.ValidGithubLink(match) {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Content:   "Le lien github est invalide et ou ne comporte aucune lignes à montrer",
			Ephemeral: true,
		})

		return
	}

	message, ok := util.FormatGithubLines(match)

	if !ok {
		_, _ = interaction.Respond(cmd.MessageResponse{
			Content:   "Le repo' semble être inaccessible ou le fichier comporte trop peu de lignes",
			Ephemeral: true,
		})

		return
	}

	_, _ = interaction.Respond(cmd.MessageResponse{
		Content: message,
	})
}
