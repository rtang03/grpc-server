package core

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

var DownloadCommand = cli.Command{
	Name:   "download",
	Usage:  "download a file",
	Action: downloadAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Value: "localhost:1313",
			Usage: "address of the server to connect to",
		},
		&cli.StringFlag{
			Name:  "file",
			Usage: "filename to download",
		},
		&cli.StringFlag{
			Name:  "cacert",
			Usage: "path of a certifcate to add to the root CAs",
		},
		&cli.StringFlag{
			Name:  "servername-override",
			Usage: "use serverNameOverride for tls ca cert",
		},
	},
}

func downloadAction(c *cli.Context) (err error) {
	var (
		address            = c.String("address")
		file               = c.String("file")
		rootCertificate    = c.String("cacert")
		serverNameOverride = c.String("servername-override")
		client             Client
	)

	if address == "" {
		must(errors.New("address"))
	}

	if file == "" {
		must(errors.New("file must be set"))
	}

	if rootCertificate == "" {
		must(errors.New("cacert must be set"))
	}

	grpcClient, err := NewClientGRPC(ClientGRPCConfig{
		Address:            address,
		RootCertificate:    rootCertificate,
		ServerNameOverride: serverNameOverride,
		Filename:           file,
		UsePublicFolder:    true,
	})
	must(err)
	client = &grpcClient

	err = client.DownloadFile(file)
	must(err)
	defer client.Close()

	fmt.Printf("\nsuccessfully downloaded: %s\n", file)
	return
}
