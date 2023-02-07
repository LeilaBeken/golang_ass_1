package pck

type Item struct {
    Name string
    Price float64
    Rating float64
}

type User struct {
    Name string
    Password string
    Items []Item
}

type Database struct {
    Users []User
	Items []Item
}