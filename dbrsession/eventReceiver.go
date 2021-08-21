package dbrsession

import "fmt"

type eventReceiver struct{}

const (
	logPrefix = "[DBR EVENT RECEIVER]"
)

func (e *eventReceiver) Event(eventName string) {
	fmt.Printf("%s called Event() event=%s\n", logPrefix, eventName)
}

func (e *eventReceiver) EventKv(eventName string, kvs map[string]string) {}

func (e *eventReceiver) EventErr(eventName string, err error) error {
	return err
}

func (e *eventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Printf("%s called EventErrKv() event=%s, err=%v, kvs=%v\n", logPrefix, eventName, err, kvs)
	return err
}
func (e *eventReceiver) Timing(eventName string, nanoseconds int64) {}

func (e *eventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {}
