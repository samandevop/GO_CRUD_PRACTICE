package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"crud/models"

	"github.com/gin-gonic/gin"
)

// CreateActor godoc
// @ID create_actor
// @Router /actor [POST]
// @Summary Create Actor
// @Description Create Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param actor body models.CreateActor true "CreateActorRequestBody"
// @Success 201 {object} models.Actor "GetactorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateActor(c *gin.Context) {
	var actor models.CreateActor

	err := c.ShouldBindJSON(&actor)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Actor().Create(context.Background(), &actor)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.Actor().GetByPKey(
		context.Background(),
		&models.ActorPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdActor godoc
// @ID get_by_id_actor
// @Router /actor/{id} [GET]
// @Summary Get By Id Actor
// @Description Get By Id Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Actor "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetActorById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Actor().GetByPKey(
		context.Background(),
		&models.ActorPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListActor godoc
// @ID get_list_actor
// @Router /actor [GET]
// @Summary Get List Actor
// @Description Get List Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListActorResponse "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetActorList(c *gin.Context) {
	var (
		limit  int
		offset int
		err    error
	)

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.Actor().GetList(
		context.Background(),
		&models.GetListActorRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateActor godoc
// @ID update_actor
// @Router /actor/{id} [PUT]
// @Summary Update Actor
// @Description Update Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param actor body models.UpdateActor true "CreateActorRequestBody"
// @Success 200 {object} models.Actor "GetactorsBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateActor(c *gin.Context) {

	var (
		actor models.UpdateActor
	)

	id := c.Param("id")

	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required actor id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required actor id").Error())
		return
	}

	err := c.ShouldBindJSON(&actor)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Actor().Update(
		context.Background(),
		id,
		&actor,
	)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
		return
	}

	resp, err := h.storage.Actor().GetByPKey(
		context.Background(),
		&models.ActorPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdActor godoc
// @ID delete_by_id_actor
// @Router /actor/{id} [DELETE]
// @Summary Delete By Id Actor
// @Description Delete By Id Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Actor "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteActor(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required actor id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required actor id").Error())
		return
	}

	err := h.storage.Actor().Delete(
		context.Background(),
		&models.ActorPrimarKey{
			Id: id,
		},
	)

	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
