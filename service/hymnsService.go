package service

import (
	"context"
	"log"
	"math"
	"math/rand"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/pojos"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/samber/lo"
)

// RANDOM数
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// 計算マップ1
var termToIndex map[string]int

// 計算マップ2
var docFreq map[string]int

// コーパスサイズ
var corpusSize int

func CountHymnsAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
		).Count(ctx)
}

func GetHymnById(id int64) (*ent.Hymn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().Where(
		hymn.VisibleFlg(true),
		hymn.ID(id),
	).Only(ctx)
}

func CountHymnsByKeyword(keyword string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).Count(ctx)
}

func GetHymnsByKeyword(keyword string, pageNum int) ([]pojos.HymnDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	offset := (pageNum - 1) * int(common.DefaultPageSize)
	hymns, err := EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).
		Order(hymn.ByID()).
		Limit(int(common.DefaultPageSize)).
		Offset(offset).All(ctx)
	return map2DTOs(hymns, pojos.LineNumber(5)), err
}

func GetHymnsRandomFive(keyword string) ([]pojos.HymnDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for _, strange := range common.StrangeArray {
		if strings.Contains(strings.ToLower(keyword), strange) || len(keyword) >= 100 {
			hymns, err := EntCore.Hymn.Query().
				Where(hymn.VisibleFlg(true)).
				Order(hymn.ByID()).
				Limit(int(common.DefaultPageSize)).All(ctx)
			log.Printf("怪しいキーワード: %v\n", keyword)
			return map2DTOs(hymns, pojos.LineNumber(5)), err
		}
	}
	if keyword == common.EmptyString {
		hymns, err := EntCore.Hymn.Query().
			Where(hymn.VisibleFlg(true)).All(ctx)
		hymnDtos := map2DTOs(hymns, pojos.LineNumber(5))
		return randomFiveLoop2(hymnDtos), err
	}
	hymns, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpEQ(keyword),
				hymn.NameKrEQ(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains("["+keyword+"]"),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	withName := map2DTOs(hymns, pojos.LineNumber(1))
	hymnDtos := slices.Clone(withName)
	withNameIds := lo.Map(withName, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	hymns2, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns2 = lo.Filter(hymns2, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID)
	})
	withNameLike := map2DTOs(hymns2, pojos.LineNumber(2))
	hymnDtos = append(hymnDtos, withNameLike...)
	if len(hymnDtos) >= int(common.DefaultPageSize) {
		return randomFiveLoop2(hymnDtos), err
	}
	withNameLikeIds := lo.Map(withNameLike, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	keyword = tools.GetDetailKeyword(keyword)
	hymns3, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.SerifContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns3 = lo.Filter(hymns3, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID) && !lo.Contains(withNameLikeIds, a.ID)
	})
	withSerifLike := map2DTOs(hymns3, pojos.LineNumber(3))
	hymnDtos = append(hymnDtos, withSerifLike...)
	if len(hymnDtos) >= int(common.DefaultPageSize) {
		return randomFiveLoop2(hymnDtos), err
	}
	withSerifLikeIds := lo.Map(withSerifLike, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	hymns4, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true)).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns4 = lo.Filter(hymns4, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID) && !lo.Contains(withNameLikeIds, a.ID) && !lo.Contains(withSerifLikeIds, a.ID)
	})
	filteredRecords := map2DTOs(hymns4, pojos.LineNumber(5))
	hymnDtos = randomFiveLoop(hymnDtos, filteredRecords)
	return hymnDtos, err
}

