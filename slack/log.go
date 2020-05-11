package slack

import "github.com/nlopes/slack"

func LogInfo(channel, message string) error {
	return defaultObj.PostTextMessage(channel, message)
}

func (s *SlackObj)  LogInfo(channel, message string) error {

	api := slack.New(s.Token)

	_, _, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		return err
	}

	return nil
}