/*
 * zhihu_caffe
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
// package swagger

package model

type User struct {
 
	Id int64 `json:"id,omitempty"`
 
	Password string `json:"password,omitempty"`
 
	//Articles []Article `json:"articles,omitempty"`
 
	Email string `json:"email,omitempty"`
}