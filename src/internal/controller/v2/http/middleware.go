package http

import (
	"fmt"
	"net/http"
	"src/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
	adminCtx            = "adminID"
	keySecret           = "keySecret"
)

func (a *AuthController) Identity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "empty auth header"})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid auth header"})
	}

	userID, err := a.authService.ParseToken(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	user, err := a.userService.GetUserByID(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	if user.Role == model.UserRoleCustomer {
		c.Set(userCtx, userID)
	} else {
		c.Set(adminCtx, userID)
	}
}

func (a *AuthController) UserIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "empty auth header"})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid auth header"})
	}

	userID, err := a.authService.ParseToken(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	c.Set(userCtx, userID)
}

func (a *AuthController) AdminIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "empty auth header"})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid auth header"})
	}

	adminID, err := a.authService.ParseToken(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	admin, err := a.userService.GetUserByID(c, adminID)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	if admin.Role != model.UserRoleAdmin {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "not an admin"})
	}

	c.Set(adminCtx, adminID)
}

func getUserID(c *gin.Context) (int, error) {

	id, ok := c.Get(userCtx)
	if !ok {
		return 0, fmt.Errorf("%s", "userID not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, fmt.Errorf("%s", "userID is of invalid type")
	}

	return idInt, nil
}

func getKeySecret(c *gin.Context) (string, error) {

	curKeySecret, ok := c.Get(keySecret)
	if !ok {
		return "", fmt.Errorf("%s", "userID not found")
	}

	curKeySecretString, ok := curKeySecret.(string)
	if !ok {
		return "", fmt.Errorf("%s", "userID is of invalid type")
	}

	return curKeySecretString, nil
}