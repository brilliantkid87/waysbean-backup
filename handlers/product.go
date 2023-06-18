package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	productdto "waysbean/dto/product"
	dto "waysbean/dto/result"
	"waysbean/models"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProduct(c echo.Context) error {
	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	price, _ := strconv.Atoi(c.FormValue("price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	request := productdto.CreateProductRequest{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Price:       price,
		Image:       dataFile,
		Stock:       stock,
	}

	fmt.Println(dataFile)

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	})
	// // }

	// userLogin := c.Get("userLogin")
	// userRole := userLogin.(jwt.MapClaims)["role"].(string)

	// if userRole != "admin" {
	// 	return c.JSON(http.StatusUnauthorized, dto.ErrorResult{
	// 		Code:    http.StatusUnauthorized,
	// 		Message: "Anda bukan admin!",
	// 	})
	// }

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Stock:       request.Stock,
		Image:       dataFile,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseProduct(data),
	})

}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseProduct(product),
	})
}

func convertResponseProduct(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		ID:          u.ID,
		Name:        u.Name,
		Price:       u.Price,
		Description: u.Description,
		Stock:       u.Stock,
		Image:       u.Image,
	}
}
