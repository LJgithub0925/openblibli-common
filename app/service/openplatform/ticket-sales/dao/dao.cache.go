// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -nullcache=&model.Promotion{PromoID:-1} -check_null_code=$!=nil&&$.PromoID==-1
		Promo(c context.Context, promoID int64) (*model.Promotion, error)
		// cache: -nullcache=&model.PromotionGroup{GroupID:-1} -check_null_code=$!=nil&&$.GroupID==-1
		PromoGroup(c context.Context, groupID int64) (*model.PromotionGroup, error)
		// cache: -nullcache=&model.PromotionOrder{OrderID:-1} -check_null_code=$!=nil&&$.OrderID==-1
		PromoOrder(c context.Context, orderID int64) (*model.PromotionOrder, error)
		// cache: -nullcache=[]*model.PromotionOrder{{GroupID:-1}} -check_null_code=len($)==1&&$[0].GroupID==-1
		PromoOrders(c context.Context, groupID int64) ([]*model.PromotionOrder, error)

		// cache: -nullcache=[]*model.OrderMain{{OrderID:-1}} -check_null_code=len($)==1&&$[0].OrderID==-1
		Orders(context.Context, *model.OrderMainQuerier) ([]*model.OrderMain, error)
		// cache: -nullcache=-1 -check_null_code=$==-1
		OrderCount(context.Context, *model.OrderMainQuerier) (int64, error)
		// cache: -nullcache=&model.OrderDetail{OrderID:0} -check_null_code=$!=nil&&$.OrderID==0
		OrderDetails(context.Context, []int64) (map[int64]*model.OrderDetail, error)
		// cache: -nullcache=[]*model.OrderSKU{{OrderID:-1}} -check_null_code=len($)==1&&$[0].OrderID==-1
		OrderSKUs(context.Context, []int64) (map[int64][]*model.OrderSKU, error)
		// cache: -nullcache=&model.OrderPayCharge{ChargeID:""} -check_null_code=$!=nil&&$.ChargeID==""
		OrderPayCharges(context.Context, []int64) (map[int64]*model.OrderPayCharge, error)

		// cache:
		SkuByItemID(c context.Context, itemID int64) (map[string]*model.SKUStock, error)
		// cache:
		GetSKUs(c context.Context, skuIds []int64, withNewStock bool) (map[int64]*model.SKUStock, error)
		// cache:
		Stocks(c context.Context, keys []int64, isLocked bool) (res map[int64]int64, err error)

		// cache: -nullcache=[]*model.Ticket{{ID:-1}} -check_null_code=len($)==1&&$[0].ID==-1
		TicketsByOrderID(c context.Context, orderID int64) (res []*model.Ticket, err error)
		// cache: -nullcache=[]*model.Ticket{{ID:-1}} -check_null_code=len($)==1&&$[0].ID==-1
		TicketsByScreen(c context.Context, screenID int64, UID int64) (res []*model.Ticket, err error)
		// cache: -nullcache=&model.Ticket{ID:-1} -check_null_code=$!=nil&&$.ID==-1
		TicketsByID(c context.Context, id []int64) (res map[int64]*model.Ticket, err error)
		// cache: -nullcache=&model.TicketSend{ID:-1} -check_null_code=$!=nil&&$.ID==-1
		TicketSend(c context.Context, SendTID []int64, TIDType string) (res map[int64]*model.TicketSend, err error)
	}
*/

package dao

import (
	"context"

	"go-common/app/service/openplatform/ticket-sales/model"
	"go-common/library/net/metadata"
	"go-common/library/stat/prom"
)

var _ _cache

// Promo get data from cache if miss will call source method, then add to cache.
func (d *Dao) Promo(c context.Context, id int64) (res *model.Promotion, err error) {
	addCache := true
	res, err = d.CachePromo(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.PromoID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("Promo")
		return
	}
	prom.CacheMiss.Incr("Promo")
	res, err = d.RawPromo(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Promotion{PromoID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCachePromo(metadata.WithContext(c), id, miss)
	})
	return
}

// PromoGroup get data from cache if miss will call source method, then add to cache.
func (d *Dao) PromoGroup(c context.Context, id int64) (res *model.PromotionGroup, err error) {
	addCache := true
	res, err = d.CachePromoGroup(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.GroupID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("PromoGroup")
		return
	}
	prom.CacheMiss.Incr("PromoGroup")
	res, err = d.RawPromoGroup(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.PromotionGroup{GroupID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCachePromoGroup(metadata.WithContext(c), id, miss)
	})
	return
}

// PromoOrder get data from cache if miss will call source method, then add to cache.
func (d *Dao) PromoOrder(c context.Context, id int64) (res *model.PromotionOrder, err error) {
	addCache := true
	res, err = d.CachePromoOrder(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.OrderID == -1 {
			res = nil
		}
	}()
	if res != nil {
		prom.CacheHit.Incr("PromoOrder")
		return
	}
	prom.CacheMiss.Incr("PromoOrder")
	res, err = d.RawPromoOrder(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.PromotionOrder{OrderID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCachePromoOrder(metadata.WithContext(c), id, miss)
	})
	return
}

