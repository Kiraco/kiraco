package middleware

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

const (
	//Low priority
	Low = iota
	//Medium priority
	Medium
	//High priority
	High
)

// FilePathName - path anf file name of the file
var FilePathName = "/api/middleware/domain.txt"

// Repository - repo (file, db) to read the mocks
type Repository interface {
	Read() []*Queue
}

// Queue - message
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + FilePathName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	queues := []*Queue{}
	queue := &Queue{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			queues = append(queues, queue)
			queue = &Queue{}
			continue
		}
		if strings.Contains(strings.ToLower(text), "weight") {
			weight, err := strconv.Atoi(strings.Split(text, ":")[1])
			if err != nil {
				log.Fatal(err)
			}
			queue.Weight = weight

		} else if strings.Contains(strings.ToLower(text), "priority") {
			priority, err := strconv.Atoi(strings.Split(text, ":")[1])
			if err != nil {
				log.Fatal(err)
			}
			queue.Priority = priority
		} else {
			queue.Domain = text
		}
	}
	return queues
}

// QueueList - list of messages
var QueueList []string

// ProxyMiddleware - proxy middleware
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")

	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "domain error"})
		return
	}
	insertElement(domain)

	c.Next()
}

func insertElement(domain string) {
	domainPriority := getPriorityLevel(domain)
	var newQueueList []string
	elementAdded := false
	if len(QueueList) > 0 {
		for index, element := range QueueList {
			elementPriority := getPriorityLevel(element)
			if elementPriority > domainPriority {
				newQueueList = append(newQueueList, element)
			} else {
				elementAdded = true
				newQueueList = append(newQueueList, domain)
				newQueueList = append(newQueueList, QueueList[index:]...)
				break
			}
		}
	} else {
		elementAdded = true
		newQueueList = append(newQueueList, domain)
	}
	QueueList = newQueueList
	if !elementAdded {
		QueueList = append(QueueList, domain)
	}
}

func getPriorityLevel(domain string) int {
	var requestPriority int
	var repo Repository
	repo = &Queue{}
	for _, row := range repo.Read() {
		if row.Domain == domain {
			if row.Priority >= 5 && row.Weight >= 5 {
				requestPriority = High
			} else if row.Priority < 5 && row.Weight < 5 {
				requestPriority = Low
			} else {
				requestPriority = Medium
			}
		}
	}
	return requestPriority
}
