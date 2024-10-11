package notification

type NotificationMessage struct{
	Message 	string		`json:"message"`
}

type SendNotificationRequest struct{
	User_id		string		`json:"user_id"`
	Message 	string		`json:"message"`
}

type SendNotificationResponse struct{
	Success 	bool 		`json:"success"`
}

type SubscribeRequest struct{
	User_id 	string		`json:"user_id"`
}