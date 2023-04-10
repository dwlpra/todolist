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

type TodoHandler struct {
	service *usecase.TodoService
}

func NewTodoHandler(service *usecase.TodoService) *TodoHandler {
	return &TodoHandler{service}
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	activityGroupID := c.Query("activity_group_id")
	var todos = new([]entity.Todo)
	var err error
	if todos, err = h.service.GetAllTodos(&activityGroupID); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Success",
			Message: "Success",
			Data:    utils.EmptyResp{},
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (h *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var todo = new(entity.Todo)
	var err error
	if todo, err = h.service.GetTodoByID(id); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
			Data:    utils.EmptyResp{},
		})

	}
	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {

	var todo = new(entity.Todo)
	var err error
	if err = c.BodyParser(todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: err.Error(),
		})

	}

	if todo.Title == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
	}

	if todo.ActivityGroupID == 0 {
		return c.Status(http.StatusBadRequest).JSON(utils.Resp{
			Status:  "Bad Request",
			Message: "activity_group_id cannot be null",
		})
	}

	h.service.CreateTodo(todo)

	return c.Status(http.StatusCreated).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}
func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	var todo = new(entity.Todo)
	var err error
	if err = c.BodyParser(todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.EmptyResp{})

	}

	if todo, err = h.service.UpdateTodo(id, todo); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
			Data:    utils.EmptyResp{},
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}
func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var err error
	if err = h.service.DeleteTodo(id); err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.Resp{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %v Not Found", id),
		})

	}

	return c.Status(http.StatusOK).JSON(utils.Resp{
		Status:  "Success",
		Message: "Success",
		Data:    utils.EmptyResp{},
	})
}
