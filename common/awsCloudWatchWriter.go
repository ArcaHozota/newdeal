package common

//
//import (
//	"context"
//	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
//	cwTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
//	"sync"
//	"time"
//)
//
//type CloudWatchWriter struct {
//	Client        *cloudwatchlogs.Client
//	LogGroupName  string
//	LogStreamName string
//	sequenceToken *string
//	mu            sync.Mutex
//}
//
//func (w *CloudWatchWriter) Write(p []byte) (n int, err error) {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//	now := time.Now().UnixMilli()
//	input := &cloudwatchlogs.PutLogEventsInput{
//		LogGroupName:  aws.String(w.LogGroupName),
//		LogStreamName: aws.String(w.LogStreamName),
//		LogEvents: []cwTypes.InputLogEvent{
//			{
//				Timestamp: aws.Int64(now),
//				Message:   aws.String(string(p)),
//			},
//		},
//		SequenceToken: w.sequenceToken,
//	}
//	resp, err := w.Client.PutLogEvents(context.TODO(), input)
//	if err != nil {
//		return 0, err
//	}
//	w.sequenceToken = resp.NextSequenceToken
//	return len(p), nil
//}
