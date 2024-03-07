CREATE TABLE IF NOT EXISTS goods (
    Id Int64,
    ProjectId Int64,
    Name String,
    Description String,
    Priority Int64,
    Removed UInt8,
    EventTime DateTime
) ENGINE = MergeTree()
ORDER BY (Id);