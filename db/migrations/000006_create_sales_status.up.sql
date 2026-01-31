DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'sales_status'
    ) THEN
        CREATE TYPE sales_status AS ENUM (
            'unpaid',
            'partial',
            'paid',
            'cancelled'
        );
    END IF;
END $$;
