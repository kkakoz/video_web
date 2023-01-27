package bootstrap

import (
	"context"
	"fmt"
	"github.com/kkakoz/pkg/logger"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/ossx"
	"video_web/pkg/timex"
	"video_web/pkg/videos"
)

func uploadDirVideo(name, brief, uploadDir string, categoryId int64, publishAt *timex.Time) (err error) {

	if uploadDir == "" {
		return errors.New("dir is empty")
	}
	fs, err := os.ReadDir(uploadDir)
	if err != nil {
		panic(err)
	}

	bucketPath := "https://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/"

	fileMap := map[string]string{} // key name value

	for _, f := range fs {
		filePath := uploadDir + "\\" + f.Name()
		split := strings.Split(f.Name(), ".")
		name := split[0]

		fileMap[name] = filePath
	}

	ossDirName := uuid.NewV4().String() + "/"

	coverPath, ok := fileMap["cover"]
	if !ok {
		return errors.New("cover file not found")
	}

	err = ossx.Bucket.PutObjectFromFile(ossDirName+"cover", coverPath)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			curErr := ossx.Bucket.DeleteObject(ossDirName)
			if curErr != nil {
				logger.Info("delete dir err")
			}
		}

	}()

	resources := make([]*dto.Resource, 0)
	for i := 1; ; i++ {
		key := fmt.Sprintf("第%d集", i)
		file, ok := fileMap[key]
		if !ok {
			break
		}
		fileInfo, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		duration, err := videos.GetMP4Duration(fileInfo)
		err = ossx.Bucket.PutObjectFromFile(ossDirName+key, file)
		if err != nil {
			return err
		}
		resources = append(resources, &dto.Resource{
			Url:      bucketPath + ossDirName + key,
			Name:     key,
			Duration: int64(duration),
		})
	}

	ctx := local.WithUserLocal(context.TODO(), &entity.User{
		ID: 1,
	})

	_, err = logic.Video().Add(ctx, &dto.VideoAdd{
		Name:       name,
		Type:       entity.VideoTypeAnime,
		CategoryId: categoryId,
		Cover:      bucketPath + ossDirName + "cover",
		Brief:      brief,
		PublishAt:  publishAt,
		Resources:  resources,
		State:      entity.VideoStateNormal,
	})
	if err != nil {
		return err
	}
	return nil
}
