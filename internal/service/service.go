package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/rorast/blog-service/global"
	"github.com/rorast/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
