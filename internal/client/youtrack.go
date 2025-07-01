package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pmdroid/jira-youtrack-proxy/internal/config"
	"github.com/pmdroid/jira-youtrack-proxy/internal/model"
)

func CreateYouTrackIssue(
	req *model.YouTrackCreateIssueRequest,
	requestCtx *model.RequestContext,
	cfg *config.Config,
) (*model.YouTrackResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling YouTrack request: %v", err)
	}

	youtrackURL := fmt.Sprintf("%s/api/issues", strings.TrimRight(cfg.YouTrackURL, "/"))
	httpReq, err := http.NewRequest("POST", youtrackURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", requestCtx.YouTrackToken))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request to YouTrack: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading YouTrack response: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("YouTrack API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var youtrackResponse model.YouTrackResponse
	if err := json.Unmarshal(respBody, &youtrackResponse); err != nil {
		return nil, fmt.Errorf("error parsing YouTrack response: %v", err)
	}

	return &youtrackResponse, nil
}