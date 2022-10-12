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
var hawadeet = []hadoota{}

func isEmpty(s string)(bool){
	spaces := 0
	for i := 0; i < len(s); i++{
		if s[i] == ' '{
			spaces+=1
		}
	}
	return (spaces == len(s))
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

	r.GET("/",func (c *gin.Context)  {
		
	})

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
		if isEmpty(newHadoota.Body){
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