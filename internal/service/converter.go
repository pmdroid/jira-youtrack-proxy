package service

import (
	"github.com/pmdroid/jira-youtrack-proxy/internal/config"
	"github.com/pmdroid/jira-youtrack-proxy/internal/model"
)

func ConvertJiraToYouTrack(jiraReq model.JiraCreateIssueRequest, cfg *config.Config) (*model.YouTrackCreateIssueRequest, error) {
	youtrackReq := &model.YouTrackCreateIssueRequest{
		Summary:     jiraReq.Fields.Summary,
		Description: jiraReq.Fields.Description,
		Project: model.YouTrackProject{
			ID: jiraReq.Fields.Project.Key,
		},
		CustomFields: []model.YouTrackCustomField{},
	}

	if jiraReq.Fields.IssueType.Name != "" {
		if typeFieldID, exists := cfg.TypeFieldMap[jiraReq.Fields.Project.Key]; exists {
			youtrackReq.CustomFields = append(youtrackReq.CustomFields, model.YouTrackCustomField{
				ID:   typeFieldID,
				Type: "SingleEnumIssueCustomField",
				Value: model.YouTrackFieldValue{
					Name: mapIssueType(jiraReq.Fields.IssueType.Name),
				},
			})
		}
	}

	if jiraReq.Fields.Priority.Name != "" {
		if priorityFieldID, exists := cfg.PriorityFieldMap[jiraReq.Fields.Project.Key]; exists {
			youtrackReq.CustomFields = append(youtrackReq.CustomFields, model.YouTrackCustomField{
				ID:   priorityFieldID,
				Type: "SingleEnumIssueCustomField",
				Value: model.YouTrackFieldValue{
					Name: mapPriority(jiraReq.Fields.Priority.Name),
				},
			})
		}
	}

	if jiraReq.Fields.Assignee.Name != "" {
		if assigneeFieldID, exists := cfg.AssigneeFieldMap[jiraReq.Fields.Project.Key]; exists {
			youtrackReq.CustomFields = append(youtrackReq.CustomFields, model.YouTrackCustomField{
				ID:   assigneeFieldID,
				Type: "SingleUserIssueCustomField",
				Value: model.YouTrackFieldValue{
					Name: jiraReq.Fields.Assignee.Name,
				},
			})
		}
	}

	return youtrackReq, nil
}

func mapIssueType(jiraType string) string {
	typeMap := map[string]string{
		"Task":    "Task",
		"Story":   "Feature",
		"Epic":    "Epic",
		"Feature": "Feature",
	}

	if mapped, exists := typeMap[jiraType]; exists {
		return mapped
	}

	return "Task"
}

func mapPriority(jiraPriority string) string {
	priorityMap := map[string]string{
		"Highest": "Critical",
		"High":    "Major",
		"Medium":  "Normal",
		"Low":     "Minor",
		"Lowest":  "Minor",
	}

	if mapped, exists := priorityMap[jiraPriority]; exists {
		return mapped
	}

	return "Normal"
}