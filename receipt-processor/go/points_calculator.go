/*
 * points_calculator.go - Receipt Points Calculation Logic
 *
 * This file implements the `calculatePoints` function to compute points for a receipt
 * based on the following 7 rules:
 * 1. One point for each alphanumeric character in the retailer's name.
 * 2. 50 points if the total is a round dollar amount with no cents.
 * 3. 25 points if the total is a multiple of 0.25.
 * 4. 5 points for every two items on the receipt.
 * 5. If an item's description length is a multiple of 3, 20% of the item's price is awarded as points.
 * 6. 6 points if the purchase day is odd.
 * 7. 10 points if the purchase time is between 2 PM and 4 PM.
 *
 * Logs are added to trace each rule's contribution to the total points.
 */
package openapi

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	for _, char := range receipt.Retailer {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			points++
		}
	}
	fmt.Println("Points after Retailer Name:", points)

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("Invalid total amount format: %v", err)
		return 0
	}

	if math.Mod(total, 1.0) == 0 {
		points += 50
		fmt.Println("Added 50 points for round dollar total.")
	} else {
		fmt.Println("No points added for round dollar total.")
	}
	fmt.Println("Points after Round Dollar Check:", points)

	if math.Mod(total, 0.25) == 0 {
		points += 25
		fmt.Println("Added 25 points for total being multiple of 0.25.")
	} else {
		fmt.Println("No points added for total being multiple of 0.25.")
	}
	fmt.Println("Points after 0.25 Check:", points)

	itemPairs := (len(receipt.Items) / 2) * 5
	points += itemPairs
	fmt.Printf("Added %d points for every two items.\n", itemPairs)
	fmt.Println("Points after Item Count Check:", points)

	for _, item := range receipt.Items {
		descLen := utf8.RuneCountInString(strings.TrimSpace(item.ShortDescription))
		itemPrice, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			log.Printf("Invalid item price format for '%s': %v", item.ShortDescription, err)
			continue
		}

		if descLen%3 == 0 {
			itemPoints := int(math.Ceil(itemPrice * 0.2))
			points += itemPoints
			fmt.Printf("Added %d points for item '%s' (desc length %d).\n", itemPoints, item.ShortDescription, descLen)
		} else {
			fmt.Printf("No points added for item '%s' (desc length %d).\n", item.ShortDescription, descLen)
		}
	}
	fmt.Println("Points after Item Description Check:", points)

	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("Invalid purchase date format: %v", err)
		return 0
	}
	if purchaseDate.Day()%2 != 0 {
		points += 6
		fmt.Println("Added 6 points for odd purchase day.")
	} else {
		fmt.Println("No points added for even purchase day.")
	}
	fmt.Println("Points after Purchase Date Check:", points)

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("Invalid purchase time format: %v", err)
		return 0
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
		fmt.Println("Added 10 points for purchase between 2 PM and 4 PM.")
	} else {
		fmt.Println("No points added for purchase outside 2 PM - 4 PM window.")
	}
	fmt.Println("Final Points after Purchase Time Check:", points)

	return points
}
