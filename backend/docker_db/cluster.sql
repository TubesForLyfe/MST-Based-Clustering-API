Use cluster_db;

Create table `cluster_log` (
    datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(), 
    filename VARCHAR(200) NOT NULL,
    cluster_amount INT NOT NULL,
    cluster_result VARCHAR(5000) NOT NULL,
    PRIMARY KEY (datetime)
);