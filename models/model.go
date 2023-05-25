package models

type Categories struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	ParentId  int    `json:"parentId"`
	Icon      string `json:"icon"`
	UserId    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
type Types struct {
	Type           string `json:"type"`
	ParentCategory ParentCategory
}

type ParentCategory struct {
	ParentID    int `json:"parentid"`
	SubCategory SubCategory
}

type SubCategory struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	UserID     int    `json:"userid"`
	CreatedAt  string `json:"createdat"`
}


// for get categories
// type{
// 		parentid {
// 				  subcategory{
// 							  category_id
// 							  name
// 						   	  icon
// 							  userid
// 							  createdat
// 						     }
// 			 }
// }