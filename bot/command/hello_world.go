package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
)

type HelloWorld struct{}

func (HelloWorld) Run(interaction *cmd.Interaction) {
	followup, err := interaction.DeferResponse()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = followup.Create(cmd.MessageResponse{
		Content: "Hello, world!",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
