package gomemgo

import (
	"github.com/warrior-labs/gomemgo/memgopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type memgoClient struct {
	c memgopb.MemgoServiceClient
}

type MemgoSecureOptions struct {
	X509CertificatePath string
	X509KeyPath         string
}

func NewClient(dsn string, secureOptions *MemgoSecureOptions) (*memgoClient, error) {
	var client memgopb.MemgoServiceClient
	if secureOptions != nil {
		creds, err := credentials.NewClientTLSFromFile(secureOptions.X509CertificatePath, secureOptions.X509KeyPath)
		if err != nil {
			return nil, err
		}
		conn, err := grpc.NewClient(dsn, grpc.WithTransportCredentials(creds))
		if err != nil {
			return nil, err
		}
		client = memgopb.NewMemgoServiceClient(conn)
	} else {
		conn, err := grpc.NewClient(dsn, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		client = memgopb.NewMemgoServiceClient(conn)
	}

	return &memgoClient{
		c: client,
	}, nil
}
