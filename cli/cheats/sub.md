# To subscribe to messages, in a queue group and acknowledge any JetStream ones
ms-client sub source.subject --queue work --ack

# To subscribe to a randomly generated inbox
ms-client sub --inbox

# To dump all messages to files, 1 file per message
ms-client sub --inbox --dump /tmp/archive

# To process all messages using xargs 1 message at a time through a shell command
ms-client sub subject --dump=- | xargs -0 -n 1 -I "{}" sh -c "echo '{}' | wc -c"

# To receive new messages received in a stream with the subject ORDERS.new
ms-client sub ORDERS.new --next

# To report the number of subjects with message and byte count. The default `--report-top` is 10
ms-client sub ">" --report-subjects --report-top=20

# To base64 decode message bodies before rendering them
ms-client sub 'encoded.sub' --translate "base64 -d"
