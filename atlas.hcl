env "local" {
  src = "file://db/schemas/schema.sql"

  format = "sql"

  url = getenv("DB_DSN")

  dev = getenv("DB_DSN_DEV")

  migration {
    dir = "file://db/migrations"
  }
}

env "staging" {
  src = "file://db/schemas"

  format = "sql"

  url = getenv("DB_DSN_STAGING")

  dev = getenv("DB_DSN_DEV")

  migration {
    dir = "file://db/migrations"
  }
}
