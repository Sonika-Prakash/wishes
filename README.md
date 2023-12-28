# wishes
It is a simple website that allows you to send new year wishes to your friends and family.
You can generate a new year wish with your name in it and share this with anyone you want by sharing the link with them.

This is specifically for new year wishes. But you can be creative and you can use this as a base for any other wishes by tweaking the index.html file.

## How to run and use:
- Clone this repository.
- Run "go run cmd/wishes/main.go".
- A HTTP server will be started at 0.0.0.0:10000 by default. You can change these values if you want in main.go.
- Open the webpage in any browser with the query parameter n=<name>. So the link will be something like this: 0.0.0.0:10000?n=sonika. You will see the name "sonika" in the wishes.
- In the input box below, type your name and click on Generate. A new template with your name will be generated.
- Now click on Copy link to copy this link to your clipboard. Now you can share it with anyone by sending them this link.
