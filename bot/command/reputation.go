package command

import (
	"fmt"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/utils/json"
	"github.com/opendev-co/discord-bot/bot"
	"github.com/opendev-co/discord-bot/util"
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
	Target cmd.User `description:"L'utilisateur dont vous souhaitez ajouter un point de réputation"`
}

func (r ReputationAdd) Run(interaction *cmd.Interaction) {
	target := discord.UserID(r.Target).String()
	user := interaction.Member().User.ID.String()

	if target == user {
		_, err := interaction.Respond(cmd.MessageResponse{
			Embeds:    []discord.Embed{util.EmbedError("Vous ne pouvez pas ajouter de réputation à vous-même")},
			Ephemeral: true,
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
				Embeds:    []discord.Embed{util.EmbedError(fmt.Sprintf("Vous avez déjà ajouté un point de réputation à <@%v>", r.Target))},
				Ephemeral: true,
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
		Embeds: []discord.Embed{util.EmbedSuccess(fmt.Sprintf("Vous avez ajouté un point de réputation à <@%v>", r.Target))},
	})
	if err != nil {
		fmt.Println(err)
	}
}

type ReputationShow struct {
	Target cmd.Optional[cmd.User] `description:"L'utilisateur dont vous souhaitez connaître la réputation"`
}

func (r ReputationShow) Run(interaction *cmd.Interaction) {
	target := interaction.Member().User.ID.String()
	if r.Target.Provided() {
		target = discord.UserID(r.Target.Get()).String()
	}
	var p = "point"
	pts := reputations(target)
	if len(pts) > 1 {
		p += "s"
	}
	_, err := interaction.Respond(cmd.MessageResponse{
		Embeds: []discord.Embed{util.EmbedSuccess(fmt.Sprintf("<@%v> a %v %s de réputation", target, len(pts), p))},
	})
	if err != nil {
		fmt.Println(err)
	}
}

type ReputationTop struct{}

func (r ReputationTop) Run(interaction *cmd.Interaction) {
	page := 1
	embed, components, maxPage := util.GetLeaderboardMessageArguments(reps, page)

	_, err := interaction.Respond(cmd.MessageResponse{
		Embeds:     []discord.Embed{embed},
		Components: components,
	})

	interaction.API().AddHandler(func(m *gateway.InteractionCreateEvent) {
		if i, ok := m.Data.(*discord.ButtonInteraction); ok {
			if i.CustomID == "before" {
				if page > 1 {
					page--
				}
			} else if i.CustomID == "next" {
				if page < maxPage {
					page++
				}
			}

			embed, components, _ := util.GetLeaderboardMessageArguments(reps, page)

			if err != nil {
				fmt.Println(err)
			}

			_, err = interaction.EditResponse(api.EditInteractionResponseData{
				Embeds:     &[]discord.Embed{embed},
				Components: &components,
			})

			_ = bot.S.RespondInteraction(m.ID, m.Token,
				api.InteractionResponse{
					Type: api.DeferredMessageUpdate,
				},
			)

			if err != nil {
				fmt.Println(err)
			}
		}
	})

	if err != nil {
		fmt.Println(err)
	}
}
