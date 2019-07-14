package actions

func Dispatch (action string, payload map[string]interface{}) *interface{}{

	switch action {
	case "login":
		return Login(payload)
	case "register":
		return CreateUser(payload)
	default:
		return nil
	}
}
