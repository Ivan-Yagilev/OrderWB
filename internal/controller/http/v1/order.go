package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order/internal/entity"
	"order/internal/service"
)

type orderRoutes struct {
	orderService service.Order
}

func newOrderRoutes(g *echo.Group, orderService service.Order) {
	r := &orderRoutes{
		orderService: orderService,
	}

	g.POST("/create", r.create)
}

func (r *orderRoutes) create(c echo.Context) error {
	var input entity.Order

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	// В канал может прилететь что угодно!
	//if err := c.Validate(input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return err
	//}

	err := r.orderService.CreateOrder(c.Request().Context(), input)
	if err != nil {
		if err == service.ErrOrderAlreadyExists {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return nil
}
