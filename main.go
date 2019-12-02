package main

import (
	"github.com/NguyenHuy1812/forum/api"
	"github.com/NguyenHuy1812/forum/elastic"
	
)

func main() {
	elastic.Run()

	api.Run()

}
