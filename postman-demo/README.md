**Copyright (c) 2018 BitMart Exchange. All Rights Reserved.**

## BitMart Exchange Postman Sample API

### How to use it

Please import ```BitMart API.postman_collection.json``` as project and ```Bitmart API.postman_environment.json``` as environment.

The requests are ordered as you can see the first requests are public requests then authenticated requests. If you are familiar with Newman, you can run this project in command line. All necessary post-validation and prescript are added in each request.

Since postman doesn't support RSA at this moment, you will have to sign the message outside and replace ```client_secret``` variable then call [OAuth Token](rest/authenticated/oauth.md) to get bearer token OR integrated with NewMan to make it automated.


