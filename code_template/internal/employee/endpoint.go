package employee

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-tutorial/code_template/internal"
	"go-tutorial/code_template/internal/models"
	"net/http"
)

type employeeService interface {
	GetEmployeeById(ctx context.Context, empId string) ([]models.Employee, error)
	//GetEmployeeByFirstName(firstName string)
	//GetEmployeeByLastName(lastName string)
}

type Endpoint struct {
	cv  *internal.Configs
	srv employeeService //Service Interface Tier
}

func NewEndpoint(cv *internal.Configs) *Endpoint {
	return &Endpoint{cv: cv, srv: NewService(cv)}
}

func (e Endpoint) GetEmployeeById(c echo.Context) error {
	var request struct {
		EmployeeId string `json:"employee_id"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "mismatch json request format",
			Data:    nil,
		})
	}

	data, err := e.srv.GetEmployeeById(c.Request().Context(), request.EmployeeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	})
}
