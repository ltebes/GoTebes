package domain

type User struct {
	id    int
	Name  string
	email string
	pass  string
	nick  string
}

func NewUser(id int, name, email, pass, nick string) *User {

	user := User{
		id,
		name,
		email,
		pass,
		nick,
	}
	return &user
}

func (t *User) GetId() int {
	return t.id
}

func (t *User) SetId(id int) {
	t.id = id
}
