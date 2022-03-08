for i in {1..50}; do curl -XPOST -d'{"senderId":"bob123", "receiverId":"alice456", "amount":1}' localhost:8080/v1/transfer; done
for i in {1..50}; do curl -XPOST -d'{"senderId":"alice456", "receiverId":"bob123", "amount":1}' localhost:8080/v1/transfer; done

curl -XPOST -d'{"senderId":"alice456", "receiverId":"bob123", "amount":49}' localhost:8080/v1/transfer

curl -XPOST -d'{"id":"john789", "balance":10, "currency":"AUD"}' localhost:8080/v1/accounts

curl -XPOST -d'{"senderId":"john789", "receiverId":"bob123", "amount":5}' localhost:8080/v1/transfer
