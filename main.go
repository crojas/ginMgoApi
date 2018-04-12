package main

import (
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

type Product struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"na"`
	brand    string        `json:"brand" bson:"ba"`
	Variants []Variant     `json:"variants" bson:"va"`
}

type Variant struct {
	Description string  `json:"description" bson:"ds"`
	Sku         string  `json:"sku" bson:"sk"`
	Price       float32 `json:"price" bson:"pr"`
	Stock       float32 `json:"stock" bson:"st"`
}

func (db *DB) GetProductByID(c *gin.Context) {
	producID := bson.ObjectIdHex(c.Param("id"))
	var product Product
	err := db.Collection.Find(bson.M{"_id": producID}).One(&product)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func (db *DB) GetAllProduct(c *gin.Context) {
	var products []Product
	err := db.Collection.Find(nil).All(&products)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func (db *DB) CreateProduct(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	if err == nil {
		db.Collection.Insert(product)
		c.JSON(http.StatusCreated, product)
	} else {
		panic(err)
	}
}

func (db *DB) UpdateProduct(c *gin.Context) {
	producID := bson.ObjectIdHex(c.Param("_id"))
	var product Product
	err := c.BindJSON(&product)
	if err == nil {
		err = db.Collection.Update(bson.M{"_id": producID}, bson.M{"$set": &product})
		if err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, product)
		}
	} else {
		panic(err)
	}

}

func (db *DB) DeleteProduct(c *gin.Context) {
	producID := bson.ObjectIdHex(c.Param("id"))
	err := db.Collection.Remove(bson.M{"_id": producID})
	if err != nil {
		c.JSON(http.StatusNoContent, nil)
	}
}

func main() {
	session, err := mgo.Dial(GetConnectionString())
	c := session.DB("ecommerce").C("products")
	db := &DB{Session: session, Collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := gin.Default()
	router.GET("/products", db.GetAllProduct)
	router.GET("/products/:id", db.GetProductByID)
	router.POST("/products", db.CreateProduct)
	router.PUT("/products/:id", db.UpdateProduct)
	router.DELETE("/products/:id", db.DeleteProduct)
	router.Run(Port())

}

func GetConnectionString() string {
	connectionString := os.Getenv("DB")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}
	return connectionString
}

func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"

	}
	return ":" + port
}
