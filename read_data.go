package main

import (
	"encoding/json"
	"github.com/andreychuk/num2word"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func GetData(path string) (data Data) {
	file, _ := ioutil.ReadFile(path)

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatalln(err)
	}
	data.prepareData()

	return data
}

func (data *Data) prepareData() {
	for i, jobs := range data.Jobs {
		amount := jobs.Price * jobs.Quantity
		data.Jobs[i].Amount = amount
		data.Jobs[i].Id = i + 1
		data.JobsAmount += amount
	}
	data.StrJobsAmount = num2word.UaAmount(data.JobsAmount, false)
}

func (data *Data) setSignOn(str string) {
	if str == "" {
		data.SignOn = time.Now()
	} else {
		data.SignOn, _ = time.Parse("2006-01-02", str)
	}
	data.setSingOnUA()
}

func (data *Data) setSingOnUA() {
	result := strconv.Itoa(data.SignOn.Day()) + " "
	m := data.SignOn.Month()
	switch m {
	case time.Month(time.January):
		result += "січня "
	case time.Month(time.February):
		result += "лютого "
	case time.Month(time.March):
		result += "березня "
	case time.Month(time.April):
		result += "квітня "
	case time.Month(time.May):
		result += "травня "
	case time.Month(time.June):
		result += "червня "
	case time.Month(time.July):
		result += "липня "
	case time.Month(time.August):
		result += "серпня "
	case time.Month(time.September):
		result += "вересня "
	case time.Month(time.October):
		result += "жовтня "
	case time.Month(time.November):
		result += "листопада "
	case time.Month(time.December):
		result += "грудня "
	}
	result += strconv.Itoa(data.SignOn.Year()) + " р."

	data.SignOnUA = result
}
