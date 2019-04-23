package main

import "time"

type (

	// FOP structure
	FOP struct {
		Name   string `json:"name"`
		DRFO   string `json:"drfo"`
		RR     string `json:"rr"`
		Bank   string `json:"bank"`
		MFO    string `json:"mfo"`
		Person string `json:"person"`
	}

	// JobItems structure
	JobItems struct {
		Job      string  `json:"job"`
		Quantity float64 `json:"quantity"`
		Price    float64 `json:"price"`
		Amount   float64 `json:"amount"`
		Id       int     `json:"id"`
	}

	// Data structure
	Data struct {
		Provider      FOP        `json:"provider"`
		Recipient     FOP        `json:"recipient"`
		Jobs          []JobItems `json:"jobs"`
		JobsAmount    float64    `json:"jobs_amount"`
		Act           string     `json:"act"`
		StrJobsAmount string     `json:"str_jobs_amount"`
		SignOn        time.Time  `json:"sign_on"`
		SignOnUA      string     `json:"sign_on_ua"`
	}
)
