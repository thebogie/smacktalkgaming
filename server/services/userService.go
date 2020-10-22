package services

import (
	//"errors"

	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thebogie/smacktalkgaming/config"
	"github.com/thebogie/smacktalkgaming/repos"
	"github.com/thebogie/smacktalkgaming/types"
)

const apisecret string = "LETMEINNOW"

// UserService interface
type UserService interface {
	GetUserByObjectID(*types.User) bool
	GetUserByEmail(*types.User) bool
	AddUser(*types.User)
	ValidateJWT(*types.User) bool
	CreateJWT(*types.User) string
}

type userService struct {
	UserRepo repos.UserRepo
}

// NewUserService will instantiate User Service
func NewUserService(
	userRepo repos.UserRepo) UserService {

	return &userService{
		UserRepo: userRepo,
	}
}

//Claims fish
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (us *userService) CreateJWT(player *types.User) (newtoken string) {

	var jwtKey = []byte(apisecret)

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(15 * time.Minute)
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
	newtoken, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		//c.JSON(http.StatusInternalServerError, gin.H{"msg": "JWT ERROR"})
		return
	}

	return newtoken

}

func (us *userService) ValidateJWT(player *types.User) bool {
	var tokenString string
	//normally Authorization the_token_xxx
	strArr := strings.Split(player.Token, " ")
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apisecret), nil
	})

	config.Apex.Infof("TOKEN?" + token.Raw)

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return false
	}

	return true

}

func (us *userService) AddUser(in *types.User) {

	if !us.GetUserByEmail(in) {
		us.UserRepo.AddUser(in)
	} else {
		config.Apex.Infof("User already exists: %+v", in)
	}

	return
}

func (us *userService) GetUserByObjectID(in *types.User) bool {
	//if id == 0 {
	//	return nil, errors.New("id param is required")
	//}

	return us.UserRepo.FindUserByObjectID(in)
}

func (us *userService) GetUserByEmail(in *types.User) bool {

	return us.UserRepo.FindUserByEmail(in)
}
