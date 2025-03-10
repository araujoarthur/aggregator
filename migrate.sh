if [ "$#" -ne 1 ]; then
   echo "Usage: $0 <(u)p/(d)own>"
   exit 1
fi

migration_type="$1"
DB_URL="postgresql://arthur@/gator?host=/tmp" 
MIGRATIONS_DIR="sql/schema"


if [ "$migration_type" = "up" ] || [ "$migration_type" = "u" ]; then
   echo "Running migrations"
   goose -dir "$MIGRATIONS_DIR" postgres "$DB_URL" up;
elif [ "$migration_type" = "down" ] || [ "$migration_type" = "d" ]; then
   echo "Running migrations"
   goose -dir "$MIGRATIONS_DIR" postgres "$DB_URL" down;
else
    echo "Invalid argument: $migration_type"
    echo "Usage: $0 <(u)p/(d)own>"
    exit 1
fi