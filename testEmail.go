package main

import (
	"fmt"
	"io"
	"net/smtp"
	"text/template"

	"github.com/go-sql-driver/mysql"
	gomail "gopkg.in/gomail.v2"
)

//NewEmail struct
type NewEmail struct {
	from    string
	to      []string
	subject string
	body    string
}

//Email struct
type Email struct {
	Id                int            `json:"id"`
	ReservationTypeId int            `json:"reservation_type_id"`
	UserId            int            `json:"user_id"`
	Status            int            `json:"status"`
	Created_at        mysql.NullTime `json:"created_at"`
	// Updated_at            mysql.NullTime `json:"updated_at"`
	// Deleted_at            mysql.NullTime `json:"deleted_at"`
}

//NewEmail function to create new email
func (newEmail NewEmail) SendNewEmail() (bool, error) {
	auth := smtp.PlainAuth("", "rajkaran85@gmail.com", "hjofhireljfnxvum", "smtp.gmail.com")
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: Email testing!\n"
	msg := []byte(subject + mime + "\n" + "Ting Tong Ting Tong, email testingggg!!")
	addr := "smtp.gmail.com:587"
	to := []string{"karan957singh@gmail.com"}

	if err := smtp.SendMail(addr, auth, "gurdial957singh@gmail.com", to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "gurdial957singh@gmail.com")
	m.SetHeader("To", "richu.mahajan03@gmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "bhoolleeee bhandarri!")
	// m.SetBody("text/html", "Hello <b>bhraa</b> Jiii!!!")
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		t, err := template.ParseFiles("./template.html")
		fmt.Println("Hereeeeeeeeeeeeeeee222222222222222222", err)
		data := map[string]string{"title": "Mongambooooo!"}
		return t.Execute(w, data)
	})

	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "rajkaran85@gmail.com", "hjofhireljfnxvum")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email Sent!")
	// newEmail := &NewEmail{
	// 	to:      []string{"test"},
	// 	subject: "test",
	// 	body:    "test",
	// }
	// success, err := newEmail.SendNewEmail()
	// if success {
	// 	fmt.Println("Email Sent!")
	// } else {
	// 	fmt.Println("Email failed!", err)
	// }
}
