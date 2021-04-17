package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v35/github"
	"golang.org/x/mod/modfile"
	"golang.org/x/oauth2"
)

var (
	modFilePath = flag.String("p", "./go.mod", "Tha path to go.mod")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	dat, err := ioutil.ReadFile(*modFilePath)
	if err != nil {
		log.Fatal("cannot find go.mod...")
	}

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("please set GITHUB_AUTH_TOKEN to your env")
	}

	m, err := modfile.Parse("go.mod", dat, nil)
	if err != nil {
		log.Fatal("cannot parse go.mod...")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	for _, r := range m.Require {
		modpath := r.Mod.Path
		if !strings.HasPrefix(modpath, "github.com") {
			continue
		}
		if err := star(ctx, client, modpath); err != nil {
			log.Fatalf("failed to star, reason: %v", err.Error())
		}
	}
}

// star stars to the repository
func star(ctx context.Context, client *github.Client, modpath string) error {
	repo := strings.Split(modpath[11:], "/")
	resp, err := client.Activity.Star(ctx, repo[0], repo[1])
	if err != nil {
		return fmt.Errorf("request to github for star: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("response from github is not 204")
	}
	log.Printf("Starred! %s \n", modpath)
	return nil
}
