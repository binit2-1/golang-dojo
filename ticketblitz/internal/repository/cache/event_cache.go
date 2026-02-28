package cache

import (
	"context"
	"encoding/json"
	"ticketblitz/internal/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

type CachedEventRepo struct{
	BaseRepo domain.EventRepository
	Client *redis.Client
}

func NewCachedEventRepo(BaseRepo domain.EventRepository, Client *redis.Client) domain.EventRepository{
	return &CachedEventRepo{
		BaseRepo: BaseRepo,
		Client: Client,
	}
}


func(c *CachedEventRepo) CreateEvent(event *domain.Event) error{
	return c.BaseRepo.CreateEvent(event)
}


func (c *CachedEventRepo) GetEventByID(id string)(*domain.Event, error){
	ctx := context.Background()
	cacheKey := "event:"+id
	
	
	//try getting it from redis
	cachedData, err := c.Client.Get(ctx, cacheKey).Result()
	if err == nil{
		//Got in cache 
		var event domain.Event
		json.Unmarshal([]byte(cachedData), &event)
		return &event, nil
	}

	//not in cache 
	event, err := c.BaseRepo.GetEventByID(id)
	if err != nil{
		return nil, err
	}

	//save result to redis
	eventJSON, _ := json.Marshal(event)
	c.Client.Set(ctx, cacheKey, eventJSON, 30*time.Second)

	return event, nil
}

func(c *CachedEventRepo) PurchaseTicket(userID string, eventID string) error{
	err := c.BaseRepo.PurchaseTicket(userID, eventID)
	if err != nil{
		return err
	}

	ctx := context.Background()
	delKey := "event:" + eventID
	
	return  c.Client.Del(ctx, delKey).Err()
}