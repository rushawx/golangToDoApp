package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTData struct {
	SessionId string
	Code      string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"session_id": data.SessionId,
		"code":       data.Code,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	sessionId := t.Claims.(jwt.MapClaims)["session_id"]
	code := t.Claims.(jwt.MapClaims)["code"]
	return t.Valid, &JWTData{SessionId: sessionId.(string), Code: code.(string)}
}
