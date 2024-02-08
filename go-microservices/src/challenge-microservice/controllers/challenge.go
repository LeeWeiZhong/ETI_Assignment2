/*
 * @File: controllers.challenge.go
 * @Description: Implements challenge API logic functions
 * 
 */
package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"challenge-microservice/common"
	"challenge-microservice/daos"
	"challenge-microservice/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// challenge manages challenge CRUD
type Challenge struct {
	challengeDAO daos.Challenge
}

// Login godoc
// @Summary Log in to the service
// @Description Log in to the service
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param user formData string true "Username"
// @Param password formData string true "Password"
// @Failure 401 {object} models.Error
// @Success 200 {object} models.Token
// @Router /login [post]
func (m *Challenge) Login(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	formData := url.Values{
		"user":     {username},
		"password": {password},
	}

	var authAddr string = common.Config.AuthAddr + "/api/v1/admin/auth"
	resp, err := http.PostForm(authAddr, formData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		var token models.Token
		json.NewDecoder(resp.Body).Decode(&token)
		ctx.JSON(http.StatusOK, token)
	} else {
		var e models.Error
		json.NewDecoder(resp.Body).Decode(&e)
		ctx.JSON(resp.StatusCode, e)
	}
}

// Addchallenge godoc
// @Summary Add a new challenge
// @Description Add a new challenge
// @Tags challenge
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.Addchallenge true "Add challenge"
// @Failure 401 {object} models.Error
// @Success 200 {object} models.Message
// @Router /challenges [post]
func (m *Challenge) AddChallenge(ctx *gin.Context) {
	var challenge models.Challenge
	if err := ctx.BindJSON(&challenge); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	challenge.ID = primitive.NewObjectID()
	err := m.challengeDAO.Insert(challenge)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusForbidden, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// Listchallenges godoc
// @Summary List all existing challenges
// @Description List all existing challenges
// @Tags challenge
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 404 {object} models.Error
// @Success 200 {object} models.challenge
// @Router /challenges/list [get]
func (m *Challenge) ListChallenges(ctx *gin.Context) {
	var challenges []models.Challenge
	var err error
	challenges, err = m.challengeDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, challenges)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{common.StatusCodeUnknown, "Cannot retrieve challenge information"})
		log.Debug("[ERROR]: ", err)
	}
}
