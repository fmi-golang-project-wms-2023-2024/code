package httpgateway

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	orderv1 "github.com/nikola-enter21/wms/backend/api/order/v1"
	productv1 "github.com/nikola-enter21/wms/backend/api/product/v1"
	invoicev1 "github.com/nikola-enter21/wms/backend/api/user/v1"
	userv1 "github.com/nikola-enter21/wms/backend/api/user/v1"
	"github.com/nikola-enter21/wms/backend/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	log = logging.MustNewLogger()
)

func Serve(ctx context.Context, httpURL, grpcURL string) {
	conn, err := grpc.DialContext(
		ctx,
		grpcURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	marshaller := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		},
	}

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaller),
		runtime.WithForwardResponseOption(responseHeaderMatcher),
	)
	err = userv1.RegisterUserServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}
	err = productv1.RegisterProductServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}
	err = orderv1.RegisterOrderServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}
	err = invoicev1.RegisterUserServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    httpURL,
		Handler: allowCORS(gwmux),
	}

	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func responseHeaderMatcher(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	headers := w.Header()
	if location, ok := headers["Grpc-Metadata-Location"]; ok {
		w.Header().Set("Location", location[0])
		w.WriteHeader(http.StatusFound)
	}

	return nil
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}
