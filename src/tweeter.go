package main

import (
	"github.com/abiosoft/ishell"
	"github.com/ltebes/GoTebes/src/domain"
	"github.com/ltebes/GoTebes/src/service"
)

func main() {

	service.InitializeService()

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Register user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Name: ")
			name := c.ReadLine()
			c.Print("Pass: ")
			pass := c.ReadLine()
			c.Print("Email: ")
			email := c.ReadLine()
			c.Print("Nick: ")
			nick := c.ReadLine()
			user := domain.NewUser(name, email, pass, nick)
			service.Register(user)
			c.Print("Usuario registrado\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			user := service.GetUser()

			// c.Print("Write your userName: ")
			// user := c.ReadLine()

			if user == nil {
				c.Print("No hay usuario registrado - Utilice 'login' para iniciar sesion\n")
				return
			}

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			err := service.PublishTweet(tweet)

			if err == nil {
				c.Print("Tweet send\n")
			} else {
				c.Print(err)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.Run()

}
