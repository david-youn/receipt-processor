package api

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (r *Receipt) CalculatePoints() (int, error) {
	var points int

	fmt.Println("Breakdown:")

	// 1. One point for every alphanumeric character in the retailer name.
	alnumCount := len(regexp.MustCompile(`[A-Za-z0-9]`).FindAllString(r.Retailer, -1))
	fmt.Println(alnumCount, " points - retailer name has ", alnumCount, "characters")
	points += alnumCount

	// 2. 50 points if the total is a round dollar amount with no cents.
	if total, err := strconv.ParseFloat(r.Total, 64); err == nil {
		if total == math.Floor(total) {
			fmt.Println(50, " points - total is a round dollar amount")
			points += 50
		}
		// 3. 25 points if the total is a multiple of 0.25.
		if math.Mod(total, 0.25) == 0 {
			fmt.Println(25, " points - total is a multiple of 0.25")
			points += 25
		}
	} else {
		return 0, errors.New("invalid total format")
	}

	// 4. 5 points for every two items on the receipt.
	fmt.Println((len(r.Items)/2)*5, " points - ", len(r.Items), " items (", (len(r.Items) / 2), " pairs @ 5 points each)")
	points += (len(r.Items) / 2) * 5

	// 5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range r.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				fmt.Println(int(math.Ceil(price*0.2)), " points - ", desc, " is ", len(desc), " characters (a multiple of 3)")
				fmt.Println("item price of ", price, " * 0.2 = ", price*0.2, ", rounded up is ", int(math.Ceil(price*0.2)), " points")
				points += int(math.Ceil(price * 0.2))
			} else {
				return 0, errors.New("invalid item price format")
			}
		}
	}

	// 6. 6 points if the day in the purchase date is odd.
	if date, err := time.Parse("2006-01-02", r.PurchaseDate); err == nil {
		if date.Day()%2 != 0 {
			points += 6
			fmt.Println(6, " points - purchase day is odd")
		}
	} else {
		return 0, errors.New("invalid purchase date format")
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if purchaseTime, err := time.Parse("15:04", r.PurchaseTime); err == nil {
		if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() == 0) {
			points += 10
			fmt.Println(10, " points - ", purchaseTime, " is between 2:00pm and 4:00pm")
		}
	} else {
		return 0, errors.New("invalid purchase time format")
	}
	fmt.Println("--------------")
	fmt.Println("Total Points: ", points)
	fmt.Println()
	return points, nil
}
