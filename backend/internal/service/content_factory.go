package service

import (
	"errors"
	"fmt"
)

type PostContent interface {
	Validate() error
	GenerateSummary() string
	GetAdditionalFields() map[string]interface{}
}

type ArticleContent struct {
	Introduction string `json:"introduction"`
	Conclusion   string `json:"conclusion"`
}

func (c *ArticleContent) Validate() error {
	if c.Introduction == "" {
		return errors.New("introduction is required for articles")
	}
	return nil
}

func (c *ArticleContent) GenerateSummary() string {
	return fmt.Sprintf("Article with introduction and conclusion. Introduction: %s", c.Introduction)
}

func (c *ArticleContent) GetAdditionalFields() map[string]interface{} {
	return map[string]interface{}{
		"type":         "article",
		"introduction": c.Introduction,
		"conclusion":   c.Conclusion,
	}
}

type TutorialContent struct {
	Prerequisites string   `json:"prerequisites"`
	SkillLevel    string   `json:"skill_level"`
	Steps         []string `json:"steps"`
}

func (c *TutorialContent) Validate() error {
	if len(c.Steps) == 0 {
		return errors.New("at least one step is required for tutorials")
	}
	return nil
}

func (c *TutorialContent) GenerateSummary() string {
	return fmt.Sprintf("Tutorial for %s level with %d steps", c.SkillLevel, len(c.Steps))
}

func (c *TutorialContent) GetAdditionalFields() map[string]interface{} {
	return map[string]interface{}{
		"type":          "tutorial",
		"prerequisites": c.Prerequisites,
		"skill_level":   c.SkillLevel,
		"steps":         c.Steps,
	}
}

type ReviewContent struct {
	Rating    int      `json:"rating"`
	Product   string   `json:"product"`
	Pros      []string `json:"pros"`
	Cons      []string `json:"cons"`
	Recommend bool     `json:"recommend"`
}

func (c *ReviewContent) Validate() error {
	if c.Rating < 1 || c.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	if c.Product == "" {
		return errors.New("product name is required for reviews")
	}
	return nil
}

func (c *ReviewContent) GenerateSummary() string {
	recommendText := "not recommended"
	if c.Recommend {
		recommendText = "recommended"
	}
	return fmt.Sprintf("%s of %s with rating %d (%s)",
		c.Product, "review", c.Rating, recommendText)
}

func (c *ReviewContent) GetAdditionalFields() map[string]interface{} {
	return map[string]interface{}{
		"type":      "review",
		"rating":    c.Rating,
		"product":   c.Product,
		"pros":      c.Pros,
		"cons":      c.Cons,
		"recommend": c.Recommend,
	}
}

type ContentFactory struct{}

func (f *ContentFactory) CreateContent(contentType string, fields map[string]interface{}) (PostContent, error) {
	switch contentType {
	case "article":
		return &ArticleContent{
			Introduction: getStringField(fields, "introduction"),
			Conclusion:   getStringField(fields, "conclusion"),
		}, nil
	case "tutorial":
		return &TutorialContent{
			Prerequisites: getStringField(fields, "prerequisites"),
			SkillLevel:    getStringField(fields, "skill_level"),
			Steps:         getStringSliceField(fields, "steps"),
		}, nil
	case "review":
		return &ReviewContent{
			Rating:    getIntField(fields, "rating"),
			Product:   getStringField(fields, "product"),
			Pros:      getStringSliceField(fields, "pros"),
			Cons:      getStringSliceField(fields, "cons"),
			Recommend: getBoolField(fields, "recommend"),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
}

// Helper functions for field extraction
func getStringField(fields map[string]interface{}, key string) string {
	if val, ok := fields[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getStringSliceField(fields map[string]interface{}, key string) []string {
	if val, ok := fields[key]; ok {
		if slice, ok := val.([]string); ok {
			return slice
		}
	}
	return []string{}
}

func getIntField(fields map[string]interface{}, key string) int {
	if val, ok := fields[key]; ok {
		if i, ok := val.(int); ok {
			return i
		}
	}
	return 0
}

func getBoolField(fields map[string]interface{}, key string) bool {
	if val, ok := fields[key]; ok {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}
