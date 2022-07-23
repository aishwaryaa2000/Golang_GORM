package notification

type Notification struct{
	msg string
}

func (n Notification)SendCreditNotification() string{
	n.msg = "Amount has been credited.Notifcation sent"
	return n.msg
}

func (n Notification)SendDebitNotification() string{ 
	n.msg = "Amount has been debited.Notifcation sent"
	return n.msg
}