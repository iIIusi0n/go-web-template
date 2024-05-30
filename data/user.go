package data

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	RealName    string `json:"real_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

func (m *manager) CreateUser(user *User) (int, error) {
	query := "INSERT INTO users (username, real_name, email, phone_number) VALUES (?, ?, ?, ?)"
	result, err := m.db.Exec(query, user.Username, user.RealName, user.Email, user.PhoneNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadUser(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := m.db.QueryRow(query, id)

	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.RealName, &user.Email, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *manager) UpdateUser(user *User) error {
	query := "UPDATE users SET username = ?, real_name = ?, email = ?, phone_number = ? WHERE id = ?"
	_, err := m.db.Exec(query, user.Username, user.RealName, user.Email, user.PhoneNumber, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
