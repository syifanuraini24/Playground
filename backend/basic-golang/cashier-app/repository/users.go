package repository

import (
	"errors"
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedin, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedin,
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	// 1. load data, check error
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	// 2. loop data => user.username == username && user.password == password => existingUser
	var existingUser *User
	for _, user := range users {
		if user.Username == username && user.Password == password {
			if user.Loggedin {
				return &username, nil
			}
			existingUser = &user
			break
		}
	}
	// 3. if user not exists => return nil, errors.New("unknown user")
	if existingUser == nil {
		return nil, errors.New("Login Failed")
	}
	// 4. if user exists, logout all, changeStatus(username, true)
	u.LogoutAll()
	u.changeStatus(username, true)
	return &username, nil
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	// load data
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	// loop data, user.loggedin == true
	for _, user := range users {
		if user.Loggedin {
			// return user.name, nil
			return &user.Username, nil
		}
	}
	return nil, errors.New("no user is logged in")
}

func (u *UserRepository) Logout(username string) error {
	users, err := u.SelectAll()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Username == username {
			if !u.Loggedin {
				return errors.New("unauthorized")
			}
			break
		}
	}

	// changeStatus(username, false)
	return u.changeStatus(username, false)
}

func (u *UserRepository) Save(users []User) error {
	records := [][]string{
		{"username", "password", "loggedin"},
	}
	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
		})
	}
	return u.db.Save("users", records)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}
	// search user with specified username
	newUsers := users
	for i, user := range users {
		if user.Username == username {
			user.Loggedin = status
			newUsers[i] = user
			break
		}
	}

	// return Save()
	return u.Save(newUsers)
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	newUsers := make([]User, len(users))
	for i, user := range users {
		user.Loggedin = false
		newUsers[i] = user
	}

	return u.Save(newUsers)
}
