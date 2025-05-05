# to create a replicated KV bucket
ms-client kv add CONFIG --replicas 3

# to store a value in the bucket
ms-client kv put CONFIG username bob

# to read just the value with no additional details
ms-client kv get CONFIG username --raw

# view an audit trail for a key if history is kept
ms-client kv history CONFIG username

# to see the bucket status
ms-client kv status CONFIG

# observe real time changes for an entire bucket
ms-client kv watch CONFIG
# observe real time changes for all keys below users
ms-client kv watch CONFIG 'users.>''

# create a bucket backup for CONFIG into backups/CONFIG
ms-client kv status CONFIG
ms-client stream backup <stream name> backups/CONFIG

# restore a bucket from a backup
ms-client stream restore <stream name> backups/CONFIG

# list known buckets
ms-client kv ls
