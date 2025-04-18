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

const filePath = "/data/visited_posts.txt"

const ChatID = -4748240323

var visitedPosts = make(map[string]struct{})

func main() {

	var err error

	visitedPosts, err = LoadVisitedPosts()
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = false

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

			time.Sleep(5 * time.Second)
		}

		time.Sleep(5 * time.Minute)
	}

	// }

}

func LoadVisitedPosts() (map[string]struct{}, error) {
	visited := make(map[string]struct{})

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		return visited, nil
	}

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

	posts1, err := divar.Search(1)
	if err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	posts2, err := divar.Search(2)
	if err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	posts3, err := divar.Search(3)
	if err != nil {
		return nil, err
	}

	posts := append(posts1, posts2...)
	posts = append(posts, posts3...)

	newPosts := make([]divar.PostRowData, 0, 4)

	for _, post := range posts {
		if _, exists := visitedPosts[post.Token]; !exists {
			newPosts = append(newPosts, post)
		}
	}

	return newPosts, nil
}

func AppendToFile(text string) error {
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
	err := AppendToFile(token)
	if err != nil {
		return err
	}

	return nil

}
