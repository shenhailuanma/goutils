package slack

import "github.com/nlopes/slack"

func PostTextMessage(channel, message string) error {
	return defaultObj.PostTextMessage(channel, message)
}

func (s *SlackObj)  PostTextMessage(channel, message string) error {

	api := slack.New(s.Token)

	_, _, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		return err
	}

	return nil
}