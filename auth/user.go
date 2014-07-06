package auth

type AuthenticatedUser interface {
	Id() string
	IsAuthenticated() bool
	Authenticate(password string) bool
}

type AuthUser struct {
	doc           authUserDoc
	authenticated bool
}

type authUserDoc struct {
	id   string `bson:"_id"`
	hash []byte
}

var _ AuthenticatedUser = (*AuthUser)(nil)

func (a AuthUser) Id() string                        { return a.doc.id }
func (a AuthUser) IsAuthenticated() bool             { return a.authenticated }
func (a AuthUser) Authenticate(password string) bool { return TestPassword(password, a.doc.hash) }
