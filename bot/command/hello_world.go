package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
)

type HelloWorld struct{}

func (HelloWorld) Run(interaction *cmd.Interaction) {
	_, err := interaction.Respond(cmd.MessageResponse{
		Content:   fmt.Sprintf("Hello, %s", interaction.Member().User.Username),
		Ephemeral: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
