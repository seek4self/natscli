# To set up a responder that runs an external command with the 3rd subject token as argument
ms-client reply "service.requests.>" --command "service.sh {{2}}"

# To set up basic responder
ms-client reply service.requests "Message {{Count}} @ {{Time}}"
ms-client reply service.requests --echo --sleep 10
