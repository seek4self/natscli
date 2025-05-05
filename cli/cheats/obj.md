# to create a replicated bucket
ms-client obj add FILES --replicas 3

# store a file in the bucket
ms-client obj put FILES image.jpg

# store contents of STDIN in the bucket
cat x.jpg|ms-client obj put FILES --name image.jpg

# retrieve a file from a bucket
ms-client obj get FILES image.jpg -O out.jpg

# delete a file
ms-client obj del FILES image.jpg

# delete a bucket
ms-client obj del FILES

# view bucket info
ms-client obj info FILES

# view file info
ms-client obj info FILES image.jpg

# list known buckets
ms-client obj ls

# view all files in a bucket
ms-client obj ls FILES

# prevent further modifications to the bucket
ms-client obj seal FILES

# create a bucket backup for FILES into backups/FILES
ms-client obj status FILES
ms-client stream backup <stream name> backups/FILES

# restore a bucket from a backup
ms-client stream restore <stream name> backups/FILES
