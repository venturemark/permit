package resolver

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/resolver/message"
	"github.com/venturemark/permission/pkg/resolver/timeline"
	"github.com/venturemark/permission/pkg/resolver/update"
	"github.com/venturemark/permission/pkg/resolver/user"
	"github.com/venturemark/permission/pkg/resolver/venture"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Resolver struct {
	message  *message.Resolver
	timeline *timeline.Resolver
	update   *update.Resolver
	user     *user.Resolver
	venture  *venture.Resolver
}

func New(config Config) (*Resolver, error) {
	var err error

	var messageResolver *message.Resolver
	{
		c := message.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		messageResolver, err = message.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var timelineResolver *timeline.Resolver
	{
		c := timeline.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		timelineResolver, err = timeline.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var updateResolver *update.Resolver
	{
		c := update.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		updateResolver, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var userResolver *user.Resolver
	{
		c := user.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		userResolver, err = user.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var ventureResolver *venture.Resolver
	{
		c := venture.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		ventureResolver, err = venture.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	r := &Resolver{
		message:  messageResolver,
		timeline: timelineResolver,
		update:   updateResolver,
		user:     userResolver,
		venture:  ventureResolver,
	}

	return r, nil
}

func (r *Resolver) Message() permission.Resolver {
	return r.message
}

func (r *Resolver) Timeline() permission.Resolver {
	return r.timeline
}

func (r *Resolver) Update() permission.Resolver {
	return r.update
}

func (r *Resolver) User() permission.Resolver {
	return r.user
}

func (r *Resolver) Venture() permission.Resolver {
	return r.venture
}
