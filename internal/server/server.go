package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/popeskul/product-store/internal/handler"
	"github.com/popeskul/product-store/pkg/domain"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer     *grpc.Server
	productsServer domain.ProductsServiceServer
}

func NewServer(handlers *handler.Handler, interceptor grpc.UnaryServerInterceptor) *Server {
	return &Server{
		grpcServer:     grpc.NewServer(grpc.UnaryInterceptor(interceptor)),
		productsServer: handlers,
	}
}

func (s *Server) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	domain.RegisterProductsServiceServer(s.grpcServer, s.productsServer)

	go func() {
		log.Fatalln(s.grpcServer.Serve(lis))
	}()

	log.Println("Server started on port", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	s.grpcServer.GracefulStop()

	return nil
}
