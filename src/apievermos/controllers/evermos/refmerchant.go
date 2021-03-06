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

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func GetMerchant(c *gin.Context) {
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
	merchantsCollection := evermosDatabase.Collection("merchants")
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

	var txt Mevermos.FindMerchantJSON
	c.BindJSON(&txt)
	merchantID := txt.ID

	if merchantID == "" {
		var merchants []Mevermos.Merchant
		merchantCursor, err := merchantsCollection.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		//fmt.Println(merchants)
		if err = merchantCursor.All(ctx, &merchants); err != nil {
			panic(err)
		}
		defer merchantCursor.Close(ctx)
		if err != nil {
			c.JSON(404, gin.H{"message": "merchants Not Found"})
		} else {
			c.JSON(200, gin.H{"data": merchants})
		}
	} else {
		id, _ := primitive.ObjectIDFromHex(merchantID)
		var merchants []Mevermos.Merchant

		merchantCursor, err := merchantsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if err = merchantCursor.All(ctx, &merchants); err != nil {
			panic(err)
		}
		fmt.Println(merchants)
		// c.JSON(merchants)
		defer merchantCursor.Close(ctx)
		if merchants == nil {
			c.JSON(404, gin.H{"message": "merchants Not Found"})
		} else {
			c.JSON(200, gin.H{"data": merchants})
		}
	}
}

func InsertMerchant(c *gin.Context) {
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
	merchantsCollection := evermosDatabase.Collection("merchants")

	if errdb != nil {
		c.JSON(http.StatusOK, gin.H{"statusload": http.StatusInternalServerError, "statusdb": errdb,
			"result": "Missing Connection"})
		logger.WithFields(logger.Fields{
			"detail": errdb,
		}).Error("Missing Connection")
		logkoe.Info("Missing Connection", "statusdb:", errdb, "statusload :", http.StatusInternalServerError)
		return
	}

	var txtMerchant Mevermos.InsertMerchantJSON
	c.BindJSON(&txtMerchant)

	var MerchantName = txtMerchant.MerchantName
	var Address = txtMerchant.Address
	var Phone = txtMerchant.Phone
	var Owner = txtMerchant.Owner
	var IsActive = txtMerchant.IsActive
	var currentTimeCrt = time.Now()
	var currentTimeUpd = time.Now()

	if MerchantName == "" || Address == "" || IsActive == 0 {
		c.JSON(404, gin.H{"Message": "Something went wrong, please check your data Merchant!"})
	} else {

		merchantsResult, err := merchantsCollection.InsertOne(ctx, bson.D{
			{Key: "merchantname", Value: MerchantName},
			{Key: "address", Value: Address},
			{Key: "phone", Value: Phone},
			{Key: "owner", Value: Owner},
			{Key: "isactive", Value: IsActive},
			{Key: "dtmcrt", Value: currentTimeCrt},
			{Key: "dtmupd", Value: currentTimeUpd},
		})
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"merchantID": merchantsResult.InsertedID, "Message": "Your Merchant has been inserted"})
	}
}
