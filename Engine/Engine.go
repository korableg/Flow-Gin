package Engine

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/korableg/mini-gin/Config"
	fl "github.com/korableg/mini-gin/flow"
	"github.com/korableg/mini-gin/flow/errs"
	"github.com/korableg/mini-gin/flow/leveldb"
	"net/http"
	"strconv"
)

var engine *gin.Engine
var flow *fl.Flow

func init() {

	if Config.Debug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.New()
	engine.Use(defaultHeaders())

	engine.NoRoute(pageNotFound)
	engine.NoMethod(methodNotAllowed)

	engine.GET("/node", getAllNodes)
	engine.GET("/node/:name", getNode)
	engine.POST("/node/:name", newNode)
	engine.DELETE("/node/:name", deleteNode)

	engine.GET("/hub", getAllHubs)
	engine.GET("/hub/:name", getHub)
	engine.POST("/hub/:name", newHub)
	engine.PATCH("/hub/:action/:nameHub/:nameNode", patchHub)
	engine.DELETE("/hub/:name", deleteHub)

	engine.POST("/message/:nameNode/:nameHub", sendMessage)
	engine.GET("/message/:name", getMessage)
	engine.DELETE("/message/:name", deleteMessage)

	factory := leveldb.New(".")
	flow = fl.New(factory)

}

func Run() {
	go func() {
		err := engine.Run(Config.Address())
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func Close() {
	flow.Close()
}

func defaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", fmt.Sprintf("Flow:%s", Config.Version()))
	}
}

func pageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, errs.New(errs.ERR_PAGE_NOT_FOUND))
}

func methodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, errs.New(errors.New("method is not allowed")))
}

func getAllNodes(c *gin.Context) {
	c.JSON(http.StatusOK, flow.GetAllNodes())
}

func getNode(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, flow.GetNode(name))
}

func newNode(c *gin.Context) {

	name := c.Params.ByName("name")
	careful := c.Query("careful") == "true"

	if n, err := flow.NewNode(name, careful); err == nil {
		c.JSON(http.StatusCreated, n)
	} else {
		c.JSON(http.StatusBadRequest, errs.New(err))
	}
}

func deleteNode(c *gin.Context) {

	name := c.Params.ByName("name")
	flow.DeleteNode(name)

	c.Status(http.StatusOK)

}

func getAllHubs(c *gin.Context) {
	c.JSON(http.StatusOK, flow.GetAllHubs())
}

func getHub(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, flow.GetHub(name))
}

func newHub(c *gin.Context) {
	name := c.Params.ByName("name")
	if n, err := flow.NewHub(name); err == nil {
		c.JSON(http.StatusCreated, n)
	} else {
		c.JSON(http.StatusBadRequest, errs.New(err))
	}
}

func patchHub(c *gin.Context) {

	nameHub := c.Params.ByName("nameHub")
	nameNode := c.Params.ByName("nameNode")
	action := c.Params.ByName("action")

	var err error

	switch action {
	case "addnode":
		err = flow.AddNodeToHub(nameHub, nameNode)
	case "deletenode":
		err = flow.DeleteNodeFromHub(nameHub, nameNode)
	default:
		err = errs.New(errors.New("action not allowed"))
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err))
		return
	}

	c.JSON(http.StatusOK, flow.GetHub(nameHub))

}

func deleteHub(c *gin.Context) {

	name := c.Params.ByName("name")
	flow.DeleteHub(name)

	c.Status(http.StatusOK)

}

func sendMessage(c *gin.Context) {

	nameNode := c.Params.ByName("nameNode")
	nameHub := c.Params.ByName("nameHub")

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	_, err = flow.SendMessage(nameNode, nameHub, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err))
	}
	c.Status(http.StatusOK)

}

func getMessage(c *gin.Context) {

	name := c.Params.ByName("name")

	m, err := flow.GetMessage(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err))
		return
	}
	if m == nil {
		c.Status(http.StatusNoContent)
		return
	}

	contentLength, err := c.Writer.Write(m.Data())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	c.Header("Message-ID", strconv.FormatInt(m.ID(), 16))
	c.Header("Message-From", m.From())
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", strconv.Itoa(contentLength))

	c.Status(http.StatusOK)

}

func deleteMessage(c *gin.Context) {

	name := c.Params.ByName("name")

	err := flow.RemoveMessage(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err))
		return
	}

	c.Status(http.StatusOK)

}
