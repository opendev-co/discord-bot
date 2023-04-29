package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json"
	"os"
	"sync"
)

var (
	respMu = sync.Mutex{}
	reps   = make(map[string][]string)
)

func SaveReputation() {
	data, err := json.Marshal(reps)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("data/reputation.json", data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func reputations(id string) []string {
	respMu.Lock()
	defer respMu.Unlock()
	return reps[id]
}

func init() {
	data, err := os.ReadFile("data/reputation.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &reps)
}

type ReputationAdd struct {
	Target cmd.User `description:"The user to add reputation to"`
}

func (c ReputationAdd) Run(interaction *cmd.Interaction) {
	target := discord.UserID(c.Target).String()
	user := interaction.Member().User.ID.String()

	if target == user {
		_, err := interaction.Respond(cmd.MessageResponse{
			Content: fmt.Sprintf("Vous ne pouvez pas ajouter de réputation à vous-même"),
		})
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	pts := reputations(target)
	for _, v := range pts {
		if v == user {
			_, err := interaction.Respond(cmd.MessageResponse{
				Content: fmt.Sprintf("Vous avez déjà ajouté un point de réputation à <@%v>", c.Target),
			})
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
	respMu.Lock()
	defer respMu.Unlock()
	reps[target] = append(reps[target], user)

	_, err := interaction.Respond(cmd.MessageResponse{
		Content: fmt.Sprintf("Vous avez ajouté un point de réputation à <@%v>", c.Target),
	})
	if err != nil {
		fmt.Println(err)
	}
}

type ReputationShow struct {
	Target cmd.Optional[cmd.User] `description:"The user to show reputation of"`
}

func (c ReputationShow) Run(interaction *cmd.Interaction) {
	target := interaction.Member().User.ID.String()
	if c.Target.Provided() {
		target = discord.UserID(c.Target.Get()).String()
	}
	pts := reputations(target)
	_, err := interaction.Respond(cmd.MessageResponse{
		Content: fmt.Sprintf("<@%v> a %v points de réputation", target, len(pts)),
	})
	if err != nil {
		fmt.Println(err)
	}
}
