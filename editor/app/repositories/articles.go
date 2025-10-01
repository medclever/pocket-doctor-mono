package repositories

import (
	"editor/app/domain"
	"editor/app/types"
	"editor/app/utils"
	"encoding/json"
	"fmt"
	"os"
)

type articleRepository struct {
	dataDir string
	path    string
}

func NewArticleRepo(dataDir string) types.ArticleRepository {
	repo := articleRepository{
		dataDir: dataDir,
		path:    fmt.Sprintf("%s/articles.json", dataDir),
	}
	return &repo
}

func (a *articleRepository) RestoreAll() (articles types.Articles, err error) {
	if !utils.CheckFile(a.path) {
		return domain.InitArticles([]types.Article{}), nil
	}
	data, err := os.ReadFile(a.path)
	if err != nil {
		return
	}
	itemsData := []types.ArticleData{}
	err = json.Unmarshal(data, &itemsData)
	if err != nil {
		return
	}
	articles = domain.InitArticles(nil)
	for _, itemData := range itemsData {
		item := domain.InitArticle(itemData)
		articles.Add(item)
	}
	return articles, nil
}

func (a *articleRepository) PersistAll(articles types.Articles) (err error) {
	utils.CreateFileIfNotExist(a.path)
	data, err := json.MarshalIndent(articles.ExportData(), "", "\t")
	if err != nil {
		return
	}
	err = os.WriteFile(a.path, data, 0644)
	return
}
