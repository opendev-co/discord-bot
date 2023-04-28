package main

import (
	"context"
	"github.com/andreashgk/go-interactions/cmd"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joho/godotenv"
	"github.com/opendev-co/discord-bot/bot/command"
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

	// Create a new command handler.
	h := cmd.NewHandler(nil).WithCommands(
		cmd.New("hello", "Hello, world!").WithExecutor(command.HelloWorld{}),
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

	if err := s.Close(); err != nil {
		log.Println("cannot close:", err)
	}
}
