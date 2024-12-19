package connections

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"os"

	"github.com/Akuzike8/siem_api/dto"
	"github.com/Velocidex/yaml/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Function to load the configuration file
func LoadConfig(configPath string) (*dto.VelociraptorConfig, error) {
	// Read the configuration file
	y, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Define the struct to store the config
	var config ApiClientConfig
	// Unmarshal the YAML file into the config struct
	err = yaml.Unmarshal(y, &config)
	if err != nil {
		return nil, err
	}

	// Return a new config object from the DTO package
	return &dto.VelociraptorConfig{
		ClientCert:          config.ClientCert,
		ClientPrivateKey:    config.ClientPrivateKey,
		CaCertificate:       config.CaCertificate,
		ApiConnectionString: config.ApiConnectionString,
	}, nil
}

// Function to establish a gRPC connection with TLS credentials
func CreateGRPCConnection(config *dto.VelociraptorConfig) (*grpc.ClientConn, error) {
	// Load the certificates
	cert, err := tls.X509KeyPair([]byte(config.ClientCert), []byte(config.ClientPrivateKey))
	if err != nil {
		return nil, err
	}

	// Prepare the CA certificate pool
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(config.CaCertificate))

	// Create the TLS credentials
	tlsCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   "VelociraptorServer",
	})

	// Establish the gRPC connection
	conn, err := grpc.Dial(config.ApiConnectionString, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Function to execute a VQL query and return the query client
func ExecuteVQLQuery(conn *grpc.ClientConn, query string) (API_QueryClient, error) {
	client := NewAPIClient(conn)

	// Create the VQL request
	request := &VQLCollectorArgs{
		Query: []*VQLRequest{
			{
				Name: "Query",
				VQL:  query,
			},
		},
	}

	// Create a context and execute the query
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Execute the query and return the client
	queryClient, err := client.Query(ctx, request)
	if err != nil {
		return nil, err
	}

	return queryClient, nil
}

// Function to process and log the responses
func ProcessResponses(queryClient API_QueryClient) error {
	for {
		resp, err := queryClient.Recv()
		if err != nil {
			if err == io.EOF {
				break // End of the stream, nothing more to receive
			}
			return err // Return error instead of logging panic
		}

		// Log the response or log messages
		if resp.Response != "" {
			log.Print("Response: ", resp.Response)
		} else if resp.Log != "" {
			log.Print(resp.Log)
		}
	}
	return nil // Return nil if no errors occurred
}
