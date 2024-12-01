package main

import "sync"

type DbConections struct {
	Host     string
	DbName   string
	UserId   string
	Password string
}

var connectionPool sync.Pool

func main() {

}
