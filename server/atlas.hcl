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
  url = getenv("LOCAL_DATABASE_URL")
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

// atlas schema diff --env "local" --to "env://src" --from "postgresql://..."
// Hosted supabase dev database
// info? connections from SMU wifi won't work
env "dev" {
  src = data.external_schema.gorm.url
  url = getenv("DEV_DATABASE_URL")
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

