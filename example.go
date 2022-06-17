package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )
// type Companies struct{
// CompanyName string `json:"name"`
// CatchPhrase string `json:"catchPhrase"`
// Bs string `json:"bs"`
// }
// type Geos struct {
// Lat string `json:"lat"`
// Lng string `json:"lng"`
// }
// type Addresses struct{
// Street string `json:"street"`
// Suite string `json:"suite"`
// City string `json:"city"`
// ZipCode string `json:"zipcode"`
// Geo Geos `json:"geo"`
// }
// type User struct {
// 	Id     int64 `json:"id"`
// 	Name  string `json:"name"`
// 	Username   string `json:"username"`
// 	Email string `json:"email"`
// 	Address Addresses `json:"address"`
// 	Phone string `json:"phone"`
// 	Website string `json:"website"`
// 	Company Companies `json:"company"`
// }

func main2() {
	// 	url := "https://jsonplaceholder.typicode.com/users/1"
	// 	resp, getErr := http.Get(url)
	// 	if getErr != nil {
	//         log.Fatal(getErr)
	//       }
	// 	  body, readErr := ioutil.ReadAll(resp.Body)
	//       if readErr != nil {
	//         log.Fatal(readErr)
	//       }
	// 	  //fmt.Println(body)
	// 	  post_obj := User{}

	// 	  jsonErr := json.Unmarshal(body,&post_obj)
	// 	  if jsonErr != nil {
	// 		log.Fatal(jsonErr)
	// 	 }
	// 	 r := gin.Default()
	// 	r.GET("/ping", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 				"data":post_obj,
	// 		})
	// 	})
	// 	r.Run()

}
