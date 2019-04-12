# Renai

## A /r/openttd build chain tool

<sup><sup><sup><sup>_se no-_</sup></sup></sup></sup>

**Renai** is a build tool which does the following:

* Listens for a HTTP REST command, with a target build version
* Builds that target version using the local Docker host (and the Dockerfile contained in our docker_openttd repository)
* Pushes it to the Docker Hub

This builder is designed to be nudged by a HTTP client.