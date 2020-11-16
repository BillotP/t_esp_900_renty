// Package models contains all db schemes
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/BillotP/gorenty"
)

// JobStatus represent a queued job status
type JobStatus string

const (
	// JobStatusTodo is a job todo
	JobStatusTodo = JobStatus("TODO")
	// JobStatusDone is a finished job
	JobStatusDone = JobStatus("DONE")
	// JobStatusError is an error job
	JobStatusError = JobStatus("ERROR")
)

var (
	// DefaultDoneExpiracy is the validity for the done jobs queue (2 weeks)
	DefaultDoneExpiracy = time.Duration(14 * (24 * time.Hour))
	// DefaultTodoExpiracy is the validity for the todo jobs queue (8 days)
	DefaultTodoExpiracy = time.Duration(8 * (24 * time.Hour))
	// DefaultErrorExpiracy is the validity for error in job (2 weeks)
	DefaultErrorExpiracy = time.Duration(14 * (24 * time.Hour))
)

// Key return a formatted queue key for a job
func (j JobStatus) Key(annonceID string) string {
	return annonceID + "-" + string(j)
}

// RequestError is an error after a get phonenumber error
type RequestError struct {
	Status   int    `json:"status"`
	Response string `json:"response"`
}

func (r RequestError) Error() string {
	return fmt.Sprintf("Got status %v, res : [%s]", r.Status, r.Response)
}

// Job is a job todo or with error
type Job struct {
	ID                 string        `json:"annonceId"`
	Phone              string        `json:"phone"`
	Type               string        `json:"annonceType"`
	Summary            string        `json:"summary"`
	Departement        string        `json:"department"`
	CollectedAt        time.Time     `json:"collectedAt"`
	ScrappingRequestID uuid.UUID     `json:"scrappingRequestId"`
	Error              *RequestError `json:"error,omitempty"`
}

