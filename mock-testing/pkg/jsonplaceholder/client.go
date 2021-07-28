package jsonplaceholder

type Client struct {
	Post PostService
}

func NewClient() Client {
	return Client{
		Post: &PostServiceClient{},
	}
}
