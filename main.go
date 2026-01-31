package main

import (
	"encoding/json"
	"fmt"
	"service2service_poc/models"
	"service2service_poc/srv/commentapi"
	"service2service_poc/srv/orderapi"
)

func main() {
	// ---------- PHASE 1: construct all services ----------
	commentService := commentapi.NewService()
	orderService := orderapi.NewService()

	// ---------- PHASE 2: write inter-service dependencies ----------
	commentService.WithServices(orderService)
	orderService.WithServices(commentService)

	// ---------- PHASE 3: operational phase; run some tests ----------

	// print pre-test order state
	fmt.Printf("PRE-TEST ORDER\n%#v\n\n", orderService.GetOrder("order-1"))

	// runs some actions
	commentService.WriteComment(models.Comment{ // also trigger order-api.GetOrder() to get some info
		CommentId: "",
		OrderId:   "order-1",
		Text:      "Some manual comment test",
	})

	orderService.UpdateOrderStatus("order-1", "DELIVERED") // also trigger comment-api.WriteComment

	// print each services' state
	fmt.Printf("FINAL ORDER\n%#v\n\n", orderService.GetOrder("order-1"))

	ps, _ := json.MarshalIndent(commentService.DB, "", "  ")
	fmt.Printf("COMMENT DB's CONTENTS\n%s\n\n", ps)
}
