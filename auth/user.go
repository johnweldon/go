package auth

type AuthenticatedUser interface {
	ID() string
	IsAuthenticated() bool
	Authenticate(password string) bool
}

type authUser struct {
	doc           authUserDoc
	authenticated bool
}

type authUserDoc struct {
	id   string `bson:"_id"`
	hash []byte
}

var _ AuthenticatedUser = (*authUser)(nil)

func (a authUser) ID() string                        { return a.doc.id }
func (a authUser) IsAuthenticated() bool             { return a.authenticated }
func (a authUser) Authenticate(password string) bool { return TestPassword(password, a.doc.hash) }
