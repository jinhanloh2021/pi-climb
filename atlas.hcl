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

// Local Development Environment
// supabase start
env "local" {
  src = data.external_schema.gorm.url   
  url = "postgresql://postgres:postgres@127.0.0.1:54322/postgres?sslmode=disable" // local supabase DB
  dev = "docker://postgres/15/dev"
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
env "dev" {
  src = data.external_schema.gorm.url
  url = "postgres://[YOUR_SUPABASE_DEV_USER]:[YOUR_SUPABASE_DEV_PASSWORD]@[YOUR_SUPABASE_DEV_HOST]:5432/[YOUR_SUPABASE_DEV_DBNAME]?sslmode=require"
  dev = "docker://postgres/15/dev"
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

