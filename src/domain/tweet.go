package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	id   int
	User *User
	Text string
	Date *time.Time
}

type Tweet interface {
	GetId() int
	GetName() string
	GetText() string
	GetQuote() Tweet
	GetUrlImage() string
	GetDate() *time.Time
}

// hay que definir los metodos sobre tweetmanager

type ImageTweet struct {
	TextTweet
	UrlPath string
}

type QuoteTweet struct {
	TextTweet
	EmbebidedTweet Tweet
}

func NewTweet(id int, user *User, text string) *TextTweet {

	date := time.Now()

	tweet := TextTweet{
		id,
		user,
		text,
		&date,
	}
	return &tweet
}

func NewImageTweet(id int, user *User, text string, UrlPath string) *ImageTweet {

	date := time.Now()

	tweet := ImageTweet{
		TextTweet{
			id,
			user,
			text,
			&date,
		},
		UrlPath,
	}
	return &tweet
}

func NewQuoteTweet(id int, user *User, text string, EmbebidedTweet Tweet) *QuoteTweet {

	date := time.Now()

	tweet := QuoteTweet{
		TextTweet{
			id,
			user,
			text,
			&date,
		},
		EmbebidedTweet,
	}
	return &tweet
}

func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%v: %v", t.User.Name, t.Text)
}

func (t *TextTweet) String() string {
	return fmt.Sprintf("@%v: %v", t.User.Name, t.Text)
}

func (t *TextTweet) GetId() int {
	return t.id
}
