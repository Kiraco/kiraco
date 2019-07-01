package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	FilePathName = "/domain_test.txt"
}

func TestRead(t *testing.T) {
	var repo Repository
	repo = &Queue{}
	for index, row := range repo.Read() {
		switch index {
		case 0:
			{
				assert.Equal(t, "alpha", row.Domain)
				assert.Equal(t, 5, row.Weight)
				assert.Equal(t, 7, row.Priority)
			}
		case 1:
			{
				assert.Equal(t, "omega", row.Domain)
				assert.Equal(t, 1, row.Weight)
				assert.Equal(t, 1, row.Priority)
			}
		case 2:
			{
				assert.Equal(t, "beta", row.Domain)
				assert.Equal(t, 1, row.Weight)
				assert.Equal(t, 5, row.Priority)
			}
		}
	}
}

func TestProxyMiddleware(t *testing.T) {
	//Already tested in main_test.go
}

func TestInsertElement(t *testing.T) {
	assert.Equal(t, 0, len(QueueList))
	insertElement("omega")
	assert.Equal(t, 1, len(QueueList))
	assert.Equal(t, "omega", QueueList[0])

	insertElement("beta")
	assert.Equal(t, 2, len(QueueList))
	assert.Equal(t, "beta", QueueList[0])

	insertElement("alpha")
	assert.Equal(t, 3, len(QueueList))
	assert.Equal(t, "alpha", QueueList[0])
}

func TestGetPriorityLevel(t *testing.T) {
	level := getPriorityLevel("alpha")
	assert.Equal(t, High, level)

	level = getPriorityLevel("beta")
	assert.Equal(t, Medium, level)

	level = getPriorityLevel("omega")
	assert.Equal(t, Low, level)
}
