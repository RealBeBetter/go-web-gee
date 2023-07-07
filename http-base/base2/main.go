/**
 * @author wei.song
 * @since 2023/7/8 0:49
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
)

// Engine is the uni handler for all requests
// 用来处理所有请求的 Handler
type Engine struct {
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		_, err := fmt.Fprintf(writer, "Receive Http Request, URL is : %q", request.URL.Path)
		if err != nil {
			return
		}
	case "/hello":
		header := request.Header
		headerKeyMap := make([]string, 0)
		for k := range header {
			headerKeyMap = append(headerKeyMap, k)
		}

		sort.Strings(headerKeyMap)

		for headerKey := range headerKeyMap {

			headerVal := header.Get(string(headerKey))
			_, err := fmt.Fprintf(writer, "Header [%q] = %q \n", headerKey, headerVal)
			if err != nil {
				return
			}
		}

		fmt.Fprintf(writer, "Original Sort......\n")

		for k, v := range header {
			_, err := fmt.Fprintf(writer, "Header [%q] = %q \n", k, v)
			if err != nil {
				return
			}
		}
	default:
		_, err := fmt.Fprintf(writer, "Request is Error, 404 Not Found: %s \n", request.URL)
		if err != nil {
			return
		}
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
