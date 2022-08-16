package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CreateNewContainer(image string) (container.ContainerCreateCreatedBody, error) {
	cli, err := client.NewEnvClient()

	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	//	IPAM := network.EndpointIPAMConfig{
	//		IPv4Address: "172.18.0.2",
	//	}
	//	endPoint := network.EndpointSettings{
	//		IPAMConfig: &IPAM,
	//		NetworkID:  "e2ed2a08594c", // ID hardcoded, should be a variable in the future
	//		IPAddress:  "172.18.0.2",
	//	}

	hostBinding := nat.PortBinding{ //
		HostIP: "0.0.0.0",
		//HostIP:   "172.18.0.2",
		HostPort: "80",
	}

	containerPort, err := nat.NewPort("tcp", "80")

	if err != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			PortBindings: portBinding,
			//			AutoRemove:   true,
			//			NetworkMode:  "--net vulnNet",
		}, nil, //&network.NetworkingConfig{
		//EndpointsConfig: endPoint,
		//}
		nil, "")

	if err != nil {
		panic(err)
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started \n", cont.ID)

	return cont, nil
}

func ListAllContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}

// Stop a container
func StopContainer(containerID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		panic(err)
	}
	fmt.Print("Container ", containerID[:10], " is stopped")

}

func StopAllContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}

}

func LogsSpecificContainer(container string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, container, options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

}

func ListAllImages() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	// list all images with all the details
	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.RepoTags)
	}

}

func PullImage(image string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out) // Redirecting stdout from os to the GO app stdout
}

func GetAllContainerState() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// types.ContainerListOptions{All: true}
	// types.ContainerState{}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.Names) // Autogen names
		fmt.Println(container.Image) // Name of the image
		fmt.Println(container.Status)
	}

}

func GetContainerState(containerID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.ID == containerID {
			fmt.Println(container.Names) // Autogen names
			fmt.Println(container.Image) // Name of the image
			fmt.Println(container.Status)
		}

	}

}

func createNetwork(networkName string) (networkID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	// Create a network
	net, err := cli.NetworkCreate(ctx, networkName, types.NetworkCreate{})
	if err != nil {
		panic(err)
	}

	newNetID := net.ID

	fmt.Printf("Created network with ID: %s \n", newNetID)

	return newNetID
	// Return the netword id here

}

func inspectNetwork(networkName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Listing and displaying information for a selected network named vulnNet
	// Filters list options
	networkOpts := types.NetworkListOptions{ // Fetches the vulnNet network
		Filters: filters.NewArgs(
			filters.Arg("name", networkName),
		),
	}

	// List networks
	nett, err2 := client.NetworkAPIClient.NetworkList(cli, ctx, networkOpts)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Name of the network: ", nett[0].Name)

	fmt.Printf("Network ID: %s\n", nett[0].ID)

	fmt.Println("The subnet is: ", nett[0].IPAM.Config[0].Subnet)

}

// connect container to network
func connectContainerToNetwork(containerID string, networkID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Connecting a container to a network
	err = cli.NetworkConnect(ctx, networkID, containerID, &network.EndpointSettings{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected container: %s to network : %s\n", containerID, networkID)
}

func removeNetwork(networkID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	// Remove a network
	err = cli.NetworkRemove(ctx, networkID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Removed network with ID: %s \n", networkID)
}

// get IP address of a container
func getContainerIPAddress(containerID string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Get the IP address of a container
	cont, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("IP address of container: %s is: %s\n", containerID, cont.NetworkSettings.IPAddress)
}

//
// netID := createNetwork("vulnNet2")
// time.Sleep(10 * time.Second)
// inspectNetwork("vulnNet2")

// create a new container and return the container ID
// container, err := CreateNewContainer("nginx")
// if err != nil {
// 	panic(err)
// }

// time.Sleep(10 * time.Second)

// ListAllContainers()

// LogsSpecificContainer(container.ID)

// StopAllContainers()

// ListAllImages()

// Pull an image from DockerHub
// PullImage("hello-world")

// Get the states of all containers in the docker environment
// GetAllContainerState()

// GetContainerState(container.ID)

// connectContainerToNetwork(container.ID, netID)
// time.Sleep(10 * time.Second)

// getContainerIPAddress(container.ID)

// StopContainer(container.ID)
// time.Sleep(10 * time.Second)

// removeNetwork(netID)
