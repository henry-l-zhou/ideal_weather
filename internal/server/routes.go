package server

func (s *Server) setupRoutes() {
	s.router.Get("/weather/{city}", s.getWeather)
}
