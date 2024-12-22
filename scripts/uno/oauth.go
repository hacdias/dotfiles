package uno

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(length int) (string, error) {
	b := make([]byte, length)

	_, err := seededRand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func InteractiveLogin(oauth2 *oauth2.Config, port int) (*oauth2.Token, error) {
	state, err := randomString(10)
	if err != nil {
		return nil, err
	}

	oauth2.RedirectURL = fmt.Sprintf("http://localhost:%d/callback", port)

	url := oauth2.AuthCodeURL(state)
	fmt.Printf("Please open the following URL, authenticate, and close the tab:\n%s\n", url)

	request := make(chan *http.Request, 1)

	handler := func(w http.ResponseWriter, r *http.Request) {
		request <- r

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html><html><head></head><body><h1>Please close this page and go back to the CLI.</h1></body></html>`))
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		_ = server.ListenAndServe()
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

	return oauth2.Exchange(ctx, code)
}
