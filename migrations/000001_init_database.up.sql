CREATE TABLE IF NOT EXISTS folders (
    folder_id SERIAL PRIMARY KEY,
    folde_name VARCHAR(255) NOT NULL,
    folder_date TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks(
    task_id SERIAL PRIMARY KEY,
    task_name VARCHAR(255) NOT NULL,
    task_text TEXT,
    task_date TIMESTAMP,
    status BOOLEAN,
    folder_id INT,
    FOREIGN KEY (folder_id) REFERENCES folders(folder_id)
);