package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	beef "github.com/napakornsk/go-algorithm-api/beef/proto"
	"google.golang.org/grpc"
)

func fetchFromExternal() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (s *BeefServiceServer) GetAllBeef(req *empty.Empty, stream grpc.ServerStreamingServer[beef.GetAllBeefResponse]) error {
	fmt.Println("The GetAllBeef service was invoked...")
	res, err := fetchFromExternal()
	if err != nil {
		return fmt.Errorf("failed to fetch data from external API: %v", err)
	}
	words := strings.Fields(res)

	for _, w := range words {
		if err := stream.Send(&beef.GetAllBeefResponse{
			Data: strings.Trim(w, ",."),
		}); err != nil {
			return fmt.Errorf("failed to send word throuh stream: %v", err)
		}
	}

	return nil
}
