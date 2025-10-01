package app

import (
	"editor/app/domain"
	"editor/app/repositories"
	"editor/app/types"
	"editor/app/utils"
	lib_errors "errors"
	"fmt"
	"os"
)

type App struct {
	timeNow            types.TimeNow
	articles           types.Articles
	articleRepository  types.ArticleRepository
	messages           types.Messages
	messageRepository  types.MessageRepository
	languages          types.Languages
	languageRepository types.LanguageRepository
}

func New() types.App {
	rootPath := "./data"
	articleRepository := repositories.NewArticleRepo(rootPath)
	messageRepository := repositories.NewMessageRepo(rootPath)
	languageRepository := repositories.NewLanguageRepo(rootPath)
	app := App{
		timeNow:            utils.NewTimeNowDefault(),
		articles:           domain.InitArticles([]types.Article{}),
		messages:           domain.InitMessages([]types.Message{}),
		languages:          domain.InitLanguages([]types.Language{}),
		articleRepository:  articleRepository,
		messageRepository:  messageRepository,
		languageRepository: languageRepository,
	}
	return &app
}

func (a *App) SetTimeNow(timeNow types.TimeNow) types.App {
	a.timeNow = timeNow
	return a
}

func (a *App) Init() {
	articles, err := a.articleRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.articles = articles

	messages, err := a.messageRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.messages = messages

	languages, err := a.languageRepository.RestoreAll()
	PrintErrorAndExit(err)
	a.languages = languages
}

func (a *App) Persist() (err error) {
	errs := []error{}

	err = a.articleRepository.PersistAll(a.articles)
	errs = append(errs, err)

	err = a.messageRepository.PersistAll(a.messages)
	errs = append(errs, err)

	err = a.languageRepository.PersistAll(a.languages)
	errs = append(errs, err)

	if len(errs) > 0 {
		return lib_errors.Join(errs...)
	}
	return
}

func PrintErrorAndExit(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
