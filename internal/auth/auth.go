package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func Init() error {
	ctxBg := context.Background()
	provider, err := oidc.NewProvider(ctxBg, "http://localhost:8081/")
	if err != nil {
		return err
	}
	config := oauth2.Config{
		ClientID:     "goPlay",
		ClientSecret: "27KljKqKrNEZD3BPsQDzsqt1SVctH6Gc",
		RedirectURL:  "http://localhost/oidc/redirect",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	oidcConfig := oidc.Config{
		ClientID: "goPlay",
	}

	fmt.Print(config)

	return nil
}
