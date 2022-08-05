package logic_test

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"testing"
	"time"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/timex"
)

func TestVideo(t *testing.T) {

	ctx := local.WithUserLocal(context.TODO(), &entity.User{
		ID:   1,
		Name: "zhangsan",
	})
	lo.Must0(logic.Collection().Add(ctx, &dto.CollectionAdd{
		Name:       "合集1",
		Type:       1,
		CategoryId: 1,
		Cover:      "conver",
		Brief:      "",
		PublishAt:  &timex.Time{time.Now()},
		Videos: []dto.VideoEasy{
			{
				Url:  "http://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/2-1%20%E7%AB%A0%E8%8A%82%E7%AE%80%E4%BB%8B%5B%E5%A4%A9%E4%B8%8B%E6%97%A0%E9%B1%BC%5D%5Bshikey.com%5D.mp4",
				Name: "第一集",
			},
		},
	}))

	lo.Must0(logic.Collection().Add(ctx, &dto.CollectionAdd{
		Name:       "合集2",
		Type:       1,
		CategoryId: 1,
		Cover:      "conver",
		Brief:      "",
		PublishAt:  &timex.Time{time.Now()},
		Videos: []dto.VideoEasy{
			{
				Url:  "http://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/2-1%20%E7%AB%A0%E8%8A%82%E7%AE%80%E4%BB%8B%5B%E5%A4%A9%E4%B8%8B%E6%97%A0%E9%B1%BC%5D%5Bshikey.com%5D.mp4",
				Name: "第一集",
			},
		},
	}))

	collections, count, err := logic.Video().GetPageCollections(context.TODO(), &dto.BackCollectionList{Pager: dto.Pager{
		Page:     1,
		PageSize: 10,
	}})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(collections, count)

	collection, err := logic.Video().GetCollection(context.TODO(), &dto.CollectionId{CollectionId: 1})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(collection)

}