func GetHymnsKanumi(id int64) ([]pojos.HymnDTO, error) {
	hymnById, err := GetHymnById(id)
	if err != nil {
		return nil, err
	}
	hymnDtos := make([]pojos.HymnDTO, 0)
	hymnDto := pojos.HymnDTO{
		ID:          strconv.FormatInt(hymnById.ID, 10),
		NameJP:      hymnById.NameJp,
		NameKR:      hymnById.NameKr,
		Serif:       hymnById.Serif,
		Link:        hymnById.Link,
		Score:       nil,
		Biko:        common.EmptyString,
		UpdatedUser: strconv.FormatInt(hymnById.UpdatedUser, 10),
		UpdatedTime: hymnById.UpdatedTime,
		LineNumber:  pojos.LineNumber(2),
	}
	hymnDtos = append(hymnDtos, hymnDto)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	hymns, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.IDNEQ(id),
		).All(ctx)
	matchHymns := findMatches(hymnById.Serif, hymns)
	matchDtos := map2DTOs(matchHymns, pojos.LineNumber(3))
	hymnDtos = append(hymnDtos, matchDtos...)
	return hymnDtos, err
}

// 任意の５つの賛美歌を選択する
func randomFiveLoop(hymnsRecords, totalRecords []pojos.HymnDTO) []pojos.HymnDTO {
	idSet := make(map[string]struct{})
	for _, h := range hymnsRecords {
		idSet[h.ID] = struct{}{}
	}
	var filteredRecords []pojos.HymnDTO
	for _, item := range totalRecords {
		if _, exists := idSet[item.ID]; !exists {
			filteredRecords = append(filteredRecords, item)
		}
	}
	concernList1 := slices.Clone(hymnsRecords) // コピー
	if len(hymnsRecords) < int(common.DefaultPageSize) {
		sagaku := int(common.DefaultPageSize) - len(hymnsRecords)
		for range sagaku {
			index := random.Intn(len(filteredRecords))
			concernList1 = append(concernList1, filteredRecords[index])
		}
	}
	concernList2 := distinctHymnDtos(concernList1)
	if len(concernList2) == int(common.DefaultPageSize) {
		return concernList2
	}
	return randomFiveLoop(concernList2, filteredRecords)
}

// 任意の５つの賛美歌を選択する
func randomFiveLoop2(hymnsRecords []pojos.HymnDTO) []pojos.HymnDTO {
	var concernList1 []pojos.HymnDTO
	for range int(common.DefaultPageSize) {
		index := random.Intn(len(hymnsRecords))
		concernList1 = append(concernList1, hymnsRecords[index])
	}
	concernList2 := distinctHymnDtos(concernList1)
	if len(concernList2) == int(common.DefaultPageSize) {
		return concernList2
	}
	return randomFiveLoop(concernList2, hymnsRecords)
}

// 重複したエレメントを除外する
func distinctHymnDtos(input []pojos.HymnDTO) []pojos.HymnDTO {
	seen := make(map[string]struct{})
	var result []pojos.HymnDTO
	for _, h := range input {
		if _, exists := seen[h.ID]; !exists {
			seen[h.ID] = struct{}{}
			result = append(result, h)
		}
	}
	return result
}

// 韓国語単語を取得する
func analyzeKorean(koreanText string) ([]string, error) {
	execDir, _ := getExecutableDir() // main.go と同じ階層のスクリプトのパスを取得
	scriptPath := filepath.Join(execDir, "komoran.py")
	if err != nil {
		return nil, err
	}
	out, err := exec.Command("python3", scriptPath, koreanText).Output()
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(out)), nil
}

// DTOへマップする
func map2DTOs(hymns []*ent.Hymn, lineNo pojos.LineNumber) []pojos.HymnDTO {
	return lo.Map(hymns, func(hm *ent.Hymn, _ int) pojos.HymnDTO {
		return pojos.HymnDTO{
			ID:          strconv.FormatInt(hm.ID, 10),
			NameJP:      hm.NameJp,
			NameKR:      hm.NameKr,
			Serif:       hm.Serif,
			Link:        hm.Link,
			Score:       nil,
			Biko:        common.EmptyString,
			UpdatedUser: strconv.FormatInt(hm.UpdatedUser, 10),
			UpdatedTime: hm.UpdatedTime,
			LineNumber:  lineNo,
		}
	})
}

