package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rezazareiii/divarbot/divar"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const BotToken = "8052520945:AAHNrbby4hZIKUhRCDZtJZ7z-519MgYc7Vk"

const ChatID = -4748240323

var visitedPosts = make(map[string]struct{})

func main() {

	var err error

	visitedPosts, err = LoadVisitedPosts("visited_posts.txt")
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	for {

		newPosts, err := fetchNewlyAddedPosts()
		if err != nil {
			log.Println("Error fetching posts:", err)
			continue
		}

		for _, post := range newPosts {
			msg := tgbotapi.NewMessage(ChatID, fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\nhttps://divar.ir/v/%s", post.Title, post.Rent, post.Credit, post.Location, post.Token))

			if _, err := bot.Send(msg); err != nil {
				fmt.Println("Error sending message:", err)
				continue
			}

			err = SaveVisitedPosts(post.Token)
			if err != nil {
				fmt.Println("Error saving visited post:", err)
				continue
			}

			time.Sleep(10 * time.Second)
		}

		time.Sleep(10 * time.Second)
	}

	// }

}

func LoadVisitedPosts(filePath string) (map[string]struct{}, error) {
	visited := make(map[string]struct{})

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			visited[line] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return visited, nil
}

func fetchNewlyAddedPosts() ([]divar.PostRowData, error) {

	postTokens, err := divar.Search()
	if err != nil {
		return nil, err
	}

	newPosts := make([]divar.PostRowData, 0, 4)

	for _, post := range postTokens {
		if _, exists := visitedPosts[post.Token]; !exists {
			newPosts = append(newPosts, post)
		}
	}

	return newPosts, nil
}

func AppendToFile(filePath, text string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text + "\n")
	if err != nil {
		return err
	}

	return nil
}

func SaveVisitedPosts(token string) error {

	visitedPosts[token] = struct{}{}
	err := AppendToFile("visited_posts.txt", token)
	if err != nil {
		return err
	}

	return nil

}
