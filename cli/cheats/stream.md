# Adding, Removing, Viewing a Stream
ms-client stream add
ms-client stream info STREAMNAME
ms-client stream rm STREAMNAME

# Editing a single property of a stream
ms-client stream edit STREAMNAME --description "new description"
# Editing a stream configuration in your editor
EDITOR=vi ms-client stream edit -i STREAMNAME

# Show a list of streams, including basic info or compatible with pipes
ms-client stream list
ms-client stream list -n

# Find all empty streams or streams with messages
ms-client stream find --empty
ms-client stream find --empty --invert

# Creates a new Stream based on the config of another, does not copy data
ms-client stream copy ORDERS ARCHIVE --description "Orders Archive" --subjects ARCHIVE

# Get message 12344, delete a message, delete all messages
ms-client stream get ORDERS 12345
ms-client stream rmm ORDERS 12345

# Purge messages from streams
ms-client stream purge ORDERS
# deletes up to, but not including, 1000
ms-client stream purge ORDERS --seq 1000
ms-client stream purge ORDERS --keep 100
ms-client stream purge ORDERS --subject one.subject

# Page through a stream
ms-client stream view ORDERS
ms-client stream view --id 1000
ms-client stream view --since 1h
ms-client stream view --subject one.subject

# Backup and restore
ms-client stream backup ORDERS backups/orders/$(date +%Y-%m-%d)
ms-client stream restore ORDERS backups/orders/$(date +%Y-%m-%d)

# Marks a stream as read only
ms-client stream seal ORDERS

# Force a cluster leader election
ms-client stream cluster ORDERS down

# Evict the stream from a node
stream cluster peer-remove ORDERS nats1.example.net
