package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func addValue(cli *clientv3.Client, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		_, err := cli.Put(ctx, "/logagent/conf/", string(i)) //设置
		cancel()

		if err != nil {
			switch err {
			case context.Canceled:
				log.Fatalf("ctx 被其他携程取消 : %v", err)
			case context.DeadlineExceeded:
				log.Fatalf("ctx 已超时: %v", err)
			case rpctypes.ErrEmptyKey:
				log.Fatalf("client-side error: %v", err)
			default:
				log.Fatalf("连接集群节点, which are not etcd servers: %v", err)
			}
		}
	}

	wg.Done()

}

func watchEtcd(cli *clientv3.Client, wg *sync.WaitGroup) {

	fmt.Println("watch 节点 /logagent/conf/")
	rch := cli.Watch(context.Background(), "/logagent/conf/")

	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s,%q :%q \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	wg.Done()

}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("连接ectd服务器集群失败")
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go addValue(cli, &wg)
	go watchEtcd(cli, &wg)

	wg.Wait()
	// defer cli.Close()

}
