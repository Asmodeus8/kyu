package main
import("context";"fmt";"time";"github.com/Asmodeus8/kyu/queue")
func main(){ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second);defer cancel();q:=queue.New(32,func(_ context.Context,j queue.Job)error{fmt.Printf("processed %s attempt=%d\n",j.ID,j.Attempts);return nil});q.Start(ctx,4);for i:=1;i<=10;i++{q.Enqueue(queue.Job{ID:fmt.Sprintf("job-%d",i),Payload:[]byte("demo")})};<-ctx.Done();q.Wait()}
