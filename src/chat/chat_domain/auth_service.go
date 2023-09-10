package chat_domain

type AuthService interface {
	SignIn(username string, password string) (*User, string, error)
	SignUp(user User, userCredentials UserCredentials) (*User, string, error)
	SignOut(session Session) error
}
