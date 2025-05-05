# To see all available schemas using regular expressions
ms-client schema search 'response|request'

# To view a specific schema
ms-client schema info io.ms-client.jetstream.api.v1.stream_msg_get_request --yaml

# To validate a JSON input against a specific schema
ms-client schema validate io.ms-client.jetstream.api.v1.stream_msg_get_request request.json
