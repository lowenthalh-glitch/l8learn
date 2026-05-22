/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 *
 * Reads uploaded files from FileStore storage.
 * Follows the same pattern as FileStorePut.go — read encrypted file, decrypt via security provider.
 */
package evalimports

import (
	"fmt"
	"github.com/saichler/l8services/go/services/filestore"
	"github.com/saichler/l8types/go/ifs"
	"os"
	"path/filepath"
	"strings"
)

// WriteCleanedText encrypts and saves the cleaned evaluation text,
// replacing the original PDF. Returns the new file path.
func WriteCleanedText(originalPath string, cleanedText string, evalId string, vnic ifs.IVNic) (string, error) {
	dir := filepath.Dir(originalPath)
	newPath := filepath.Join(dir, "cleaned_"+evalId+".txt")

	// Encrypt the cleaned text before writing
	encrypted, err := vnic.Resources().Security().Encrypt([]byte(cleanedText))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt cleaned text: %w", err)
	}

	if err := os.WriteFile(newPath, []byte(encrypted), 0600); err != nil {
		return "", fmt.Errorf("failed to write cleaned text: %w", err)
	}

	return newPath, nil
}

// DeleteOriginalFile removes the original uploaded file from disk.
func DeleteOriginalFile(storagePath string) {
	if storagePath == "" {
		return
	}
	os.Remove(filepath.Clean(storagePath))
}

// ReadUploadedFile reads and decrypts a file from FileStore storage.
func ReadUploadedFile(storagePath string, vnic ifs.IVNic) ([]byte, error) {
	if storagePath == "" {
		return nil, fmt.Errorf("storage path is empty")
	}

	cleanPath := filepath.Clean(storagePath)
	if !strings.HasPrefix(cleanPath, filestore.StorageRoot) {
		return nil, fmt.Errorf("access denied: path outside storage root")
	}

	encryptedData, err := os.ReadFile(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", storagePath)
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	data, err := vnic.Resources().Security().Decrypt(string(encryptedData))
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt file: %w", err)
	}

	return data, nil
}
