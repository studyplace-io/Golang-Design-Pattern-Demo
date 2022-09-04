package IteratorPattern

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
