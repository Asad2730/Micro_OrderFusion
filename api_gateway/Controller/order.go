package controller

import (
	"strconv"

	pt "github.com/Asad2730/Micro_OrderFusion/proto/order"
	"github.com/gin-gonic/gin"
)

type OrderClient struct {
	gRPCClient pt.OrderServiceClient
}

func NewOrderClient(client pt.OrderServiceClient) *OrderClient {
	return &OrderClient{gRPCClient: client}
}

func (s *OrderClient) CreateOrder(c *gin.Context) {
	var order *pt.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}
	req := &pt.CreateOrderRequest{
		UserId: order.UserId,
		Total:  order.Total,
		Status: order.Status,
	}

	res, err := s.gRPCClient.CreateOrder(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}
	c.JSON(200, res.Order)
}

func (s *OrderClient) CreateOrderItem(c *gin.Context) {
	var orderItem *pt.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(500, gin.H{"Error binding data": err.Error()})
		return
	}

	req := &pt.CreateOrderItemRequest{
		OrderId:   orderItem.Id,
		ProductId: orderItem.ProductId,
		Price:     orderItem.Price,
	}

	res, err := s.gRPCClient.CreateOrderItem(c, req)
	if err != nil {
		c.JSON(402, gin.H{"Error Creating": err.Error()})
		return
	}

	c.JSON(200, res.OrderItem)
}

func (s *OrderClient) OrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := s.gRPCClient.OrderByID(c, &pt.OrderByIDRequest{Id: int32(id)})
	if err != nil {
		c.JSON(402, gin.H{"Error Getting Data": err.Error()})
		return
	}

	c.JSON(200, res)
}

func (s *OrderClient) OrderList(c *gin.Context) {
	res, err := s.gRPCClient.OrderList(c, &pt.OrderListRequest{})
	if err != nil {
		c.JSON(402, gin.H{"Error Getting Data": err.Error()})
		return
	}

	c.JSON(200, res.OrderList)
}
