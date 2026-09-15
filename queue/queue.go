package queue

import("context";"errors";"sync";"time")
type Job struct{ID string; Payload []byte; Attempts int; MaxAttempts int}
type Handler func(context.Context,Job)error
type Queue struct{jobs chan Job; dead chan Job; wg sync.WaitGroup; handler Handler; retryDelay time.Duration}
func New(buffer int,h Handler)*Queue{return &Queue{jobs:make(chan Job,buffer),dead:make(chan Job,buffer),handler:h,retryDelay:100*time.Millisecond}}
func(q *Queue)Enqueue(j Job)error{if j.ID==""{return errors.New("job id required")};if j.MaxAttempts<1{j.MaxAttempts=3};q.jobs<-j;return nil}
func(q *Queue)Start(ctx context.Context,workers int){if workers<1{workers=1};for i:=0;i<workers;i++{q.wg.Add(1);go q.worker(ctx)}}
func(q *Queue)worker(ctx context.Context){defer q.wg.Done();for{select{case<-ctx.Done():return;case j:=<-q.jobs:j.Attempts++;if err:=q.handler(ctx,j);err!=nil{if j.Attempts>=j.MaxAttempts{select{case q.dead<-j:default:}}else{select{case<-ctx.Done():return;case<-time.After(q.retryDelay):q.jobs<-j}}}}}}
func(q *Queue)DeadLetters()<-chan Job{return q.dead}
func(q *Queue)Wait(){q.wg.Wait()}
