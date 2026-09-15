# Kyu

A concurrent background-job queue implemented in Go.

## Current capabilities
- buffered job ingestion
- configurable worker pool
- retry attempts with delay
- maximum-attempt policy
- dead-letter channel
- context-driven shutdown
- thread-safe worker lifecycle

```bash
go run .
```

`queue/queue.go` is reusable as a package. The root executable demonstrates a four-worker queue. Implemented independently by Adewale Babalola as a backend concurrency exercise.
