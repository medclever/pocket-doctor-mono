package repositories

import (
	"editor/app/domain"
	"editor/app/types"
	"editor/app/utils"
	"encoding/json"
	"fmt"
	"os"
)

type languageRepository struct {
	path string
}

func NewLanguageRepo(dataDir string) types.LanguageRepository {
	repo := languageRepository{
		path: fmt.Sprintf("%s/languages.json", dataDir),
	}
	return &repo
}

func (r *languageRepository) RestoreAll() (languages types.Languages, err error) {
	entyties := []types.Language{}
	if !utils.CheckFile(r.path) {
		return domain.InitLanguages(entyties), nil
	}
	data, err := os.ReadFile(r.path)
	if err != nil {
		return
	}
	languagesData := []types.LanguageData{}
	err = json.Unmarshal(data, &languagesData)
	if err != nil {
		return
	}
	for _, languageData := range languagesData {
		entyties = append(entyties, domain.InitLanguage(languageData))
	}
	return domain.InitLanguages(entyties), nil
}

func (r *languageRepository) PersistAll(languages types.Languages) (err error) {
	utils.CreateFileIfNotExist(r.path)
	data, err := json.MarshalIndent(languages.ExportData(), "", "\t")
	if err != nil {
		return
	}
	err = os.WriteFile(r.path, data, 0644)
	return
}
