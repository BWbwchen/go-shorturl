package api

import (
	"net/http"
	"os"
	"shorturl/database"
	"shorturl/model"
	"shorturl/util"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load environment variable
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func HandleRedirect(c *gin.Context) {
	shortName := c.Param("shortname")

	URL, stateCode := database.Find(shortName)
	if stateCode == model.NotFound {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	//c.JSON(http.StatusOK, "Hello %s", URL)
	c.Redirect(http.StatusMovedPermanently, URL)
}

func AddShortName(c *gin.Context) {
	shortName := c.PostForm("shortname")
	URL := c.PostForm("url")

	if shortName == "" {
		shortName = util.GetNewShortName()
	} else if util.CheckValid(shortName) == false {
		shortName = util.GetNewShortName()
	}

	newShort := model.ShorturlSturct{
		Shortname: shortName,
		URL:       URL,
	}

	database.Insert(newShort)
	//c.String(http.StatusOK, "shortName is : %s, URL is : %s", shortName, URL)
	c.JSON(http.StatusOK, util.SendResponse(200, newShort.Shortname))
}

func LoginHandler(c *gin.Context) {
	var body struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		/*
			fmt.Println(c.GetRawData())
			fmt.Println(c.PostForm("account"), c.PostForm("password"))
			fmt.Println(err.Error(), body)
		*/
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if checkAccount(body.Account, body.Password) {
		jwtSecret := []byte(os.Getenv("KEY"))

		now := time.Now()
		jwtId := body.Account + strconv.FormatInt(now.Unix(), 10)

		// set claims and sign
		claims := model.Claims{
			Account: body.Account,
			StandardClaims: jwt.StandardClaims{
				Audience:  body.Account,
				ExpiresAt: now.Add(20 * time.Second).Unix(),
				Id:        jwtId,
				IssuedAt:  now.Unix(),
				Issuer:    "ginJWT",
				Subject:   body.Account,
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "fail",
	})
}

func checkAccount(account string, password string) bool {
	if account == "abcd" && password == "1234" {
		return true
	}
	return false
}

func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not login yet, please login",
		})
		c.Abort()
		return
	}
	token := strings.Split(auth, "Bearer ")[1]
	jwtSecret := []byte(os.Getenv("KEY"))

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*model.Claims); ok && tokenClaims.Valid {
		c.Set("account", claims.Account)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
