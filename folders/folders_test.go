package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {

	const testOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
	t.Run("test for a prestored organization", func(t *testing.T) {
		mockRequest := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(testOrgID),
		}

		response, err := folders.GetAllFoldersPaginated(mockRequest)

		assert.Nil(t, err, "Expected no error, got an error")
		assert.NotNil(t, response, "Expected a non-nil response")
		assert.NotNil(t, response.Folders, "Expected non-nil folders in the response")
	})


}


