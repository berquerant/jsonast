package jsonast

type (
	StringLike interface {
		~string
	}

	IntLike interface {
		~int
	}

	StringOrIntLike interface {
		StringLike | IntLike
	}
)
