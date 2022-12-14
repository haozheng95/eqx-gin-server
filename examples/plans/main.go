package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"sigs.k8s.io/yaml"

	metal "github.com/haozheng95/eqx-gin-server/metal/v1"
)

func main() {
	// TODO Because the swagger2 spec can not declare that facilities can be
	// returned with plan results, including available_in results in:
	// available_in:
	// - href: ""
	//include := []string{}
	//exclude := []string{"available_in"}
	configuration := metal.NewConfiguration()
	configuration.Debug = true
	//configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	configuration.AddDefaultHeader("X-Auth-Token", "h9FYNNnc3sbdcSr3PUyptriEyRQPqwhg")
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		}}}
	configuration.HTTPClient = client

	api_client := metal.NewAPIClient(configuration)
	//resp, r, err := api_client.PlansApi.FindPlans(context.Background()).Include(include).Exclude(exclude).Execute()
	//resp, r, err := api_client.DevicesApi.FindProjectDevices(context.Background(), "42207fc3-dda2-471e-8c84-179908f64f7b").Include(include).Exclude(exclude).Execute()
	resp, r, err := api_client.EventsApi.FindEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.FindPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	//response from `FindPlans`: PlanList
	s, err := yaml.Marshal(resp)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `PlansApi.FindPlans`: %v\n", string(s))
}
