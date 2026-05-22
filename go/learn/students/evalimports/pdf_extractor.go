/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * PDF text extraction using pure Go library — no system dependencies.
 */
package evalimports

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

// ExtractTextFromPDF extracts text from all pages of a PDF document.
// Returns empty string with error if the PDF has no extractable text layer.
func ExtractTextFromPDF(fileData []byte) (string, error) {
	reader := bytes.NewReader(fileData)
	pdfReader, err := pdf.NewReader(reader, int64(len(fileData)))
	if err != nil {
		return "", fmt.Errorf("open PDF: %w", err)
	}

	pageCount := pdfReader.NumPage()
	if pageCount == 0 {
		return "", fmt.Errorf("PDF has 0 pages")
	}

	var sb strings.Builder
	for i := 1; i <= pageCount; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}
		if text != "" {
			if sb.Len() > 0 {
				sb.WriteString("\n--- PAGE BREAK ---\n")
			}
			sb.WriteString(text)
		}
	}

	result := strings.TrimSpace(sb.String())
	if result == "" {
		return "", fmt.Errorf("PDF has no extractable text layer (scanned image)")
	}
	return result, nil
}
