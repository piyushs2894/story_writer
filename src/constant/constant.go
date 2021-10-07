package constant

const (
	// Environment
	ENV_DEVELOPMENT = "development"
	ENV_STAGING     = "staging"
	ENV_PRODUCTION  = "production"

	//DB
	DriverMysql     = "mysql"
	DB_STORY_WRITER = "story_writer"

	//HTTP Client
	HTTPClientTimeout = 10

	// Length Configurations
	MAX_TITLE_LENGTH      = 2
	MAX_PARAGRAPHS_LENGTH = 7
	MAX_SENTENCE_LENGTH   = 10
	MAX_WORDS_LENGTH      = 15

	//Query Params
	LIMIT           = 100
	OFFSET          = 0
	DEFAULT_SORT_BY = "created_at"
	DEFAULT_ORDER   = "ASC"
)
