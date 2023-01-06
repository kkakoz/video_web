package logic

import (
	"github.com/kkakoz/pkg/redisx"
	"sync"
	"time"
	"video_web/internal/model/dto"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/keys"
)

type statisticsLogic struct{}

var statisticsOnce sync.Once
var _statistics *statisticsLogic

func Statistics() *statisticsLogic {
	statisticsOnce.Do(func() {
		_statistics = &statisticsLogic{}
	})
	return _statistics
}

func (s *statisticsLogic) CalculateUV(req *dto.CalculateUV) (*vo.StatisticsUV, error) {
	redisKeys := make([]string, 0)
	cur := req.StartAt.Time
	for cur.Before(req.EndAt.Time.Add(time.Hour * 24)) {
		redisKeys = append(redisKeys, keys.UniqueVisitorKey(cur))
		cur = cur.Add(time.Hour * 24)
	}

	mergeKey := keys.UniqueVisitorRangeKey(req.StartAt.Time, req.EndAt.Time)
	_, err := redisx.Client().PFMerge(mergeKey, redisKeys...).Result()
	if err != nil {
		return nil, err
	}
	result, err := redisx.Client().PFCount(mergeKey).Result()
	if err != nil {
		return nil, err
	}
	return &vo.StatisticsUV{Count: result}, nil
}

func (s *statisticsLogic) CalculateDAU(req *dto.CalculateDAU) (*vo.StatisticsUV, error) {
	redisKeys := make([]string, 0)
	cur := req.StartAt.Time
	for cur.Before(req.EndAt.Time.Add(time.Hour * 24)) {
		redisKeys = append(redisKeys, keys.UniqueVisitorKey(cur))
		cur = cur.Add(time.Hour * 24)
	}

	mergeKey := keys.DailyActiveUserRangeKey(req.StartAt.Time, req.EndAt.Time)
	_, err := redisx.Client().BitOpOr(mergeKey, redisKeys...).Result()
	if err != nil {
		return nil, err
	}
	result, err := redisx.Client().PFCount(mergeKey).Result()
	if err != nil {
		return nil, err
	}
	return &vo.StatisticsUV{Count: result}, nil
}
