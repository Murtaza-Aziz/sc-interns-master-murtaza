package folders

import (
	"github.com/gofrs/uuid"
)
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	
	var (
		err error
		fs  []*Folder
		f  []Folder
	)


	for _, f1 := range f {
		tempF1 := f1 
		fs = append(fs, &tempF1)
	}
	
	var ffr *FetchFolderResponse = &FetchFolderResponse{Folders: fs}
	return ffr, err
}
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	
	return resFolder, nil
}
