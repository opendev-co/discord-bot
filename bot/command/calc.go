package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mnogu/go-calculator"
	"github.com/opendev-co/discord-bot/util"
	"strings"
)

type Calculate struct {
	Expression string `description:"The expression to calculate"`
}

var expressionReplacer = strings.NewReplacer("x", "*", "÷", "/")

func (c Calculate) Run(interaction *cmd.Interaction) {
	res, err := calculator.Calculate(expressionReplacer.Replace(strings.ToLower(c.Expression)))
	if err != nil {
		_, err = interaction.Respond(cmd.MessageResponse{
			Embeds:    []discord.Embed{util.EmbedError("L'expression que vous avez spécifié est invalide")},
			Ephemeral: true,
		})
		fmt.Println(err)
		return
	}
	_, err = interaction.Respond(cmd.MessageResponse{
		Embeds: []discord.Embed{util.EmbedSuccessTitle(fmt.Sprintf("Résultat : %v", res), c.Expression)},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
