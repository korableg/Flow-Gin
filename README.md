# Flow-Gin
RESTful implementation of Flow message queue framework (https://github.com/korableg/flow) by Gin web framework (https://github.com/gin-gonic/gin)

# Methods
|METHOD|PATH|DESCRIPTION|
| ------ | ------ | ------ |
|GET|/node|Getting all nodes|
|GET|/node/:name|Getting node by name|
|POST|/node/:name?careful=false or true|Add new node|
|DELETE|/node/:name|Delete node by name|
|GET|/hub|Get all hubs|
|GET|/hub/:name|Get hub by name|
|POST|/hub/:name|Add new hub|
|PATCH|/hub/:action/:nameHub/:nameNode|Action must be addnode or deletenode|
|DELETE|/hub/:name|Delete hub|
|POST|/message/tohub/:nodeFrom/:hubTo|Post message, body in query-payload|
|POST|/message/tonode/:nodeFrom/:nodeTo|Post message, body in query-payload|
|GET|/message/:name|Get message from node|
|DELETE|/message/:name|Delete message|
