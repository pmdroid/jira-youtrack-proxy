package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/pmdroid/jira-youtrack-proxy/internal/client"
	"github.com/pmdroid/jira-youtrack-proxy/internal/config"
	"github.com/pmdroid/jira-youtrack-proxy/internal/model"
	"github.com/pmdroid/jira-youtrack-proxy/internal/service"
)

func HandleCreateIssue(c echo.Context, cfg *config.Config) error {
	requestCtx := &model.RequestContext{}

	auth := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Basic ") {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Basic Authentication required (email:token)",
		})
	}

	encoded := strings.TrimPrefix(auth, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Basic Auth encoding",
		})
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Basic Auth format (expected email:token)",
		})
	}

	requestCtx.YouTrackToken = credentials[1]

	var jiraReq model.JiraCreateIssueRequest
	if err := c.Bind(&jiraReq); err != nil {
		log.
			Error().
			Err(err).
			Msg("Error parsing Jira request")

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	youtrackRequest, err := service.ConvertJiraToYouTrack(jiraReq, cfg)
	if err != nil {
		log.
			Error().
			Err(err).
			Msg("Error converting request")

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	log.
		Debug().
		Interface("request", youtrackRequest).
		Msg("YouTrack Request")

	youtrackResponse, err := client.CreateYouTrackIssue(youtrackRequest, requestCtx, cfg)
	if err != nil {
		log.
			Error().
			Err(err).
			Msg("Error creating YouTrack issue")

		return c.NoContent(http.StatusInternalServerError)
	}

	log.
		Debug().
		Interface("response", youtrackResponse).
		Msg("YouTrack Response")

	jiraResponse := model.JiraResponse{
		Key:  youtrackResponse.ID,
		ID:   youtrackResponse.ID,
		Self: fmt.Sprintf("%s/api/issues/%s", cfg.YouTrackURL, youtrackResponse.ID),
	}

	return c.JSON(http.StatusCreated, jiraResponse)
}

func HandleProjectDetails(c echo.Context, cfg *config.Config) error {
	projectID := c.Param("id")
	if projectID == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	log.
		Debug().
		Str("projectID", projectID).
		Msg("Returning fake project details for project ID")

	host := c.Request().Host
	project := model.JiraProject{
		Id:          projectID,
		Key:         projectID,
		Name:        fmt.Sprintf("Project %s", projectID),
		Self:        fmt.Sprintf("https://%s/rest/api/2/project/%s", host, projectID),
		Description: fmt.Sprintf("Fake project details for YouTrack project %s", projectID),
	}

	return c.JSON(http.StatusOK, project)
}