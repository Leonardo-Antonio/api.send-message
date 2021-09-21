# SEND MESSAGE 

> Change version - heroku config:set GOVERSION=go1.16

### USE 

#### send
- https://send-message-go.herokuapp.com/api/v1/users/message [POST]
```json
{
    "name": "<string>",
    "text": "<string>",
    "email": "<string|email>"
}
```