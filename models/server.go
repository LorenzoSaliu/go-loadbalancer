
package models

type Server struct{
	address string
	isAlive bool
}

func CreateNewServer(address string) *Server {
	return &Server{
		address: address,
		isAlive: true,
	}	
}
func (s *Server) GetAdd() string{
	return s.address
}
func (s *Server) SetAdd(add string){
	s.address = add
}

func (s *Server) GetStatus() bool{
	return s.isAlive
}
func (s *Server) SetStatus(status bool){
	s.isAlive = status
}
