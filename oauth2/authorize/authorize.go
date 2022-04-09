package authorize

import (
	"context"
	"time"

	"github.com/codeslala/gotil/oauth2"
	pb "github.com/codeslala/gotil/proto/authorize"
	"github.com/codeslala/gotil/util/must"
	"google.golang.org/grpc"
)

func authorize(authorization string) error {
	conn, err := grpc.Dial(oauth2.AuthAddress(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer must.Close(conn)

	cli := pb.NewAuthorizeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = cli.Authorize(ctx, &pb.Request{
		Authorization: authorization,
	})
	if err != nil {
		return err
	}
	return nil
}
