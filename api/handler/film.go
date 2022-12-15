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

// CreateFilm godoc
// @ID create_film
// @Router /film [POST]
// @Summary Create Film
// @Description Create Film
// @Tags Film
// @Accept json
// @Produce json
// @Param film body models.CreateFilm true "CreateFilmRequestBody"
// @Success 201 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateFilm(c *gin.Context) {
	var film models.CreateFilm

	err := c.ShouldBindJSON(&film)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Film().Create(context.Background(), &film)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.Film().GetByPKey(
		context.Background(),
		&models.FilmPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdFilm godoc
// @ID get_by_id_film
// @Router /film/{id} [GET]
// @Summary Get By Id Film
// @Description Get By Id Film
// @Tags Film
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetFilmById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Film().GetByPKey(
		context.Background(),
		&models.FilmPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListFilm godoc
// @ID get_list_film
// @Router /film [GET]
// @Summary Get List Film
// @Description Get List Film
// @Tags Film
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListFilmResponse "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetFilmList(c *gin.Context) {
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

	resp, err := h.storage.Film().GetList(
		context.Background(),
		&models.GetListFilmRequest{
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

// UpdateFilm godoc
// @ID update_film
// @Router /film/{id} [PUT]
// @Summary Update Film
// @Description Update Film
// @Tags Film
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param film body models.UpdateFilm true "CreateFilmRequestBody"
// @Success 200 {object} models.Film "GetFilmsBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateFilm(c *gin.Context) {

	var (
		film models.UpdateFilm
	)

	id := c.Param("id")

	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required film id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required film id").Error())
		return
	}

	err := c.ShouldBindJSON(&film)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Film().Update(
		context.Background(),
		id,
		&film,
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

	resp, err := h.storage.Film().GetByPKey(
		context.Background(),
		&models.FilmPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdFilm godoc
// @ID delete_by_id_film
// @Router /film/{id} [DELETE]
// @Summary Delete By Id Film
// @Description Delete By Id Film
// @Tags Film
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteFilm(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required film id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required film id").Error())
		return
	}

	err := h.storage.Film().Delete(
		context.Background(),
		&models.FilmPrimarKey{
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
