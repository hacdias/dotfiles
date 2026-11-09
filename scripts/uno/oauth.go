package uno

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

func InteractiveLogin(config *oauth2.Config, port int) (*oauth2.Token, error) {
	ln, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	state := rand.Text()
	verifier := oauth2.GenerateVerifier()
	config.RedirectURL = fmt.Sprintf("http://localhost:%d/callback", port)

	url := config.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Please open the following URL, authenticate, and close the tab:\n%s\n", url)

	request := make(chan *http.Request, 1)

	handler := func(w http.ResponseWriter, r *http.Request) {
		request <- r

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html><html><head></head><body><h1>Please close this page and go back to the CLI.</h1></body></html>`))
	}

	server := &http.Server{
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		_ = server.Serve(ln)
	}()

	defer func() {
		_ = server.Shutdown(context.Background())
	}()

	r := <-request

	if s := r.URL.Query().Get("state"); s != state {
		return nil, errors.New("state does not match")
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("code was empty")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	return config.Exchange(ctx, code, oauth2.VerifierOption(verifier))
}
