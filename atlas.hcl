env "local" {
  src = "file://db/schemas"

  url = getenv("DB_DSN")

  dev = getenv("DB_DSN_DEV")

  migration {
    dir = "file://db/migrations"
  }
}

env "staging" {
  src = "file://db/schemas"

  url = getenv("DB_DSN_STAGING")

  dev = getenv("DB_DSN_DEV")

  migration {
    dir = "file://db/migrations"
  }
}
