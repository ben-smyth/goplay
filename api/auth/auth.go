package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func OidcCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	state := "random"
	if r.URL.Query().Get("state") != state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		http.Error(w, "failed to get userinfo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		OAuth2Token *oauth2.Token
		UserInfo    *oidc.UserInfo
	}{
		oauth2Token, userInfo,
	}

	// For demonstration purposes; in a real application, you'd probably use the user information more securely
	fmt.Fprintf(w, "UserInfo: %#v\n", resp)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}
