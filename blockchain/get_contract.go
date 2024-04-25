package blockchain

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/sahil-cloud/backend/constants"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

var orgConfig OrgSetup = OrgSetup{
	OrgName:      "Org1",
	MSPID:        "Org1MSP",
	CertPath:     constants.CERT_PATH,
	KeyPath:      constants.KEY_PATH,
	TLSCertPath:  constants.TLS_CERT_PATH,
	PeerEndpoint: "localhost:7051",
	GatewayPeer:  "peer0.org1.example.com",
}

var Contract *client.Contract

func Initialize() {
	log.Printf("Initializing connection for %s...\n", orgConfig.OrgName)
	clientConnection := newGrpcConnection()
	id := newIdentity()
	sign := newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	orgConfig.Gateway = *gateway
	network := orgConfig.Gateway.GetNetwork(constants.CHANNEL_NAME)
	Contract = network.GetContract(constants.CHAINCODE_NAME)
	log.Println("Successfully initialized blockchain connection")
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(orgConfig.TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, orgConfig.GatewayPeer)

	connection, err := grpc.Dial(orgConfig.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(orgConfig.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(orgConfig.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := ioutil.ReadDir(orgConfig.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(orgConfig.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
