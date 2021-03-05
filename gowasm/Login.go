package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"

	"github.com/vugu/vugu"
)

//{"_id":"5f32bea6d9bd97944bcb3cae","email":"mitch@gmail.com","firstName":"FISH123","lastName":"TOAD","password":"$argon2id$v=19$m=65536,t=1,p=4$hQjYT4if/EyrTmVxkL83Ng$wKPD3678H5MxABXzeSF4RmkdFvLCmuPaHS+3/Ek5Vi8","birthdate":"0001-01-01T00:00:00Z","nickname":"",
//"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1pdGNoQGdtYWlsLmNvbSIsImV4cCI6MTYxNDU2MDcxNX0.Xc-JNqznn0aSYZnhEnYvlDCRZS0EPhBaI0e-uv9hOIk"}

//Login : Login
type Login struct {
	ID        string `json:"_id"`
	Email     string `json:"email"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

var user Login

//HandleClick : handleclick
func (c *Login) HandleClick(event vugu.DOMEvent) {

	// c.PriceData = bpi{}

	ee := event.EventEnv()

	user.Email = c.Email
	user.Password = c.Password

	//fmt.Printf("\n\n USER  START:::: %+v", user)

	go func() {

		//Encode the data
		postBody, _ := json.Marshal(map[string]string{
			"email":    c.Email,
			"password": c.Password,
		})
		responseBody := bytes.NewBuffer(postBody)
		//Leverage Go's HTTP Post function to make request
		resp, err := http.Post("http://192.168.86.45:5000/api/login", "application/json", responseBody)
		//Handle Error
		if err != nil {
			fmt.Errorf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&user)
		if err != nil {
			fmt.Errorf("An Error Occured %v", err)
		}

		ee.Lock()
		defer ee.UnlockRender()
		//c.Username = username
		//c.Password = password

		fmt.Printf("\n\n json object:::: %+v", user)
		//if c.ValidateUser() {
		//	fmt.Printf("\n\n LOGGED IN")
		//c.LoginResponse = "LOGGED IN"
		//c.Navigate("/profile", nil)
		//} else {
		//	fmt.Printf("\n\n NOT LOGGED IN")
		//c.LoginResponse = "NOT LOGGED IN"
		//}

	}()
}
