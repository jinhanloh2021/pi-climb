// Schema single source of truth from /models
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres",
  ]
}

// Local dev env
// source .env && atlas schema apply --env local
env "local" {
  src = data.external_schema.gorm.url   
  url = getenv("DATABASE_URL")
  dev = "docker://postgres/17/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  exclude = [
    "auth",
    "storage",
    "vault",
    "graphql",
    "graphql_public",
    "extensions",
    "realtime",
    "_realtime",
    "pgbouncer",
    "supabase_functions"
  ]
}

