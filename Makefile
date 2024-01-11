.PHONY: all client server


client:
	cd client && npm run dev

server:
	cd server && go run ./cmd