// MustGetQueue return the whole job queue
func MustGetQueue(redisClient *redis.Client) (queue []Job) {
	var err error
	var elems []Job
	var ctx = context.Background()
	var allJobs = "*-*"
	iter := redisClient.Scan(ctx, 0, allJobs, 0).Iterator() // Get all value in queue
	for iter.Next(ctx) {
		var item Job
		var jobData string
		jobKey := iter.Val()
		if jobData, err = redisClient.Get(ctx, jobKey).Result(); err != nil {
			fmt.Printf("Error(MustGetQueue): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		if err = json.Unmarshal([]byte(jobData), &item); err != nil {
			fmt.Printf("Error(MustGetQueue): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		if item.ScrappingRequestID == uuid.Nil &&
			item.Summary == "" {
			if goscrappy.Debug {
				fmt.Printf("Info(MustGetQueue): %+v is not a todo or done job\n", item)
			}
			continue
		}
		elems = append(elems, item)
	}
	if goscrappy.Debug {
		fmt.Printf("Info(GetDoneJobs): Got %v job(s) from redis queue\n", len(elems))
	}
	return elems
}

// GetValid from the phone number slice
func GetValid(redisClient *redis.Client, phones []Job) (valid []Job) {
	valid = []Job{}
	var ctx = context.Background()
	for el := range phones {
		if phones[el].Phone != "" &&
			phones[el].ID != "" {
			todoJobForKey := JobStatusTodo.Key(phones[el].ID)
			doneJobForKey := JobStatusDone.Key(phones[el].ID)
			if exist, err := redisClient.Get(
				ctx,
				todoJobForKey,
			).Result(); (err != nil && err != redis.Nil) || exist != "" { // Skip if already done
				if goscrappy.Debug && err == nil {
					fmt.Printf(
						"Info(GetValid): Job key [%s] is already in queue [%s]\n",
						todoJobForKey,
						exist,
					)
				}
				if err != nil {
					fmt.Printf("Error(GetValid): Failed to get job(s) : %s\n", err.Error())
				}
				continue
			} else if exist, err := redisClient.Get(
				ctx,
				doneJobForKey,
			).Result(); (err != nil && err != redis.Nil) || exist != "" { // Skip if already done
				if goscrappy.Debug && err == nil {
					fmt.Printf(
						"Info(GetValid): Job key [%s] is already DONE [%s]\n",
						doneJobForKey,
						exist,
					)
				}
				if err != nil {
					fmt.Printf("Error(GetValid): Failed to get job(s) : %s\n", err.Error())
				}
				continue
			} else if goscrappy.Debug {
				fmt.Printf("Info(GetValid): Job [%s] doesn't exist yet\n", doneJobForKey)
			}

			valid = append(valid, phones[el])
		}
	}
	return valid
}

// GetDoneJobs from redis queue
func GetDoneJobs(redisClient *redis.Client) []DoneJobItem {
	var err error
	var elems []DoneJobItem
	var ctx = context.Background()
	var allDONEJobs = JobStatusDone.Key("*")
	iter := redisClient.Scan(ctx, 0, allDONEJobs, 0).Iterator() // Get all value in queue
	for iter.Next(ctx) {
		var item DoneJob
		var jobData string
		jobKey := iter.Val()
		if jobData, err = redisClient.Get(ctx, jobKey).Result(); err != nil {
			fmt.Printf("Error(GetDoneJobs): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		if err = json.Unmarshal([]byte(jobData), &item); err != nil {
			fmt.Printf("Error(GetDoneJobs): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		elems = append(elems, DoneJobItem{
			DoneJob: item,
		})
	}
	if goscrappy.Debug {
		fmt.Printf("Info(GetDoneJobs): Got %v done job(s) from redis queue\n", len(elems))
	}
	return elems
}

// GetFailedJobs from redis queue (filtered by time range)
func GetFailedJobs(redisClient *redis.Client, fromDate, toDate *time.Time) []Job {
	var err error
	var elems []Job
	var ctx = context.Background()
	var allERRORJobs = JobStatusError.Key("*")
	iter := redisClient.Scan(ctx, 0, allERRORJobs, 0).Iterator() // Get all value in queue
	for iter.Next(ctx) {
		var item Job
		var jobData string
		jobKey := iter.Val()
		if jobData, err = redisClient.Get(ctx, jobKey).Result(); err != nil {
			fmt.Printf("Error(GetFailedJobs): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		if err = json.Unmarshal([]byte(jobData), &item); err != nil {
			fmt.Printf("Error(GetFailedJobs): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		if fromDate != nil && toDate != nil {
			if !(item.CollectedAt.Before(*toDate) &&
				item.CollectedAt.After(*fromDate)) {
				if goscrappy.Debug {
					fmt.Printf("Info(GetFailedJobs): Job key %s is not in date range\n", jobKey)
				}
				continue
			}
		}
		elems = append(elems, item)
	}
	if goscrappy.Debug {
		fmt.Printf("Info(GetDoneJobs): Got %v error job(s) from redis queue\n", len(elems))
	}
	return elems
}

// GetErrors from the phone number slice
func GetErrors(redisClient *redis.Client, phones []Job) (errors []Job) {
	errors = []Job{}
	var ctx = context.Background()
	for el := range phones {
		if phones[el].Phone == "" ||
			phones[el].Error != nil {
			todoJobForKey := JobStatusTodo.Key(phones[el].ID)
			doneJobForKey := JobStatusDone.Key(phones[el].ID)
			if exist, err := redisClient.Get(
				ctx,
				todoJobForKey,
			).Result(); (err != nil && err != redis.Nil) || exist != "" { // Skip if already done
				if goscrappy.Debug && err == nil {
					fmt.Printf(
						"Info(GetErrors): Job key [%s] is already in queue [%s]\n",
						todoJobForKey,
						exist,
					)
				}
				if err != nil {
					fmt.Printf("Error(GetErrors): Failed to get job(s) : %s\n", err.Error())
				}
				continue
			} else if exist, err := redisClient.Get(
				ctx,
				doneJobForKey,
			).Result(); (err != nil && err != redis.Nil) || exist != "" { // Skip if already done
				if goscrappy.Debug && err == nil {
					fmt.Printf(
						"Info(GetErrors): Job key [%s] is already DONE [%s]\n",
						doneJobForKey,
						exist,
					)
				}
				if err != nil {
					fmt.Printf("Error(GetErrors): Failed to get job(s) : %s\n", err.Error())
				}
				continue
			} else if goscrappy.Debug {
				fmt.Printf("Info(GetErrors): Job [%s] doesn't exist yet\n", doneJobForKey)
			}

			errors = append(errors, phones[el])
		}
	}
	return errors
}

// CountBlocked request (403) in error queue
func CountBlocked(jobs []Job) int64 {
	cnt := int64(0)
	elemscount := len(jobs)
	for el := range jobs {
		if jobs[el].Error == nil {
			if goscrappy.Debug {
				fmt.Printf("Info(CountBlocked): Job %v / %v has no error\n",
					el,
					elemscount)
			}
			continue
		} else if jobs[el].Error.Status == http.StatusForbidden {
			cnt++
		}
	}
	if goscrappy.Debug {
		fmt.Printf("Info(CountBlocked): Found %v 403'd request on %v total\n", cnt, elemscount)
	}
	return cnt
}

// CountValid count valid phone numbers
func CountValid(phones []Job) int {
	cnt := 0
	for el := range phones {
		if phones[el].Phone != "" {
			cnt++
		}
	}
	return cnt
}

// CountInvalid count error for 400 or 410 errors
func CountInvalid(phones []Job) (cnt int) {
	cnt = 0
	for el := range phones {
		if phones[el].Phone == "" &&
			phones[el].Error != nil &&
			phones[el].Error.Status != http.StatusForbidden {
			cnt++
		}
	}
	return cnt
}

// ClearTodo clear queue for done jobs (represented by their keys in done variable)
func ClearTodo(redisClient *redis.Client, done []string) {
	cleared := 0
	var ctx = context.Background()
	for _, key := range done {
		if _, err := redisClient.Del(ctx, key).Result(); err != nil {
			fmt.Printf("Error(ClearTodo): %s\n", err.Error())
		}
		if goscrappy.Debug {
			fmt.Printf("Info(ClearTodo): Job key [%s] successfully removed\n", key)
		}
		cleared++
	}
	if goscrappy.Debug {
		fmt.Printf("Info(ClearTodo): Successfully removed %v TODO job(s)\n", cleared)
	}
}

// ClearAllQueue clear all the redis queue items (!USE WITH CAUTION)
func ClearAllQueue(redisClient *redis.Client) (err error) {
	allQueue := "*"
	var deleted int64
	var ctx = context.Background()
	iter := redisClient.Scan(ctx, 0, allQueue, 0).Iterator() // Get all value in queue
	for iter.Next(ctx) {
		var ok int64
		jobKey := iter.Val()
		if ok, err = redisClient.Del(ctx, jobKey).Result(); err != nil {
			fmt.Printf("Error(ClearAllQueue): Failed to get key %s: %s\n", jobKey, err.Error())
			iter.Next(ctx)
			continue
		}
		deleted += ok
	}
	if goscrappy.Debug {
		fmt.Printf("Info(ClearAllQueue): Having dropped %v items from redis queue\n", deleted)
	}
	return nil
}

// ClearAllDoneJobTestQueue clear all the test doneJob in redis queue items
func ClearAllDoneJobTestQueue(redisClient *redis.Client) (err error) {
	var cleared = 0
	var ctx = context.Background()
	allDonejobs := GetDoneJobs(redisClient)
	for el := range allDonejobs {
		if allDonejobs[el].IsTest() {
			jobKey := JobStatusDone.Key(allDonejobs[el].AnnonceID)
			if err = redisClient.Del(ctx, jobKey).Err(); err != nil {
				return err
			}
			cleared++
		}
	}
	if goscrappy.Debug {
		fmt.Printf("Info(ClearAllDoneJobTestQueue): Having dropped %v items from redis queue\n", cleared)
	}
	return nil
}

// SaveDone save queue for done jobs
func SaveDone(redisClient *redis.Client, done map[string]DoneJob) {
	saved := 0
	var ctx = context.Background()
	for jobID, job := range done {
		var err error
		var saveByte []byte
		if saveByte, err = json.Marshal(job); err != nil {
			fmt.Printf("Error(SaveDone): %s\n", err.Error())
			continue
		}
		key := JobStatusDone.Key(jobID)
		if _, err = redisClient.Set(
			ctx,
			key,
			string(saveByte),
			DefaultDoneExpiracy,
		).Result(); err != nil {
			fmt.Printf("Error(SaveDone): %s\n", err.Error())
			continue
		}
		if goscrappy.Debug {
			fmt.Printf("Info(SaveDone): Job key [%s] successfully saved\n", key)
		}
		saved++
	}
	if goscrappy.Debug {
		fmt.Printf("Info(SaveDone): Successfully saved %v DONE job(s)\n", saved)
	}
}

// SaveTodo queue for all gathered phone numbers
func SaveTodo(redisClient *redis.Client, valid []Job) {
	saved := 0
	var ctx = context.Background()
	for el := range valid { // Save new jobs in queue
		todoKey := JobStatusTodo.Key(valid[el].ID)
		datas, _ := json.Marshal(valid[el]) // Serialize the phone numbers
		if _, err := redisClient.Set(       // Save it
			ctx,
			todoKey,
			datas,
			DefaultTodoExpiracy,
		).Result(); err != nil {
			fmt.Printf("Error(Handle): Failed to save job queue : %s\n", err.Error())
			os.Exit(1)
		}
		saved++
	}
	if goscrappy.Debug {
		fmt.Printf("Info(SaveQueue): Successfully saved %v job(s)\n", saved)
	}
}

// SaveErrrors for all blocked or invalid phone numbers request jobs
func SaveErrrors(redisClient *redis.Client, errors []Job) {
	saved := 0
	var ctx = context.Background()
	for el := range errors { // Save new jobs in queue
		errorKey := JobStatusError.Key(errors[el].ID)
		datas, _ := json.Marshal(errors[el]) // Serialize the phone numbers
		if _, err := redisClient.Set(        // Save it
			ctx,
			errorKey,
			datas,
			DefaultErrorExpiracy,
		).Result(); err != nil {
			fmt.Printf("Error(Handle): Failed to save job queue : %s\n", err.Error())
			os.Exit(1)
		}
		saved++
	}
	if goscrappy.Debug {
		fmt.Printf("Info(SaveQueue): Successfully saved %v job(s)\n", saved)
	}
}

// FilterAnnonceResults filter annonce scrapping results
func FilterAnnonceResults(queue []Job, results []Job) (todo []Job) {
	filtered := 0
	todo = []Job{}
	for el := range results {
		if IsInQueue(queue, results[el].ID) {
			filtered++
			continue
		}
		todo = append(todo, results[el])
	}
	if goscrappy.Debug {
		fmt.Printf("Info(FilterAnnonceResults): Filtering %v existing job(s)\n", filtered)
	}
	return todo
}

// IsInQueue check if annonceID is allready in queue
func IsInQueue(queue []Job, annonceID string) bool {
	for el := range queue {
		if queue[el].ID == annonceID {
			if queue[el].Error != nil &&
				queue[el].Error.Status != http.StatusForbidden {
				break
			}
			return true
		}
	}
	return false
}
