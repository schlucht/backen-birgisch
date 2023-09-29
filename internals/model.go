package internals

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

type User struct {
	Name     string
	Mail     string
	Username string
	Todos    []Todo
	Id       int
}

type Users []*User
	


func NewUsers() Users {
	var usrs Users
	
	return usrs
}

func (m *Users) AllUsers()Users {
	url := "https://jsonplaceholder.typicode.com/todos"
	user := "https://jsonplaceholder.typicode.com/users"
	
	req, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer req.Body.Close()

	var users []*User
	var todos []*Todo

	if err = json.NewDecoder(req.Body).Decode(&todos); err != nil {
		return nil
	}
	reqU, err := http.Get(user)
	if err != nil {
		return nil
	}
	if err = json.NewDecoder(reqU.Body).Decode(&users); err != nil {
		return nil
	}

	for _, usr := range users {
		for _, to := range todos {
			if to.UserId == usr.Id {
				usr.Todos = append(usr.Todos, *to)
			}
		}
	}
	return users 
}

func (m *Users) CompletedTodos()Users
