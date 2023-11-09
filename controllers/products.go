package controllers

import (
	"apirediscache/db"
	"apirediscache/db/redis"
	"apirediscache/models"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	rdb := redis.ClientInstance()
	products := []models.Product{}
	db.DB.Find(&products)
	bytes, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}

	keyExist := rdb.Exists("getAllProducts").Val()
	if keyExist == 0 {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{"data": products})
		rdb.Set("getAllProducts", string(bytes), 0)
		rdb.Set("getAllProducts:validation", "false", 2*time.Second)
	} else if keyExist == 1 {
		productFromCache, err := rdb.Get("getAllProducts").Result()
		if err != nil {
			panic(err)
		}
		isCacheStale, _ := rdb.Get("getAllProducts:validation").Result()

		if isCacheStale != "false" {
			wg := sync.WaitGroup{}
			wg.Add(3)
			go println("cache is stale - refreshing...")
			wg.Done()
			go rdb.Set("getAllProducts", string(bytes), 0)
			wg.Done()
			go rdb.Set("getAllProducts:validation", "false", 2*time.Second)
			wg.Done()

			wg.Wait()
		}
		var data []models.Product
		if err := json.Unmarshal([]byte(productFromCache), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func CreateProduct(c *gin.Context) {
	input := models.CreateProductInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{Name: input.Name, Value: input.Value}
	db.DB.Create(&product)
}
