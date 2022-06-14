package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //for database
	"github.com/toshkentov01/task/data_service/config"
	"github.com/toshkentov01/task/data_service/pkg/errs"
	"github.com/toshkentov01/task/data_service/pkg/logger"
)

var (
	loggerTest  logger.Logger
	cfg         *config.Config
	dataService *DataService
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	path = path[0 : len(path)-7]
	path += ".env"

	if info, err := os.Stat(path); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(path)
			if err != nil {
				fmt.Println("Err:", err)
			}
		}
	} else {
		fmt.Println("Not exists")
	}

	cfg = config.Get()
	loggerTest = logger.New(cfg.LogLevel, "data_service")
	dataService = NewDataService(loggerTest)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetPost(t *testing.T) {
	type testCase struct {
		postID      int
		expectedErr error
	}

	testCases := []testCase{
		{
			postID:      0,
			expectedErr: errs.ErrNotFound,
		},
		{
			postID:      1111,
			expectedErr: nil,
		},
	}

	for _, testData := range testCases {
		t.Run("Testing GetPost method", func(t *testing.T) {
			_, err := dataService.storage.Data().GetPost(testData.postID)

			if err != testData.expectedErr {
				t.Error("Error while testing GetPost method: ", err.Error())
			}
		})
	}
}

func TestListPosts(t *testing.T) {
	type testCase struct {
		Limit, Page uint32
		ExpectedErr error
	}

	testCases := []testCase{
		{
			Limit:       100,
			Page:        3,
			ExpectedErr: nil,
		},
	}

	for _, testData := range testCases {
		t.Run("Testing ListPosts method", func(t *testing.T) {
			_, err := dataService.storage.Data().ListPosts(testData.Limit, testData.Page)

			if err != testData.ExpectedErr {
				t.Error("Error while testing ListPosts method: ", err.Error())
			}
		})
	}

}

func TestUpdatePost(t *testing.T) {
	type testCase struct {
		Title, Body string
		PostID      int
		ExpectedErr error
	}

	testCases := []testCase{
		{
			PostID:      1111,
			Title:       "Updated",
			Body:        "This is updated post",
			ExpectedErr: nil,
		},
		{
			PostID:      1000000,
			Title:       "Updated",
			Body:        "Updated",
			ExpectedErr: errs.ErrNotFound,
		},
	}

	for _, testData := range testCases {
		t.Run("Testing UpdatePost nethod", func(t *testing.T) {
			err := dataService.storage.Data().UpdatePost(testData.PostID, testData.Title, testData.Body)

			if err != testData.ExpectedErr {
				t.Error("Error while testing UpdatePost method: ", err.Error())
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	err := dataService.storage.Data().DeletePost(1111)

	if err != nil {
		t.Error("Error while testing DeletePost method: ", err.Error())
	}
}

func TestCheckForOwnership(t *testing.T) {
	_, err := dataService.storage.Data().CheckForOwnership(1111, 2251)
	if err != nil {
		t.Error("Error while testing CheckForOwnership method, error: ", err.Error())
	}
}