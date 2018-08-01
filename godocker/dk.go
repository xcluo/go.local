package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	DefaultDockerHost = "uninx:///var/run/docker2.sock"
)

func main() {
	// Ignore "imported and not used"
	func(vals ...interface{}) {}(
		context.Background, log.New, http.Client{}, time.Saturday,
		docker.NewClient, types.ContainerListOptions{},
		api.DefaultVersion,
	) // Delete before commit
	uuu := func(i ...interface{}) {}
	ppp := func(s string) { fmt.Println("<< " + s) } // For testing
	var _, _ = uuu, ppp

	fmt.Println("<< Init Client")
	// Init client by socket
	// host, version := "http://127.0.0.1:2376", "18.03.1-ce"
	// host, version := "tcp://127.0.0.1:2375", api.DefaultVersion
	// host, version := "http://192.168.2.192:2375", api.DefaultVersion
	// host, version := "unix:///var/run/docker.sock", api.DefaultVersion
	// client := http.Client{Transport: &http.Transport{
	// 	MaxIdleConns:       10,
	// 	IdleConnTimeout:    30 * time.Second,
	// 	DisableCompression: true}}
	// httpHeaders := map[string]string{}
	// c, err := docker.NewClient(host, version, &client, httpHeaders)
	// fmt.Println(c, err)

	// Init client by env
	c, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Client: ", c)
	}

	fmt.Println("<< Check Client")
	if res := docker.IsErrConnectionFailed(err); res != false {
		fmt.Println("Failed Connection", res)
	} else {
		fmt.Println("Success")
	}

	ppp("Image List")
	options := types.ImageListOptions{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	images, err := c.ImageList(ctx, options)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%T", images)
		for _, image := range images {
			fmt.Printf("\nID=%s \t labels=%d \n", image.ID, image.Size)
		}
	}

	fmt.Println("<< Container List")
	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	} else {
		for _, container := range containers {
			fmt.Printf("ID=%s, Image=%s\n", container.ID[:10], container.Image)
		}
	}

	fmt.Println("<< Container Stop")
	d := time.Duration(3000)
	err = c.ContainerStop(context.Background(), "ab4f", &d)
	if err != nil {
		panic(err)
	}

	fmt.Println("<< Container Start")
	err = c.ContainerStart(context.Background(), "ab4f", types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("<< Container Restart")
	err = c.ContainerRestart(context.Background(), "ab4f", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("<< Container Create")
	containerConfig := container.Config{
		Hostname:   "bch",
		Domainname: "bch.local",
		// Cmd:        "ls",
		Image: "ubuntu",
	}
	hostConfig := container.HostConfig{}
	networkingConfig := network.NetworkingConfig{}
	mycon, err := c.ContainerCreate(context.Background(), &containerConfig,
		&hostConfig, &networkingConfig, "bch")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("BMBM = ", mycon)
	}
}
