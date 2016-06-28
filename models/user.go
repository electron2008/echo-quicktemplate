package models

type User struct {
	Id        int    `db:"bigserial primary key" json:"id"`
	FirstName string `db:"varchar(80) NOT NULL" json:"first_name"`
	LastName  string `db:"varchar(80) NOT NULL" json:"last_name"`
	Email     string `db:"varchar(80) NOT NULL UNIQUE" json:"email"`
}
