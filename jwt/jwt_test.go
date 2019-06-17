package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

type Company struct {
	CompanyId int `json:"company_id"`
	jwt.StandardClaims
}

func TestDdf(t *testing.T) {
	s, e := Generate(Company{CompanyId: 1234, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 5).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}, []byte("32434234"))
	fmt.Println(s, e)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		claims, e := Analysis(&Company{}, s, []byte("32434234"))
		if e == nil {
			company, _ := claims.(*Company)
			fmt.Println(company.CompanyId)
		}
	}
}


