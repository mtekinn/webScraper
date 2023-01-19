package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	var err error
	var urllist = []string{"https://jsonplaceholder.typicode.com/posts/1", "https://jsonplaceholder.typicode.com/posts/2", "https://jsonplaceholder.typicode.com/posts/3"}
	// Create an Array of Structs to store the data
	var posts []Post
	// Create a new http client with a timeout of 10 seconds
	client := &http.Client{Timeout: time.Second * 10}

	for _, url := range urllist {
		// Make a GET request to the website
		response, err := client.Get(url)
		if err != nil {
			fmt.Println("Error while making GET request: ", err)
			continue // move to next url if error occurs
		}
		defer response.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error while reading response body: ", err)
			continue // move to next url if error occurs
		}

		// Unmarshal the JSON
		var post Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("Error while unmarshaling JSON: ", err)
			continue // move to next url if error occurs
		}

		// Extract the specific data you want
		fmt.Println("Title: ", post.Title)
		fmt.Println("User ID: ", post.UserId)
		fmt.Println("Body: ", post.Body)
		posts = append(posts, post) // append the post to the posts array
	}

	// Create a new CSV file
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Error while creating CSV file: ", err)
		return
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the data to the CSV file
	err = writer.Write([]string{"title", "userId", "body"})
	if err != nil {
		fmt.Println("Error while writing to cvs file: ", err)
		return
	}

	// loop through the array of posts and write to csv file
	for _, post := range posts {
		err = writer.Write([]string{post.Title, fmt.Sprintf("%d", post.UserId), post.Body})
		if err != nil {
			fmt.Println("Error while writing to cvs file: ", err)
			return
		}
	}
}
