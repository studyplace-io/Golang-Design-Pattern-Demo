package IteratorPattern

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(t *testing.T) {

	user1 := NewUser("aaa", 10)
	user2 := NewUser("bbb", 20)

	userCollection := &UserCollection{
		Users: []*User{user1, user2},
	}

	// 返回迭代器对象
	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

}
