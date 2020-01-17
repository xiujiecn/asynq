// Copyright 2020 Kentaro Hibino. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package asynq provides a framework for background task processing.

Asynq uses Redis as a message broker. To connect to redis server,
specify the options using one of RedisConnOpt types.

    redis = &asynq.RedisClientOpt{
        Addr:     "localhost:6379",
        Password: "secretpassword",
        DB:       3,
    }

The Client is used to register a task to be processed at the specified time.

Task is created with two parameters: its type and payload.

    client := asynq.NewClient(redis)

    t := asynq.NewTask(
        "send_email",
        map[string]interface{}{"user_id": 42})

    // Schedule the task t to be processed a minute from now.
    err := client.Schedule(t, time.Now().Add(time.Minute))

The Background is used to run the background task processing with a given
handler.
    bg := asynq.NewBackground(redis, &asynq.Config{
        Concurrency: 10,
    })

    bg.Run(handler)

Handler is an interface with one method ProcessTask which
takes a task and returns an error. Handler should return nil if
the processing is successful, otherwise return a non-nil error.
If handler panics or returns a non-nil error, the task will be retried in the future.

Example of a type that implements the Handler interface.
    type TaskHandler struct {
        // ...
    }

    func (h *TaskHandler) ProcessTask(task *asynq.Task) error {
        switch task.Type {
        case "send_email":
            id, err := task.Payload.GetInt("user_id")
            // send email
        //...
        default:
            return fmt.Errorf("unexpected task type %q", task.Type)
        }
        return nil
    }
*/
package asynq
