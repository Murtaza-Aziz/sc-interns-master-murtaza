package folders

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofrs/uuid"
)

func GetAllFoldersPaginated(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	orgID := req.OrgID
	allFolders := GetSampleData()

	var folders []*Folder
	for _, folder := range allFolders {
		if folder.OrgId == orgID && !folder.Deleted {
			folders = append(folders, folder)
		}
	}

	fetchResponse := &FetchFolderResponse{Folders: folders}
	return fetchResponse, nil
}

func FetchPaginatedFoldersByOrgID(orgID uuid.UUID, page int, pageSize int) (*FetchFolderResponse, string, error) {
	allFolders := GetSampleData()

	var folders []*Folder
	for _, folder := range allFolders {
		if folder.OrgId == orgID && !folder.Deleted {
			folders = append(folders, folder)
		}
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(folders) {
		return &FetchFolderResponse{}, "", nil
	}

	if end > len(folders) {
		end = len(folders)
	}

	paginatedFolders := folders[start:end]

	fetchResponse := &FetchFolderResponse{Folders: paginatedFolders}

	var nextPageToken string
	if end < len(folders) {
		nextPageToken = generateNextPageToken(page + 1 + (page-1)/pageSize, orgID)
	}

	return fetchResponse, nextPageToken, nil
}

func generateNextPageToken(page int, orgID uuid.UUID) string {
	tokenString := fmt.Sprintf("%d%s", page, orgID.String())

	hash := sha256.New()
	hash.Write([]byte(tokenString))
	hashed := hash.Sum(nil)

	token := hex.EncodeToString(hashed)

	return token
}
