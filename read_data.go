package main

import (
	"encoding/json"
	"github.com/andreychuk/num2word"
	"io/ioutil"
	"time"
)

func GetData(path string) (data Data) {
	file, _ := ioutil.ReadFile(path)

	_ = json.Unmarshal([]byte(file), &data)
	data.prepareData()

	return data
}

func (data *Data) prepareData() {
	for i, jobs := range data.Jobs {
		amount := jobs.Price * jobs.Quantity
		data.Jobs[i].Amount = amount
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
}
