package discord

import (
	"bytes"
	"github.com/diamondburned/arikawa/v3/discord"
	"net/http"
	"strconv"
)

func RemoveMessagesIntegrations(token string, message discord.Message) (err error) {
	guildID := strconv.Itoa(int(message.GuildID))
	channelID := strconv.Itoa(int(message.ChannelID))
	messageID := strconv.Itoa(int(message.ID))

	url := "https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID
	payload := []byte(`{"flags":4}`)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("authorization", token)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("referer", "https://discord.com/channels/"+guildID+"/"+channelID)

	client := &http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
