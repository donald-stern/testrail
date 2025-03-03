package testrail

import "fmt"

// CaseHistory represents the change history of a Test Case
type CaseHistory struct {
	History []UpdateHistory `json:"history"`
}

// UpdateHistory represents all changes saved at the same time
type UpdateHistory struct {
	ID        int                  `json:"id"`
	TypeID    int                  `json:"type_id"`
	CreatedOn int                  `json:"created_on"`
	UserID    int                  `json:"user_id"`
	Changes   []CaseHistoryChanges `json:"changes"`
	Comments  []string             `json:"comments"`
}

type CaseHistoryChanges struct {
	TypeID   int                      `json:"type_id"`
	OldText  string                   `json:"old_text"`
	NewText  string                   `json:"new_text"`
	Label    string                   `json:"label"`
	Options  CaseHistoryChangeOptions `json:"options"`
	Field    string                   `json:"field"`
	OldValue interface{}              `json:"old_value"`
	NewValue interface{}              `json:"new_value"`
}

type CaseHistoryChangeOptions struct {
	IsRequired   bool   `json:"is_required"`
	DefaultValue string `json:"default_value"`
}

// GetCaseHistory returns the history for an existing Test Case
func (c *Client) GetCaseHistory(caseID int) (CaseHistory, error) {
	returnCaseHistory := CaseHistory{}
	err := c.sendRequest("GET", fmt.Sprintf("get_history_for_case/%d", caseID), nil, &returnCaseHistory)
	return returnCaseHistory, err
}
