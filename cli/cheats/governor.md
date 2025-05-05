# to create governor with 10 slots and 1 minute timeout
ms-client governor add cron 10 1m

# to view the configuration and state
ms-client governor view cron

# to reset the governor, clearing all slots
ms-client governor reset cron

# to run long-job.sh when a slot is available, giving up after 20 minutes without a slot
ms-client governor run cron $(hostname -f) --max-wait 20m long-job.sh'
