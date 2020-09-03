package cmd

import (
	"errors"
	"fmt"

	"github.com/rtang03/grpc-server/core"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

var Upload = cli.Command{
	Name:   "upload",
	Usage:  "upload a file",
	Action: uploadAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Value: "localhost:1313",
			Usage: "address of the server to connect to",
		},
		&cli.IntFlag{
			Name:  "chunk-size",
			Usage: "size of chunk messages (grpc only)",
			Value: 1 << 12,
		},
		&cli.StringFlag{
			Name:  "file",
			Usage: "file to upload",
		},
		&cli.StringFlag{
			Name:  "root-certificate",
			Usage: "path of a certifcate to add to the root CAs",
		},
		&cli.BoolFlag{
			Name:  "compress",
			Usage: "whether or not to enable payload compression",
		},
	},
}

func uploadAction(c *cli.Context) (err error) {
	var (
		chunkSize       = c.Int("chunks-zie")
		address         = c.String("address")
		file            = c.String("file")
		rootCertificate = c.String("root-certificate")
		compress        = c.Bool("compress")
		client          core.Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	if file == "" {
		must(errors.New("file must be set"))
	}

	grpcClient, err := core.NewClientGRPC(core.ClientGRPCConfig{
		Address:         address,
		RootCertificate: rootCertificate,
		Compress:        compress,
		ChunkSize:       chunkSize,
	})
	must(err)
	client = &grpcClient

	stat, err := client.UploadFile(context.Background(), file)
	must(err)
	defer client.Close()

	fmt.Printf("%d\n", stat.FinishedAt.Sub(stat.StartedAt).Nanoseconds())

	return
}
