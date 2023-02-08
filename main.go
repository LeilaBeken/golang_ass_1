package main
import "fmt"
type Item struct {   
	Name   string
    Price  float64   
    Rating float64
}
type User struct {   
	Name     string
    Password string}
type ItemService struct {
   Items []Item
}
func (s *ItemService) RegisterItem(item Item) {
	s.Items = append(s.Items, item)
}
func (s *ItemService) SearchByName(name string) []Item {
   result := make([]Item, 0)   
   for _, item := range s.Items {
	if item.Name == name {result = append(result, item)}
   }   
   return result
}
func (s *ItemService) FilterByPrice(price float64) []Item {   
	result := make([]Item, 0)
	for _, item := range s.Items {      
		if item.Price <= price {result = append(result, item)}
	}   
    return result
}
func (s *ItemService) FilterByRating(rating float64) []Item {   
	result := make([]Item, 0)
    for _, item := range s.Items {      
	if item.Rating >= rating {result = append(result, item)}
    }   
	return result
}
func (s *ItemService) GiveRating(item Item, rating float64) {   
	for i, it := range s.Items {
    if it.Name == item.Name {   
		s.Items[i].Rating = rating
        break
	}
}}

type UserService struct {
	Users []User
}

func (s *UserService) RegisterUser(user User) {
	s.Users = append(s.Users, user)
}
func (s *UserService) Authorize(name, password string) bool {
   for _, user := range s.Users {      
	if user.Name == name && user.Password == password {return true}
   }   
   return false
}
func main() {   
	fmt.Println()
}