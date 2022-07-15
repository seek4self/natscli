# STHG MS Client

- [功能](#功能)
  - [高效发布订阅](#高效发布订阅)
  - [高效请求回复](#高效请求回复)
  - [持久化](#持久化)
  - [性能测试](#性能测试)

消息服务客户端测试工具

## 功能

- 支持高效发布/订阅
- 支持可靠发布/订阅
- 支持主题通配符 `*`、`>`
- 支持请求/回复
- 支持持久化流（stream）的增删改查
- 支持消费者（consumer）的增删改查
- 支持账户资源查询
- 支持性能测试

```bash
$ ./ms-client      
usage: ms-client [<flags>] <command> [<args> ...]

STHG-MS Utility

STHG-MS and JetStream administration.

See 'ms-client cheat' for a quick cheatsheet of commands

Flags:
  -h, --help                    Show context-sensitive help (also try --help-long and --help-man).
      --version                 Show application version.
  -s, --server=STHG_URL         STHG-MS server urls
      --user=STHG_USER          Username or Token
      --password=STHG_PASSWORD  Password
      --connection-name=NAME    Nickname to use for the underlying STHG-MS Connection
      --creds=FILE              User credentials
      --nkey=FILE               User NKEY
      --tlscert=FILE            TLS public certificate
      --tlskey=FILE             TLS private key
      --tlsca=FILE              TLS certificate authority chain
      --timeout=DURATION        Time to wait on responses from STHG-MS
      --js-api-prefix=PREFIX    Subject prefix for access to JetStream API
      --js-event-prefix=PREFIX  Subject prefix for access to JetStream Advisories
      --js-domain=DOMAIN        JetStream domain to access
      --inbox-prefix=PREFIX     Custom inbox prefix to use for inboxes
      --trace                   Trace API interactions
```

### 高效发布订阅

```bash
# 订阅 test.* 主题消息
$ ms-client sub test.* -s 192.168.1.123

# 每 1s 向主题 test.123 主题发布一条消息
$ ms-client pub test.123 -s 192.168.1.123 "hello world {{Count}}" --count 100 --sleep 1s

# 更多信息查看帮助
$ ms-client pub -h
$ ms-client sub -h
```

> `-s` 指定消息服务器地址，若消息服务端口号不为 4222，需要指定端口号

### 高效请求回复

```bash
# 订阅 test.* 主题消息，并回复消息 ‘hi’
$ ms-client reply 'test.*' -s 192.168.1.123 "hi"

# 向主题 test.123 主题发布一条消息并等待回复
$ ms-client request test.123 -s 192.168.1.123 "hello world"

# 更多信息查看帮助
$ ms-client reply -h
$ ms-client request -h
```

### 持久化

- 添加持久化流

> 创建流时，参数不明确时可以键入 ? 查询参数描述

```bash
# 添加监听 test.* 主题的文件流 test_str
$ ./ms-client str add test_str -s 192.168.1.123     # 流的名称，不能包含空格、制表符、句点 (`.`)、大于 (`>`) 或星号 (`*`)               
? Subjects test.*                                   # 消费主题列表
? Storage file                                      # 存储方式：file / memory
? Replication 1                                     # 在集群 JetStream 中为每条消息保留多少个副本，最多 5 个
? Retention Policy Limits                           # 消息保留策略，Limits:超出限制删除，Interest：所有消费者都消费过才删除，WorkQueue：被消费一次删除
? Discard Policy Old                                # 丢弃消息策略，Old: 超出限制删除旧消息，New：超出限制丢弃新消息
? Stream Messages Limit -1                          # 最大存储消息数量，-1 为没有限制
? Per Subject Messages Limit -1                     # 每个主题存储消息数量
? Total Stream Size -1                              # 流最大存储空间
? Message TTL -1                                    # 流生命周期 单位： (s)econds, (m)inutes, (h)ours, (y)ears, (M)onths, (d)ays.
? Max Message Size -1                               # 单个消息最大字节
? Duplicate tracking time window 2m0s               # 跟踪重复消息的窗口，以纳秒表示
? Allow message Roll-ups No                         # 
? Allow message deletion Yes
? Allow purging subjects or the entire stream Yes
Stream test_str was created

Information for Stream test_str created 2022-07-15T16:57:53+08:00

Configuration:

             Subjects: test.*
     Acknowledgements: true
            Retention: File - Limits
             Replicas: 1
       Discard Policy: Old
     Duplicate Window: 2m0s
    Allows Msg Delete: true
         Allows Purge: true
       Allows Rollups: false
     Maximum Messages: unlimited
        Maximum Bytes: unlimited
          Maximum Age: unlimited
 Maximum Message Size: unlimited
    Maximum Consumers: unlimited


State:

             Messages: 0
                Bytes: 0 B
             FirstSeq: 0
              LastSeq: 0
     Active Consumers: 0

```

- 添加消费者

```bash
# 向文件流 test_str 中添加一个 pull 消费者
$ ./ms-client con add test_str -s 192.168.1.123
? Consumer name test_con                                        # 消费者名称
? Delivery target (empty for Pull Consumers)                    # 消费者类型，为空时为 pull 消费者，填入主题时为 push 消费者，将消息转发到指定主题
? Start policy (all, new, last, subject, 1h, msg sequence) all  # 从什么位置/时间开始消费的策略
? Acknowledgement policy explicit                               # 响应策略
? Replay policy instant                                         # 
? Filter Stream by subject (blank for all) 
? Maximum Allowed Deliveries -1
? Maximum Acknowledgements Pending 0
? Deliver headers only without bodies No
? Add a Retry Backoff Policy No
Information for Consumer test_str > test_con created 2022-07-15T17:33:05+08:00

Configuration:

        Durable Name: test_con
           Pull Mode: true
      Deliver Policy: All
          Ack Policy: Explicit
            Ack Wait: 30s
       Replay Policy: Instant
     Max Ack Pending: 20,000
   Max Waiting Pulls: 512

State:

   Last Delivered Message: Consumer sequence: 0 Stream sequence: 0
     Acknowledgment floor: Consumer sequence: 0 Stream sequence: 0
         Outstanding Acks: 0 out of maximum 20,000
     Redelivered Messages: 0
     Unprocessed Messages: 0
            Waiting Pulls: 0 of maximum 512

```

- 发布消息存储到流

```bash
# 将 100 条消息存储在流 test_str 中
$ ms-client pub test.111 -s 192.168.1.123 "hello world {{Count}}" --count 100
```

- 查看流信息

```bash
$ ./ms-client str info test_str
Information for Stream test_str created 2022-07-15T16:57:53+08:00

Configuration:

             Subjects: test.*
     Acknowledgements: true
            Retention: File - Limits
             Replicas: 1
       Discard Policy: Old
     Duplicate Window: 2m0s
    Allows Msg Delete: true
         Allows Purge: true
       Allows Rollups: false
     Maximum Messages: unlimited
        Maximum Bytes: unlimited
          Maximum Age: unlimited
 Maximum Message Size: unlimited
    Maximum Consumers: unlimited


State:

             Messages: 100
                Bytes: 5.1 KiB
             FirstSeq: 1 @ 2022-07-15T09:44:45 UTC
              LastSeq: 100 @ 2022-07-15T09:44:45 UTC
     Active Consumers: 1

```

- 消费流中的一条消息，

```bash
$ ./ms-client con sub test_str test_con
[17:46:45] subj: test.111 / tries: 1 / cons seq: 1 / str seq: 1 / pending: 99

hello world 1

Acknowledged message
```

或者使用 `next`， 更多用法使用 `ms-client con next -h` 查看帮助

```bash
$ ./ms-client con next test_str test_con
[17:47:51] subj: test.111 / tries: 1 / cons seq: 2 / str seq: 2 / pending: 98

hello world 2

Acknowledged message
```

### 性能测试

一些示例：

```bash
$ ./ms-client cheat bench
# benchmark core ms-client publish and subscribe with 10 publishers and subscribers
# 使用 10 个发布者和订阅者 测试发布订阅
ms-client bench testsubject --pub 10 --sub 10 --msgs 10000 --size 512

# benchmark core ms-client request-reply with queuing
# 测试和请求回复性能
ms-client bench testsubject --sub 4 --reply
ms-client bench testsubject --pub 4 --request --msgs 20000

# benchmark JetStream synchronously acknowledged publishing purging the data first
# 测试发布消息到持久化流中
ms-client bench testsubject --js --syncpub --pub 10  --msgs 10000 --purge

# benchmark JS publish and push consumers at the same time purging the data first
ms-client bench testsubject --js --pub 4 --sub 4 --purge

# benchmark JS stream purge and async batched publishing to the stream
ms-client bench testsubject --js --pub 4 --purge

# benchmark JS stream get replay from the stream using a push consumer
ms-client bench testsubject --js --sub 4

# benchmark JS stream get replay from the stream using a pull consumer
ms-client bench testsubject --js --sub 4 --pull

# simulate a message processing time (for reply mode and pull JS consumers) of 50 microseconds
ms-client bench testsubject --reply --sub 1 --acksleep 50us

# generate load by publishing messages at an interval of 100 nanoseconds rather than back to back
ms-client bench testsubject --pub 1 --pubsleep 100ns

# remember when benchmarking JetStream
# Once you are finished benchmarking, remember to free up the resources (i.e. memory and files) consumed by the stream using 'ms-client stream rm'

```

> **注意**：完成基准测试后，请记住使用 `ms-client stream rm` 释放流消耗的资源（即内存和文件）
