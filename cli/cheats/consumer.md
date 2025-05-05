# Adding, Removing, Viewing a Consumer
ms-client consumer add
ms-client consumer info ORDERS NEW
ms-client consumer rm ORDERS NEW

# Editing a consumer
ms-client consumer edit ORDERS NEW --description "new description"

# Get messages from a consumer
ms-client consumer next ORDERS NEW --ack
ms-client consumer next ORDERS NEW --no-ack
ms-client consumer sub ORDERS NEW --ack

# Force leader election on a consumer
ms-client consumer cluster down ORDERS NEW
