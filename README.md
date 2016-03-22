# Little Bastard

Simple [Docker image](https://hub.docker.com/r/equivalent/little_bastard/) with a script in [golang](https://golang.org/) to
execute GET request on an endpoint.

This is usefull if you cannot do cron jobs in your application settup.

Just expose a certain route in your web application to execute the job
(or to schedule backround job) and trigger
`little_bastard` image to execute request on it.


## Usage

```bash
docker pull equivalent/little_bastard
```


Specify `URL` env variable


```bash
docker run -e "URL=http://www.my-app.dot/execute-something.html" equivalent/little_bastard
```

Default timeout is 1000ms (1 second), if you need different timeout:

```bash
# 2 second timeout
docker run -e "SLEEPFOR=2000" -e "URL=http://www.my-app.dot/execute-something.html" equivalent/little_bastard
```

To execute on localhost of host image:

```bash
docker run -e "SLEEPFOR=2000" -e "URL=http://localhost:3000/execute-something.html" --net="host" equivalent/little_bastard
```


You want authentication ? How about passing token param.

```bash
docker run -e "URL=https://www.my-app.dot/execute-something.html?token=1234556" --net="host" equivalent/little_bastard
```

... if you need something more sofisticated, sorry this image is really basic.

## Kill the container

```bash
docker kill $(docker ps | grep -e little_bastard | awk "{print \$1}")

# sudo version
sudo docker kill $(sudo docker ps | grep -e little_bastard | awk "{print \$1}")
```

