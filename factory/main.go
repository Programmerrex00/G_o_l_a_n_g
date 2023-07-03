package main

import "fmt"

//Solo notificacion SMS e Email

// Interfece principal
type INotificationFactory interface {
	sendNotification()
	GetSender() ISender
}

//Interface para ISender
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//structs para los SMS
type SMSNotification struct {
}

//Metodo de la interface principal para los SMS
func (SMSNotification) sendNotification() {
	fmt.Println("Sending Notification SMS")
}

//Metodo para la interface principal para los SMS, retorna un SMSNotificationSender, que es el struct, que implementa la interface ISender que le da vida a dos metodos de esa interface
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// struct para implementar los metodos de la interface  ISender
type SMSNotificationSender struct {
}

//metodos de la interface ISender siendo implementados por SMSNotificationSender
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

//Para Email
type EmailNotification struct {
}

func (EmailNotification) sendNotification() {
	fmt.Println("Sending Notification Via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

//Ahora crearemos una funcion para que retorne las clases para enviar las notificaciones, retornara la interface principal o un error
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type!!")
}

//Funciones para enviar la notification y resivirla
func sendNotification(f INotificationFactory) {
	f.sendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getChannel(smsFactory)
	fmt.Println("")
	getMethod(emailFactory)
	getChannel(emailFactory)
}
