package main

import (
	"context"
	"fmt"

	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func main() {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("连接ectd服务器集群失败")
		return
	}

	defer etcdClient.Close()

	//设置5秒超时，访问etcd有超时控制
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)

	_, err = etcdClient.Put(ctx, "/lock/l1/l2/l3", "1")
	cancle()

	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Printf("ctx 被其他携程取消 : %v", err)
		case context.DeadlineExceeded:
			fmt.Printf("ctx 已超时: %v", err)
		case rpctypes.ErrEmptyKey:
			fmt.Printf("client-side error: %v", err)
		default:
			fmt.Printf("连接集群节点, which are not etcd servers: %v", err)
		}

		return

	}

	session, err := concurrency.NewSession(etcdClient, concurrency.WithTTL(5))

	if err != nil {
		fmt.Println("获取session失败")
		return
	}

	defer session.Close()

	fmt.Println("获取session成功")

	// 谨慎选择锁的pfx的值，确定除了锁之外不允许在此之下创建或修改其他key
	//​ 最开始在lock下建立的一个key，它的创建时间比创建的所有关于锁的key的创建时间要早，导致永远获取不到锁。

	go func() {

		time.Sleep(time.Second * 3)
		locker := concurrency.NewLocker(session, "/newlock/l1")
		fmt.Println("L1 locking...")
		locker.Lock()
		fmt.Println("L1  locked")
		fmt.Println("L1  unlocking...")
		locker.Unlock()
		fmt.Println("L1  unlocked")

	}()

	locker := concurrency.NewLocker(session, "/lock/l1")
	fmt.Println("L2  locking...")
	locker.Lock()
	fmt.Println("L2  locked")
	fmt.Println("L2  unlocking...")
	locker.Unlock()
	fmt.Println("L2  unlocked")

}
