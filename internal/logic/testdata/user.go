package testdata

import (
	gen2 "video_web/internal/logic/gen"
)

type userIn struct {
	Name         string
	IdentifyType int32
	Identifier   string
	Credential   string
}

type userExpected struct {
	RegisterCode int
	LoginCode    int
	Token        string
	UserId       int
}

var UserTests = []struct {
	In       *userIn
	Expected *userExpected
}{
	{
		In:       &userIn{Name: gen2.GetName(), IdentifyType: 1, Identifier: gen2.GetString(10) + "@qq.com", Credential: gen2.GetString(10)},
		Expected: &userExpected{RegisterCode: 200, LoginCode: 200},
	},
}
