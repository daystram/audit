package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/config"
	"github.com/daystram/audit/audit-be/constants"
	"github.com/daystram/audit/audit-be/datatransfers"
)

func AuthMiddleware(c *gin.Context) {
	var accessToken string
	if accessToken = strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "); accessToken == "" {
		if accessToken == "" {
			c.Set(constants.ContextKeyIsAuthenticated, false)
			c.Next()
			return
		}
	}
	var err error
	var tokenInfo datatransfers.AccessTokenInfo
	if tokenInfo, err = verifyAccessToken(accessToken); err != nil || !tokenInfo.Active {
		log.Printf("[AuthMiddleware] invalid access_token. %+v\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: "invalid access_token"})
		return
	}
	c.Set(constants.ContextKeyIsAuthenticated, true)
	c.Set(constants.ContextKeySubject, tokenInfo.Subject)
	c.Next()
}

func verifyAccessToken(accessToken string) (info datatransfers.AccessTokenInfo, err error) {
	var response *http.Response
	if response, err = http.Post(fmt.Sprintf("%s/oauth/introspect", config.AppConfig.RatifyIssuer),
		"application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(fmt.Sprintf(
			"token=%s&client_id=%s&client_secret=%s&token_type_hint=access_token",
			accessToken, config.AppConfig.RatifyClientID, config.AppConfig.RatifyClientSecret,
		))),
	); err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	err = json.Unmarshal(body, &info)
	return
}
