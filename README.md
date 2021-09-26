# go-url-shortener

Simple url shortener on GO.

Here is the link where u can see postman request examples:
https://documenter.getpostman.com/view/8653169/UUxxhUAC

expired_on parameter in POST request is optional.

When u send GET request with uuid there are two possible responses:
1) message about expired_on date of the link u requested is expired
2) redirect to original url
