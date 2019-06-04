package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/idalin/govel/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSRouter(r *gin.Engine) {
	ws := r.Group("/ws/v1/book")
	ws.GET("/search", WSSearchBooks)
}

func WSSearchBooks(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	ch := make(chan *models.Book, 10)
MainLoop:
	for {
		// 读取 ws 中的数据
		mt, message, err := ws.ReadMessage()
		fmt.Printf("getting ws request.key is %s.\n", message)
		if err != nil {
			break
		}
		go func() {
			for i := range models.BSCache.Items() {
				if b, ok := models.BSCache.Get(i); ok {
					bs, ok := b.(models.BookSource)
					if ok {
						searchResult := bs.SearchBook(string(message))
						if searchResult != nil {
							for _, sr := range searchResult {
								ch <- sr
							}
						}
					} else {
						fmt.Println("not book source.")
					}
				}
			}
			close(ch)
		}()
		for i := range ch {
			book, _ := json.Marshal(i)
			err = ws.WriteMessage(mt, book)
			if err != nil {
				break MainLoop
			}
		}
	}
}
