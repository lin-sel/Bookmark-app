package models

// BookmarkResponse Structure Contain totalpage, totalcount and listofbookmark
type BookmarkResponse struct {
	TotalCount        int64       `json:"totalcount"`
	TotalPage         int64       `json:"totalpage"`
	PageNumber        int64       `json:"-"`
	PageSize          int64       `json:"-"`
	ListOfBookmarkDTO *[]Bookmark `json:"listofbookmark"`
}

// NewResponseBookmark Return New  Instance of BookmarksResponse
func NewResponseBookmark(bookmarks *[]Bookmark, pagenumber int64, pagesize int64) *BookmarkResponse {
	return &BookmarkResponse{
		ListOfBookmarkDTO: bookmarks,
		PageNumber:        pagenumber,
		PageSize:          pagesize,
	}
}

// CategoryResponse Structure Contain totalpage, totalcount and listofcategory
type CategoryResponse struct {
	TotalCount        int64       `json:"totalcount"`
	TotalPage         int64       `json:"totalpage"`
	PageSize          int64       `json:"-"`
	PageNumber        int64       `json:"-"`
	ListOfCategoryDTO *[]Category `json:"listofcategory"`
}

// NewResponseCategory Return New  Instance of CategoryResponse
func NewResponseCategory(categorys *[]Category, pagenumber int64, pagesize int64) *CategoryResponse {
	return &CategoryResponse{
		ListOfCategoryDTO: categorys,
		PageNumber:        pagenumber,
		PageSize:          pagesize,
	}
}

// TokenResponse Return After Successful Login
type TokenResponse struct {
	User
	Token string `json:"token"`
}
