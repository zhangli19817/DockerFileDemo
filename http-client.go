//Demo
package main
import (
"net/http"
)
import "io/ioutil"
import "fmt"
func main() {
url := "https://raw.githubusercontent.com/" +
"dockerinaction/ch8_multi_stage_build/master/http-client.go"
resp, err := http.Get(url)
if err != nil {
panic(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
fmt.Println("response:\n", string(body))
}
