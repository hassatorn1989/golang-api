package models

type SubjectTypeAnswerModel struct {
	ID        string `gorm:"primaryKey" json:"id"`
	SubjectID string `json:"subject_id"`
	Answer    string `json:"answer"`
}

func (SubjectTypeAnswerModel) TableName() string {
	return "subject_type_answers"
}
