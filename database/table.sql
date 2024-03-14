CREATE TABLE credit_statistics (
    id INT PRIMARY KEY IDENTITY (1, 1),
    assignments_made INT NOT NULL,
    successful_assignments INT NOT NULL,
    failed_assignments INT NOT NULL,
    average_successful_assignments FLOAT NOT NULL,
    average_failed_assignments FLOAT NOT NULL
);