// テキストによって韓国語単語を取得する
func tokenizeKoreanTextWithFq(originalText string) map[string]int {
	// \p{Hangul} に相当するGoの正規表現
	regex := regexp.MustCompile(`\p{Hangul}`)
	var builder strings.Builder
	for _, ch := range originalText {
		if regex.MatchString(string(ch)) {
			builder.WriteRune(ch)
		}
	}
	koreanText := builder.String()
	koreanTokens, err := analyzeKorean(koreanText)
	if err != nil {
		log.Println(err)
		return nil
	}
	return lo.CountValues(koreanTokens)
}

// コーパスを取得する
func preprocessCorpus(originalTexts []string) {
	termToIndex = map[string]int{}
	docFreq = map[string]int{}
	corpusSize = len(originalTexts)
	// 第一遍：建立文档频率
	for _, doc := range originalTexts {
		termFreq := tokenizeKoreanTextWithFq(doc)
		keySet := lo.Keys(termFreq)
		for _, term := range keySet {
			docFreq[term] += 1
		}
	}
	// 第二遍：建立词汇表索引
	docKeySet := lo.Keys(docFreq)
	for index, doc := range docKeySet {
		termToIndex[doc] = index
	}
}

// TF-IDFベクターを計算する
func computeTFIDFVector(originalText string) []float64 {
	termFreq := tokenizeKoreanTextWithFq(originalText)
	// 総単語数の計算（TF の分母）
	termVals := lo.Values(termFreq)
	totalTerms := lo.Reduce(termVals, func(agg int, val int, _ int) int {
		return agg + val
	}, 0)
	// 結果ベクトルの初期化（全部 0.0）
	vector := make([]float64, len(termToIndex))
	// TF-IDF の計算と格納
	for term, count := range termFreq {
		index, ok := termToIndex[term]
		if !ok {
			continue
		}
		tf := float64(count) / float64(totalTerms)
		df := docFreq[term] // getOrDefault 相当で 0 に初期化される
		idf := math.Log(float64(corpusSize) / float64(df+1))
		vector[index] = tf * idf
	}
	return vector
}

// コサイン類似度を計算する
func cosineSimilarity(vectorA []float64, vectorB []float64) float64 {
	dotProduct := 0.00
	normA := 0.00
	normB := 0.00
	for i := range vectorA {
		dotProduct += vectorA[i] * vectorB[i]
		normA += math.Pow(vectorA[i], 2)
		normB += math.Pow(vectorB[i], 2)
	}
	if normA == 0 || normB == 0 {
		return 0.00
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// 歌詞が似てる三つの賛美歌を取得する
func findMatches(target string, hymns []*ent.Hymn) []*ent.Hymn {
	serifs := lo.Map(hymns, func(hm *ent.Hymn, _ int) string {
		return hm.Serif
	})
	preprocessCorpus(serifs)
	targetVector := computeTFIDFVector(target)
	elementVectors := lo.Map(hymns, func(hm *ent.Hymn, _ int) []float64 {
		return computeTFIDFVector(hm.Serif)
	})
	// mapを定義する
	heapMap := map[*ent.Hymn]float64{}
	for i := range hymns {
		similarity := cosineSimilarity(targetVector, elementVectors[i])
		heapMap[hymns[i]] = similarity
	}
	pairs := lo.Entries(heapMap)
	slices.SortFunc(pairs, func(a, b lo.Entry[*ent.Hymn, float64]) int {
		if a.Value > b.Value {
			return -1
		} else if a.Value < b.Value {
			return 1
		}
		return 0 // 降順なので `-` をつける
	})
	matches := []*ent.Hymn{}
	for _, pair := range pairs {
		matches = append(matches, pair.Key)
	}
	return lo.Slice(matches, 0, 3)
}

func getExecutableDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return common.EmptyString, err
	}
	return filepath.Dir(exePath), nil
}
