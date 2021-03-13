package api

import (
	"net/http"
	"shorturl/database"
	"shorturl/model"
	"shorturl/util"

	"github.com/gin-gonic/gin"
)

func HandleRedirect(c *gin.Context) {
	shortName := c.Param("shortname")

	URL, stateCode := database.Find(shortName)
	if stateCode == model.NotFound {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	//c.JSON(http.StatusOK, "Hello %s", URL)
	c.Redirect(http.StatusMovedPermanently, URL)
}

func AddShortName(c *gin.Context) {
	shortName := c.PostForm("shortname")
	URL := c.PostForm("url")

	if shortName == "" {
		shortName = util.GetNewShortName()
	} else if util.CheckValid(shortName) == false {
		shortName = util.GetNewShortName()
	}

	newShort := model.ShorturlSturct{
		Shortname: shortName,
		URL:       URL,
	}

	database.Insert(newShort)
	//c.String(http.StatusOK, "shortName is : %s, URL is : %s", shortName, URL)
	c.JSON(http.StatusOK, util.SendResponse(200, newShort.Shortname))
}
