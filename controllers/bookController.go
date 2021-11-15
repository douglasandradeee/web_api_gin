package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/douglasandradeee/web_api_gin/database"
	"github.com/douglasandradeee/web_api_gin/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	db := database.GetDB()

	// Faz o parse da requisicao recebida no formato JSON para a struct/objeto.
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func FindBookByID(c *gin.Context) {
	id := c.Param("id")

	// Convertendo o ID de string para int.
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Erro ao converter Id")
		log.Fatalln(err.Error())
		return
	}

	db := database.GetDB()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func FindAllBooks(c *gin.Context) {
	db := database.GetDB()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find results: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)

}

func UpdadeBook(c *gin.Context) {
	var book models.Book
	db := database.GetDB()

	// Faz o parse da requisicao recebida no formato JSON para a struct/objeto.
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func DeleteBookByID(c *gin.Context) {
	id := c.Param("id")

	// Convertendo o ID de string para Uint.
	newId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Erro ao converter Id")
		log.Fatalln(err.Error())
		return
	}

	db := database.GetDB()

	err = db.Delete(&models.Book{}, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}
	c.JSON(204, gin.H{
		"sucess": "livro deletado com sucesso",
	})
}
