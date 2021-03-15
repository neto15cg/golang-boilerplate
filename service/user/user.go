package user

type UserType struct {
	UserId string	`json:"userId"`
	Name   string	`json:"name"`

}


type Getter struct {}

func (g *Getter) Run() (*UserType, error) {
	 user := UserType{
		 UserId: "123",
		 Name: "João das Couves",
	 }

	 return &user, nil
}