package pop3

import (
	"net"
	"net/textproto"
	"smtp-server/pkg/logging"
	"strings"
)

type Request struct {
	Server *Server
	Conn net.Conn
	TextProto *textproto.Conn
	RemoteAddr string
	AuthUser string
	Authable bool
	Quit     bool
	Line []string
}

func NewRequest(conn net.Conn,srv *Server) (req * Request,err error){
	req = new(Request)
	req.Server = srv
	req.RemoteAddr = conn.RemoteAddr().String()
	req.Conn = conn
	req.TextProto = textproto.NewConn(conn)
	req.Authable = true
	req.Line = []string{}

	return req,nil
}

func (req *Request) Serve(){
	defer func(){
		req.TextProto.Close()
		req.Conn.Close()
	}()

	logging.Info("pop3 client in ",req.RemoteAddr)
	err := req.TextProto.PrintfLine("+OK simple POP3 server %s powered by Go",req.Server.Host)

	if err != nil{
		logging.Debug(err)
	}

	for !req.Quit && err == nil {
		err = req.Process()
		if err != nil {
			logging.Debug(err)
			return
		}
	}
}

func (req *Request) Process() error{
	s,err := req.TextProto.ReadLine()
	if err != nil{
		logging.Debug(err)
	}
	req.Line = strings.Split(s," ")
	logging.Info(len(req.Line),req.Line)

	if len(req.Line) <= 0 {
		return req.TextProto.PrintfLine("%d %s (%s)", 500, "Command not recognized", s)
	}

	if req.Server.Processors == nil {
		req.Server.Processors = DefaultProcessors
	}

	req.Line[0] = strings.ToUpper(req.Line[0])

	processor,found := req.Server.Processors[req.Line[0]]
	
	if !found {
		return req.TextProto.PrintfLine("%d %s (%s)", 500, "Command not recognized", req.Line[0])
	}
	
	return processor(req)
}