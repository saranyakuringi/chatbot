package api

import (
	"bot/model"
	"bot/util"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func POST_NewUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error in json request body", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in json request body"})
			return
		}

		//unmarshal the data
		var data util.Userlist
		err = json.Unmarshal(id, &data)
		if err != nil {
			log.Println("Error in unmarshal of json data", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in unmarshal of json data"})
			return
		}
		fmt.Println("data", data)
		//c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in unmarshal of json data"})

		userlist, err := model.Insert_User(db, data)
		fmt.Println("userlist", userlist)
		if err != nil {
			fmt.Println("Error in user insert query", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in user insert query"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"Response": userlist})
	}
}

func HandleQuestionRequest(c *gin.Context) {
	var req util.Question

	id, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in json request body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in json request body"})
		return
	}

	//unmarshal the data

	err = json.Unmarshal(id, &req)
	if err != nil {
		log.Println("Error in unmarshal of json data", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in unmarshal of json data"})
		return
	}
	fmt.Println("data", req)
	// Call GetResponse function with the provided question
	answer, err := GetResponse(req.Question)
	if err != nil {
		log.Println("Error in json request body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in json request body"})
		return
	}

	// Return the response in JSON format
	c.IndentedJSON(http.StatusOK, gin.H{"Response": answer})
}

func GetResponse(question string) (string, error) {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apikey := viper.GetString("API_KEY")
	if apikey == "" {
		panic("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	err := client.CompletionStreamWithEngine(ctx, "gpt-3.5-turbo-instruct", gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(100),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	result := fmt.Sprint("result", err)

	if err != nil {
		log.Println("Error in receiving response form the gpt api", err)
		return result, err
	}

	fmt.Printf("\n")
	return result, nil
}
