package universal

type User struct {
	id              int64
	name            string
	followCount     int64
	followerCount   int64
	isFollow        bool
	avatar          string
	backgroundImage string
	signature       string
	totalFavorited  int64
	workCount       int64
	favoritesCount  int64
}

type Video struct {
	id            int64
	author        User
	playUrl       string
	coverUrl      string
	favoriteCount int64
	commentCount  int64
	isFavorite    bool
	title         string
}

type Response struct {
	statusCode int32
	statusMsg  string
}

type Comment struct {
	id        int64
	user      User
	content   string
	creatData string
}
