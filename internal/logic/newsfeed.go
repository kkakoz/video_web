package logic

import "context"

type newsfeed struct {
}

func Newsfeed() *newsfeed {
	return &newsfeed{}
}

func (newsfeed) Add(ctx context.Context) {

}
