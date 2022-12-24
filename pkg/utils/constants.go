package utils

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
		dislikes INTEGER DEFAULT 0
	);`
	TableForComments = `CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY,
		postID INTEGER,
		owner TEXT,
		content TEXT,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0
	);`
)
