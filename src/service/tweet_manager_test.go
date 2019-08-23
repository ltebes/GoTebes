package service_test

import (
	"fmt"
	"testing"

	"github.com/ltebes/GoTebes/src/domain"
	"github.com/ltebes/GoTebes/src/service"
)

func TestPublichedTweetIsSaver(t *testing.T) {

	um := service.GetInstance()

	tm := service.NewTweetManager()

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
	tm.PublishTweet(tweet, um)

	// Validation
	publishedTweet := tm.GetTweet()
	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %v: %s \nbut is %v: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization

	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet *domain.Tweet

	var id int

	user := &domain.User{Name: ""}
	um.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet, um)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet *domain.Tweet

	var id int

	user := &domain.User{
		Name: "Lean",
	}
	um.Register(user)
	var text string
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet, um)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}
func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization

	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet *domain.Tweet

	var id int

	user := &domain.User{
		Name: "Lean",
	}
	um.Register(user)
	text := "This is my first tweetaaaaaaaaaaasasdasdasdasdasdasdasdadasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdadasdasdasdasdasdasdasdasdasdasdadgasdfawdhfjaldjkfhadhjsjhkfahjkfhkjfhkfadhklfadhkfdshjlkafdshjkfadh"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet, um)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text must not exceed 140 characters" {
		t.Error("Expected error is ext must not exceed 140 characters")
	}
}

func isValidTweet(t *testing.T, publishedTweet *domain.Tweet, id int, user *domain.User, text string) bool {
	if publishedTweet.User.Name == user.Name &&
		publishedTweet.Text == text &&
		publishedTweet.GetId() == id {
		return true
	}
	return false
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	um := service.GetInstance()

	tm := service.NewTweetManager()

	// service.InitializeService()          // initalaiz
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data

	var id int

	user := &domain.User{Name: "Lean"}
	um.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	secondUser := &domain.User{Name: "Lean"}
	um.Register(secondUser)
	secondText := "This is my second tweet"
	secondTweet = domain.NewTweet(id, secondUser, secondText)

	// Operation
	tm.PublishTweet(tweet, um)
	tm.PublishTweet(secondTweet, um)

	// Validation
	publishedTweets := tm.GetTweets()
	fmt.Print(publishedTweets)
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, id, secondUser, secondText) {
		return
	}
	// Same for secondPublishedTweet
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	// service.InitializeService()
	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet *domain.Tweet
	id := 1

	user := &domain.User{Name: "Lean"}
	um.Register(user)
	text := "This is my first tweet"
	tweet = domain.NewTweet(id, user, text)

	// Operation
	id, _ = tm.PublishTweet(tweet, um)

	// Validation
	publishedTweet := tm.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	// service.InitializeService()

	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	id := 1

	user := &domain.User{Name: "Lean"}
	user.SetId(1)
	um.Register(user)

	secondUser := &domain.User{Name: "Tebes"}
	secondUser.SetId(2)
	um.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(id, user, text)

	secondTweet = domain.NewTweet(id, user, secondText)

	thirdTweet = domain.NewTweet(id, secondUser, text)

	tm.PublishTweet(tweet, um)
	tm.PublishTweet(secondTweet, um)
	tm.PublishTweet(thirdTweet, um)

	// Operation
	count := tm.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	// service.InitializeService()
	um := service.GetInstance()

	tm := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	id := 1

	user := &domain.User{Name: "Lean"}
	user.SetId(1)
	um.Register(user)

	secondUser := &domain.User{Name: "Tebes"}
	secondUser.SetId(2)
	um.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(id, user, text)

	secondTweet = domain.NewTweet(id, user, secondText)

	thirdTweet = domain.NewTweet(id, secondUser, text)
	// publish the 3 tweets

	tm.PublishTweet(tweet, um)
	tm.PublishTweet(secondTweet, um)
	tm.PublishTweet(thirdTweet, um)

	// Operation
	tweets := tm.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
		t.Error("Expected asdasd")
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	isValidTweet(t, firstPublishedTweet, id, user, text)
	isValidTweet(t, secondPublishedTweet, id, user, secondText)

}
func TestCanGetAStringFromATweet(t *testing.T) {

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
	textToPrint := tweet.String()

	// Validation
	expectedText := fmt.Sprintf("@%v: %v", tweet.User.Name, tweet.Text)
	if textToPrint != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

    // Initialization
    tweet := domain.NewImageTweet("grupoesfera", "This is my image",
                "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
    // Operation
    text := tweet.PrintableTweet()
    // Validation
    expectedText := "@grupoesfera: This is my image
                    http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
    if text != expectedText {...}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
    // Initialization
    quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
    tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)
    // Validation
    expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
    if text != expectedText {...}
}
