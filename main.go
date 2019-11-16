//main

package main

import (
	"gomod/router"
)


func main() {
	router := router.InitRouter()
	router.Run(":2048") // listen and serve on 0.0.0.0:2048
}
