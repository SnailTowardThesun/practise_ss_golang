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
	"flag"
	"fmt"
	"os"
)

var _logger *ssLog
var _version string = "1.0.0"
var _author string = "ME_KUN_HAN"
var _email string = "hanvskun@hotmail.com"

func Initialize() (err error) {
		if _logger, err = NewSSLog(LOG_TANK_CONSOLE, ""); err != nil {
				return err
		}

		return nil
}

func ProInfo() {
		_logger.Trace(nil, fmt.Sprintf("SS_GOLANG %v Copyright(c) 2017", _version))
		_logger.Trace(nil, fmt.Sprintf("author: %v", _author))
		_logger.Trace(nil, fmt.Sprintf("contact: %v", _email))
}

func main() {
		var err error
		if err = Initialize(); err != nil {
				fmt.Println("initailize the project failed.")
				os.Exit(-1)
		}

		ProInfo()

		var confPath string
		flag.StringVar(&confPath, "conf", "./conf/ss.conf", "the path of config file")

		flag.Usage = func() {
				_logger.Trace(nil, fmt.Sprintf("Usage: %v [--conf=string] [-h|--help]", os.Args[0]))
				flag.PrintDefaults()
				_logger.Trace(nil, fmt.Sprintf("For example:"))
				_logger.Trace(nil, fmt.Sprintf("	%v --conf=./conf/ss.conf", os.Args[0]))
		}
		flag.Parse()

		var conf *SSConfig
		if conf, err = NewSSConfig(confPath); err != nil {
				_logger.Error(nil, fmt.Sprintf("Initialize the configure failed. err is %v", err))
				os.Exit(-1)
		}

		var server *Sock5Server
		if server, err = NewSock5Server(conf); err != nil {
				_logger.Error(nil, fmt.Sprintf("Create Sock5 server failed. err is %v", err))
				os.Exit(-1)
		}

		os.Exit(server.Run())
}
