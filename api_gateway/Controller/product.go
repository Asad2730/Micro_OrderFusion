package controller

import (
	"strconv"

	pt "github.com/Asad2730/Micro_OrderFusion/proto/product"
	"github.com/gin-gonic/gin"
)

type ProductClient struct {
	gRPCClient pt.ProductServiceClient
}

func NewProductClient(client pt.ProductServiceClient) *ProductClient {
	return &ProductClient{gRPCClient: client}
}

func (client *ProductClient) CreateProduct(c *gin.Context) {
	var product *pt.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}

	req := &pt.CreateProductRequest{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
	}

	res, err := client.gRPCClient.CreateProduct(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}

	c.JSON(201, res.Message)
}

func (client *ProductClient) ProductList(c *gin.Context) {
	res, err := client.gRPCClient.ProductList(c, &pt.ListProductRequest{})
	if err != nil {
		c.JSON(402, gin.H{"Error Getting List": err.Error()})
		return
	}

	c.JSON(200, res.Products)
}

func (client *ProductClient) ProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := client.gRPCClient.ProductByID(c, &pt.RequestProductID{Id: int32(id)})
	if err != nil {
		c.JSON(402, gin.H{"Error Getting Data": err.Error()})
	}

	c.JSON(200, res.Product)
}

func (client *ProductClient) UpdateProduct(c *gin.Context) {
	var product *pt.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}

	req := &pt.RequestProductUpdate{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
	}

	res, err := client.gRPCClient.UpdateProduct(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Getting Data": err.Error()})
	}

	c.JSON(200, res.Product)
}

func (client *ProductClient) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := client.gRPCClient.DeleteProduct(c, &pt.RequestProductID{Id: int32(id)})
	if err != nil {
		c.JSON(402, gin.H{"Error Getting Data": err.Error()})
	}

	c.JSON(200, res.Message)

}
