package gogithub

import (
	"context"
	"github.com/google/go-github/v29/github"
	"strings"
)

type PullRequest struct {
	//ID                  *int64     `json:"id,omitempty"`
	Number int    `json:"number"`
	State  string `json:"state"`
	Title  string `json:"title"`
	//Body                *string    `json:"body,omitempty"`
	//CreatedAt           *time.Time `json:"created_at,omitempty"`
	//UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	//ClosedAt            *time.Time `json:"closed_at,omitempty"`
	//MergedAt            *time.Time `json:"merged_at,omitempty"`
	//Labels              []*Label   `json:"labels,omitempty"`
	//User                *User      `json:"user,omitempty"`
	Merged         bool   `json:"merged"`
	Mergeable      bool   `json:"mergeable"`
	MergeableState string `json:"mergeable_state"`
	//MergedBy            *User      `json:"merged_by,omitempty"`
	//MergeCommitSHA      *string    `json:"merge_commit_sha,omitempty"`
	//Comments            *int       `json:"comments,omitempty"`
	//Commits             *int       `json:"commits,omitempty"`
	//Additions           *int       `json:"additions,omitempty"`
	//Deletions           *int       `json:"deletions,omitempty"`
	//ChangedFiles        *int       `json:"changed_files,omitempty"`
	//URL                 *string    `json:"url,omitempty"`
	//HTMLURL             *string    `json:"html_url,omitempty"`
	//IssueURL            *string    `json:"issue_url,omitempty"`
	//StatusesURL         *string    `json:"statuses_url,omitempty"`
	//DiffURL             *string    `json:"diff_url,omitempty"`
	//PatchURL            *string    `json:"patch_url,omitempty"`
	//CommitsURL          *string    `json:"commits_url,omitempty"`
	//CommentsURL         *string    `json:"comments_url,omitempty"`
	//ReviewCommentsURL   *string    `json:"review_comments_url,omitempty"`
	//ReviewCommentURL    *string    `json:"review_comment_url,omitempty"`
	//Assignee            *User      `json:"assignee,omitempty"`
	//Assignees           []*User    `json:"assignees,omitempty"`
	//Milestone           *Milestone `json:"milestone,omitempty"`
	//MaintainerCanModify *bool      `json:"maintainer_can_modify,omitempty"`
	//AuthorAssociation   *string    `json:"author_association,omitempty"`
	//NodeID              *string    `json:"node_id,omitempty"`
	//RequestedReviewers  []*User    `json:"requested_reviewers,omitempty"`
}

type CommitFile struct {
	SHA      string `json:"sha"`
	Filename string `json:"filename"`
	Changes  int    `json:"changes"`
	Status   string `json:"status"`
	Patch    string `json:"patch"`
}

func PullRequests(repo, owner string) ([]PullRequest, error) {
	return defaultObj.PullRequests(repo, owner)
}

func (g *Gogithub) PullRequests(repo, owner string) ([]PullRequest, error) {
	var output = []PullRequest{}

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	if owner == "" {
		owner = g.Username
	}

	pullrequests, _, err := client.PullRequests.List(context.Background(), owner, repo, nil)
	if err != nil {
		return output, err
	}

	for _, one := range pullrequests {
		var itemOne = PullRequest{}

		if one.Number != nil {
			itemOne.Number = *one.Number
		}
		if one.Title != nil {
			itemOne.Title = *one.Title
		}
		if one.State != nil {
			itemOne.State = *one.State
		}
		if one.Merged != nil {
			itemOne.Merged = *one.Merged
		}
		if one.Mergeable != nil {
			itemOne.Mergeable = *one.Mergeable
		}
		if one.MergeableState != nil {
			itemOne.MergeableState = *one.MergeableState
		}

		output = append(output, itemOne)
	}

	return output, nil
}

func PullRequestsFilter(repo, owner, filterHead, filterBase, state string) ([]PullRequest, error) {
	return defaultObj.PullRequestsFilter(repo, owner, filterHead, filterBase, state)
}

func (g *Gogithub) PullRequestsFilter(repo, owner, filterHead, filterBase, state string) ([]PullRequest, error) {

	var output = []PullRequest{}

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	if owner == "" {
		owner = g.Username
	}
	if state == "" {
		state = "open"
	}

	var opt = github.PullRequestListOptions{}
	opt.State = state
	opt.Head = filterHead
	opt.Base = filterBase

	pullrequests, _, err := client.PullRequests.List(context.Background(), owner, repo, &opt)
	if err != nil {
		return output, err
	}

	for _, one := range pullrequests {
		var itemOne = PullRequest{}

		if one.Number != nil {
			itemOne.Number = *one.Number
		}
		if one.Title != nil {
			itemOne.Title = *one.Title
		}
		if one.State != nil {
			itemOne.State = *one.State
		}
		if one.Merged != nil {
			itemOne.Merged = *one.Merged
		}
		if one.Mergeable != nil {
			itemOne.Mergeable = *one.Mergeable
		}
		if one.MergeableState != nil {
			itemOne.MergeableState = *one.MergeableState
		}

		output = append(output, itemOne)
	}

	return output, nil
}

func PullRequestListFiles(repo, owner string, number int) ([]CommitFile, error) {
	return defaultObj.PullRequestListFiles(repo, owner, number)
}

func (g *Gogithub) PullRequestListFiles(repo, owner string, number int) ([]CommitFile, error) {

	var output = []CommitFile{}

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	if owner == "" {
		owner = g.Username
	}

	files, _, err := client.PullRequests.ListFiles(context.Background(), owner, repo, number, nil)
	if err != nil {
		return output, err
	}

	for _, one := range files {

		var itemOne = CommitFile{}

		if one.SHA != nil {
			itemOne.SHA = *one.SHA
		}
		if one.Status != nil {
			itemOne.Status = *one.Status
		}
		if one.Filename != nil {
			itemOne.Filename = *one.Filename
		}
		if one.Changes != nil {
			itemOne.Changes = *one.Changes
		}
		if one.Patch != nil {
			itemOne.Patch = *one.Patch
		}

		output = append(output, itemOne)
	}

	return output, nil
}

func PullRequestUpdate(repo, owner string, number int, state string) error {
	return defaultObj.PullRequestUpdate(repo, owner, number, state)
}

func (g *Gogithub) PullRequestUpdate(repo, owner string, number int, state string) error {

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(g.Username),
		Password: strings.TrimSpace(g.Password),
	}

	client := github.NewClient(tp.Client())

	if owner == "" {
		owner = g.Username
	}

	pull := github.PullRequest{}

	pull.State = &state

	_, _, err := client.PullRequests.Edit(context.Background(), owner, repo, number, &pull)
	if err != nil {
		return err
	}

	return nil
}
