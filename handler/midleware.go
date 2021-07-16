package handler

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"

	"go-api-game-boot-camp/service/loger"

	"github.com/dgrijalva/jwt-go"

	"go-api-game-boot-camp/service/auth"

	"github.com/gin-gonic/gin"

	"go-api-game-boot-camp/app"
)

type Midleware struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func NewMidleware(conf *app.Configs, em *app.ErrorMessage) *Midleware {
	return &Midleware{
		CV: conf,
		EM: em,
	}
}

const apikey = "8af8e5a938ec9e8162ec532b77c3a0c3e3dbc1f61710ce5dbe7f51cf4018137a"

// ValidateRequestHeader validate for request header
func (m *Midleware) ValidateRequestHeader(c *gin.Context) {
	//return func(c *gin.Context) {
	//log := logController.NewLogController()
	apiKey := c.Request.Header.Get("x-api-key")
	if apiKey == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "x-api-key Not Found")
		return
	}

	m.ValidateAPIKEY(apiKey, c)
	m.authWithClaims(c)
	return
	//}
}

func (m *Midleware) ValidateAPIKEY(k string, c *gin.Context) {
	//return func(c *gin.Context) {
	log := loger.NewLogController()

	if k != apikey {
		if k == "" {
			log.Error("no x-api-key")
			c.Abort()
			c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"Status":             http.StatusNotImplemented,
				"MessageCode":        "00000",
				"MessageDescription": "access denied",
			})
		} else {
			log.Error("x-api-key not correct")
			c.Abort()
			c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"Status":             http.StatusNotImplemented,
				"MessageCode":        "00000",
				"MessageDescription": "access denied",
			})
		}
		return
	}
	return
}

func (m *Midleware) authWithClaims(c *gin.Context) {

	token, ok := c.Request.Header["Authorization"]
	if len(token) == 0 || !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, m.EM.Auth.InvalidToken)
		return
	}

	claims := &auth.Claims{}
	var verifyKey *rsa.PublicKey

	verifyBytes, err := ioutil.ReadFile(auth.PubKeyPath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	tkn, err := jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if tkn == nil || !tkn.Valid {
		//log.Error("Authorization not correct")
		c.AbortWithStatusJSON(http.StatusUnauthorized, m.EM.Auth.AuthorizationExpiration)
		return
	}
	c.Set("claim", claims)
	//check session
	sect, err := auth.CheckSection(claims.Section, claims.CharacterId)
	if err != nil || sect != true {
		//log.Error("session not match")
		c.AbortWithStatusJSON(http.StatusUnauthorized, m.EM.Auth.GetSessionFail)
		return
	}
	return

}
