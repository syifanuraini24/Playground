package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	fmt.Println(products)

	encoder.Encode(ProductListSuccessResponse{Products: productToResponse(products)})
}

func productToResponse(products []repository.Product) []Product {
	results := make([]Product, len(products))
	for i, product := range products {
		results[i] = Product{
			Name:     product.ProductName,
			Price:    product.Price,
			Category: product.Category,
		}
	}
	return results
}
