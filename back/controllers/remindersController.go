package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type resposeGetReminders struct {
	Remindrers []responseGetReminder `json:"reminders"`
}

type responseGetReminder struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	UserName string    `json:"user_name"`
	Date     time.Time `json:"date"`
}

func (controller *Controller) GetReminders(c *gin.Context) {
	reminders := []responseGetReminder{
		responseGetReminder{
			ID:       1,
			Title:    "hoge",
			Body:     "hogehoge",
			UserName: "User1",
			Date:     time.Date(2020, 5, 20, 23, 59, 59, 0, time.Local),
		},
		responseGetReminder{
			ID:       2,
			Title:    "hoge2",
			Body:     "hogehoge2",
			UserName: "User2",
			Date:     time.Date(2020, 5, 21, 23, 59, 59, 0, time.Local),
		},
	}

	res := resposeGetReminders{reminders}
	c.JSON(http.StatusOK, res)
}

func (controller *Controller) PostReminders(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) GetReminder(c *gin.Context) {
	id := c.Param("id")
	res := responseGetReminder{}
	if id == "1" {
		res = responseGetReminder{
			ID:       1,
			Title:    "hoge",
			Body:     "hogehoge",
			UserName: "User1",
			Date:     time.Date(2020, 5, 20, 23, 59, 59, 0, time.Local),
		}

		c.JSON(http.StatusOK, res)
	} else if id == "2" {
		res = responseGetReminder{
			ID:       2,
			Title:    "hoge2",
			Body:     "hogehoge2",
			UserName: "User2",
			Date:     time.Date(2021, 5, 20, 23, 59, 59, 0, time.Local),
		}

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, nil)
	}

}
