# Tsubasa, The /r/openttd Infrastructure

**Tsubasa** is all of /r/openttd's custom infrastructure in one handy monorepository.

It is built using the microservices model, which means individual components can be modified / screwed with without bring the rest of the service down.

Mostly written in Go / Docker, though some _Make_ magic happens occasionally. 
Other languages are perfectly acceptable.

Intraservice communication and management happens using the following:
* _Protobuf / GRPC_ - RPC between services.
* _Kubernetes_ - Service discovery, deployment, etcetera.