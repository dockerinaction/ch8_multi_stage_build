# Overview #

This repository contains the multi-stage build example for Docker In Action 2ed's Chapter 8.

Execute the image build with:

```sh
docker build -t dockerinaction/http-client -f http-client.df .
```

Run the resulting image with:

```sh
docker container run --rm -it dockerinaction/http-client:latest
```

The http-client program will search google for 'docker' and print the initial output for the page.
