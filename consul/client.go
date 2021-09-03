package consul

import (
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/hashicorp/consul/api"
	"strconv"
	"strings"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type KVClient struct {
	consul *api.KV
	folder string
}

func NewKVClient(host, datacenter, folder string) (*KVClient, error) {

	consulClient, err := api.NewClient(&api.Config{
		Address:    host,
		Datacenter: datacenter,
	})
	if err != nil {
		return nil, err
	}

	// remove trailing slash
	folder = strings.TrimRight(folder, "/")

	return &KVClient{consul: consulClient.KV(), folder: folder}, nil
}

func (k *KVClient) GetInt(key string) (int, error) {

	kv, _, err := k.Get(key, nil)
	if err != nil {
		return 0, err
	}

	parsedValue, err := strconv.Atoi(string(kv.Value))
	if err != nil {
		return 0, err
	}

	return parsedValue, nil
}

func (k *KVClient) GetString(key string) (string, error) {

	kv, _, err := k.Get(key, nil)
	if err != nil {
		return "", err
	}

	return string(kv.Value), nil
}

func (k *KVClient) PutInt(key string, value int) error {

	_, err := k.Put(key, []byte(strconv.Itoa(value)), nil)
	if err != nil {
		return err
	}

	return nil
}

func (k *KVClient) PutString(key, value string) error {

	_, err := k.Put(key, []byte(value), nil)
	if err != nil {
		return err
	}
	return nil
}

func (k *KVClient) Get(key string, queryOptions *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {

	kv, meta, err := k.consul.Get(k.addFolderToKey(key), queryOptions)
	if err != nil {
		incGetErrorCounter()
		return nil, nil, err
	}
	if kv == nil {
		return nil, meta, ErrKeyNotFound
	}

	incGetCounter()
	return kv, meta, err
}

func (k *KVClient) Put(key string, value []byte, writeOptions *api.WriteOptions) (*api.WriteMeta, error) {

	kvPair := &api.KVPair{
		Key:   k.addFolderToKey(key),
		Value: value,
	}

	meta, err := k.consul.Put(kvPair, writeOptions)
	if err != nil {
		incPutErrorCounter()
		return nil, err
	}

	incPutCounter()
	return meta, nil
}

func (k *KVClient) addFolderToKey(key string) string {

	if strings.HasPrefix(key, k.folder+"/") { // the folder is already included in the key
		return key
	}

	return fmt.Sprintf("%s/%s", k.folder, key)
}
