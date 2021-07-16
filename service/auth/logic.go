package auth

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm"
	"go-api-game-boot-camp/app"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

var srv *authService

type authService struct {
	conf   *app.Configs
	em     *app.ErrorMessage
	repo   *authRepo
	claims *Claims
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &authService{
		conf: conf,
		em:   em,
		repo: &authRepo{conf: conf},
	}
}

func NewAuthService(conf *app.Configs, em *app.ErrorMessage) *authService {
	repo := authRepo{}

	return &authService{
		conf: conf,
		em:   em,
		repo: repo.InitAuthRepo(conf, em),
	}
}

func encryptPassword(password string) string {
	s := password
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashPassword : "+sha1_hash)
	return sha1_hash
}

func compilePassword(password string, loginUuid string) string {
	//request password + loginUuid[15:20]
	//uuid Start 0
	detailPass := password + loginUuid[15:20]
	fmt.Println("compilePassword : "+detailPass)
	return detailPass
}

func (srv *authService) EnCode(request credentials) (result parseCode, err error) {
	var signKey *rsa.PrivateKey

	signBytes, err := ioutil.ReadFile(PriKeyPath)
	if err != nil {
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return
	}

	log.Infof("Check Character In Database ...")
	var detail character
	detail, err = srv.repo.getCharacter(request)
	if err != nil && err.Error() == "record not found" {
		err = srv.em.Auth.LoginNotFound
		log.Error(err)
		return
	}
	log.Infof("Detail : %+v", detail)

	log.Info("[Hashing Password ...]")

	detailPass := compilePassword(request.Password, detail.LoginUuid)

	hashPass := encryptPassword(detailPass)

	log.Info("CheckPassword in database ...")

	// expect Password && Username
	if detail.Password != hashPass || request.Username != detail.Username {
		err = srv.em.Auth.InvalidUsernamePassword
		log.Error(err)
		//return err
		return
	}
	section := uuid.New().String()
	expirationTime := time.Now().Add(1 * time.Hour)
	//expirationTime := time.Now().Add(time.Duration(viper.GetInt("expirationTime.expHour")) * time.Hour)

	claims := &Claims{
		CharacterId: detail.CharacterId,
		Section:     section,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		err = srv.em.Auth.StatusInternalServerError
		log.Error(err)
		// If there is an error in creating the JWT return an internal server error
		return
	}
	log.Info("[Update Time ...]")

	timeNow := time.Now()
	//expiration Token +1Hour
	timeToken := timeNow.Add(1 * time.Hour)

	// update lastLogin_date , section , exp_section
	update, err := srv.repo.updateLastLogin(request, timeNow, section, timeToken)
	if err != nil && err.Error() != "record not found" {
		err = srv.em.Auth.LoginNotFound
		log.Error(err)
		return
		//return
	}
	log.Info("LastLoginDate : ", update.LastLogin)
	log.Info("Expiration Token : ", update.ExpSection)

	result = parseCode{
		Value: tokenString,
	}

	return
}

func CheckSection(section string, characterId int) (checkSect bool, err error) {
	//validate Section
	var detail character
	detail, err = getSection(characterId)
	if err != nil && err.Error() == "record not found" {
		checkSect = false
		log.Error("Login not found.")
		return
	}
	log.Infof("Detail : %+v", detail)
	if detail.Section != section {
		checkSect = false
		log.Error("An unexpected error occurred on get session.")
		return
	}
	checkSect = true
	return
}

func (srv *authService) GetClaimCurrent(c *gin.Context) {
	if claims, ok := c.Get("claim"); ok {
		srv.claims = claims.(*Claims)
	}
}

func (srv *authService) GetCharacterIdInClaim() int {
	return srv.claims.CharacterId
}
