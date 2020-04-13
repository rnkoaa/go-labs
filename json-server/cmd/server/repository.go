package main

//type Repository interface {
//}

type RepositoryOp struct {
	db      *Database
	Album   AlbumService
	Comment CommentService
	Photo   PhotoService
	Post    PostService
	Todo    TodoService
	User    UserService
}

//var _ Repository = &RepositoryOp{}

func NewRepository(db *Database) *RepositoryOp {
	var repo = &RepositoryOp{}
	repo.User = &UserServiceOp{
		db: db,
	}
	repo.Album = &AlbumServiceOp{
		db: db,
	}
	repo.Comment = &CommentServiceOp{
		db: db,
	}
	repo.Photo = &PhotoServiceOp{
		db: db,
	}
	repo.Post = &PostServiceOp{
		db: db,
	}
	repo.Todo = &TodoServiceOp{
		db: db,
	}

	return repo
}
