package service

import (
// model "my-dms/model"
)

// func storePdf(pdfFile []byte, metaData, mmodel.DocumentMeta) {
//
//
// }

type DocumentDTO struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	FileName string   `json:"fileName"`
	Keywords []string `json:"keywords"`
}
