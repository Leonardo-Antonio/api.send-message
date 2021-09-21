# SEND MESSAGE 

> Change version - heroku config:set GOVERSION=go1.16

### USE 
#### create .env 
```shell
EMAIL=<mail>
PASSWORD_EMAIL=<password_application>

PORT=8080
```
#### send
- https://send-message-go.herokuapp.com/api/v1/users/message [POST]
```json
{
    "name": "<string>",
    "text": "<string>",
    "email": "<string|email>"
}
```