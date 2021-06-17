package model

type ClickUpTaskRequest struct {
	Title       string   `form:"title"`
	Description string   `form:"description"`
	Tags 		string   `form:"tags"`
	Priority  	int      `form:"priority"`
}

type ClickUpTask struct {
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Assignees      []int      `json:"assignees"`
	Priority	   int		  `json:"priority"`
	Tags		   []string	  `json:"tags"`
}

type TagsList struct {
	BaseResponse
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Name    string `json:"name"`
	TagFG   string `json:"tag_fg"`
	TagBG   string `json:"tag_bg"`
	Creator int    `json:"creator"`
}