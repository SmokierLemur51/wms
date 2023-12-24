package routes

import (
	"strconv"
	"log"
)

// need to validate forms 

func ConvertStrInt(f string) (int, error) {
    conversion, err := strconv.Atoi(f)
    if err != nil {
        log.Printf("Error converting %s to int.", f)
        return 0, err
    }
    return conversion, nil
}

func ConvertStrFloat64(f string) (float64, error) {
	conversion, err := strconv.ParseFloat(f, 64)
    if err != nil {
        log.Printf("Error converting %s to float64.", f)
        return 0.0, err
    }
    return conversion, nil
}

