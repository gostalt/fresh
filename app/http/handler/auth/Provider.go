package auth

import (
	"errors"
	"gostalt/app/entity"
	"gostalt/app/entity/user"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
	"golang.org/x/crypto/bcrypt"
)

type Provider struct {
	store *sessions.CookieStore
}

func NewProvider(store *sessions.CookieStore) Provider {
	return Provider{
		store: store,
	}
}

// DefaultRedirect is the path to redirect to when a user
// successfully logs in or registers a new account.
func (p Provider) DefaultRedirect() string {
	return "/home"
}

// ProcessLogin persists the user details to the session.
func (p Provider) ProcessLogin(w http.ResponseWriter, r *http.Request, user interface{}) error {
	session, err := p.store.Get(r, "gostalt")
	if err != nil {
		return err
	}

	session.Values["user"] = user
	session.Save(r, w)

	return nil
}

// CreateUser ...
func (p Provider) CreateUser(r *http.Request) (*entity.User, error) {
	client := di.Get(r, "entity-client").(*entity.Client)

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	u, err := client.User.Create().SetUsername(username).SetPassword(encrypted).Save(r.Context())
	if u == nil || err != nil {
		return &entity.User{}, errors.New("unable to create user")
	}

	return u, nil
}

// RetrieveUser attempts to find the database record for the provided
// login credentials. If the record cannot be found, or if the
// password is incorrect for the user, an error is returned.
func (p Provider) RetrieveUser(r *http.Request) (*entity.User, error) {
	client := di.Get(r, "entity-client").(*entity.Client)

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	u, err := client.User.Query().Where(user.UsernameEQ(username)).First(r.Context())
	if err != nil {
		return &entity.User{}, errors.New("user does not exist")
	}

	if err := bcrypt.CompareHashAndPassword(u.Password, []byte(password)); err != nil {
		return &entity.User{}, errors.New("invalid password")
	}

	return u, nil
}
