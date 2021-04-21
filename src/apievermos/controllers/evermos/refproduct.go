package evermos

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexcesaro/log/stdlog"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	Mevermos "apigoevermos/src/apievermos/model/evermos"
	Conf "apigoevermos/src/config"

	logger "github.com/sirupsen/logrus"
)

func GetProduct(c *gin.Context) {
	logkoe := stdlog.GetFromFlags()

	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}

	var ctx = func() context.Context {
		return context.Background()
	}()
	evermosDatabase, errdb := Conf.Connectmongo()
	productsCollection := evermosDatabase.Collection("products")
	//fmt.Println(errdb)

	if errdb != nil {
		c.JSON(http.StatusOK, gin.H{"statusload": http.StatusInternalServerError, "statusdb": errdb,
			"result": "Missing Connection"})
		logger.WithFields(logger.Fields{
			"detail": errdb,
		}).Error("Missing Connection")
		logkoe.Info("Missing Connection", "statusdb:", errdb, "statusload :", http.StatusInternalServerError)
		return
	}

	var txt Mevermos.FindProductJSON
	c.BindJSON(&txt)
	ProductID := txt.ID

	if ProductID == "" {
		var products []Mevermos.Product
		ProductCursor, err := productsCollection.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		//fmt.Println(products)
		if err = ProductCursor.All(ctx, &products); err != nil {
			panic(err)
		}
		defer ProductCursor.Close(ctx)
		if err != nil {
			c.JSON(404, gin.H{"message": "products Not Found"})
		} else {
			c.JSON(200, gin.H{"data": products})
		}
	} else {
		id, _ := primitive.ObjectIDFromHex(ProductID)
		var products []Mevermos.Product

		ProductCursor, err := productsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if err = ProductCursor.All(ctx, &products); err != nil {
			panic(err)
		}
		//fmt.Println(products)
		// c.JSON(products)

		defer ProductCursor.Close(ctx)
		if products == nil {
			c.JSON(404, gin.H{"message": "products Not Found"})
		} else {
			c.JSON(200, gin.H{"data": products})
		}
	}
}

func InsertProduct(c *gin.Context) {
	logkoe := stdlog.GetFromFlags()

	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}

	var ctx = func() context.Context {
		return context.Background()
	}()
	evermosDatabase, errdb := Conf.Connectmongo()
	productsCollection := evermosDatabase.Collection("products")

	if errdb != nil {
		c.JSON(http.StatusOK, gin.H{"statusload": http.StatusInternalServerError, "statusdb": errdb,
			"result": "Missing Connection"})
		logger.WithFields(logger.Fields{
			"detail": errdb,
		}).Error("Missing Connection")
		logkoe.Info("Missing Connection", "statusdb:", errdb, "statusload :", http.StatusInternalServerError)
		return
	}

	var txtProduct Mevermos.InsertProductJSON
	c.BindJSON(&txtProduct)

	var ProductName = txtProduct.ProductName
	var Price = txtProduct.Price
	var Category = txtProduct.Category
	var Stock = txtProduct.Stock
	var MerchantID = txtProduct.MerchantID
	var MerchantName = txtProduct.MerchantName
	var IsActive = txtProduct.IsActive
	var currentTimeCrt = time.Now()
	var currentTimeUpd = time.Now()

	if ProductName == "" || Price == "" || Category == "" {
		c.JSON(404, gin.H{"Message": "Something went wrong, please check your data Product!"})
	} else {

		productsResult, err := productsCollection.InsertOne(ctx, bson.D{
			{Key: "merchantid", Value: MerchantID},
			{Key: "merchantname", Value: MerchantName},
			{Key: "productname", Value: ProductName},
			{Key: "price", Value: Price},
			{Key: "category", Value: Category},
			{Key: "stock", Value: Stock},
			{Key: "isactive", Value: IsActive},
			{Key: "dtmcrt", Value: currentTimeCrt},
			{Key: "dtmupd", Value: currentTimeUpd},
		})
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"ProductID": productsResult.InsertedID, "Message": "Your Product has been inserted"})
	}
}
