package orderapi

import (
	"service2service_poc/models"
)

// ICommentService is a service this package depends on
type ICommentService interface {
	WriteComment(in models.Comment)
}

type Service struct {
	db             map[string]*models.Order // in-memory DB. map from orderID -> Order
	commentService ICommentService
}

func NewService() *Service {
	initDb := map[string]*models.Order{
		"order-1": &models.Order{
			OrderId:        "order-1",
			GlobalEntityId: "FP_SG",
			Status:         "ACCEPTED",
		},
	}

	return &Service{
		db: initDb,
	}
}

// WithServices allows late-binding of dependencies after all services have been constructed
func (s *Service) WithServices(commentSvc ICommentService) {
	s.commentService = commentSvc
}

func (s *Service) GetOrder(orderId string) models.Order {
	return *s.db[orderId]
}

func (s *Service) UpdateOrderStatus(orderId string, newStatus string) {
	// update the order
	s.db[orderId].Status = newStatus

	// write a comment about it
	s.commentService.WriteComment(models.Comment{
		OrderId: orderId,
		Text:    "status updated to " + newStatus,
	})
}
