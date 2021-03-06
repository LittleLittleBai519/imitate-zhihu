package enum

const (
	ByTime   = "time"
	ByHeat   = "heat"
	ByUpvote = "upvote"
)

const (
	Question = iota
	Answer
	Comment
)

const (
	ViewCount = "ViewCount"
	UpvoteCount = "UpvoteCount"
	CommentCount = "CommentCount"
	AnswerCount = "AnswerCount"
	DownvoteCount = "DownvoteCount"
)

const (
	StateNoVote = 0
	StateUpVote = 1
	StateDownVote = 2
)

