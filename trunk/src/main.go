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
	if err := Initialize(); err != nil {
		fmt.Println("initailize the project failed.")
		os.Exit(-1)
	}

	ProInfo()

	var listen string
	flag.StringVar(&listen, "listen", ":1080", "the http server listen at")

	flag.Usage = func() {
		_logger.Trace(nil, fmt.Sprintf("Usage: %v [--listen=string] [-h|--help]", os.Args[0]))
		flag.PrintDefaults()
		_logger.Trace(nil, fmt.Sprintf("For example:"))
		_logger.Trace(nil, fmt.Sprintf("	%v --listen=:2033", os.Args[0]))
	}
	flag.Parse()
	
	conn := NewSock5Conn()
	
	_logger.Trace(conn, fmt.Sprintf("the server is listening at %v", listen))
	if err := conn.Listen(listen); err != nil {
		_logger.Error(conn, "")
		os.Exit(-1)
	}
}
