package resource

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/resource/message"
	"github.com/venturemark/permission/pkg/resource/timeline"
	"github.com/venturemark/permission/pkg/resource/update"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Resource struct {
	message  *message.Resource
	timeline *timeline.Resource
	update   *update.Resource
}

func New(config Config) (*Resource, error) {
	var err error

	var messageResource *message.Resource
	{
		c := message.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		messageResource, err = message.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var timelineResource *timeline.Resource
	{
		c := timeline.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		timelineResource, err = timeline.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var updateResource *update.Resource
	{
		c := update.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		updateResource, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	r := &Resource{
		message:  messageResource,
		timeline: timelineResource,
		update:   updateResource,
	}

	return r, nil
}

func (g *Resource) Message() permission.Resolver {
	return g.message
}

func (g *Resource) Timeline() permission.Resolver {
	return g.timeline
}

func (g *Resource) Update() permission.Resolver {
	return g.update
}
