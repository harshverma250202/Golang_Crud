package repository

import (
	"errors"
	"log"

	"github.com/harsh/project/internal/db"
)



type Repository struct {
}

var (
	Users Repository = Repository{}
)

func (r Repository) FindAll() ([]User, error) {
	users := make([]User, 0)
	rows, err := db.Client.Query("select id, name, email, password from users")
	if err != nil {
		log.Println("error while fetching from database", err.Error())
		return nil, errors.New("error while fetching from database")
	}

	for rows.Next() {
		c := new(User)
		err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.Password)
		if err != nil {
			log.Println("error while scanning from rows", err.Error())
			return nil, errors.New("error while scanning users")
		}
		users = append(users, *c)
	}
	return users, nil
}

func (r Repository) FindById(id int) (User, error) {
	user := User{}
	err := db.Client.QueryRow("select id, name, email, password from users where id=$1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("error while fetching from database", err.Error())
		return User{}, errors.New("error while fetching from database")
	}
	return user, nil
}
func (r Repository) Create(user User) error {
	// id := rand.Intn(1000000)
	_, err := db.Client.Exec("insert into users (name, email, password) values ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("error while inserting into database", err.Error())
		return errors.New("error while inserting into database")
	}
	return nil
}

func (r Repository) Update(user User, id int) error {
	olduser, err := r.FindById(id)
	if err != nil {
		return err
	}
	olduser.Name = user.Name
	olduser.Email = user.Email
	olduser.Password = user.Password
	_, err = db.Client.Exec("update users set name=$1, email=$2, password=$3 where id=$4", olduser.Name, olduser.Email, olduser.Password, id)
	if err != nil {
		log.Println("error while updating into database", err.Error())
		return errors.New("error while updating into database")
	}
	return nil

}

func (r Repository) Delete(id int) error {
	_, err := db.Client.Exec("delete from users where id=$1", id)
	if err != nil {
		log.Println("error while deleting from database", err.Error())
		return errors.New("error while deleting from database")
	}
	return nil
}
