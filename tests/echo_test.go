package tests

import (
	"testing"
)

const (
	echoP = "1323/tcp"
)

func Test_Echo(t *testing.T) {
	// You need to connect the container of the server with the container of the db.
	// server <-> mongodb
	// server <-> mysql

	// ctx := context.Background()

	// req := testcontainers.ContainerRequest{
	// 	FromDockerfile: testcontainers.FromDockerfile{
	// 		Context: "../",
	// 	},
	// 	ExposedPorts: []string{echoP},
	// 	WaitingFor:   wait.ForListeningPort(nat.Port(echoP)),
	// }

	// echoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	// 	ContainerRequest: req,
	// 	Started:          true,
	// })
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer echoC.Terminate(ctx)
	// ip, err := echoC.Host(ctx)
	// if err != nil {
	// 	t.Error(err)
	// }
	// port, err := echoC.MappedPort(ctx, echoP)
	// if err != nil {
	// 	t.Error(err)
	// }
	// resp, err := http.Get(fmt.Sprintf("http://%s:%s", ip, port.Port()))
	// if resp.StatusCode != http.StatusOK {
	// 	t.Errorf("Expected status code %d. Got %d.", http.StatusOK, resp.StatusCode)
	// }
}
