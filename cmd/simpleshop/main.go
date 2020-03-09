package main

import (
	// "database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/database"

	// "github.com/lmnzr/simpleshop/cmd/simpleshop/database/filter"
	// "github.com/lmnzr/simpleshop/cmd/simpleshop/database/group"
	// "github.com/lmnzr/simpleshop/cmd/simpleshop/database/order"
	_ "github.com/lmnzr/simpleshop/cmd/simpleshop/docs"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/env"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/helper/jwt"

	// reflectutil "github.com/lmnzr/simpleshop/cmd/simpleshop/helper/model"
	logutil "github.com/lmnzr/simpleshop/cmd/simpleshop/helper/log"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/middleware"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/models"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
)

type auth struct {
	Token string `json:"token" xml:"token" example:"eyJhbGciOiJIU"`
}

//City :
type City struct {
	ID          models.NullInt    `json:"id" field:"id" type:"int" increment:"auto" pkey:"true"`
	Name        models.NullString `json:"name" field:"name" type:"string"`
	IsDeleted   models.NullBool   `json:"is_deleted" field:"is_deleted" type:"boolean" hidden:"true"`
	LastUpdate  models.NullTime   `json:"lastupdate" field:"lastupdate" type:"datetime"`
	Remark      models.NullString `json:"remark" field:"remark" type:"string"`
	Count       models.NullInt    `json:"count" field:"count" type:"int"`
	IsKnown     models.NullBool   `json:"is_known" field:"is_known" type:"boolean"`
	DeletedTime models.NullTime   `json:"deletetime" field:"deletetime" type:"datetime"`
}

//SetName :
func (c *City) SetName(val string) *City {
	c.Name = models.NewNullString(val)
	return c
}

//SetIsKnown :
func (c *City) SetIsKnown(val bool) *City {
	c.IsKnown = models.NewNullBool(val)
	return c
}

//SetDeletedTime :
func (c *City) SetDeletedTime(val time.Time) *City {
	c.DeletedTime = models.NewNullTime(val)
	return c
}

//SetID :
func (c *City) SetID(val int64) *City {
	c.ID = models.NewNullInt(val)
	return c
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logutil.Logger(nil).Error("Error loading .env file")
	}

	environment := env.Getenv("ENVIRONMENT", "development")

	if environment != "development" {
		log.SetLevel(log.InfoLevel)

		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(&lumberjack.Logger{
			Filename:   "var/log/app.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	} else {
		log.SetLevel(log.DebugLevel)
	}

}

// @title Simpleshop Swagger API
// @version 1.0
// @description Swagger API for Golang Project Simpleshop.
// @termsOfService http://swagger.io/terms/
// @BasePath
func main() {
	db, _ := database.OpenDbConnection()
	errping := db.Ping()

	if errping != nil {
		logutil.LoggerDB().Panic("failed to connect to database")
	}

	var city City
	city.SetIsKnown(false)
	city.SetName("Groningen")
	city.SetDeletedTime(time.Now())
	city.SetID(1)

	// fmt.Println(reflectutil.GetFieldTag(city,city.ID,"field"))

	citymodel := database.NewTableQuery(db, "city", city)

	// var filters []filter.Filter
	// filters = append(filters, filter.NewAndFilter("id", "2"))
	// filters = append(filters, filter.NewAndFilter("unknown", "unknown"))
	// citymodel.QueryModel.SetFilters(filters)

	// var groups []group.Group
	// groups = append(groups, group.NewGroup("id"))
	// citymodel.SetGroups(groups)

	// var orders []order.Order
	// orders = append(orders, order.NewOrderDescending("id"))
	// citymodel.SetOrders(orders)

	res, querr := citymodel.RetrieveAll()
	// citymodel.Retrieve()

	if querr != nil {
		logutil.Logger(nil).Error(querr)
	} else {
		for  i := 0; i < len(res); i++ {
			fmt.Println(string(res[:][i]))
		}
			
	}

	// res2,querr2 := citymodel.Retrieve()

	// if querr2 != nil {
	// 	logutil.Logger(nil).Error(querr2)
	// } else {
	// 	fmt.Println(string(res2))
	// }

	router := echo.New()
	middleware.Setup(router)

	port := env.GetenvI("PORT", 9000)

	router.GET("/", public)
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.GET("/forbidden/", forbidden)
	router.GET("/protected/", protected, middleware.JwtMiddleware)
	router.GET("/credential/", credential)

	hello.Routes(router)

	lock := make(chan error)
	go func(lock chan error) { lock <- router.Start(fmt.Sprintf(":%d", port)) }(lock)

	err := <-lock
	if err != nil {
		logutil.Logger(nil).Panic("failed to start application")
	}
}

func public(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to simpleshop")
}

func forbidden(c echo.Context) error {
	return echo.NewHTTPError(500, "Forbidden Land")
}

func protected(c echo.Context) error {
	name := c.Get("name").(string)
	return c.String(http.StatusOK, "Welcome "+name)
}

func credential(c echo.Context) error {
	cred := jwt.Credential{
		Name:  "Almas",
		UUID:  "11037",
		Admin: true,
	}
	pl := jwt.NewPayload(cred)
	token, _ := jwt.Signing(pl)
	a := auth{
		Token: token,
	}
	return c.JSON(http.StatusOK, a)
}
