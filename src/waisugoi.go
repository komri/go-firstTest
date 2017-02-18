package main

import "fmt"

type Friends struct {
	skill string
}

func newFriends(skill string) *Friends {
	friends := new(Friends)
	friends.skill = skill
	return friends
}

func (f Friends) wai() string {
	return "わーい\n"
}

func (f Friends) sugoi() string {
	return "すごーい\n"
}

func (f Friends) tanoshii() string {
	return "たのしー\n"
}

func (f Friends) beFriend() string {
	return f.wai() + f.sugoi() + f.tanoshii() + "君は" + f.skill + "がとくいな、フレンズなんだね!!" + "\n"
}

func main() {
	var friends *Friends = newFriends("クソリプ")
	fmt.Println(friends.beFriend())
}
