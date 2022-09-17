// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	model "github.com/leapord/prometheusx/internal/model/do"
	entity "github.com/leapord/prometheusx/internal/model/entity"
)

type IGroup interface {
	AddGroup(ctx context.Context, group *model.Group) (err error)
	UpdateGroup(ctx context.Context, group *model.Group) (err error)
	DeleteById(ctx context.Context, id int) (group entity.Group, err error)
	Detail(ctx context.Context, id int) (group entity.Group, err error)
	Page(ctx context.Context, group model.Group, pageNo int, pageSize int) (list []entity.Group, total int, err error)
	List(ctx context.Context) (groups []entity.Group, err error)
}

var localGroup IGroup

func Group() IGroup {
	if localGroup == nil {
		panic("implement not found for interface IGroup, forgot register?")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
