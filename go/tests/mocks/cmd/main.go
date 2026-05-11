/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/saichler/l8learn/go/tests/mocks"
)

type httpClient struct {
	address  string
	user     string
	password string
	token    string
	client   *http.Client
}

func (c *httpClient) Post(endpoint string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := c.address + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("POST %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("POST %s: status %d", endpoint, resp.StatusCode)
	}
	return nil
}

func main() {
	address := flag.String("address", "https://localhost:2773", "Server address")
	user := flag.String("user", "admin", "Username")
	password := flag.String("password", "admin", "Password")
	insecure := flag.Bool("insecure", false, "Skip TLS verification")
	flag.Parse()

	transport := &http.Transport{}
	if *insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &httpClient{
		address:  *address,
		user:     *user,
		password: *password,
		client:   &http.Client{Transport: transport},
	}

	// TODO: Authenticate and get bearer token
	// client.token = authenticate(client, *user, *password)

	fmt.Printf("L8Learn Mock Data Generator\n")
	fmt.Printf("Target: %s\n\n", *address)

	mocks.RunAllPhases(client)

	fmt.Println("\nDone!")
	os.Exit(0)
}
