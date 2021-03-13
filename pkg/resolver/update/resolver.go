package update

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Resolver struct {
	logger logger.Interface
	redigo redigo.Interface
}

func New(config Config) (*Resolver, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	r := &Resolver{
		logger: config.Logger,
		redigo: config.Redigo,
	}

	return r, nil
}

func (r *Resolver) Role(met map[string]string) (string, error) {
	var rok *key.Key
	{
		rok = key.Role(update(met))
	}

	var usi string
	{
		usi = met[metadata.UserID]
	}

	var rol *schema.Role
	{
		k := rok.List()
		s := usi

		str, err := r.redigo.Sorted().Search().Index(k, s)
		if err != nil {
			return "", tracer.Mask(err)
		}

		if str == "" {
			return "", nil
		}

		rol = &schema.Role{}
		err = json.Unmarshal([]byte(str), rol)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return rol.Obj.Metadata[metadata.RoleKind], nil
}

func (r *Resolver) Visibility(met map[string]string) (string, error) {
	var mek *key.Key
	{
		mek = key.Update(met)
	}

	var upd *schema.Update
	{
		k := mek.List()
		s := mek.ID().F()

		str, err := r.redigo.Sorted().Search().Score(k, s, s)
		if err != nil {
			return "", tracer.Mask(err)
		}

		if len(str) == 0 {
			return "", nil
		}

		upd = &schema.Update{}
		err = json.Unmarshal([]byte(str[0]), upd)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return upd.Obj.Metadata[metadata.ResourceVisibility], nil
}

func update(met map[string]string) map[string]string {
	cop := map[string]string{}

	for k, v := range met {
		cop[k] = v
	}

	cop[metadata.ResourceKind] = "update"

	return cop
}
