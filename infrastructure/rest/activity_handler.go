package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dwlpra/todolist/domain/entity"
	"github.com/dwlpra/todolist/usecase"
	"github.com/dwlpra/todolist/utils"
	fiber "github.com/gofiber/fiber/v2"
)

type ActivityHandler struct {
	service *usecase.ActivityService
}

func NewActivityHandler(service *usecase.ActivityService) *ActivityHandler {
	return &ActivityHandler{service}
}

func (h *ActivityHandler) GetAllActivities(c *fiber.Ctx) error {
	var activities = new([]entity.Activity)
	var err error
	if activities, err = h.service.GetAllActivities(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.EmptyResp{})
	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (h *ActivityHandler) GetActivityByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var activity = new(entity.Activity)
	var err error
	if activity, err = h.service.GetActivityByID(id); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %v Not Found", id),
			Data:    utils.EmptyResp{},
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (h *ActivityHandler) CreateActivity(c *fiber.Ctx) error {
	var activity = new(entity.Activity)
	var err error
	if err = c.BodyParser(activity); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: err.Error(),
		})

	}

	if activity.Title == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	if err = h.service.CreateActivity(activity); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (h *ActivityHandler) UpdateActivity(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var activity = new(entity.Activity)
	var err error
	if err := c.BodyParser(activity); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})

	}

	data, err := h.service.UpdateActivity(id, activity)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %v Not Found", id),
			Data:    utils.EmptyResp{},
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}

func (h *ActivityHandler) DeleteActivity(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteActivity(id); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %v Not Found", id),
			Data:    utils.EmptyResp{},
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    utils.EmptyResp{},
	})
}
