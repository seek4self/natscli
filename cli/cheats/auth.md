# Basic steps for setting up decentralized authentication
All configuration changes are stored locally until `ms-client auth account push`ed to a ms-client cluster.
    
# Create a new operator and set as working context
ms-client auth operator add sysopp

# Generate a template server configuration file from an operator
ms-client server generate server.conf

# Create a new account
ms-client auth account add MyAccount

# Create a new user in an account
ms-client auth user add MyUser

# Create an admin user in system account
ms-client auth user add admin SYSTEM

# Export credentials for a user
ms-client auth user credential sys_admin.cred admin SYSTEM

# Push an account or its changes from a specific operator to a specific server, using system account credentials. 
ms-client auth account push MyAccount --server ms://localhost:4222 --operator sysopp --creds sys_admin.cred  

# Use `ms-client context` and `ms-client auth operator select` to set defaults
ms-client context add sysadmin --description "System Account" --server ms://localhost:4222 --creds sys_admin.cred

ms-client auth operator select sysopp

# Push account with default settings
ms-client auth account push MyAccount


