package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	todo "to_do_list"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("SIGNUP HANDLER STARTED") //
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("SIGNUP HANDLER STARTED") //добавил сам
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
