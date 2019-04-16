# service.build.docker

This service is a builder that builds the requested images and pushes them to the requested repository.

## Optimizations that could be made

* This would probably work way better as a queue worker, pulling jobs from a queue and then emitting a message on completion.