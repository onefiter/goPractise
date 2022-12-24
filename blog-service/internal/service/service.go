package service

import (
	"context"
	"log"

	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	log.Println("svc: ", svc)
	svc.dao = dao.New(global.DBEngine)
	return svc
}
