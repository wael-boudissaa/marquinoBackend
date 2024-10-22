package user

import (
	"database/sql"
	"fmt"

	"github.com/wael-boudissaa/marquinoBackend/services/auth"
	"github.com/wael-boudissaa/marquinoBackend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(user types.UserLogin) (*types.User, error) {
	query := `Select * from profile where email = ? `
	rows, err := s.db.Query(query, user.Email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.Id == "" {
		return nil, fmt.Errorf("user not found ")
	}
	return u, nil
}

func (s *Store) GetUserById(user types.User) (*types.User, error) {
	query := `SELECT * FROM profile where idProfile= ?`
	rows, err := s.db.Query(query, user.Id)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (s *Store) CreateUser(user types.User, idUser string, token string, hashedPassword string) error {
	query := `INSERT INTO profile (idProfile, firstName, lastName, email, password, address, createdAt, lastLogin, refreshToken, type)
			  VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?)`

	_, err := s.db.Exec(query, idUser, user.FirstName, user.LastName, user.Email, hashedPassword, user.Address, token, user.Type)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	queryCustomer := `INSERT into customer(idCustomer, idProfile) values(?, ?)`
	idCustomer, err := auth.CreateAnId()
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	_, err = s.db.Exec(queryCustomer, idCustomer, idUser)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Address,

		&user.Phone,
		&user.CreatedAt,
		&user.Type,
		&user.LastLogin,
		&user.Refreshtoken,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
