package users

import (
	"database/sql"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Repository struct {
    DB *sql.DB
}

func (repo *Repository) GetAll() ([]User, error) {
    rows, err := repo.DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        var user User

        if err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
        ); err != nil {
            return nil, err
        }

        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (repo *Repository) Store(user *User) error {
    query := "INSERT INTO users (name, email) VALUES (?, ?)"
    res, err := repo.DB.Exec(query, user.Name, user.Email)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err == nil{
        user.ID = int(id)
    }
    return nil
}