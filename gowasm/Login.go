package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

//{"_id":"5f32bea6d9bd97944bcb3cae","email":"mitch@gmail.com","firstName":"FISH123","lastName":"TOAD","password":"$argon2id$v=19$m=65536,t=1,p=4$hQjYT4if/EyrTmVxkL83Ng$wKPD3678H5MxABXzeSF4RmkdFvLCmuPaHS+3/Ek5Vi8","birthdate":"0001-01-01T00:00:00Z","nickname":"",
//"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1pdGNoQGdtYWlsLmNvbSIsImV4cCI6MTYxNDU2MDcxNX0.Xc-JNqznn0aSYZnhEnYvlDCRZS0EPhBaI0e-uv9hOIk"}

//Login : Login
type Login struct {
	User User
	vgrouter.NavigatorRef
}

//User : User
type User struct {
	ID        string `json:"_id"`
	Email     string `json:"email"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

var login Login

//HandleClick : handleclick
func (c *Login) HandleClick(event vugu.DOMEvent) {

	// c.PriceData = bpi{}

	ee := event.EventEnv()

	login.User.Email = c.User.Email
	login.User.Password = c.User.Password

	fmt.Printf("\n\n USER START:::: %+v", login.User)

	go func() {
		fmt.Printf("\n\n IN GOFUNC:::: %+v", login.User)

		values := map[string]string{
			"email":    login.User.Email,
			"password": login.User.Password}
		jsonData, err := json.Marshal(values)

		if err != nil {
			fmt.Printf("\n\n*** Marshal error: %v", err)
		}

		fmt.Printf("\n\n*** jsonData %+v", jsonData)

		//Leverage Go's HTTP Post function to make request
		resp, err := http.Post("http://192.168.86.45:5000/api/login",
			"application/json",
			bytes.NewBuffer(jsonData))

		//Handle Error

		if err != nil {
			fmt.Printf("\n\n*** Post error: %v", err)

		}

		fmt.Printf("\n\n*** RESP %+v", resp)
		//err = json.NewDecoder(resp.Body).Decode(&login.User)
		//if err != nil {
		//	fmt.Printf("\n\n*** Decode error: %v", err)
		//}

		defer resp.Body.Close()

		ee.Lock()
		defer ee.UnlockRender()
	}()

}
