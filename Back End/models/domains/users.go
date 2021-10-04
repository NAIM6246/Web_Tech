package domains

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	AddressID int
}

func UserTableName() string {
	return "users"
}