// PromoOrders get data from cache if miss will call source method, then add to cache.
func (d *Dao) PromoOrders(c context.Context, id int64) (res []*model.PromotionOrder, err error) {
	addCache := true
	res, err = d.CachePromoOrders(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if len(res) == 1 && res[0].GroupID == -1 {
			res = nil
		}
	}()
	if len(res) != 0 {
		prom.CacheHit.Incr("PromoOrders")
		return
	}
	prom.CacheMiss.Incr("PromoOrders")
	res, err = d.RawPromoOrders(c, id)
	if err != nil {
		return
	}
	miss := res
	if len(miss) == 0 {
		miss = []*model.PromotionOrder{{GroupID: -1}}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCachePromoOrders(metadata.WithContext(c), id, miss)
	})
	return
}

// Orders get data from cache if miss will call source method, then add to cache.
func (d *Dao) Orders(c context.Context, id *model.OrderMainQuerier) (res []*model.OrderMain, err error) {
	addCache := true
	res, err = d.CacheOrders(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if len(res) == 1 && res[0].OrderID == -1 {
			res = nil
		}
	}()
	if len(res) != 0 {
		prom.CacheHit.Incr("Orders")
		return
	}
	prom.CacheMiss.Incr("Orders")
	res, err = d.RawOrders(c, id)
	if err != nil {
		return
	}
	miss := res
	if len(miss) == 0 {
		miss = []*model.OrderMain{{OrderID: -1}}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheOrders(metadata.WithContext(c), id, miss)
	})
	return
}

// OrderCount get data from cache if miss will call source method, then add to cache.
func (d *Dao) OrderCount(c context.Context, id *model.OrderMainQuerier) (res int64, err error) {
	addCache := true
	res, err = d.CacheOrderCount(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res == -1 {
			res = 0
		}
	}()
	if res != 0 {
		prom.CacheHit.Incr("OrderCount")
		return
	}
	prom.CacheMiss.Incr("OrderCount")
	res, err = d.RawOrderCount(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == 0 {
		miss = -1
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheOrderCount(metadata.WithContext(c), id, miss)
	})
	return
}

