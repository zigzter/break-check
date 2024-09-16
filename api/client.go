package api

import (
	"net/http"

	"github.com/google/go-github/v64/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

var ghClient *github.Client

func GetGHClient() *github.Client {
	if ghClient == nil {
		// TODO: Set proper path for cache
		transport := httpcache.NewTransport(diskcache.New("./cache/"))
		transport.MarkCachedResponses = true
		client := &http.Client{
			Transport: transport,
		}
		ghClient = github.NewClient(client)
	}
	return ghClient
}
