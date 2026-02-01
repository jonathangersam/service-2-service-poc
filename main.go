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

	// ---------- PHASE 2: wire inter-service dependencies ----------
	commentService.WithServices(orderService)
	orderService.WithServices(commentService)

	// ---------- PHASE 3: operational phase; run some tests ----------

	// print orderService's pre-test state
	fmt.Printf("PRE-TEST ORDER\n%#v\n\n", orderService.GetOrder("order-1")) // OUTPUT: models.Order{OrderId:"order-1", GlobalEntityId:"FP_SG", Status:"ACCEPTED"}

	// runs some actions that trigger inter-service calls
	commentService.WriteComment(models.Comment{OrderId: "order-1", Text: "Some manual comment test"}) // also triggers order-api.GetOrder() to retrieve some info
	orderService.UpdateOrderStatus("order-1", "DELIVERED")                                            // also triggers comment-api.WriteComment()

	// print orderService's final state
	fmt.Printf("FINAL ORDER\n%#v\n\n", orderService.GetOrder("order-1")) // OUTPUT: models.Order{OrderId:"order-1", GlobalEntityId:"FP_SG", Status:"DELIVERED"}

	// print commentService's final state
	ps, _ := json.MarshalIndent(commentService.DB, "", "  ")
	fmt.Printf("COMMENT DB's CONTENTS\n%s\n\n", ps) // OUTPUT: (there will be 2 comments, one manual comment, and another made thru order-api)
}