// OrderDetails get data from cache if miss will call source method, then add to cache.
func (d *Dao) OrderDetails(c context.Context, keys []int64) (res map[int64]*model.OrderDetail, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheOrderDetails(c, keys)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("OrderDetails", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if v != nil && v.OrderID == 0 {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.OrderDetail
	prom.CacheMiss.Add("OrderDetails", int64(len(miss)))
	missData, err = d.RawOrderDetails(c, miss)
	if res == nil {
		res = make(map[int64]*model.OrderDetail)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if res[key] == nil {
			if missData == nil {
				missData = make(map[int64]*model.OrderDetail, len(keys))
			}
			missData[key] = &model.OrderDetail{OrderID: 0}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheOrderDetails(metadata.WithContext(c), missData)
	})
	return
}

// OrderSKUs get data from cache if miss will call source method, then add to cache.
func (d *Dao) OrderSKUs(c context.Context, keys []int64) (res map[int64][]*model.OrderSKU, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheOrderSKUs(c, keys)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (len(res[key]) == 0) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("OrderSKUs", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if len(v) == 1 && v[0].OrderID == -1 {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64][]*model.OrderSKU
	prom.CacheMiss.Add("OrderSKUs", int64(len(miss)))
	missData, err = d.RawOrderSKUs(c, miss)
	if res == nil {
		res = make(map[int64][]*model.OrderSKU)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if len(res[key]) == 0 {
			if missData == nil {
				missData = make(map[int64][]*model.OrderSKU, len(keys))
			}
			missData[key] = []*model.OrderSKU{{OrderID: -1}}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheOrderSKUs(metadata.WithContext(c), missData)
	})
	return
}

// OrderPayCharges get data from cache if miss will call source method, then add to cache.
func (d *Dao) OrderPayCharges(c context.Context, keys []int64) (res map[int64]*model.OrderPayCharge, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheOrderPayCharges(c, keys)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("OrderPayCharges", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if v != nil && v.ChargeID == "" {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.OrderPayCharge
	prom.CacheMiss.Add("OrderPayCharges", int64(len(miss)))
	missData, err = d.RawOrderPayCharges(c, miss)
	if res == nil {
		res = make(map[int64]*model.OrderPayCharge)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if res[key] == nil {
			if missData == nil {
				missData = make(map[int64]*model.OrderPayCharge, len(keys))
			}
			missData[key] = &model.OrderPayCharge{ChargeID: ""}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheOrderPayCharges(metadata.WithContext(c), missData)
	})
	return
}

// SkuByItemID get data from cache if miss will call source method, then add to cache.
func (d *Dao) SkuByItemID(c context.Context, id int64) (res map[string]*model.SKUStock, err error) {
	addCache := true
	res, err = d.CacheSkuByItemID(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("SkuByItemID")
		return
	}
	prom.CacheMiss.Incr("SkuByItemID")
	res, err = d.RawSkuByItemID(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheSkuByItemID(metadata.WithContext(c), id, miss)
	})
	return
}

// GetSKUs get data from cache if miss will call source method, then add to cache.
func (d *Dao) GetSKUs(c context.Context, keys []int64, withNewStock bool) (res map[int64]*model.SKUStock, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheGetSKUs(c, keys, withNewStock)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("GetSKUs", int64(len(keys)-len(miss)))
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.SKUStock
	prom.CacheMiss.Add("GetSKUs", int64(len(miss)))
	missData, err = d.RawGetSKUs(c, miss, withNewStock)
	if res == nil {
		res = make(map[int64]*model.SKUStock)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheGetSKUs(metadata.WithContext(c), missData, withNewStock)
	})
	return
}

// Stocks get data from cache if miss will call source method, then add to cache.
func (d *Dao) Stocks(c context.Context, keys []int64, isLocked bool) (res map[int64]int64, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheStocks(c, keys, isLocked)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if _, ok := res[key]; !ok {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("Stocks", int64(len(keys)-len(miss)))
	if len(miss) == 0 {
		return
	}
	var missData map[int64]int64
	prom.CacheMiss.Add("Stocks", int64(len(miss)))
	missData, err = d.RawStocks(c, miss, isLocked)
	if res == nil {
		res = make(map[int64]int64)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheStocks(metadata.WithContext(c), missData, isLocked)
	})
	return
}

// TicketsByOrderID get data from cache if miss will call source method, then add to cache.
func (d *Dao) TicketsByOrderID(c context.Context, id int64) (res []*model.Ticket, err error) {
	addCache := true
	res, err = d.CacheTicketsByOrderID(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if len(res) == 1 && res[0].ID == -1 {
			res = nil
		}
	}()
	if len(res) != 0 {
		prom.CacheHit.Incr("TicketsByOrderID")
		return
	}
	prom.CacheMiss.Incr("TicketsByOrderID")
	res, err = d.RawTicketsByOrderID(c, id)
	if err != nil {
		return
	}
	miss := res
	if len(miss) == 0 {
		miss = []*model.Ticket{{ID: -1}}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheTicketsByOrderID(metadata.WithContext(c), id, miss)
	})
	return
}

// TicketsByScreen get data from cache if miss will call source method, then add to cache.
func (d *Dao) TicketsByScreen(c context.Context, id int64, UID int64) (res []*model.Ticket, err error) {
	addCache := true
	res, err = d.CacheTicketsByScreen(c, id, UID)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if len(res) == 1 && res[0].ID == -1 {
			res = nil
		}
	}()
	if len(res) != 0 {
		prom.CacheHit.Incr("TicketsByScreen")
		return
	}
	prom.CacheMiss.Incr("TicketsByScreen")
	res, err = d.RawTicketsByScreen(c, id, UID)
	if err != nil {
		return
	}
	miss := res
	if len(miss) == 0 {
		miss = []*model.Ticket{{ID: -1}}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheTicketsByScreen(metadata.WithContext(c), id, miss, UID)
	})
	return
}

// TicketsByID get data from cache if miss will call source method, then add to cache.
func (d *Dao) TicketsByID(c context.Context, keys []int64) (res map[int64]*model.Ticket, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheTicketsByID(c, keys)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("TicketsByID", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if v != nil && v.ID == -1 {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.Ticket
	prom.CacheMiss.Add("TicketsByID", int64(len(miss)))
	missData, err = d.RawTicketsByID(c, miss)
	if res == nil {
		res = make(map[int64]*model.Ticket)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if res[key] == nil {
			if missData == nil {
				missData = make(map[int64]*model.Ticket, len(keys))
			}
			missData[key] = &model.Ticket{ID: -1}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheTicketsByID(metadata.WithContext(c), missData)
	})
	return
}

// TicketSend get data from cache if miss will call source method, then add to cache.
func (d *Dao) TicketSend(c context.Context, keys []int64, TIDType string) (res map[int64]*model.TicketSend, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	res, err = d.CacheTicketSend(c, keys, TIDType)
	if err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("TicketSend", int64(len(keys)-len(miss)))
	defer func() {
		for k, v := range res {
			if v != nil && v.ID == -1 {
				delete(res, k)
			}
		}
	}()
	if len(miss) == 0 {
		return
	}
	var missData map[int64]*model.TicketSend
	prom.CacheMiss.Add("TicketSend", int64(len(miss)))
	missData, err = d.RawTicketSend(c, miss, TIDType)
	if res == nil {
		res = make(map[int64]*model.TicketSend)
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	for _, key := range keys {
		if res[key] == nil {
			if missData == nil {
				missData = make(map[int64]*model.TicketSend, len(keys))
			}
			missData[key] = &model.TicketSend{ID: -1}
		}
	}
	if !addCache {
		return
	}
	d.cache.Save(func() {
		d.AddCacheTicketSend(metadata.WithContext(c), missData, TIDType)
	})
	return
}
