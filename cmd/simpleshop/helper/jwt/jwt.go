package jwt

import (
	"errors"
	"time"
	b64 "encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/config"
)

//Credential :
type Credential struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"isAdmin"`
}

//CustomPayload :
type CustomPayload struct {
	jwt.StandardClaims
	Credential
}

//NewPayload :
func NewPayload(credential Credential) CustomPayload {
	config, conferr := config.GetConfig()

	var jwtaudience string
	var jwtsubject string
	var jwtissuer string
	var jwtexpires int

	if conferr != nil {
		jwtaudience = "example.com"
		jwtsubject = "client"
		jwtissuer = "example.com"
		jwtexpires = 15
	} else {
		jwtaudience = config.GetString("jwtAudience")
		jwtsubject = config.GetString("jwtSubject")
		jwtissuer = config.GetString("jwtIssuer")
		jwtexpires = config.GetInt("jwtExpiresMinute")
	}

	now := time.Now()
	pl := CustomPayload{
		StandardClaims: jwt.StandardClaims{
			Audience:  jwtaudience,
			ExpiresAt: now.Add(time.Duration(jwtexpires) * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   jwtsubject,
			Issuer:    jwtissuer,
		},
		Credential: credential,
	}
	return pl
}

//Signing :
func Signing(payload CustomPayload) (string, error) {
	config, conferr := config.GetConfig()

	var jwtsignkey string

	if conferr != nil {
		jwtsignkey = "secret"
	} else {
		jwtsignkey = config.GetString("jwtSignkey")
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(jwtsignkey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedtoken, err := token.SignedString([]byte(sEnc))
	return signedtoken, err
}

//Parse :
func Parse(signedtoken string) (Credential, error) {
	config, conferr := config.GetConfig()

	var credential Credential
	var err error

	var jwtsignkey string

	if conferr != nil {
		jwtsignkey = "secret"
	} else {
		jwtsignkey = config.GetString("jwtSignkey")
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(jwtsignkey))

	token, parseerr := jwt.ParseWithClaims(signedtoken, &CustomPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sEnc), nil
	})

	if parseerr != nil {
		err = parseerr
	} else {
		if claims, ok := token.Claims.(*CustomPayload); ok && token.Valid {
			credential = claims.Credential
		} else {
			err = errors.New("Invalid JWT Payload")
		}
	}

	return credential, err
}
