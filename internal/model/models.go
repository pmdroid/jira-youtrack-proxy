package model

type JiraCreateIssueRequest struct {
	Fields JiraFields `json:"fields"`
}

type RequestContext struct {
	YouTrackToken string
}

type JiraFields struct {
	Project     JiraProject  `json:"project"`
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	IssueType   JiraType     `json:"issuetype"`
	Priority    JiraPriority `json:"priority,omitempty"`
	Assignee    JiraUser     `json:"assignee,omitempty"`
}

type JiraType struct {
	Name string `json:"name"`
}

type JiraPriority struct {
	Name string `json:"name"`
}

type JiraUser struct {
	Name string `json:"name"`
}

type JiraResponse struct {
	Key  string `json:"key"`
	ID   string `json:"id"`
	Self string `json:"self"`
}

type JiraProject struct {
	Self        string `json:"self"`
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type YouTrackCreateIssueRequest struct {
	Summary      string                `json:"summary"`
	Description  string                `json:"description,omitempty"`
	Project      YouTrackProject       `json:"project"`
	CustomFields []YouTrackCustomField `json:"customFields,omitempty"`
}

type YouTrackProject struct {
	ID string `json:"id"`
}

type YouTrackCustomField struct {
	ID    string             `json:"id"`
	Type  string             `json:"$type"`
	Value YouTrackFieldValue `json:"value"`
}

type YouTrackFieldValue struct {
	Name string `json:"name"`
}

type YouTrackResponse struct {
	ID   string `json:"id"`
	Type string `json:"$type"`
}