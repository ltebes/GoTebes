package domain_test

import (
	"fmt"
	"testing"

	"github.com/ltebes/GoTebes/src/domain"
	"github.com/ltebes/GoTebes/src/service"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	um := service.GetInstance()

	// Initialization
	var tweet *domain.Tweet

	var id int

	user := &domain.User{
		Name: "Lean",
	}
	um.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	textToPrint := tweet.PrintableTweet()

	// Validation
	expectedText := fmt.Sprintf("@%v: %v", tweet.User.Name, tweet.Text)
	if textToPrint != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
