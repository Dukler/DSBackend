package actions

func Dispatch (action string, payload map[string]interface{}, answer *interface{}){

	switch action {
	case "login":
		*answer = Login(payload)
	case "register":
		*answer = CreateUser(payload)
	default:
		break
	}
}
