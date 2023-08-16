package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sethgrid/pester"
)

func fetchURL(ctx context.Context, client *pester.Client, url string, r interface{}) error {
	// log.Printf("Fetching %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", "nullpraisal/v3-ccpfix peer@nullsec.au")
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
	case 404:
		return nil
	default:
		return fmt.Errorf("Error talking to esi: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(r)
	defer resp.Body.Close()
	return err
}
