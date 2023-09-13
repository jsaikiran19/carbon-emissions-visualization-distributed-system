package models

import ("time")

type Log struct {
	Message string `json:"message"`
	Username string `json:"username"`
	Timestamp time.Time `json:"timestamp"`
} 