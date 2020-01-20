package gogithub

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/v29/github"
	"strings"
)

func (g *Gogithub) TriggerRepositoryDispatch(repo, owner, eventType, payload string) error {


	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	opts := github.DispatchRequestOptions{}
	opts.EventType = eventType
	opts.ClientPayload = &json.RawMessage{}

	_, _, err := client.Repositories.Dispatch(context.Background(), owner, repo, opts)
	if err != nil {
		return err
	}

	return nil
}