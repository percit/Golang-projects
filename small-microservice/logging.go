package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct { //middleware decorator
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) { //named return
	defer func(start time.Time) { //this will run at the end and will say how long did function take
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
