# To see all servers, including their server ID and show a response graph
ms-client server ping --id --graph --user system

# To see information about a specific server
ms-client server info nats1.example.net --user system
ms-client server info NCAXNST2VH7QGBVYBEDQGX73GMBXTWXACUTMQPTNKWLOYG2ES67NMX6M --user system

# To list all servers and show basic summaries, expecting responses from 10 servers
ms-client server list 10 --user system

# To report on current connections
ms-client server report connections
ms-client server report connz --account WEATHER
ms-client server report connz --sort in-msgs
ms-client server report connz --top 10 --sort in-msgs

# To limit connections report to surveyor connections and all from a specific IP using https://expr.medv.io/docs/Language-Definition
ms-client server report connz --filter 'lower(conns.name) matches "surveyor" || conns.ip == "46.101.44.80"'

# To report on accounts
ms-client server report accounts
ms-client server report accounts --account WEATHER --sort in-msgs --top 10

# To report on JetStream usage by account WEATHER
ms-client server report jetstream --account WEATHER --sort cluster

# To generate a NATS Server bcrypt command
ms-client server password
ms-client server pass -p 'W#OZwVN-UjMb8nszwvT2LQ'
ms-client server pass -g
PASSWORD='W#OZwVN-UjMb8nszwvT2LQ' ms-client server pass

# To request raw monitoring data from servers
ms-client server request subscriptions --detail --filter-account WEATHER --cluster EAST
ms-client server req variables --name nats1.example.net
ms-client server req connections --filter-state open
ms-client server req connz --subscriptions --name nats1.example.net
ms-client server req gateways --filter-name EAST
ms-client server req leafnodes --subscriptions
ms-client server req accounts --account WEATHER
ms-client server req jsz --leader

# To manage JetStream cluster RAFT membership
ms-client server raft step-down
