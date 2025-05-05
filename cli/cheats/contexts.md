# Create or update
ms-client context add development --server ms-client.dev.example.net:4222 [other standard connection properties]
ms-client context add ngs --description "NGS Connection in Orders Account" --nsc nsc://acme/orders/new
ms-client context edit development [standard connection properties]

# View contexts
ms-client context ls
ms-client context info development --json

# Validate all connections are valid and that connections can be established
ms-client context validate --connect

# Select a new default context
ms-client context select

# Connecting using a context
ms-client pub --context development subject body
