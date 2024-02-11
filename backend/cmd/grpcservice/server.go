package grpcservice

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	bufvalidator "github.com/bufbuild/protovalidate-go"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	invoicev1 "github.com/nikola-enter21/wms/backend/api/invoice/v1"
	orderv1 "github.com/nikola-enter21/wms/backend/api/order/v1"
	productv1 "github.com/nikola-enter21/wms/backend/api/product/v1"
	userv1 "github.com/nikola-enter21/wms/backend/api/user/v1"
	"github.com/nikola-enter21/wms/backend/auth"
	"github.com/nikola-enter21/wms/backend/database"
	"github.com/nikola-enter21/wms/backend/integrations/sender"
	"github.com/nikola-enter21/wms/backend/logging"
	"google.golang.org/grpc"
)

var (
	log = logging.MustNewLogger()
)

type Server struct {
	userv1.UnsafeUserServiceServer
	productv1.UnsafeProductServiceServer
	orderv1.UnsafeOrderServiceServer
	invoicev1.UnsafeInvoiceServiceServer

	DB          database.DB
	EmailSender sender.EmailSender
}

func (s *Server) ServeForever(url string) {
	lis, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}

	v, err := bufvalidator.New()
	if err != nil {
		panic(err)
	}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpcvalidator.UnaryServerInterceptor(v),
		grpcrecovery.UnaryServerInterceptor(),
		auth.UnaryJWTInterceptor(),
	}
	streamInterceptors := []grpc.StreamServerInterceptor{
		grpcvalidator.StreamServerInterceptor(v),
		grpcrecovery.StreamServerInterceptor(),
		auth.StreamJWTInterceptor(),
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(streamInterceptors...)),
	)

	// Register services
	userv1.RegisterUserServiceServer(grpcServer, s)
	productv1.RegisterProductServiceServer(grpcServer, s)
	orderv1.RegisterOrderServiceServer(grpcServer, s)
	invoicev1.RegisterInvoiceServiceServer(grpcServer, s)

	go handleSignals(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func handleSignals(s *grpc.Server) {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-c
	log.Sync()
	go s.GracefulStop()
	<-c
	go s.Stop()
	sig := <-c
	panic(sig)
}
