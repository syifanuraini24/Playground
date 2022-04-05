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
	err = u.changeStatus(username, true)
	return &username, err
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
	return nil, errors.New("logged in user not found")
}

func (u *UserRepository) Logout(username string) error {
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
	for _, user := range users {
		if user.Username == username {
			user.Loggedin = status
			break
		}
	}

	// return Save()
	return u.Save(users)
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for _, user := range users {
		user.Loggedin = false
	}

	return u.Save(users)
}
