DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='lists' AND column_name='list_order') THEN
        ALTER TABLE lists ADD COLUMN list_order BIGINT DEFAULT 0;
    END IF;
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;