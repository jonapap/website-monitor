package mail

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"

	"gopkg.in/ini.v1"
)

//Message represents a simple email message with a header
//and body to be sent
type Message struct {
	subject string
	body    string
}

//SetSubject sets the message's subject to s
func (m *Message) SetSubject(s string) {
	m.subject = s
}

//AddLineToBody adds a new line with the message s in the body
func (m *Message) AddLineToBody(s string) {
	m.body += "\n" + s
	ini.Load("my.ini")
}

//Send sends the messsage using the settings set in the configuration
//file config/mail.ini
func (m *Message) Send() error {
	values := map[string]string{"from": "", "to": "", "password": "", "host": "", "port": "", "username": ""}
	if os.Getenv("InsideDockerContainer") != "True" {
		config, err := ini.Load("config/email.ini")

		var e *os.PathError
		if errors.As(err, &e) {
			if err := createEmailConfig(); err != nil {
				return fmt.Errorf("Error creating config file: %v", err)
			}
			return ConfigFileDidNotExistError{} //Return an error so the user can modify the default values
		} else if err != nil {
			return fmt.Errorf("Error opening config file: %v", err)
		}
		section := config.Section("")
		for k := range values {
			values[k] = section.Key(k).String()
		}
	} else { //We are inside a docker container, use the environment variables
		for k := range values {
			values[k] = os.Getenv("MAIL_" + k)
		}
	}
	message := "Subject: " + m.subject + "\n\n" + m.body

	auth := smtp.PlainAuth("", values["username"], values["password"], values["host"])
	address := values["host"] + ":" + values["port"]
	return smtp.SendMail(address, auth, values["from"], []string{values["to"]}, []byte(message))
}

//ConfigFileDidNotExistError is thrown if the config file did not exist initially
//when the function was called, and it has now been created and the user should
//modify the default values
type ConfigFileDidNotExistError struct{}

func (e ConfigFileDidNotExistError) Error() string {
	return "Config file did not exist, so it has been created for you. Please modify the default values."
}

func createEmailConfig() error {
	file := ini.Empty()
	section := file.Section("")

	section.Key("from").SetValue("example@example.com")
	section.Key("to").SetValue("example@example.com")

	section.Key("username").SetValue("example@example.com")
	section.Key("password").SetValue("example@example.com")
	section.Key("host").SetValue("example.com")
	section.Key("port").SetValue("587")

	if err := os.MkdirAll("config", 0644); err != nil {
		return fmt.Errorf("Error creating config directory: %w", err)
	}
	err := file.SaveTo("config/email.ini")
	if err != nil {
		return err
	}

	return nil
}
