package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/mnogu/go-calculator"
	"strings"
)

type Calculate struct {
	Expression string `description:"The expression to calculate"`
}

var expressionReplacer = strings.NewReplacer("x", "*", "÷", "/")

func (c Calculate) Run(interaction *cmd.Interaction) {
	res, err := calculator.Calculate(expressionReplacer.Replace(strings.ToLower(c.Expression)))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = interaction.Respond(cmd.MessageResponse{
		Content:   fmt.Sprintf("%v", res),
		Ephemeral: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
