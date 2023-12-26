CREATE TABLE Company (
                         id INTEGER,
                         name VARCHAR(255),
                         address VARCHAR(255),
                         PRIMARY KEY (id)
);

-- ----------------------------
-- Records of Company
-- ----------------------------
INSERT INTO Company VALUES (1, 'Company A', 'Address A');
INSERT INTO Company VALUES (2, 'Company B', 'Address B');

CREATE TABLE Scores (
                        id INTEGER,
                        company_id INTEGER,
                        score FLOAT,
                        calculate_date DATE,
                        score_grade VARCHAR(50),
                        PRIMARY KEY (id),
                        FOREIGN KEY (company_id) REFERENCES Company (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of Scores
-- ----------------------------
INSERT INTO Scores VALUES (1, 1, 85.0, '2022-05-01', 'B');
INSERT INTO Scores VALUES (2, 1, 90.0, '2022-06-01', 'A');
INSERT INTO Scores VALUES (3, 2, 75.0, '2022-05-01', 'C');
INSERT INTO Scores VALUES (4, 2, 80.0, '2022-06-01', 'D');
