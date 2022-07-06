package api

import "time"

const (
	MailboxWorkflowTypeAutomatic = "automatic"
	MailboxWorkflowTypeManual    = "manual"

	MailboxWorkflowStatusActive   = "active"
	MailboxWorkflowStatusInactive = "inactive"
	MailboxWorkflowStatusInvalid  = "invalid"
)

type Workflow struct {
	Id         int
	MailboxId  int
	TypeOf     string
	Status     string
	Order      int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// Creates our template Verify, send, and archive [VSA] rule to add to the central mailbox
func CreateVSAWorkflow(customerName string, forwardToList []string) Workflow {
	var newWorkflow Workflow

	// [VSA] - <insert_client_mailbox_name>
	//
	// IF
	//		customer name contains <insert_name>
	//	AND Conversation.tag contains "unprocessed"
	//
	// THEN
	// 		forward conversation to <insert_forwardTo_List>
	//	AND Add conversation.tag "valid","sent"
	//	AND Remove conversation.tag "unprocessed"
	//	AND Move to mailbox <insert_client_mailbox_name>
	//	AND Change conversation.status to "Closed"
	//
	// Check "Apply to Previous"

	return newWorkflow
}

// Calls the List Workflows route, returns an array of Workflows
// See https://developer.helpscout.com/mailbox-api/endpoints/workflows/list/
func ListWorkflows() []Workflow {
	return nil
}

// Calls the Update Workflows route to update all given workflows
// See https://developer.helpscout.com/mailbox-api/endpoints/workflows/update/
func UpdateWorkflows([]Workflow) {

	// The goal is to be able to use this in some way to add a workflow rule to an existing mailbox whenever we want

}
