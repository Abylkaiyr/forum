package migrations

const (
	TableForUsers = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email TEXT UNIQUE,
		username TEXT UNIQUE,
		password TEXT		
	);`

	TableForPosts = `CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY,
		owner TEXT,
		title TEXT,
		content TEXT,
		type TEXT,		
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0,
		createdTime DATETIME DEFAULT NULL,
		timer TEXT,
		image BLOB,
		filepath TEXT
	);`
	TableForComments = `CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY,
		postID INTEGER,
		owner TEXT,
		content TEXT,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0,
		likeState TEXT DEFAULT "" NOT NULL,
		dislikeState TEXT DEFAULT "" NOT NULL
	);`
	TableForSessions = `CREATE TABLE IF NOT EXISTS sessions (		
		id INTEGER PRIMARY KEY,
		owner TEXT, 			
		uuid TEXT,
		expireTime DATETIME DEFAULT NULL,
		status INTEGER
	);`
	TableForPostReactions = `CREATE TABLE IF NOT EXISTS reactions (		
		id INTEGER PRIMARY KEY,
		postID INTEGER, 			
		postLiker TEXT,
		postDisLiker TEXT,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0, 
		totalLikes INTEGER DEFAULT 0,
		totalDislikes INTEGER DEFAULT 0,
		likeState TEXT DEFAULT "" NOT NULL,
		dislikeState TEXT DEFAULT "" NOT NULL
	);`
)
