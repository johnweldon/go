package auth

type UserRepo interface {
	User(username string) (AuthenticatedUser, bool)
	Save(username string, hash []byte) error
}

type DefaultUserRepo struct {
	users map[string][]byte
}

var _ UserRepo = (*DefaultUserRepo)(nil)

func NewUserRepo() UserRepo {
	return DefaultUserRepo{users: make(map[string][]byte)}
}

func LoadUserRepo(users map[string][]byte) UserRepo {
	return DefaultUserRepo{users: users}
}

func (r DefaultUserRepo) User(username string) (AuthenticatedUser, bool) {
	if hash, ok := r.users[username]; ok {
		return authUser{doc: authUserDoc{id: username, hash: hash}}, true
	}
	return nil, false
}

func (r DefaultUserRepo) Save(username string, hash []byte) error {
	r.users[username] = hash
	return nil
}
