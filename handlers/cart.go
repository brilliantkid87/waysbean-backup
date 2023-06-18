package handlers

import (
	"net/http"
	"strconv"
	cartdto "waysbean/dto/cart"
	dto "waysbean/dto/result"
	"waysbean/models"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

type handlerCard struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCard {
	return &handlerCard{CartRepository}
}

func (h *handlerCard) FindCart(c echo.Context) error {
	carts, err := h.CartRepository.FindCart()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: carts})
}

func (h *handlerCard) CreateCart(c echo.Context) error {
	request := new(cartdto.CreateCartRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	cart := models.Cart{
		TransactionID: request.TransactionID,
		ProductID:     request.ProductID,
		OrderQuantity: request.OrderQuantity,
		UserID:        request.UserID,
	}

	data, err := h.CartRepository.CreateCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCart(data)})
}

func (h *handlerCard) GetCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseCart(cart),
	})
}

func (h *handlerCard) GetCardByUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// var transaction models.Transaction
	cartid, err := h.CartRepository.GetCartByUser(id)

	// fmt.Println(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cartid})
}

func convertResponseCart(u models.Cart) cartdto.CartResponse {
	return cartdto.CartResponse{
		ID:            u.ID,
		ProductID:     u.ProductID,
		OrderQuantity: u.OrderQuantity,
		UserID:        u.UserID,
	}
}
