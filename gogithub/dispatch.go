package gogithub

import (
	"context"
	"github.com/google/go-github/v29/github"
	"strings"
)

func TriggerRepositoryDispatch(repo, owner, eventType, payload string) error {
	return defaultObj.TriggerRepositoryDispatch(repo, owner, eventType, payload)
}

func (g *Gogithub) TriggerRepositoryDispatch(repo, owner, eventType, payload string) error {


	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	opts := github.DispatchRequestOptions{}
	opts.EventType = eventType
	//opts.ClientPayload = &json.RawMessage{}  // todo: payload

	_, _, err := client.Repositories.Dispatch(context.Background(), owner, repo, opts)
	if err != nil {
		return err
	}

	return nil
}