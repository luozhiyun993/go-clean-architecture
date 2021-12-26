package service

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"my-clean-rchitecture/mock"
	"testing"
	"time"
)

func Test_articleService_Fetch(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	now := time.Now()
	mockRepo := mock.NewMockIArticleRepo(ctl)

	gomock.InOrder(
		mockRepo.EXPECT().Fetch(context.TODO(), now, 10).Return(nil, nil),
	)
	
	service := NewArticleService(mockRepo)

	fetch, _ := service.Fetch(context.TODO(), now, 10)
	fmt.Println(fetch)
}
