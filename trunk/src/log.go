/*
MIT License

Copyright (c) 2017 ME_Kun_Han

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LOG_LEVEL_INFO    = "info"
	LOG_LEVEL_VERBOSE = "verbose"
	LOG_LEVEL_TRACE   = "trace"
	LOG_LEVEL_WARN    = "warn"
	LOG_LEVEL_ERROR   = "error"
)

const (
	LOG_TANK_CONSOLE = "console"
	LOG_TANK_FILE    = "file"
)

type Context interface {
	GetID() int
}

type ssLog struct {
	log *log.Logger
}

func NewSSLog(tank, path string) (*ssLog, error) {
	var logger *ssLog
	var err error

	if tank == LOG_TANK_CONSOLE {
		logger = &ssLog{
			log: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds),
		}
	} else if tank == LOG_TANK_FILE {
		if len(path) == 0 {
			err = fmt.Errorf(fmt.Sprintf("the path=%v of log file is empty", path))
			return nil, err
		}

		logger = &ssLog{
			log: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds),
		}

		// TODO:FIXME: implement tank as file
	} else {
		err = fmt.Errorf(fmt.Sprintf("the tank type=%v is invalid", tank))
		return nil, err
	}

	return logger, nil
}

func (v *ssLog) Info(ctx Context, a ...interface{}) error {
	return v.Log(ctx, LOG_LEVEL_INFO, a)
}

func (v *ssLog) Verbose(ctx Context, a ...interface{}) error {
	return v.Log(ctx, LOG_LEVEL_VERBOSE, a)
}

func (v *ssLog) Trace(ctx Context, a ...interface{}) error {
	return v.Log(ctx, LOG_LEVEL_TRACE, a)
}

func (v *ssLog) Warn(ctx Context, a ...interface{}) error {
	return v.Log(ctx, LOG_LEVEL_WARN, a)
}

func (v *ssLog) Error(ctx Context, a ...interface{}) error {
	return v.Log(ctx, LOG_LEVEL_ERROR, a)
}

func (v *ssLog) Log(ctx Context, logLevel string, a ...interface{}) error {
	curTime := time.Now().UTC().Format(time.UnixDate)
	if ctx == nil {
		a = append([]interface{}{fmt.Sprintf("[%v][%v][%v]", curTime, os.Getpid(), logLevel)}, a...)
	} else {
		a = append([]interface{}{fmt.Sprintf("[%v][%v][%v][%v]", curTime, os.Getpid(), ctx.GetID(), logLevel)}, a...)
	}

	v.log.Println(a...)
	return nil
}
