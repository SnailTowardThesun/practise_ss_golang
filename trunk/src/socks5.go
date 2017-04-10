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

type Sock5Connection struct {
	ConnID uint32
}

var Sock5ConnBasicID uint32 = 100

func NewSock5Conn() *Sock5Connection {
	conn := &Sock5Connection{
		ConnID : Sock5ConnBasicID,
	}
	
	// when the basic id is too large, roll over again.
	if Sock5ConnBasicID > 0xffffff00 {
		Sock5ConnBasicID = 100
	} else {
		Sock5ConnBasicID += 1
	}
	
	return conn
}

func (v *Sock5Connection) GetID() uint32 {
	return v.ConnID
}

func (v *Sock5Connection) Listen(port string) (err error) {
	return
}
