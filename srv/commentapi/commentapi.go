package commentapi

import (
	"math/rand"
	"time"

	"service2service_poc/models"
)

// IOrderService is a service this package depends on
type IOrderService interface {
	GetOrder(orderId string) models.Order
}

type Service struct {
	DB           []models.Comment
	orderService IOrderService
}

func NewService() *Service {
	return &Service{
		DB: make([]models.Comment, 0),
	}
}

// WithServices allows late-binding of dependencies after all services have been constructed
func (s *Service) WithServices(orderService IOrderService) {
	s.orderService = orderService
}

func (s *Service) WriteComment(in models.Comment) {
	// get the order for some info that wasn't passed in
	order := s.orderService.GetOrder(in.OrderId)
	in.GlobalEntityId = order.GlobalEntityId

	// generate new ID
	in.CommentId = genRandomId(8)

	// write the comment
	s.DB = append(s.DB, in)
}

const alphabet = "ABCDEFHJKLMNPQRTUVWXY34789" // no 0/O, 1/I, S/5, Z/2, G/6

func genRandomId(length int) string {
	rand.Seed(time.Now().UnixNano())

	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(out)
}
