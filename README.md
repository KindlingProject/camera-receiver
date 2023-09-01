# camera-receiver
摄像头接收器，定义以下GRPC接口
- SendTraceIds(context.Context, *model.TraceIds) (*model.TraceIds, error)
- QueryP9X(context.Context, *model.P9XRequest) (*model.P9XResponse, error)

打包指令
> go build -v -o camera-receiver ./cmd