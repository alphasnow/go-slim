package file_upload

type UploadChecker struct {
}

func (u *UploadChecker) Validation(rule *UploadRule, info *UploadInfo) error {
	_ = info.CorrectType()

	if err := rule.CheckSize(info.FileInfo); err != nil {
		return err
	}
	if err := rule.CheckType(info.FileInfo); err != nil {
		return err
	}

	if info.IsImage() {
		if err := rule.CheckImage(info.ImageInfo); err != nil {
			return err
		}
	}

	return nil
}
