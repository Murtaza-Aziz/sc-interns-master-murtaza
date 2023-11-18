package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
    page := 1
    pageSize := 5
    orgID := uuid.FromStringOrNil(folders.DefaultOrgID)

    for {
        res, nextPageToken, err := folders.FetchPaginatedFoldersByOrgID(orgID, page, pageSize)
        if err != nil {
            fmt.Printf("Error fetching paginated folders: %v", err)
            return
        }

        fmt.Printf("Page %d:\n", page)
        folders.PrettyPrint(res)
        fmt.Println()

        fmt.Println("Enter 'n' for next page, 'p' for previous page, or 'q' to quit:")
        var userInput string
        fmt.Scanln(&userInput)

        switch userInput  {
        case "n":
            if nextPageToken == "" {
                fmt.Println("No more pages available.")
            } else {
                page++
            }
        case "p":
            if page > 1 {
                page--
            } else {
                fmt.Println("Already at the first page.")
            }
        case "q":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid input. Please try again.")
            continue
        }
    }
}

