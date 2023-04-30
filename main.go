package main

import (
	"context"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joho/godotenv"
	"github.com/opendev-co/discord-bot/bot"
	"github.com/opendev-co/discord-bot/bot/command"
	"github.com/opendev-co/discord-bot/bot/handler"
	"log"
	"os"
	"os/signal"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		panic(err)
	}

	tok, ok := os.LookupEnv("TOKEN")
	if !ok {
		panic("missing token")
	}

	s := state.New("Bot " + tok)
	s.AddIntents(gateway.IntentGuilds | gateway.IntentGuildMessages)

	bot.S = s

	// Initialize handlers
	s.AddHandler(handler.MessageCreate)
	s.AddHandler(handler.Ready)

	// Create a new command handler.
	h := cmd.NewHandler(nil).WithCommands(
		cmd.New("hello", "Hello, world!").WithExecutor(command.HelloWorld{}),
		cmd.New("calc", "Obtenir le résultat d'une expression mathématique").WithExecutor(command.Calculate{}),
		cmd.New("github", "Obtenir les lignes d'un répertoire Github").WithExecutor(command.Github{}),
		cmd.New("clear", "Supprimer des messages").WithExecutor(command.Clear{}),
		cmd.New("reputation", "Commande de réputation").WithSubcommand("add", "Ajouter un point de réputation à un utilisateur", command.ReputationAdd{}).WithSubcommand("show", "Voir les points de réputation d'un utilisateur", command.ReputationShow{}).WithSubcommand("top", "Voir les points de réputations de tous les utilisateurs", command.ReputationTop{}),
		cmd.New("rep", "Commande de réputation").WithSubcommand("add", "Ajouter un point de réputation à un utilisateur", command.ReputationAdd{}).WithSubcommand("show", "Voir les points de réputation d'un utilisateur", command.ReputationShow{}).WithSubcommand("top", "Voir les points de réputations de tous les utilisateurs", command.ReputationTop{}),
	)

	err = h.RegisterAll(s)
	if err != nil {
		panic(err)
	}
	// Adds the interaction event handler to the bot.
	err = h.Listen(s)
	if err != nil {
		panic(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = s.Open(ctx)
	if err != nil {
		panic(err)
	}
	<-ctx.Done() // block until Ctrl+C

	command.SaveReputation()

	if err := s.Close(); err != nil {
		log.Println("cannot close:", err)
	}
}
