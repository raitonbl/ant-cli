package internal

import (
	"encoding/json"
	"github.com/raitonbl/ant/internal/project"
	"gopkg.in/yaml.v3"
	"strings"
)

type ProjectContext interface {
	GetProjectFile() *File
	GetDocument() (*project.Specification, error)
}

func GetContext(filename string) (ProjectContext, error) {
	file, err := GetFile(filename)

	if err != nil {
		return nil, err
	}

	return &DefaultContext{projectFile: file}, nil
}

type DefaultContext struct {
	projectFile       *File
	processedDocument *project.Specification
	document          *project.Specification
}

func (instance *DefaultContext) GetProjectFile() *File {
	return instance.projectFile
}

func (instance *DefaultContext) GetDocument() (*project.Specification, error) {

	if instance.document != nil {
		return instance.document, nil
	}

	if instance.GetProjectFile() == nil || instance.GetProjectFile().GetName() == "" {
		return nil, GetProblemFactory().GetConfigurationFileNotFound()
	}

	filename := instance.GetProjectFile().GetName()

	if !(strings.HasSuffix(filename, ".json") || strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml")) {
		return nil, GetProblemFactory().GetUnsupportedDescriptor()
	}

	binary := instance.GetProjectFile().GetContent()

	if strings.HasSuffix(filename, ".json") {
		return parseJson(binary)
	}

	return parseYaml(binary)
}

func parseYaml(binary []byte) (*project.Specification, error) {

	if binary == nil {
		return nil, GetProblemFactory().GetUnexpectedContext()
	}

	descriptor := project.Specification{}
	err := yaml.Unmarshal(binary, &descriptor)

	if err != nil {
		return nil, err
	}

	return &descriptor, err
}

func parseJson(binary []byte) (*project.Specification, error) {

	if binary == nil {
		return nil, GetProblemFactory().GetUnexpectedContext()
	}

	descriptor := project.Specification{}
	err := json.Unmarshal(binary, &descriptor)

	if err != nil {
		return nil, err
	}

	return &descriptor, err
}