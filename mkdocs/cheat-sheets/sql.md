---
date: 2022-12-17
draft: true
tags:
- sql
title: SQL
---

# SQL

## Good practice when running SQL query in prod

-   Have a rollback strategy
    -   For example, create a query which creates a temporary table which holds a list of ids so that this list can later be used to “roll back” changes made to other tables.
-   Run both the main query and the rollback query in staging first
-   Run the queries with a transaction without commit first, to see that it completes without errors (`BEGIN;` without `COMMIT;`)

## Repo with various ORM / raw SQL queries

[fredrikaverpil/postgres-playground: A personal PostgreSQL playground (github.com)](https://github.com/fredrikaverpil/postgres-playground)

## Run in transaction

```sql
BEGIN; -- run the below in a transaction
INSERT ... ;
COMMIT; -- will commit the changes. Comment this out to not commit the change.

```

## Avoid injection
Copied from the psycopg2 docs: [](https://www.psycopg.org/docs/sql.html)[https://www.psycopg.org/docs/sql.html](https://www.psycopg.org/docs/sql.html)
```python
query = sql.SQL("select {fields} from {table}").format(
	fields=sql.SQL(',').join([
		sql.Identifier('field1'),
		sql.Identifier('field2'),
		sql.Identifier('field3'),
	]),
	table=sql.Identifier('some_table'))
```

## SQL snippets

### Management queries

Source: [https://gist.github.com/rgreenjr/3637525](https://gist.github.com/rgreenjr/3637525)

<script src="https://gist.github.com/rgreenjr/3637525.js"></script>

```sql
-- show running queries (pre 9.2)
SELECT procpid, age(clock_timestamp(), query_start), usename, current_query 
FROM pg_stat_activity 
WHERE current_query != '<IDLE>' AND current_query NOT ILIKE '%pg_stat_activity%' 
ORDER BY query_start desc;

-- show running queries (9.2)
SELECT pid, age(clock_timestamp(), query_start), usename, query 
FROM pg_stat_activity 
WHERE query != '<IDLE>' AND query NOT ILIKE '%pg_stat_activity%' 
ORDER BY query_start desc;

-- kill running query
SELECT pg_cancel_backend(procpid);

-- kill idle query
SELECT pg_terminate_backend(procpid);

-- vacuum command
VACUUM (VERBOSE, ANALYZE);

-- all database users
select * from pg_stat_activity where current_query not like '<%';

-- all databases and their sizes
select * from pg_user;

-- all tables and their size, with/without indexes
select datname, pg_size_pretty(pg_database_size(datname))
from pg_database
order by pg_database_size(datname) desc;

-- cache hit rates (should not be less than 0.99)
SELECT sum(heap_blks_read) as heap_read, sum(heap_blks_hit)  as heap_hit, (sum(heap_blks_hit) - sum(heap_blks_read)) / sum(heap_blks_hit) as ratio
FROM pg_statio_user_tables;

-- table index usage rates (should not be less than 0.99)
SELECT relname, 100 * idx_scan / (seq_scan + idx_scan) percent_of_times_index_used, n_live_tup rows_in_table
FROM pg_stat_user_tables 
ORDER BY n_live_tup DESC;

-- how many indexes are in cache
SELECT sum(idx_blks_read) as idx_read, sum(idx_blks_hit)  as idx_hit, (sum(idx_blks_hit) - sum(idx_blks_read)) / sum(idx_blks_hit) as ratio
FROM pg_statio_user_indexes;

-- Dump database on remote host to file
$ pg_dump -U username -h hostname databasename > dump.sql

-- Import dump into existing database
$ psql -d newdb -f dump.sql
```


### Misc

Simple count:
```sql
SELECT COUNT(*) from table1 t1;
```