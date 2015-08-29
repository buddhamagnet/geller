package main

import (
	"net/http"
)

type Client struct {
	client    *http.client
	APIkey    string
	APIsecret string
}
