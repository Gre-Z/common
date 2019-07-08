package jwt_test

import (
	"fmt"
	jwt2 "github.com/Gre-Z/common/jwt"

	//jwt2 "github.com/Gre-Z/common/jwt"
	"github.com/dgrijalva/jwt-go"
	"rbac-server/models"
	"testing"
	"time"
)

type Company struct {
	CompanyId int `json:"company_id"`
	jwt.StandardClaims
}

func TestDdf(t *testing.T) {
	userJwt := models.AdminUserJwt{
		ID: 4,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	//jwt2.Generate()
	s, e := jwt2.Generate(userJwt,[]byte("32434234"))
	//s, e := Generate(Company{CompanyId: 1234, StandardClaims: jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 5).Unix(),
	//	IssuedAt:  time.Now().Unix(),
	//}}, []byte("32434234"))
	fmt.Println(s, e)
		claims, e := jwt2.Analysis(&models.AdminUserJwt{}, s, []byte("32434234"))
		fmt.Println(e)
		if e == nil {
			company, _ := claims.(*models.AdminUserJwt)
			fmt.Println(company.ID)
		}
	//}
}


