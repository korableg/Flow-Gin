package Engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/korableg/mini-gin/Config"
	"github.com/korableg/mini-gin/Mini"
	"github.com/korableg/mini-gin/Mini/Errors"
	"net/http"
	"strconv"
)

var engine *gin.Engine
var mini *Mini.Mini

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
	engine.PATCH("/hub/:nameHub/addnode/:nameNode", addNodeToHub)
	engine.PATCH("/hub/:nameHub/deletenode/:nameNode", deleteNodeFromHub)
	engine.DELETE("/hub/:name", deleteHub)

	engine.POST("/node/:nameNode/message/:nameHub", sendMessage)
	engine.GET("/node/:name/message", getMessage)
	engine.DELETE("/node/:name/message", deleteMessage)

	mini = Mini.NewMini()

}

func Run() {
	go func() {
		err := engine.Run(Config.Address())
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func defaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", fmt.Sprintf("Mini:%s", Config.Version()))
	}
}

func pageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Errors.NewError(Errors.ERR_PAGE_NOT_FOUND))
}

func methodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Errors.NewError(Errors.ERR_METHOD_NOT_ALLOWED))
}

func getAllNodes(c *gin.Context) {
	c.JSON(http.StatusOK, mini.GetAllNodes())
}

func getNode(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, mini.GetNode(name))
}

func newNode(c *gin.Context) {
	name := c.Params.ByName("name")
	if n, err := mini.NewNode(name); err == nil {
		c.JSON(http.StatusCreated, n)
	} else {
		c.JSON(http.StatusBadRequest, Errors.NewError(err))
	}
}

func deleteNode(c *gin.Context) {

	name := c.Params.ByName("name")
	mini.DeleteNode(name)

	c.Status(http.StatusOK)

}

func getAllHubs(c *gin.Context) {
	c.JSON(http.StatusOK, mini.GetAllHubs())
}

func getHub(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, mini.GetHub(name))
}

func newHub(c *gin.Context) {
	name := c.Params.ByName("name")
	if n, err := mini.NewHub(name); err == nil {
		c.JSON(http.StatusCreated, n)
	} else {
		c.JSON(http.StatusBadRequest, Errors.NewError(err))
	}
}

func addNodeToHub(c *gin.Context) {

	nameHub := c.Params.ByName("nameHub")
	nameNode := c.Params.ByName("nameNode")

	hub := mini.GetHub(nameHub)
	if hub == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_HUB_NOT_FOUND))
		return
	}
	node := mini.GetNode(nameNode)
	if node == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_NODE_NOT_FOUND))
		return
	}

	mini.AddNodeToHub(hub, node)

	c.JSON(http.StatusOK, hub)

}

func deleteNodeFromHub(c *gin.Context) {

	nameHub := c.Params.ByName("nameHub")
	nameNode := c.Params.ByName("nameNode")

	hub := mini.GetHub(nameHub)
	if hub == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_HUB_NOT_FOUND))
		return
	}
	node := mini.GetNode(nameNode)
	if node == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_NODE_NOT_FOUND))
		return
	}

	mini.DeleteNodeFromHub(hub, node)

	c.JSON(http.StatusOK, hub)

}

func deleteHub(c *gin.Context) {

	name := c.Params.ByName("name")
	mini.DeleteHub(name)

	c.Status(http.StatusOK)

}

func sendMessage(c *gin.Context) {

	nameNode := c.Params.ByName("nameNode")
	nameHub := c.Params.ByName("nameHub")

	node := mini.GetHub(nameNode)
	if node == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_NODE_NOT_FOUND))
		return
	}

	hub := mini.GetHub(nameHub)
	if hub == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_HUB_NOT_FOUND))
		return
	}

	//TODO доделать
	//mini.SendMessage(node, hub, )

}

func getMessage(c *gin.Context) {

	name := c.Params.ByName("name")
	node := mini.GetNode(name)
	if node == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_NODE_NOT_FOUND))
		return
	}

	m := mini.GetMessage(node)

	if m == nil {
		c.Status(http.StatusNoContent)
		return
	}

	contentLength, err := c.Writer.Write(m.Data())
	if err != nil {
		c.JSON(http.StatusInternalServerError, Errors.NewError(err))
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
	node := mini.GetNode(name)
	if node == nil {
		c.JSON(http.StatusBadRequest, Errors.NewError(Errors.ERR_NODE_NOT_FOUND))
		return
	}

	mini.RemoveMessage(node)

	c.Status(http.StatusOK)

}
