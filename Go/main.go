package main

import (
	"fmt"

	pb "github.com/gowtham4545/firstgoapp/Go/Proto/pb"
)

func main() {
	var b = float64(2.0)
	a := pb.Location{
		Lat: &b,
		Lng: &b,
	}
	fmt.Println(&a)
}
