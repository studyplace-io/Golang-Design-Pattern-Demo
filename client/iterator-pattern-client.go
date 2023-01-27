package main

import (
	"fmt"
	"golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Behavioral/IteratorPattern"
)




func main() {

	user1 := &IteratorPattern.User{
		Name: "a",
		Age:  30,
	}
	user2 := &IteratorPattern.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &IteratorPattern.UserCollection{
		Users: []*IteratorPattern.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
