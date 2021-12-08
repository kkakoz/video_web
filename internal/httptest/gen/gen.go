package gen

import (
	"math/rand"
	"time"
)

var (
	familiyNames = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳"}
	firstNames   = []string{"金", "木", "水", "火", "土", "春", "夏", "秋", "冬", "山", "石", "田", "天", "地", "玄", "黄", "宇", "宙", "洪", "荒"}
	// 辈分
	generationNameMap = make(map[string][]string)

	testWord = []string{}
)

func init() {
	rand.Seed(time.Now().Unix())
	for _, ln := range familiyNames {
		if ln != "欧阳" {
			generationNameMap[ln] = []string{"飞", "前", "茂", "百", "方", "书", "生", "无", "一", "用"}
		}
	}
	columna := 'a'
	for i := 0;i < 26; i++ {
		testWord = append(testWord, string(columna))
		columna++
	}
	columna = 'A'
	for i := 0;i < 26; i++ {
		testWord = append(testWord, string(columna))
		columna++
	}
}

func GetName() string {
	familiyName := familiyNames[GetInt(len(familiyNames)-1)]
	middleName := generationNameMap[familiyName][GetInt(len(generationNameMap[familiyName])-1)]
	firstName := firstNames[GetInt(len(firstNames)-1)]
	return familiyName + middleName + firstName
}

func GetInt(i int) int {
	if i == 0 {
		return rand.Int()
	}
	j := int64(i)
	return int(rand.Int63n(j))
}

func GetString(length int) string {
	s := ""
	for i := 0; i < length; i++{
		s += testWord[GetInt(len(testWord))]
	}
	return s
}

