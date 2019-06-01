package pop3

import(
	"net"
	"smtp-server/pkg/logging"
)

// Processor is a pop3 command processor
type Processor func(req *Request) error

type Server struct {
	Host string
	Addr string
	Processors map[string]Processor
}

func (srv * Server) Serve(l net.Listener) error{
	defer l.Close()

	for {
		rw,e := l.Accept()
		if e != nil{
			logging.Debug(e)
		}

		req,err := NewRequest(rw,srv)
		if err != nil{
			logging.Debug(err)
		}

		go req.Serve()
	}
}

func (srv *Server) ListenAndServe() error {
	if srv.Host == "" {
		srv.Host = "localhost"
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":smtp"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}