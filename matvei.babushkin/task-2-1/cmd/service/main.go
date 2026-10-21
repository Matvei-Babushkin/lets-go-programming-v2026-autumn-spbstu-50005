package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minTemperature = 15
	maxTemperature = 30
	invalidResult  = -1

	minCount = 1
	maxCount = 1000
)

var (
	errInvalidOperator    = errors.New("invalid operator")
	errInvalidWorkerCount = errors.New("invalid worker count")
)

type temperatureRange struct {
	min int
	max int
}

func newTemperatureRange() temperatureRange {
	return temperatureRange{
		min: minTemperature,
		max: maxTemperature,
	}
}

func (r *temperatureRange) update(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature > r.min {
			r.min = temperature
		}

	case "<=":
		if temperature < r.max {
			r.max = temperature
		}

	default:
		return fmt.Errorf("%w: %s", errInvalidOperator, operator)
	}

	return nil
}

func (r *temperatureRange) optimalTemperature() int {
	if r.min > r.max {
		return invalidResult
	}

	return r.min
}

func processDepartment() error {
	var workerCount int

	if _, err := fmt.Scan(&workerCount); err != nil {
		return fmt.Errorf("failed to read worker count: %w", err)
	}

	if workerCount < minCount || workerCount > maxCount {
		return fmt.Errorf("%w: %d", errInvalidWorkerCount, workerCount)
	}

	bounds := newTemperatureRange()

	for range workerCount {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Scan(&operator, &temperature); err != nil {
			return fmt.Errorf("failed to read temperature: %w", err)
		}

		if err := bounds.update(operator, temperature); err != nil {
			return err
		}

		if _, err := fmt.Println(bounds.optimalTemperature()); err != nil {
			return fmt.Errorf("failed to print temperature: %w", err)
		}
	}

	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		log.Print(err)

		return
	}

	if departmentCount < minCount || departmentCount > maxCount {
		log.Print("invalid department count")

		return
	}

	for range departmentCount {
		if err := processDepartment(); err != nil {
			log.Print(err)

			return
		}
	}
}
