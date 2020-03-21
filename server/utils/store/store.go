package store

import (
	"encoding/json"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

type KVStore interface {
	Load(key string) ([]byte, error)
	Store(key string, data []byte) error
	Delete(key string) error
}

type pluginStore struct {
	api plugin.API
}

var ErrNotFound = errors.New("no data")

func NewPluginStore(api plugin.API) KVStore {
	return &pluginStore{
		api: api,
	}
}

func (s *pluginStore) Load(key string) ([]byte, error) {
	data, appErr := s.api.KVGet(key)
	if appErr != nil {
		return nil, errors.WithMessage(appErr, "failed plugin KVGet")
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (s *pluginStore) Store(key string, data []byte) error {
	appErr := s.api.KVSet(key, data)
	if appErr != nil {
		return errors.WithMessagef(appErr, "failed plugin KVSet %q", key)
	}
	return nil
}

func (s *pluginStore) Delete(key string) error {
	appErr := s.api.KVDelete(key)
	if appErr != nil {
		return errors.WithMessagef(appErr, "failed plugin KVDelete %q", key)
	}
	return nil
}

func LoadJSON(s KVStore, key string, v interface{}) (returnErr error) {
	data, err := s.Load(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func StoreJSON(s KVStore, key string, v interface{}) (returnErr error) {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return s.Store(key, data)
}
