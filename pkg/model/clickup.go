package model

type ClickUpTaskRequest struct {
	Title       string   `form:"title"`      // task name
	Description string   `form:"description"`       // custom field email
	Tags 		[]string `form:"tags"` // task description
	Priority  	int      `form:"priority"`
}

type ClickUpTaskRequest2 struct {
	Issuer      string `form:"issuer"`      // task name
	Email       string `form:"email"`       // custom field email
	Description string `form:"description"` // task description
	FolderName  string `form:"folderName"`
}

type ClickUpTask struct {
	Name                      string                   `json:"name"`
	Description               string                   `json:"description"`
	Assignees                 []int                    `json:"assignees"`
	Priority			 	  int					   `json:"priority"`
	Tags				  	  []string				   `json:"tags"`
	Status                    string                   `json:"status"`
	DueDate                   int64                    `json:"due_date"`
	DueDateTime               bool                     `json:"due_date_time"`
	StartDate                 int64                    `json:"start_date"`
	StartDateTime             bool                     `json:"start_date_time"`
	NotifyAll                 bool                     `json:"notify_all"`
	CheckRequiredCustomFields bool                     `json:"check_required_custom_fields"`
	CustomFields              []ClickUpTaskCustomField `json:"custom_fields"`
}

type ClickUpTaskCustomField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type TagsList struct {
	BaseResponse
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
	TagFG string `json:"tag_fg"`
	TagBG string `json:"tag_bg"`
	Creator int `json:"creator"`
}