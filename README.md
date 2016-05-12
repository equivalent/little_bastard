# Little Bastard

Simple [Docker image](https://hub.docker.com/r/equivalent/little_bastard/) with a script in [golang](https://golang.org/) to
execute GET request on an endpoint (Request Repeater).

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


## AWS Elastic Beanstalk Dockerrun.aws.json example usage

```json
{
  "containerDefinitions": [
    {
      "name": "nginx",
      "image": "........",
      "portMappings": [
        {
          "hostPort": 80,
          "containerPort": 80
        }
      ],
    },
    {
      "name": "request_repeater",
      "image": "equivalent/little_bastard",
      "essential": true,
      "memory": 100,
      "links": [ "nginx" ],
      "environment": [
        {
          "name": "URL",
          "value": "http://nginx/sqs_pull"
        }
      ]
    }
  ]
}
```

Allocating memory to `100`MB is enough. Allocating lower than `90` MB memor will cause container to crush with `oom` (out of memory) in `/var/log/docker-events`

## Kill the container

```bash
docker kill $(docker ps | grep -e little_bastard | awk "{print \$1}")

# sudo version
sudo docker kill $(sudo docker ps | grep -e little_bastard | awk "{print \$1}")
```

