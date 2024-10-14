package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Chat struct {
	Prompt string `json:"promt"`
}

func main() {

	llm, err := openai.New()
	if err != nil {
		fmt.Println("Error initializing LLM:", err)
		return
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"message": "Hello World from LLM"},
		)
	})

	// Langchain Implementation
	app.POST("/chat", func(c *gin.Context) {
		var requestBody Chat

		// Bind JSON to the requestBody struct
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Background()
		prompt := requestBody.Prompt

		// Generate content using the language model
		content, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": content})

	})

	app.Run(":8080")
}
