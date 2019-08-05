# Tsubasa, The /r/openttd Infrastructure

**Tsubasa** is all of /r/openttd's custom infrastructure in one handy monorepository.

It is built using the microservices model, which means individual components can be modified / screwed with without bring the rest of the service down.

Mostly written in Go / Docker, though some _Make_ magic happens occasionally. 
Other languages are perfectly acceptable.

Intraservice communication and management happens using the following:
* _JSON REST_ - RPC between services.
* _Kubernetes_ - Service discovery, deployment, etcetera.
* _Kafka_ - Message / job queue.


## Directories

* __generics__ - Contains generic, importable files intended for use by other parts of the project.
* __service.*__ - Services that run completely independently from each other. These services may communicate with other **internal services only**, or with data providers (like external servers, message queues, or databases).
  * __service.*.api__ - Services which provide the external JSON REST API by communicating with their related service.
* __deploy.*__ - Dependent services with their deploy files (like Kafka, Cassandra, etc).