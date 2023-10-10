package delivery

import (
	"errors"
	"greedy/pkg/domain"
	"greedy/pkg/usecase"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	err error
)

type Handler struct {
	usecase usecase.UsecaseMethods
}

type HandlerMethods interface {
	Set(*gin.Context)
	Get(*gin.Context)
	QPush(*gin.Context)
	QPop(*gin.Context)
	BQPop(*gin.Context)
}

func NewHandlers(config *domain.Config, usecase usecase.UsecaseMethods) HandlerMethods {
	return Handler{
		usecase: usecase,
	}

}

func (h Handler) Set(c *gin.Context) {
	var input keyValueInput
	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("error parsing json :", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	value := domain.KeyValue{
		Value: input.Value.Value,
	}
	if input.Value.Expiry != 0 {
		expiry := strconv.Itoa(input.Value.Expiry)
		value.Expiry = expiry
	}
	if input.Value.Condition != "" {
		value.Condition = strings.ToUpper(input.Value.Condition)
	}
	if err = h.usecase.Set(input.Key, value); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error from server": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Value stored for key : ": input.Key})
	return
}

func (h Handler) Get(c *gin.Context) {
	var input key
	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("error parsing json :", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	if input.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error from input": errors.New("enter a valid key")})
		return
	}
	value, err := h.usecase.Get(input.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from server": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"Value for key :": input.Key, "is": value})
	return
}

func (h Handler) QPush(c *gin.Context) {
	var input QueInput
	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("error parsing json : ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	if input.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error from input": errors.New("enter a valid key")})
		return
	}
	if err = h.usecase.QPush(input.Key, input.Value...); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from server": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Value stored for key : ": input.Key})
	return
}

func (h Handler) QPop(c *gin.Context) {
	var input key
	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("error parsing json : ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	if input.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error from input": errors.New("enter a valid key")})
		return
	}
	value, err := h.usecase.QPop(input.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from server": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"Value for que :": input.Key, "is": value})
	return
}

func (h Handler) BQPop(c *gin.Context) {
	var input BQPopInput
	if err = c.ShouldBindJSON(&input); err != nil {
		log.Println("error parsing json : ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	if input.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error from input": errors.New("enter a valid key")})
		return
	}
	value, err := h.usecase.BQPop(input.Key, input.Time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from server": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"Value for que :": input.Key, "is": value, "waited seconds : ": input.Time})
	return
}
