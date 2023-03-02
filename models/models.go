package models

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`
}

func (u *User) GetByUsername(username string) error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Where("username = ?", username).First(u)
    return result.Error
}

func (u *User) Create() error {
	db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Create(u)
    return result.Error
}

func (u *User) Update() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Save(u)
    return result.Error
}

func (u *User) Delete() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Delete(u)
    return result.Error
}


type Product struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int    `json:"price"`
    Category    string `json:"category"`
    Rating        int    `json:"rating"`
}

func (p *Product) GetByID(id uint) error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.First(p, id)
    return result.Error
}

func (p *Product) Create() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Create(p)
    return result.Error
}

func (p *Product) Update() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Save(p)
    return result.Error
}

func (p *Product) Delete() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Delete(p)
    return result.Error
}