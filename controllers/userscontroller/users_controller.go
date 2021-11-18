package userscontroller

import (
	"blog/blog"
	"blog/domain/users"
	"blog/utils/errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := errors.NewBadRequest(err.Error())
		log.Error(apiErr)
		ctx.JSON(apiErr.Status, apiErr)
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		apiErr := errors.Conflict(err.Error())
		log.Error(apiErr)
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	result, saveErr := blog.CreateUser(user)
	if saveErr != nil {
		log.Error(saveErr)
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	log.Info(result)
	log.Infof("New user with id %d successfully created.", user.ID)
	ctx.JSON(http.StatusCreated, user)
}
func GetUser(ctx *gin.Context) {
	userID, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequest("Invalid user id")
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := blog.GetUser(userID)
	if getErr != nil {
		log.Error(getErr)
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}
