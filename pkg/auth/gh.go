package auth

import (
	"context"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func GitHub(token string) (client *github.Client, err error,ctx context.Context) {
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)

	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

	//test to check if token is valid
	_, _, err = client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil,err,nil
	}
	return client,nil,ctx
}