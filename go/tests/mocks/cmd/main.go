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
	"io"
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
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("POST %s: status %d: %s", endpoint, resp.StatusCode, string(respBody))
	}
	return nil
}

func (c *httpClient) authenticate() error {
	authData := fmt.Sprintf(`{"user":"%s","pass":"%s"}`, c.user, c.password)
	req, err := http.NewRequest("POST", c.address+"/auth", bytes.NewReader([]byte(authData)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("auth failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("auth failed: status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("auth response parse failed: %w", err)
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return fmt.Errorf("auth response missing token: %s", string(body))
	}

	c.token = token
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

	fmt.Printf("L8Learn Mock Data Generator\n")
	fmt.Printf("Target: %s\n\n", *address)

	// Authenticate
	fmt.Println("Authenticating...")
	if err := client.authenticate(); err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Authenticated successfully.")

	mocks.RunAllPhases(client)

	fmt.Println("\nDone!")
	os.Exit(0)
}
