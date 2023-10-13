package Config

type Config struct {
	DB struct {
		Url string `yaml:"url"`
	} `yaml:"db"`
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`
	HttpServer struct {
		Port string `yaml:"port"`
	} `yaml:"http_server"`
	GrpcServer struct {
		Port string `yaml:"port"`
	} `yaml:"grpc_server"`
}
