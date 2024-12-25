package routes

import (
	"context"
	"encoding/json"
	"findmydoc-backend/helpers"
	"findmydoc-backend/llm"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
)

type MessageParam struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type LlmParams struct {
	AccToken string         `json:"acc-token"`
	Messages []MessageParam `json:"messages"`
}

func LlmHandler(c *gin.Context) {
	var body LlmParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id = helpers.Authenticate(body.AccToken)

	if id == nil {
		c.Status(http.StatusUnauthorized)

		return
	}

	var messageContents = make([]llms.MessageContent, len(body.Messages))

	for i := 0; i < len(messageContents); i++ {
		var messageType llms.ChatMessageType

		if body.Messages[i].Type == "user" {
			messageType = llms.ChatMessageTypeHuman
		} else if body.Messages[i].Type == "system" {
			messageType = llms.ChatMessageTypeSystem
		} else {
			messageType = llms.ChatMessageTypeAI
		}

		messageContents[i] = llms.TextParts(messageType, body.Messages[i].Content)
	}

	var ctx = context.Background()
	_, err := llm.Llm.GenerateContent(ctx, messageContents, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		json, err := json.Marshal(gin.H{
			"type":    "part",
			"content": string(chunk),
		})

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return nil
		}

		fmt.Fprintln(c.Writer, string(json))
		c.Writer.Flush()

		return nil
	}))

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	json, err := json.Marshal(gin.H{
		"type": "done",
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)

		return
	}

	fmt.Fprintln(c.Writer, string(json))
	c.Writer.Flush()
}
