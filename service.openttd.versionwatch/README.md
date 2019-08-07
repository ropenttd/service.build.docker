# service.openttd.versionwatch

This service watches for new versions of OpenTTD by polling the Version Server on a regular basis.

If new versions are found (compared to what is stored in the data cache), a queue event is emitted, and the new version is cached for idempotency.

These queue events are intended to be caught by the following:

* service.build.scheduler, to schedule and build the new version into images
* service.announce, to announce the new release in the appropriate channels