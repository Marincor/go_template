package entity

type MessageAttributes struct {
	Subject  string
	Message  string
	Template string
	Args     map[string]interface{}
}
