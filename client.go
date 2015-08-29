package main

import (
	"net/http"
)

type Client struct {
	client    *http.Client
	APIkey    string
	APIsecret string
}
