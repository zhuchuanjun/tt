package main

import (
	"log"
)

func main() {
	log.Println("Starting client...")

	// 调用用户服务
	if err := runUserClient(); err != nil {
		log.Fatalf("User client error: %v", err)
	}

	//// 调用订单服务
	//if err := runOrderClient(); err != nil {
	//	log.Fatalf("Order client error: %v", err)
	//}

	log.Println("Client finished successfully")
}
