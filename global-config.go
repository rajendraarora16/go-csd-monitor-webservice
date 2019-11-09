package main

import "time"

//global Constants
var serverAddress string = "0.0.0.0:80"
var timeout = 15 * time.Second

//Routes URL
var searchRouteUrl string = "/api/monitor"

//Path
var staticPathDir string = "build"
var indexPathFile string = "index.html"
