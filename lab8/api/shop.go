package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lab8/model"
	"log"
	"net/http"
	"strings"
)

var ErrBadRequest = errors.New("bad request")
var NotFound = errors.New("not found")

func NewShopApi(baseURL string) *ShopApi {
	return &ShopApi{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

type ShopApi struct {
	baseURL string
	client  *http.Client
}

func (a *ShopApi) GetByID(id string) (model.Product, error) {
	products, err := a.ListAllProducts()
	if err != nil {
		return model.Product{}, err
	}

	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return model.Product{}, NotFound
}

func (a *ShopApi) ListAllProducts() ([]model.Product, error) {
	response, err := a.get("/api/products")
	if err != nil {
		return nil, err
	}

	var products []model.Product
	err = json.Unmarshal(response, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (a *ShopApi) DeleteProduct(id string) error {
	path := fmt.Sprintf("/api/deleteproduct?id=%s", id)
	resp, err := a.get(path)
	if err != nil {
		return err
	}

	err = a.checkError(string(resp))
	if err != nil {
		return err
	}

	var data struct {
		Status int `json:"status"`
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return err
	}

	if data.Status != 1 {
		return ErrBadRequest
	}

	return nil
}

func (a *ShopApi) AddProduct(product model.Product) (string, error) {
	resp, err := a.post("/api/addproduct", product)
	if err != nil {
		return "", err
	}

	err = a.checkError(string(resp))
	if err != nil {
		return "", ErrBadRequest
	}

	var data struct {
		ID     int `json:"id"`
		Status int `json:"status"`
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", data.ID), nil
}

func (a *ShopApi) EditProduct(product model.Product) error {
	resp, err := a.post("/api/editproduct", product)
	if err != nil {
		return err
	}

	err = a.checkError(string(resp))
	if err != nil {
		return ErrBadRequest
	}

	return nil
}

func (a *ShopApi) get(path string) ([]byte, error) {
	url := a.baseURL + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	return a.parseResponse(response)
}

func (a *ShopApi) post(path string, body any) ([]byte, error) {
	url := a.baseURL + path
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	return a.parseResponse(response)
}

func (a *ShopApi) parseResponse(response *http.Response) ([]byte, error) {
	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			log.Println(err2)
		}
	}(response.Body)
	return body, err
}

func (a *ShopApi) checkError(bodyString string) error {
	if strings.Contains(bodyString, "<h1>Произошла ошибка</h1>") {
		return fmt.Errorf("error: %s", bodyString)
	}

	return nil
}
