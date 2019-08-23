package service

import (
	"errors"

	"github.com/ltebes/GoTebes/src/domain"
)

//var Tweet *domain.Tweet
//var misTweets []*domain.Tweet

type TweetManager struct {
	misTweets []*domain.TextTweet
	Tweet     *domain.TextTweet
}

func NewTweetManager() *TweetManager {
	newManager := new(TweetManager)
	newManager.misTweets = make([]*domain.TextTweet, 0)
	return newManager
}

func (t *TweetManager) PublishTweet(tweet *domain.TextTweet, um *UserManager) (int, error) {

	if um.GetUser() == nil {
		return -1, errors.New("No hay usuarios registrados")
	}
	if tweet.User.Name == "" {
		return -1, errors.New("user is required")
	}
	if tweet.Text == "" {
		return -1, errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return -1, errors.New("text must not exceed 140 characters")
	}
	t.Tweet = tweet
	t.misTweets = append(t.misTweets, tweet)
	return tweet.GetId(), nil

}

func (t *TweetManager) GetTweet() *domain.TextTweet {
	return t.Tweet
}

func (t *TweetManager) GetTweets() []*domain.TextTweet {
	return t.misTweets
}

func (t *TweetManager) GetTweetById(id int) *domain.TextTweet {

	for _, tweet := range t.misTweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (t *TweetManager) CountTweetsByUser(user *domain.User) int {
	count := 0
	for _, tweet := range t.misTweets {
		if user.GetId() == tweet.User.GetId() {
			count++
		}
	}
	return count

}

func (t *TweetManager) GetTweetsByUser(user *domain.User) []*domain.TextTweet {
	tweets := make([]*domain.TextTweet, 0)

	for _, tweet := range t.misTweets {
		if user.GetId() == tweet.User.GetId() {
			tweets = append(tweets, tweet)
		}
	}
	return tweets
}
