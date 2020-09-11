package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"masuk/internal/grpc/service"
)

func main() {
	serverAddress := "localhost:7080"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//client := service.NewRepositoryServiceClient(conn)
	//
	//for i := range [10]int{} {
	//	repositoryModel := service.Repository{
	//		Id:        int64(i),
	//		IsPrivate: true,
	//		Name:      string("Grpc-Demo"),
	//		UserId:    1245,
	//	}
	//
	//	if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
	//		panic(fmt.Sprintf("Was not able to insert Record %v", e))
	//	} else {
	//		fmt.Println("Record Inserted..")
	//		fmt.Println(responseMessage)
	//		fmt.Println("=============================")
	//	}
	//}

	//userGetUserRequest := service.GetUserRequest{
	//	UserNum: 1,
	//}
	//
	userClient := service.NewUserServiceClient(conn)
	//resp, err := userClient.GetUser(context.Background(), &userGetUserRequest)
	//if err != nil {
	//
	//}
	//fmt.Println(resp)

	//user := &service.User{
	//	Username: "user",
	//	Email:    "user@email.com",
	//	Password: "secret",
	//}
	//userRequest := &service.CreateUserRequest{User: user}
	//respCreateUser, err := userClient.CreateUser(context.Background(), userRequest)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(respCreateUser)

	loginRequest := &service.LoginRequest{
		Email:    "syazlinsazali@gmail.com",
		Password: "secret",
	}

	token, err := userClient.Login(context.Background(), loginRequest)
	if err != nil {
		log.Println(err)
	}

	log.Println(token)
}
