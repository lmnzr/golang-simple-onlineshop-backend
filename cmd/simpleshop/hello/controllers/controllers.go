package controllers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/database"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/hello/models"
	"github.com/lmnzr/simpleshop/cmd/simpleshop/types"
)

//GetMyHello : Function Return MyHello
func GetMyHello(helloWord string, c echo.Context) *models.Hello {
	cc := c.(*types.DBContext)

	hello := new(models.Hello)

	sess, _ := session.Get("session", cc.Context)

	message := "Hello World !!!"
	if m := sess.Values["message"]; m != nil {
		message = fmt.Sprintf("%v", sess.Values["message"])
	}

	origin := "Default"
	if o := sess.Values["origin"]; o != nil {
		origin = fmt.Sprintf("%v", sess.Values["origin"])
	}

	hello.SetMessage(message).SetOrigin(origin)
	go helloLogDB(hello, cc.DB)

	return hello
}

//PostMyHello : Function Save MyHello to Session
func PostMyHello(c echo.Context) *models.Hello {
	cc := c.(*types.DBContext)

	hello := new(models.Hello)
	err := cc.Context.Bind(hello)

	if err == nil {
		sess, _ := session.Get("session", c)
		sess.Values["message"] = hello.Message
		sess.Values["origin"] = hello.Origin
		sess.Save(c.Request(), c.Response())
		go helloLogDB(hello, cc.DB)
	}

	return hello
}

func helloLogDB(hello *models.Hello, db *sql.DB) {
	log := new(models.HelloLog)
	log.SetMessage(hello.Message).SetOrigin(hello.Origin).SetIsSent(true).SetTimestamp(time.Now())
	logquery := database.NewTableQuery(db, log.GetTableName(), *log)
	logquery.Insert()
}
