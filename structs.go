package main

type Query struct {
	Key string `form:"key"`
}

type Record struct {
	Key   string `form:"key"`
	Value string `form:"value"`
}
