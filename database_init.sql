SELECT 'CREATE DATABASE cubicasa_test'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'cubicasa_test')\gexec

