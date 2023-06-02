package file_upload

import (
	"errors"
	"sync"
)

var uploadRules map[string]*UploadRule
var uploadRulesLocker sync.RWMutex

func GetUploadRule(r string) (*UploadRule, error) {
	uploadRulesLocker.RLock()
	defer uploadRulesLocker.RUnlock()
	rule, ok := uploadRules[r]
	if ok == false {
		return nil, errors.New("rule name error")
	}
	return rule, nil
}
func SetUploadRule(k string, r *UploadRule) {
	uploadRulesLocker.Lock()
	defer uploadRulesLocker.Unlock()
	uploadRules[k] = r
}

func init() {
	uploadRules = map[string]*UploadRule{}

	imageRule := &ImageRule{
		MaxLength: 4096,
		MinLength: 64,
		MaxRadio:  4,
	}
	imageValidTypes := []string{"jpeg", "jpg", "png"}

	uploadRules["image"] = &UploadRule{
		MaxSize:    10 * 1024 * 1024,
		ValidTypes: imageValidTypes,
		ImageRule:  imageRule,
	}

	uploadRules["file"] = &UploadRule{
		MaxSize: 10 * 1024 * 1024,
	}
}
