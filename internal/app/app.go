package app

import (
	"context"

	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/list"
	"github.com/ilyachelmodeev/otus_anti_bruteforce/internal/rateLimiter"
)

type app struct {
	rateLimit rateLimiter.RateLimit
	blackList list.List
	whiteList list.List
}

type App interface {
	CheckAuth(login string, password string, ip string) bool
	ResetAuth(login string, ip string)
	AddIPBlackList(ctx context.Context, ip string) error
	DeleteIPBlackList(ctx context.Context, ip string) error
	AddIPWhiteList(ctx context.Context, ip string) error
	DeleteIPWhiteList(ctx context.Context, ip string) error
}

func New(rt rateLimiter.RateLimit, whiteList list.List, blackList list.List) App {
	return &app{
		rateLimit: rt,
		blackList: blackList,
		whiteList: whiteList,
	}
}

func (a *app) CheckAuth(login string, password string, ip string) bool {
	if a.checkBlackList(ip) {
		return false
	}
	if a.checkWhiteList(ip) {
		return true
	}

	return a.rateLimit.Check(login, password, ip)
}

func (a *app) ResetAuth(login string, ip string) {
	a.rateLimit.Reset(login, ip)
}

func (a *app) AddIPBlackList(ctx context.Context, ip string) error {
	return a.blackList.Add(ctx, ip)
}

func (a *app) DeleteIPBlackList(ctx context.Context, ip string) error {
	return a.blackList.Delete(ctx, ip)
}

func (a *app) AddIPWhiteList(ctx context.Context, ip string) error {
	return a.whiteList.Add(ctx, ip)
}

func (a *app) DeleteIPWhiteList(ctx context.Context, ip string) error {
	return a.whiteList.Delete(ctx, ip)
}

func (a *app) checkWhiteList(ip string) bool {
	return a.whiteList.Check(ip)
}

func (a *app) checkBlackList(ip string) bool {
	return a.blackList.Check(ip)
}
