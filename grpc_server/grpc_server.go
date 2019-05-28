package main

import (

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	execdb()
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func execdb(){
       // 初始化資料庫：
       db, err := sql.Open("postgres", "user=postgres password=123 dbname=testdb sslmode=disable")
       defer db.Close()
       checkErr(err)

       // 新增資料列： 
       stmt, err := db.Prepare("INSERT INTO member(id,nickname,phoneNumber,email) VALUES($1,$2,$3,$4);")
       checkErr(err)
       res, err := stmt.Exec("Apple", "iPhone","0920123456","apple@gmail.com")
       checkErr(err)
       println("已新增資料列。\n")

       // 查詢資料列：
       println("查詢資料列：")
       rows, err := db.Query("SELECT id,nickname,phoneNumber,email FROM member")
       checkErr(err)
       var id, nickName, phoneNumber, email string
       for rows.Next() {
              err = rows.Scan(&id, &nickName, &phoneNumber, &email)
              checkErr(err)
              fmt.Println("\t", id, nickName, phoneNumber, email)
       }

       // 更改資料列：
       println("\n更改資料列中...")
       stmt, err = db.Prepare("UPDATE member SET phoneNumber=$1 WHERE id=$2")
       checkErr(err)
       res, err = stmt.Exec("iPhoneXR", "Apple")
       checkErr(err)
       affect, err := res.RowsAffected()
       checkErr(err)
       println("已更改", affect, "個資料列。\n")

       // 查詢資料列：
       println("查詢資料列：")
       rows, err = db.Query("SELECT id,nickname,phoneNumber,email FROM member")
       checkErr(err)
       for rows.Next() {
              err = rows.Scan(&id, &nickName, &phoneNumber, &email)
              checkErr(err)
              fmt.Println("\t", id, nickName, phoneNumber, email)
       }
	   
	   // 删除資料列：
	   /*
       stmt, err = db.Prepare("DELETE FROM member WHERE id=$1")
       checkErr(err)
       res, err = stmt.Exec("apple")
       checkErr(err)
       affect, err = res.RowsAffected()
       checkErr(err)
       println("\n已刪除", affect, "個資料列。\n")
	   */


}


func checkErr(err error) {
       if err != nil {
              panic(err)
       }
}