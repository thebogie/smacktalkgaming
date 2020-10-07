package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thebogie/smacktalkgaming/config"
	"github.com/thebogie/smacktalkgaming/services"
	"github.com/thebogie/smacktalkgaming/types"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserController interface
type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	UpdateUser(*gin.Context)
	GetUser(*gin.Context)
	GetUserStats(*gin.Context)
}

type userController struct {
	us      services.UserService
	cs      services.ContestService
	pwdhash types.PasswordConfig
}

// NewUserController instantiates User Controller
func NewUserController(
	us services.UserService, cs services.ContestService) UserController {
	return &userController{
		us: us,
		cs: cs,
		pwdhash: types.PasswordConfig{
			config.Config.Password.Time,
			config.Config.Password.Memory,
			config.Config.Password.Threads,
			config.Config.Password.Keylen},
	}
}

// @Summary Register new user
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/register [post]
func (ctl *userController) Register(c *gin.Context) {
	var rawStrings map[string]interface{}
	var player types.User

	//var jsonData map[string]interface{} // map[string]interface{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	if e := json.Unmarshal(data, &rawStrings); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	for key, value := range rawStrings {

		config.Apex.Debugf("%q is a string: %q", key, value)

		if key == "email" {
			player.Email = value.(string)
		}
		if key == "password" {
			hashed, err := services.GeneratePassword(&ctl.pwdhash, value.(string))
			if err != nil {
				config.Apex.Errorf("%s", err)
				return
			}
			player.Password = hashed
		}

	}
	ctl.us.AddUser(&player)

}

func (ctl *userController) Login(c *gin.Context) {
	var rawStrings map[string]interface{}
	var player types.User
	var attemptedpassword string

	//var jsonData map[string]interface{} // map[string]interface{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	if e := json.Unmarshal(data, &rawStrings); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	for key, value := range rawStrings {

		//config.Apex.Infof("%q is a string: %q", key, value)

		if key == "email" {
			player.Email = value.(string)
		}
		if key == "password" {
			attemptedpassword = value.(string)

		}

	}

	ctl.us.GetUserByEmail(&player)

	match, err := services.ComparePassword(attemptedpassword, player.Password)
	if err != nil {
		config.Apex.Errorf("%s", err)

		return
	}
	if match == false {
		config.Apex.Warn("WRONG PASSWORD SEND BACK TO LOGIN OR REGIESTER")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	type Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	var jwtKey = []byte("my_secret_key")

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: player.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	player.Token, err = token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error

		c.JSON(http.StatusInternalServerError, gin.H{"msg": "JWT ERROR"})
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	//c.SetCookie("token", tokenString, 604800, "/", "", false, true)

	c.JSON(http.StatusOK, player)
	config.Apex.Infof("Logged in player:%+v", player)

}

// @Summary Register new user
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/register [post]
func (ctl *userController) GetUser(c *gin.Context) {

	var player types.User

	objid, err := primitive.ObjectIDFromHex(c.Params.ByName("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "userid incorrect"})
		return
	}

	player.Userid = objid

	if !ctl.us.GetUserByObjectID(&player) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "userid not found"})
		return
	}

	c.JSON(http.StatusOK, player)
	config.Apex.Infof("Getting user now %s", player)
}

// @Summary Get a list of stats across a timeperiod
// @Produce  json
// @Param userid
// @Param timeperiod
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/:userid/stats [GET]
func (ctl *userController) GetUserStats(c *gin.Context) {

	type Userstats struct {
		Contestsplayed int
		Gamesplayed    int
		Contestswon    int
		Contestslost   int
		Conteststied   int
		Competitors    int
		Contestlist    []types.Contest
	}

	//var contestlist []types.Contest
	var ustats Userstats
	//daterange, err := time.Parse("01022006", c.Params.ByName("daterange"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "daterange incorrect"})
	//	return
	//	}

	objid, err := primitive.ObjectIDFromHex(c.Params.ByName("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "userid incorrect"})
		return
	}
	ustats.Contestlist = ctl.cs.GetContestsUserInvolved(objid)
	ustats.Contestsplayed = len(ustats.Contestlist)

	competitorList := []primitive.ObjectID{}

	for _, contest := range ustats.Contestlist {
		config.Apex.Infof("LENGTH %v", len(contest.Games))
		ustats.Gamesplayed = ustats.Gamesplayed + len(contest.Games)

		for _, stats := range contest.Outcome {

			if stats.Playerid == objid {
				if stats.Result == "won" {

					//did anybody else win...
					isthisatie := false
					for _, resulttest := range contest.Outcome {
						if resulttest.Playerid != objid && resulttest.Result == "won" {
							isthisatie = true
						}
					}

					if isthisatie {
						ustats.Conteststied = ustats.Conteststied + 1
					} else {
						ustats.Contestswon = ustats.Contestswon + 1
					}

				}
				if stats.Result == "lost" {
					ustats.Contestslost = ustats.Contestslost + 1
				}

			} else {

				if len(competitorList) == 0 {
					competitorList = append(competitorList, stats.Playerid)
				}

				alreadythere := false
				for _, entry := range competitorList {
					if entry == stats.Playerid {
						alreadythere = true
					}
				}
				if !alreadythere {

					competitorList = append(competitorList, stats.Playerid)
				}

			}
		}
	}

	//ustats.Gamesplayed = len(ustats.Contestlist)
	//ustats.Contestswon = len(ustats.Contestlist)
	//ustats.Contestslost = len(ustats.Contestlist)
	//ustats.Conteststied = len(ustats.Contestlist)
	ustats.Competitors = len(competitorList)

	config.Apex.Infof("*************USTATS %v\n", ustats)

	//jsonustats, err := json.Marshal(ustats)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "userstats failed to compile"})
		return
	}

	//isValid := json.Valid(jsonustats)

	c.JSON(http.StatusOK, ustats)
	//json.Unmarshal(jsonustats, &ustatsprint)
	//config.Apex.Infof("PRINT OUT JSON %#v\n", ustatsprint)
}

func (ctl *userController) UpdateUser(c *gin.Context) {
}
