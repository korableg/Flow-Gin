package engine

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	fl "github.com/korableg/flow"
	"github.com/korableg/flow/errs"
	"github.com/korableg/flow/leveldb"
	"github.com/korableg/flow/repo"
	"github.com/korableg/mini-gin/config"
	"net/http"
	"strconv"
)

var engine *gin.Engine
var flow *fl.Flow

func init() {

	if config.Debug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var db repo.DB
	switch config.DBProvider() {
	case "leveldb":
		db = leveldb.New(config.LevelDB().Path)
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

	engine.POST("/message/tohub/:nodeFrom/:hubTo", sendMessageToHub)
	engine.POST("/message/tonode/:nodeFrom/:nodeTo", sendMessageToNode)
	engine.GET("/message/:name", getMessage)
	engine.DELETE("/message/:name", deleteMessage)

	flow = fl.New(db)

}

func Run() {
	go func() {
		err := engine.Run(config.Address())
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func Close() error {
	return flow.Close()
}

func defaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", fmt.Sprintf("Flow:%s", config.Version()))
	}
}

func pageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, errs.New(errs.ErrPageNotFound))
}

func methodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, errs.New(errors.New("method is not allowed")))
}

func getAllNodes(c *gin.Context) {
	nodes := flow.GetAllNodes()
	c.JSON(http.StatusOK, nodes)
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
	if err := flow.DeleteNode(name); err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	c.Status(http.StatusOK)

}

func getAllHubs(c *gin.Context) {
	hubs := flow.GetAllHubs()
	c.JSON(http.StatusOK, hubs)
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
	if err := flow.DeleteHub(name); err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	c.Status(http.StatusOK)

}

func sendMessageToHub(c *gin.Context) {

	nodeFrom := c.Params.ByName("nodeFrom")
	hubTo := c.Params.ByName("hubTo")

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	_, err = flow.SendMessageToHub(nodeFrom, hubTo, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err))
	}
	c.Status(http.StatusOK)

}

func sendMessageToNode(c *gin.Context) {

	nodeFrom := c.Params.ByName("nodeFrom")
	nodeTo := c.Params.ByName("nodeTo")

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err))
		return
	}

	_, err = flow.SendMessageToNode(nodeFrom, nodeTo, data)
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
