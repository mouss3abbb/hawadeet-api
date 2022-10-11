package main

import(
	"net/http"
  	"github.com/gin-gonic/gin"
	"math/rand"
	"errors"
)


type hadoota struct{
	Status		string		`json:"status"`
	Body		string		`json:"body"`
}
var hawadeet = []hadoota{
	{Status:"Motivational",Body: "Go die"},
	{Status:"Motivational",Body: "Go die"},
	{Status:"Motivational",Body: "No no live"},
	{Status:"Love",Body: "Just stay"},
	{Status:"Sad",Body: "You died?"},
}



func getspecificHawadeet (status string) ([]hadoota,error) {
	specificHawadeet := []hadoota{}
	for _,v := range hawadeet{
		if v.Status == status{
			specificHawadeet = append(specificHawadeet,v)
		}
	}
	if len(specificHawadeet) == 0{
		return nil, errors.New("No hawadeet found")
	}else{
		return specificHawadeet, nil
	}
}
func main()  {
	r := gin.Default()

	r.GET("/",func(){})
	
	r.GET("/show-all",func (c *gin.Context){
		c.IndentedJSON(http.StatusOK, hawadeet)
	})

	r.GET("/show-random",func (c *gin.Context){
		index := rand.Intn(len(hawadeet))
		c.IndentedJSON(http.StatusOK, hawadeet[index])
	})

	r.POST("/add-hadoota",func (c *gin.Context){
		var newHadoota hadoota
		if err := c.BindJSON(&newHadoota); err != nil{
			return
		}
		hawadeet = append(hawadeet, newHadoota)
		c.IndentedJSON(http.StatusCreated, newHadoota)
	})

	r.GET("/show-random/:status",func (c *gin.Context)  {
		status := c.Param("status")
		newHawadeet, err := getspecificHawadeet(status)
		if err != nil{
			c.IndentedJSON(http.StatusNotFound,gin.H{"message":"No hawadeet found"})
		}else{
			c.IndentedJSON(http.StatusOK,newHawadeet)
		}


	})
	r.Run()
}